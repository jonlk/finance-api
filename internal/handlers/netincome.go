package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jonlk/finance-api/internal/model"
)

func NetIncome(w http.ResponseWriter, r *http.Request) {
	rev := r.URL.Query().Get(REVENUE)
	exp := r.URL.Query().Get(EXPENSES)

	result := &ApiResult{
		Operation: r.URL.Path,
	}

	revenue, err := strconv.ParseInt(rev, 0, 64)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("invalid numeric format: %v", REVENUE))
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

	res := model.NetIncome(revenue, expenses)
	output := strconv.FormatFloat(float64(res/100), 'f', 2, 64)
	result.Result = &output
	json.NewEncoder(w).Encode(result)
}
