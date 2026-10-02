package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jonlk/finance-api/internal/model"
)

func RuleOf72(w http.ResponseWriter, r *http.Request) {
	ci := r.URL.Query().Get(COMPOUND_INTEREST_RATE)

	result := &ApiResult{
		Operation: r.URL.Path,
	}

	compoundInterestRate, err := strconv.ParseFloat(ci, 64)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("invalid numeric format: %v", COMPOUND_INTEREST_RATE))
	}

	if len(result.Errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result)
		return
	}

	res, err := model.RuleOf72(compoundInterestRate)
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
