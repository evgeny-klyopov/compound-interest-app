package params

import (
	"github.com/urfave/cli/v2"
	"time"
)

type Params struct {
	DateStart                                string
	InvestmentTermInYears                    string
	PercentRate                              string
	MonthlyPayment                           string
	InitialPayment                           string
	AnnualPercentageIncreaseInMonthlyPayment string
	AvgPercentDividend                       string
	AnnualPercentageIncreaseDividend         string
}

func (p *Params) GetFlags() []cli.Flag {

	return []cli.Flag{
		&cli.StringFlag{
			Name:        "date-start",
			Required:    false,
			Value:       time.Now().Format("2006-01-02"),
			Usage:       "start date of investment",
			HasBeenSet:  true,
			Destination: &p.DateStart,
		},
		&cli.StringFlag{
			Name:        "years",
			Required:    false,
			Value:       "1",
			Usage:       "investment term in years",
			HasBeenSet:  true,
			Destination: &p.InvestmentTermInYears,
		},
		&cli.StringFlag{
			Name:        "rate",
			Required:    false,
			Value:       "10",
			Usage:       "percent rate",
			HasBeenSet:  true,
			Destination: &p.PercentRate,
		},
		&cli.StringFlag{
			Name:        "monthly-payment",
			Required:    true,
			HasBeenSet:  false,
			Usage:       "monthly payment",
			Destination: &p.MonthlyPayment,
		},
		&cli.StringFlag{
			Name:        "initial-payment",
			Required:    false,
			Value:       "0",
			HasBeenSet:  true,
			Usage:       "initial payment",
			Destination: &p.InitialPayment,
		},
		&cli.StringFlag{
			Name:        "percent-increase-monthly-payment",
			Required:    false,
			Value:       "15",
			HasBeenSet:  true,
			Usage:       "annual percentage increase in monthly payment",
			Destination: &p.AnnualPercentageIncreaseInMonthlyPayment,
		},
		&cli.StringFlag{
			Name:        "avg-percent-dividend",
			Required:    false,
			Value:       "4",
			HasBeenSet:  true,
			Usage:       "avg percent dividend",
			Destination: &p.AvgPercentDividend,
		},
		&cli.StringFlag{
			Name:        "percent-increase-dividend",
			Required:    false,
			Value:       "5",
			HasBeenSet:  true,
			Usage:       "annual percentage increase dividend",
			Destination: &p.AnnualPercentageIncreaseDividend,
		},
	}
}

func NewParams() Params {
	return Params{}
}
