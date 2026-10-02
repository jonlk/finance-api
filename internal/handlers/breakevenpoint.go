package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jonlk/finance-api/internal/model"
)

func BreakEvenPoint(w http.ResponseWriter, r *http.Request) {
	fex := r.URL.Query().Get(FIXED_EXPENSES)
	gpm := r.URL.Query().Get(GROSS_PROFIT_MARGIN_IN_PERCENTAGE)

	result := &ApiResult{
		Operation: r.URL.Path,
	}

	fixedExpenses, err := strconv.ParseInt(fex, 0, 64)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("invalid numeric format: %v", FIXED_EXPENSES))
	}

	grossProfitMarginInPercentage, err := strconv.ParseInt(gpm, 0, 64)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("invalid numeric format: %v", GROSS_PROFIT_MARGIN_IN_PERCENTAGE))
	}

	if len(result.Errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result)
		return
	}

	res, err := model.BreakEvenPoint(fixedExpenses, grossProfitMarginInPercentage)
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
