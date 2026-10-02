package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jonlk/finance-api/internal/model"
)

func PERatio(w http.ResponseWriter, r *http.Request) {
	pps := r.URL.Query().Get(PRICE_PER_SHARE)
	eps := r.URL.Query().Get(EARNINGS_PER_SHARE)

	result := &ApiResult{
		Operation: r.URL.Path,
	}

	pricePerShare, err := strconv.ParseFloat(pps, 64)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("invalid numeric format: %v", PRICE_PER_SHARE))
	}

	earningsPerShare, err := strconv.ParseFloat(eps, 64)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("invalid numeric format: %v", EARNINGS_PER_SHARE))
	}

	if len(result.Errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result)
		return
	}

	res, err := model.PERatio(pricePerShare, earningsPerShare)
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
