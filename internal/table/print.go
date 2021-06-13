package table

import (
	compoundInterest "github.com/evgeny-klyopov/compound-interest"
	"github.com/leekchan/accounting"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

type tableSetting struct {
	header      []string
	headerColor []tablewriter.Colors
	columnColor []tablewriter.Colors
	footer      []string
	footerColor []tablewriter.Colors
	total       total
	data        [][]string
}

type total struct {
	MonthlyPayment  float64
	Amount          float64
	MonthlyDividend float64
}

type Printer interface {
	Print()
	AddRow(number int, row compoundInterest.Prediction) Printer
	UpdateTotal(row compoundInterest.Prediction)
}

func NewPrinter() Printer {
	return &tableSetting{
		header: []string{
			"Month number",
			"date",
			"Avg percent dividend",
			"Monthly payment",
			"Amount",
			"Dividend",
		},
		headerColor: []tablewriter.Colors{
			{tablewriter.FgHiWhiteColor, tablewriter.Bold},
			{tablewriter.FgHiWhiteColor, tablewriter.Bold},
			{tablewriter.FgHiWhiteColor, tablewriter.Bold},
			{tablewriter.FgHiCyanColor, tablewriter.Bold},
			{tablewriter.FgHiRedColor, tablewriter.Bold},
			{tablewriter.FgHiGreenColor, tablewriter.Bold},
		},
		columnColor: []tablewriter.Colors{
			{},
			{},
			{},
			{tablewriter.Normal, tablewriter.FgHiCyanColor},
			{tablewriter.Normal, tablewriter.FgHiRedColor},
			{tablewriter.Normal, tablewriter.FgHiGreenColor},
		},
		footerColor: []tablewriter.Colors{
			{},
			{},
			{tablewriter.Bold, tablewriter.FgHiYellowColor},
			{tablewriter.Bold, tablewriter.FgHiCyanColor},
			{tablewriter.Bold, tablewriter.FgHiRedColor},
			{tablewriter.Bold, tablewriter.FgHiGreenColor},
		},
		total: total{},
	}
}

func (t *tableSetting) Print() {
	table := tablewriter.NewWriter(os.Stdout)
	t.fillFooter(t.total)

	table.SetHeader(t.header)
	table.SetHeaderColor(t.headerColor...)
	table.SetColumnColor(t.columnColor...)
	table.SetFooter(t.footer)
	table.SetFooterColor(t.footerColor...)
	table.SetBorder(false)
	table.AppendBulk(t.data)
	table.Render()

	t.cleanData().cleanTotal()
}

func (t *tableSetting) UpdateTotal(row compoundInterest.Prediction) {
	t.total.MonthlyPayment = t.total.MonthlyPayment + row.MonthlyPayment
	t.total.Amount = row.Amount
	t.total.MonthlyDividend = t.total.MonthlyDividend + row.MonthlyDividend
}

func (t *tableSetting) fillFooter(footer total) {
	t.footer = []string{
		"",
		"",
		"Total",
		accounting.FormatNumberFloat64(footer.MonthlyPayment, 2, " ", "."),
		accounting.FormatNumberFloat64(footer.Amount, 2, " ", "."),
		accounting.FormatNumberFloat64(footer.MonthlyDividend, 2, " ", "."),
	}
}

func (t *tableSetting) cleanTotal() *tableSetting {
	t.total = total{}

	return t
}
func (t *tableSetting) cleanData() *tableSetting {
	t.data = make([][]string, 12, 12)

	return t
}

func (t *tableSetting) AddRow(number int, row compoundInterest.Prediction) Printer {
	t.data = append(t.data, []string{
		strconv.Itoa(number),
		row.Date.Format("2006 Jan"),
		accounting.FormatNumberFloat64(row.AvgPercentDividend, 2, " ", "."),
		accounting.FormatNumberFloat64(row.MonthlyPayment, 2, " ", "."),
		accounting.FormatNumberFloat64(row.Amount, 2, " ", "."),
		accounting.FormatNumberFloat64(row.MonthlyDividend, 2, " ", "."),
	})

	return t
}
