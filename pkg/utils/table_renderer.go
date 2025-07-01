package utils

import (
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

func RenderTable(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetBorder(true)
	table.SetRowLine(true)

	
	for _, row := range data {
		table.Append(row)
	}

	
	table.Render()
}

func FormatStatus(status string) string {
	switch status {
	case "running":
		return color.GreenString(status)
	case "stopped":
		return color.RedString(status)
	default:
		return color.YellowString(status)
	}
}
