package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jonlk/finance-api/internal/model"
)

func CashFlow(w http.ResponseWriter, r *http.Request) {
	inc := r.URL.Query().Get(INCOME)
	exp := r.URL.Query().Get(EXPENSES)

	result := &ApiResult{
		Operation: r.URL.Path,
	}

	income, err := strconv.ParseInt(inc, 0, 64)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("invalid numeric format: %v", INCOME))
	}

	expenses, err := strconv.ParseInt(exp, 0, 64)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("invalid numeric format: %v", EXPENSES))
	}
	if len(result.Errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result)
		return
	}

	res := model.CashFlow(income, expenses)
	output := strconv.FormatInt(res, 10)
	result.Result = &output
	json.NewEncoder(w).Encode(result)
}
