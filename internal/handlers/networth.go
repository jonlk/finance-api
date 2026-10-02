package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jonlk/finance-api/internal/model"
)

func NetWorth(w http.ResponseWriter, r *http.Request) {
	a := r.URL.Query().Get(ASSETS)
	d := r.URL.Query().Get(DEBTS)

	result := &ApiResult{
		Operation: r.URL.Path,
	}

	assets, err := strconv.ParseInt(a, 0, 64)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("invalid numeric format: %v", ASSETS))
	}

	debts, err := strconv.ParseInt(d, 0, 64)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("invalid numeric format: %v", DEBTS))
	}

	if len(result.Errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result)
		return
	}

	res := model.NetWorth(assets, debts)
	output := strconv.FormatFloat(float64(res/100), 'f', 2, 64)
	result.Result = &output
	json.NewEncoder(w).Encode(result)
}
