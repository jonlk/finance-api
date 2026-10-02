package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jonlk/finance-api/internal/model"
)

func BasicLiquidityRatio(w http.ResponseWriter, r *http.Request) {
	mar := r.URL.Query().Get(MONETARY_ASSETS)
	mer := r.URL.Query().Get(MONTHLY_EXPENSES)

	result := &ApiResult{
		Operation: r.URL.Path,
	}

	monetaryAssets, err := strconv.ParseInt(mar, 0, 64)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("invalid numeric format: %v", MONETARY_ASSETS))
	}

	monthlyExpenses, err := strconv.ParseInt(mer, 0, 64)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("invalid numeric format: %v", MONTHLY_EXPENSES))
	}

	if len(result.Errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result)
		return
	}

	res, err := model.BasicLiquidityRatio(monetaryAssets, monthlyExpenses)
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
