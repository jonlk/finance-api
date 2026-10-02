package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jonlk/finance-api/internal/model"
)

func CompoundInterest(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Query().Get(PRINCIPAL)
	rbp := r.URL.Query().Get(RATE_BASIS_POINTS)
	ppy := r.URL.Query().Get(PERIODS_PER_YEAR)
	y := r.URL.Query().Get(YEARS)

	result := &ApiResult{
		Operation: r.URL.Path,
	}

	principal, err := strconv.ParseInt(p, 0, 64)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("invalid numeric format: %v", PRINCIPAL))
	}

	rateBasisPoints, err := strconv.ParseInt(rbp, 0, 64)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("invalid numeric format: %v", RATE_BASIS_POINTS))
	}

	periodsPerYear, err := strconv.ParseInt(ppy, 0, 64)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("invalid numeric format: %v", PERIODS_PER_YEAR))
	}

	years, err := strconv.ParseInt(y, 0, 64)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("invalid numeric format: %v", YEARS))
	}

	if len(result.Errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result)
		return
	}

	res, err := model.CompoundInterest(principal, rateBasisPoints, periodsPerYear, years)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		result.Errors = append(result.Errors, err.Error())
		json.NewEncoder(w).Encode(result)
		return
	}

	output := strconv.FormatInt(res, 10)
	result.Result = &output
	json.NewEncoder(w).Encode(result)
}
