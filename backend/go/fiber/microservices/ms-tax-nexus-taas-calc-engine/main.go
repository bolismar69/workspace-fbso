package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http" // <--- ADICIONE ESTA LINHA
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type CallbackRes struct {
	IDCadastroUnico  string `json:"id_cadastro_unico"`
	StatusIntegracao string `json:"status_integracao"`
}

type TaxResponse struct {
	// Parâmetros de Entrada (Echo)
	IBGE string `json:"ibge"`

	// Dados Geográficos e Temporais
	Municipio  string      `json:"municipio"`
	UF         string      `json:"uf"`
	NCM        string      `json:"ncm"`
	Ano        int         `json:"ano"`
	PIS        float64     `json:"pis"`
	COFINS     float64     `json:"cofins"`
	IPI        float64     `json:"ipi"`
	ICMS       float64     `json:"icms"`
	ISS        float64     `json:"iss"`
	CBS        float64     `json:"cbs_calculada"`
	IBS        float64     `json:"ibs_calculado"`
	IS         float64     `json:"imposto_seletivo"`
	IPVA       float64     `json:"ipva_novo"`
	ITCMD      float64     `json:"itcmd_novo"`
	CargaTotal float64     `json:"carga_total_estimada"`
	Callback   interface{} `json:"callback"`
}

func main() {
	ConnectDB()
	ConnectRedis()
	defer DB.Close()
	defer RDB.Close()

	app := fiber.New(fiber.Config{AppName: "TaxNexus Calc Engine v1"})
	app.Use(logger.New())

	app.Get("/v1/simulate/:ncm/:ibge", func(c *fiber.Ctx) error {
		ncm := c.Params("ncm")
		ibge := c.Params("ibge")
		cacheKey := "tax:" + ncm + ":" + ibge

		// 1. CACHE HIT?
		cachedData, err := RDB.Get(ctx, cacheKey).Result()
		if err == nil {
			var res TaxResponse
			json.Unmarshal([]byte(cachedData), &res)
			c.Set("X-Cache", "HIT")
			return c.JSON(res)
		}

		// 2. BUSCA NA VIEW (Toda a lógica de transição e somas está aqui)
		query := `
				SELECT 
					nome_municipio, uf,
					aliq_pis, aliq_cofins, aliq_icms, aliq_iss, aliq_ipi,
					cbs_calculada, ibs_calculado, is_calculado
				FROM tax_nexus_taas.vw_tax_simulator_reforma 
				WHERE codigo_ncm = $1 AND codigo_ibge = $2::int4 AND ano_competencia = $3
				LIMIT 1`

		var res TaxResponse
		res.IBGE = ibge             // Captura o valor da URL: "3550308"
		res.NCM = ncm               // Captura o valor da URL: "62011100"
		res.Ano = time.Now().Year() // Competência atual (pode ser parametrizado futuramente)
		var aPis, aCofins, aIcms, aIss, aIpi, aCbs, aIbs, aIs float64

		err = DB.QueryRow(context.Background(), query, ncm, ibge, res.Ano).Scan(
			&res.Municipio, &res.UF,
			&aPis, &aCofins, &aIcms, &aIss, &aIpi,
			&aCbs, &aIbs, &aIs,
		)

		if err != nil {
			log.Printf("Erro no Scan: %v", err)
			// Substituímos http.StatusNotFound por fiber.StatusNotFound
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Regra não encontrada",
				"parametros_enviados": fiber.Map{
					"ncm":  ncm,
					"ibge": ibge,
				},
			})
		}

		// 3. CÁLCULO MONETÁRIO (Simples e Direto)
		valorBase := 1000.0 // Futuramente vindo do Body/QueryParam

		// Aplicamos as alíquotas (dividindo por 100 pois no banco estão como 18.00)
		res.PIS = valorBase * (aPis / 100)
		res.COFINS = valorBase * (aCofins / 100)
		res.ICMS = valorBase * (aIcms / 100)
		res.ISS = valorBase * (aIss / 100)
		res.IPI = valorBase * (aIpi / 100)
		res.CBS = valorBase * (aCbs / 100)
		res.IBS = valorBase * (aIbs / 100)
		res.IS = valorBase * (aIs / 100)

		// Extra: IPVA/ITCMD (Funções auxiliares)
		res.IPVA = calcularNovoIPVA(valorBase, "veiculo_luxo")
		res.ITCMD = calcularNovoITCMD(valorBase, (aIbs / 100))

		// CARGA TOTAL
		res.CargaTotal = res.PIS + res.COFINS + res.ICMS + res.ISS + res.IPI + res.CBS + res.IBS + res.IS + res.IPVA + res.ITCMD

		res.Callback = CallbackRes{IDCadastroUnico: "PENDENTE", StatusIntegracao: "PROCESSADO"}

		// 4. SAVE CACHE & RETURN
		jsonData, _ := json.Marshal(res)
		RDB.Set(ctx, cacheKey, jsonData, 1*time.Hour)

		c.Set("X-Cache", "MISS")
		return c.JSON(res)
	})

	log.Fatal(app.Listen(":3000"))
}

// Funções de apoio permanecem iguais (IPVA/ITCMD)...

// A reforma permite alíquotas progressivas e diferenciação pelo impacto ambiental.
func calcularNovoIPVA(valorBase float64, tipoVeiculo string) float64 {
	// Alíquota base sugerida
	aliquota := 0.04 // 4% padrão (Ex: São Paulo)

	switch tipoVeiculo {
	case "aeronave", "embarcacao_luxo":
		// Novidade da Reforma: Incidência sobre jatinhos e iates luxuosos
		aliquota = 0.06
	case "veiculo_luxo":
		// Progressividade por valor do bem
		aliquota = 0.05
	case "eletrico", "hibrido":
		// Diferenciação por impacto ambiental (Incentivo)
		aliquota = 0.02
	}

	return valorBase * aliquota
}

// calcularNovoITCMD implementa a progressividade obrigatória da reforma (até 8%)
func calcularNovoITCMD(valorBase float64, aliquotaBase float64) float64 {
	// A reforma torna a progressividade obrigatória.
	// Exemplo de faixas simplificadas:
	aliquota := aliquotaBase

	if valorBase > 1000000 {
		aliquota = 0.08 // Teto de 8% para grandes valores
	} else if valorBase > 500000 {
		aliquota = 0.06
	} else if aliquota < 0.02 {
		aliquota = 0.02 // Piso sugerido
	}

	return valorBase * aliquota
}

// Função para consultar o Calc Engine (Fiber)
func CallCalcEngine(ncm, ibge string) (map[string]interface{}, error) {
	url := "http://localhost:3000/v1/simulate/" + ncm + "/" + ibge
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}
