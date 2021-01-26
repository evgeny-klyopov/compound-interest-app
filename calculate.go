package main

import (
	"bufio"
	"fmt"
	"github.com/evgeny-klyopov/bashColor"
	"github.com/evgeny-klyopov/compound-interest"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"
	"time"
)

type field struct {
	fieldDisplay string
	field        string
	defaultValue string
	typeValue    string
}

type calculate struct {
	params       compoundInterest.Params
	in           []field
	fields       map[string]string
	color        bashColor.Colorer
	tableSetting *tableSetting
}

func newCalculate() *calculate {
	c := &calculate{
		in: []field{
			{fieldDisplay: "date start", field: "DateStart", defaultValue: time.Now().Format("2006-01-02")},
			{fieldDisplay: "investment term in years", field: "InvestmentTermInYears", defaultValue: "1"},
			{fieldDisplay: "percent rate", field: "PercentRate", defaultValue: "10"},
			{fieldDisplay: "monthly payment", field: "MonthlyPayment", defaultValue: "30000"},
			{fieldDisplay: "initial payment", field: "InitialPayment", defaultValue: "30000"},
			{fieldDisplay: "annual percentage increase in monthly payment", field: "AnnualPercentageIncreaseInMonthlyPayment", defaultValue: "15"},
			{fieldDisplay: "avg percent dividend", field: "AvgPercentDividend", defaultValue: "4"},
			{fieldDisplay: "annual percentage increase dividend", field: "AnnualPercentageIncreaseDividend", defaultValue: "5"},
		},
		color:        bashColor.NewColor(),
		tableSetting: NewTableSetting(),
	}
	c.fields = make(map[string]string)

	return c
}

func (c *calculate) run() {
	var year *string

	compoundInterestCalc := compoundInterest.New(c.params).Calculate()
	compoundInterestCalcLength := len(compoundInterestCalc)

	for monthNumber, row := range compoundInterestCalc {
		y := row.Date.Format("2006")
		number := monthNumber + 1

		if year == nil {
			c.printHeader(y)
		} else if *year != y {
			c.tableSetting.printTable()
			c.printSeparate().printHeader(y)
		}

		c.tableSetting.addRow(number, row).updateTotal(row)
		year = &y

		if compoundInterestCalcLength == number {
			c.tableSetting.printTable()
		}
	}

}

func (c *calculate) fillParams() error {
	var err error

	c.params = compoundInterest.Params{}

	c.params.DateStart, err = time.Parse("2006-01-02", c.fields["DateStart"])

	if err != nil {
		return err
	}

	c.params.InvestmentTermInYears, err = strconv.ParseFloat(c.fields["InvestmentTermInYears"], 64)
	if err != nil {
		return err
	}
	c.params.PercentRate, err = strconv.ParseFloat(c.fields["PercentRate"], 64)
	if err != nil {
		return err
	}
	c.params.InitialPayment, err = strconv.ParseFloat(c.fields["InitialPayment"], 64)
	if err != nil {
		return err
	}
	c.params.MonthlyPayment, err = strconv.ParseFloat(c.fields["MonthlyPayment"], 64)
	if err != nil {
		return err
	}
	c.params.AnnualPercentageIncreaseInMonthlyPayment, err = strconv.ParseFloat(c.fields["AnnualPercentageIncreaseInMonthlyPayment"], 64)
	if err != nil {
		return err
	}
	c.params.AvgPercentDividend, err = strconv.ParseFloat(c.fields["AvgPercentDividend"], 64)
	if err != nil {
		return err
	}
	c.params.AnnualPercentageIncreaseDividend, err = strconv.ParseFloat(c.fields["AnnualPercentageIncreaseDividend"], 64)
	if err != nil {
		return err
	}

	return err
}
func (c *calculate) dialog() error {
	for _, f := range c.in {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter " + f.fieldDisplay + ": ")
		text, _ := reader.ReadString('\n')

		if text == "\n" {
			text = f.defaultValue
		}

		c.fields[f.field] = text
	}

	return c.fillParams()
}

func (c *calculate) setInArguments(args cli.Args) error {
	for key, f := range c.in {
		text := args.Get(key)

		if text == "" {
			text = f.defaultValue
		}

		c.fields[f.field] = text
	}

	return c.fillParams()
}

func (c *calculate) printHeader(header string) *calculate {
	c.color.Print(c.color.Magenta, header)
	c.color.Print(c.color.Info, "")
	return c
}

func (c *calculate) printSeparate() *calculate {
	c.color.Print(c.color.Info, "")
	c.color.Print(c.color.Info, "")
	c.color.Print(c.color.Info, "")

	return c
}
