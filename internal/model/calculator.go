package model

import (
	"fmt"
	"math"
)

// BasicLiquidityRatio returns how many months of expenses monetaryAssets can cover.
// monetaryAssets and monthlyExpenses are both in cents.
func BasicLiquidityRatio(monetaryAssets int64, monthlyExpenses int64) (float64, error) {
	if monthlyExpenses <= 0 {
		return 0, fmt.Errorf("BasicLiquidityRatio: monthlyExpenses must be positive, got %d", monthlyExpenses)
	}
	return (float64(monetaryAssets) / float64(monthlyExpenses)), nil
}

// BreakEvenPoint returns the revenue (in cents) needed to cover fixedExpenses,
// given a gross profit margin expressed as a whole percentage (e.g. 40 for 40%).
// fixedExpenses is in cents.
func BreakEvenPoint(fixedExpenses int64, grossProfitMarginInPercentage int64) (float64, error) {
	if grossProfitMarginInPercentage <= 0 || grossProfitMarginInPercentage > 100 {
		return 0, fmt.Errorf("BreakEvenPoint: grossProfitMarginInPercentage must be between 1 and 100, got %d", grossProfitMarginInPercentage)
	}
	marginDecimal := float64(grossProfitMarginInPercentage) / 100.0
	return float64(fixedExpenses) / marginDecimal, nil
}

// CashFlow returns income minus expenses, in cents.
func CashFlow(income int64, expenses int64) int64 {
	return income - expenses
}

// RuleOf72 estimates the number of years needed to double an investment,
// given a compound interest rate expressed as a whole percentage (e.g. 8 for 8%).
func RuleOf72(compoundInterestRate float64) (float64, error) {
	if compoundInterestRate <= 0 {
		return 0, fmt.Errorf("RuleOf72: compoundInterestRate must be positive, got %f", compoundInterestRate)
	}
	return 72 / compoundInterestRate, nil
}

// NetIncome returns revenue minus expenses, in cents.
func NetIncome(revenue int64, expenses int64) int64 {
	return revenue - expenses
}

// NetWorth returns assets minus debts, in cents.
func NetWorth(assets int64, debts int64) int64 {
	return assets - debts
}

// PERatio returns the price-to-earnings ratio.
// Unlike the rest of this package, pricePerShare and earningsPerShare are in
// whole dollars, not cents, since share prices and EPS are conventionally quoted that way.
func PERatio(pricePerShare float64, earningsPerShare float64) (float64, error) {
	if earningsPerShare <= 0 {
		return 0, fmt.Errorf("PERatio: earningsPerShare must be positive, got %f", earningsPerShare)
	}
	return pricePerShare / earningsPerShare, nil
}

// SimpleInterest returns principal * rate * time, in cents.
// principal is in cents; annualInterestRatePercentage is a whole percentage (e.g. 5 for 5%).
func SimpleInterest(principal int64, annualInterestRatePercentage int64, years int64) (float64, error) {
	if annualInterestRatePercentage <= 0 {
		return 0, fmt.Errorf("SimpleInterest: annualInterestRatePercentage must be positive, got %d", annualInterestRatePercentage)
	}
	if years <= 0 {
		return 0, fmt.Errorf("SimpleInterest: lengthBorrowedInYears must be positive, got %d", years)
	}
	rateDecimal := float64(annualInterestRatePercentage) / 100.0
	return float64(principal) * rateDecimal * float64(years), nil
}

// VariationOfInvestment returns the percentage change from purchasePrice to
// currentPrice, as a decimal fraction (e.g. 0.15 for +15%).
// currentPrice and purchasePrice are both in cents.
func VariationOfInvestment(currentPrice int64, purchasePrice int64) (float64, error) {
	if purchasePrice <= 0 {
		return 0, fmt.Errorf("VariationOfInvestment: purchasePrice must be positive, got %d", purchasePrice)
	}
	return (float64(currentPrice-purchasePrice) / float64(purchasePrice)), nil
}

// CompoundInterest returns the final balance (in cents) after compounding
// principalCents at rateBasisPoints (e.g. 500 for 5%) periodsPerYear times a
// year, for the given number of years. Intermediate math is done in float64
// and rounded once at the end to avoid per-iteration truncation error.
func CompoundInterest(principal int64, rateBasisPoints int64, periodsPerYear int64, years int64) (int64, error) {
	if principal <= 0 {
		return 0, fmt.Errorf("CompoundInterest: principalCents must be positive, got %d", principal)
	}
	if rateBasisPoints <= 0 {
		return 0, fmt.Errorf("CompoundInterest: rateBasisPoints must be positive, got %d", rateBasisPoints)
	}
	if periodsPerYear <= 0 {
		return 0, fmt.Errorf("CompoundInterest: periodsPerYear must be positive, got %d", periodsPerYear)
	}
	if years <= 0 {
		return 0, fmt.Errorf("CompoundInterest: years must be positive, got %d", years)
	}

	rate := float64(rateBasisPoints) / 10000.0
	periodRate := rate / float64(periodsPerYear)
	totalPeriods := periodsPerYear * years

	amount := float64(principal)
	for i := int64(0); i < totalPeriods; i++ {
		amount *= (1 + periodRate)
	}

	return int64(math.Round(amount)), nil
}

// FormatCents formats a cents value as a dollar.cents display string, e.g. "$1,500.00".
func FormatCents(cents int64) string {
	sign := ""
	if cents < 0 {
		sign = "-"
		cents = -cents
	}
	dollars := cents / 100
	remainder := cents % 100
	return fmt.Sprintf("%s$%d.%02d", sign, dollars, remainder)
}

// FormatPercentage formats a decimal fraction as a percentage display string, e.g. "15.00%".
func FormatPercentage(decimal float64) string {
	return fmt.Sprintf("%.2f%%", decimal*100)
}
