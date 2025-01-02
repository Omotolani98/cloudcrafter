package utils

import (
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

// RenderTable displays a table with colored rows
func RenderTable(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetBorder(true)
	table.SetRowLine(true)

	// Add data to the table
	for _, row := range data {
		table.Append(row)
	}

	// Render the table
	table.Render()
}

// FormatStatus applies colors to status values
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
