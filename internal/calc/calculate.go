package calc

import (
	"github.com/evgeny-klyopov/compound-interest-app/internal/table"

	//"bufio"
	//"fmt"
	"github.com/evgeny-klyopov/bashColor"
	compoundInterest "github.com/evgeny-klyopov/compound-interest"
	"github.com/evgeny-klyopov/compound-interest-app/internal/params"
)

type calculate struct {
	params params.Params
	color  bashColor.Colorer
	table  table.Printer
}

type Calculator interface {
	Run() error
}

func NewCalculate(params params.Params, c bashColor.Colorer) Calculator {
	return &calculate{
		color:  c,
		params: params,
		table:  table.NewPrinter(),
	}
}

func (c *calculate) Run() error {
	var year *string

	p, err := params.TransferToCompoundInterestParams(c.params)
	if err != nil {
		return err
	}
	compoundInterestCalc := compoundInterest.
		New(*p).
		Calculate()
	compoundInterestCalcLength := len(compoundInterestCalc)

	for monthNumber, row := range compoundInterestCalc {
		y := row.Date.Format("2006")
		number := monthNumber + 1

		if year == nil {
			c.printHeader(y)
		} else if *year != y {
			c.table.Print()
			c.printSeparate().printHeader(y)
		}

		c.table.AddRow(number, row).UpdateTotal(row)
		year = &y

		if compoundInterestCalcLength == number {
			c.table.Print()
		}
	}

	return err
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
