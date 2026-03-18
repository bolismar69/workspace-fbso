package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	dsn := "postgres://bolismar:admin_tax@localhost:5432/tax_nexus_db?search_path=tax_nexus_taas"
	pool, _ := pgxpool.New(context.Background(), dsn)
	DB = pool
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
