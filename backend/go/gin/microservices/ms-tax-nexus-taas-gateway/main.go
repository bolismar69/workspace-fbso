package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type TaxRequest struct {
	CNPJ              string  `json:"cnpj"`
	NCM               string  `json:"ncm"`
	IBGE              string  `json:"ibge"`
	SaldoRemanescente float64 `json:"saldo_remanescente"` 
	CallbackURL       string  `json:"callback_url"`      
}

type CallbackRes struct {
	IDCadastroUnico  string `json:"id_cadastro_unico"`
	StatusIntegracao string `json:"status_integracao"`
}

type TaxResponse struct {
	TransactionStatus string      `json:"transaction_status"`
	Calculation       interface{} `json:"calculation"`
	Callback          CallbackRes `json:"callback"`
}

func main() {
	ConnectDB()
	defer DB.Close()

	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/v1/tax/calculate", func(c *gin.Context) {
		var req TaxRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 1. Chamar o Calc Engine (Fiber)
		// CORREÇÃO AQUI: Usando req.NCM e req.IBGE (conforme definido na struct)
		calcResult, err := CallCalcEngine(req.NCM, req.IBGE) 
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Calc Engine indisponível"})
			return
		}

		// 2. Persistir Transação no Postgres (Auditoria)
		query := `INSERT INTO tax_nexus_taas.TB_TAX_TRANSACTION 
                  (cnpj_emissor, status, payload_original) 
                  VALUES ($1, 'CONCLUIDO', $2)`
		
		payload, _ := json.Marshal(req)
		_, err = DB.Exec(context.Background(), query, req.CNPJ, payload)
		if err != nil {
			// Log do erro de banco, mas não trava a resposta ao cliente
			println("Erro auditoria:", err.Error())
		}

		// 3. Retornar resposta unificada com o Callback (mesmo que placeholder)
		c.JSON(http.StatusOK, TaxResponse{
			TransactionStatus: "audited",
			Calculation:       calcResult,
			Callback: CallbackRes{
				IDCadastroUnico:  "PENDENTE",
				StatusIntegracao: "AGUARDANDO_CADASTRO_UNICO",
			},
		})
	})

	r.Run(":8080")
}
