package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jonlk/finance-api/internal/model"
)

func SimpleInterest(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Query().Get(PRINCIPAL)
	ai := r.URL.Query().Get(ANNUAL_INTEREST_RATE_PERCENTAGE)
	lb := r.URL.Query().Get(YEARS)

	result := &ApiResult{
		Operation: r.URL.Path,
	}

	principal, err := strconv.ParseInt(p, 0, 64)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("invalid numeric format: %v", PRINCIPAL))
	}

	annualInterestRatePercentage, err := strconv.ParseInt(ai, 0, 64)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("invalid numeric format: %v", ANNUAL_INTEREST_RATE_PERCENTAGE))
	}

	lengthBorrowedInYears, err := strconv.ParseInt(lb, 0, 64)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("invalid numeric format: %v", YEARS))
	}

	if len(result.Errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result)
		return
	}

	res, err := model.SimpleInterest(principal, annualInterestRatePercentage, lengthBorrowedInYears)
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
