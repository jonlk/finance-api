package main

import (
	"log"
	"net/http"

	"github.com/jonlk/finance-api/internal/handlers"
)

func main() {
	//here is a comment
	port := ":3000"

	server := &http.Server{
		Addr:    port,
		Handler: mux(),
	}

	log.Printf("service listening on port %v\n", port)
	log.Fatal(server.ListenAndServe())
}

func mux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /basicLiquidityRatio", handlers.BasicLiquidityRatio)
	mux.HandleFunc("GET /breakEvenPoint", handlers.BreakEvenPoint)
	mux.HandleFunc("GET /cashFlow", handlers.CashFlow)
	mux.HandleFunc("GET /ruleOf72", handlers.RuleOf72)
	mux.HandleFunc("GET /netIncome", handlers.NetIncome)
	mux.HandleFunc("GET /netWorth", handlers.NetWorth)
	mux.HandleFunc("GET /peRatio", handlers.PERatio)
	mux.HandleFunc("GET /simpleInterest", handlers.SimpleInterest)
	mux.HandleFunc("GET /variationOfInvestment", handlers.VariationOfInvestment)
	mux.HandleFunc("GET /compoundInterest", handlers.CompoundInterest)

	return mux
}
