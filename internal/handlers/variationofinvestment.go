package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jonlk/finance-api/internal/model"
)

func VariationOfInvestment(w http.ResponseWriter, r *http.Request) {
	cp := r.URL.Query().Get(CURRENT_PRICE)
	pp := r.URL.Query().Get(PURCHASE_PRICE)

	result := &ApiResult{
		Operation: r.URL.Path,
	}

	currentPrice, err := strconv.ParseInt(cp, 0, 64)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("invalid numeric format: %v", CURRENT_PRICE))
	}

	purchasePrice, err := strconv.ParseInt(pp, 0, 64)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("invalid numeric format: %v", PURCHASE_PRICE))
	}

	if len(result.Errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result)
		return
	}

	res, err := model.VariationOfInvestment(currentPrice, purchasePrice)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		result.Errors = append(result.Errors, err.Error())
		json.NewEncoder(w).Encode(result)
		return
	}

	output := strconv.FormatFloat(res, 'f', 4, 64)
	result.Result = &output
	json.NewEncoder(w).Encode(result)
}
