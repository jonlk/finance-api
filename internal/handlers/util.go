package handlers

type ApiResult struct {
	Operation string   `json:"operation"`
	Result    *string  `json:"result"`
	Errors    []string `json:"errors"`
}

const (
	ANNUAL_INTEREST_RATE_PERCENTAGE   = "annualInterestRatePercentage"
	ASSETS                            = "assets"
	COMPOUND_INTEREST_RATE            = "compoundInterestRate"
	CURRENT_PRICE                     = "currentPrice"
	DEBTS                             = "debts"
	EARNINGS_PER_SHARE                = "earningsPerShare"
	EXPENSES                          = "expenses"
	FIXED_EXPENSES                    = "fixedExpenses"
	GROSS_PROFIT_MARGIN_IN_PERCENTAGE = "grossProfitMarginInPercentage"
	INCOME                            = "income"
	MONETARY_ASSETS                   = "monetaryAssets"
	MONTHLY_EXPENSES                  = "monthlyExpenses"
	PERIODS_PER_YEAR                  = "periodsPerYear"
	PRICE_PER_SHARE                   = "pricePerShare"
	PRINCIPAL                         = "principal"
	PURCHASE_PRICE                    = "purchasePrice"
	RATE_BASIS_POINTS                 = "rateBasisPoints"
	REVENUE                           = "income"
	YEARS                             = "years"
)
