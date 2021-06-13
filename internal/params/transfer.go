package params

import (
	compoundInterest "github.com/evgeny-klyopov/compound-interest"
	"strconv"
	"time"
)

func TransferToCompoundInterestParams(params Params) (*compoundInterest.Params, error) {
	var err error

	p := compoundInterest.Params{}

	p.DateStart, err = time.Parse("2006-01-02", params.DateStart)
	if err != nil {
		return nil, err
	}

	p.InvestmentTermInYears, err = strconv.ParseFloat(params.InvestmentTermInYears, 64)
	if err != nil {
		return nil, err
	}

	p.PercentRate, err = strconv.ParseFloat(params.PercentRate, 64)
	if err != nil {
		return nil, err
	}

	p.InitialPayment, err = strconv.ParseFloat(params.InitialPayment, 64)
	if err != nil {
		return nil, err
	}

	p.MonthlyPayment, err = strconv.ParseFloat(params.MonthlyPayment, 64)
	if err != nil {
		return nil, err
	}

	p.AnnualPercentageIncreaseInMonthlyPayment, err = strconv.ParseFloat(
		params.AnnualPercentageIncreaseInMonthlyPayment, 64,
	)
	if err != nil {
		return nil, err
	}

	p.AvgPercentDividend, err = strconv.ParseFloat(params.AvgPercentDividend, 64)
	if err != nil {
		return nil, err
	}

	p.AnnualPercentageIncreaseDividend, err = strconv.ParseFloat(params.AnnualPercentageIncreaseDividend, 64)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
