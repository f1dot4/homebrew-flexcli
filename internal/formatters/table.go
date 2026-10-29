package formatters

import (
	"fmt"
	"strings"
)

// Table represents a simple text table.
type Table struct {
	Headers []string
	Rows    [][]string
}

// NewTable creates a new table with the given headers.
func NewTable(headers ...string) *Table {
	return &Table{
		Headers: headers,
		Rows:    [][]string{},
	}
}

// AddRow adds a row to the table.
func (t *Table) AddRow(cells ...string) {
	t.Rows = append(t.Rows, cells)
}

// Render renders the table to stdout.
func (t *Table) Render() {
	if len(t.Headers) == 0 {
		return
	}

	// Calculate column widths
	widths := make([]int, len(t.Headers))
	for i, header := range t.Headers {
		widths[i] = len(header)
	}
	for _, row := range t.Rows {
		for i, cell := range row {
			if i < len(widths) && len(cell) > widths[i] {
				widths[i] = len(cell)
			}
		}
	}

	// Render header
	headerParts := make([]string, len(t.Headers))
	for i, header := range t.Headers {
		headerParts[i] = padRight(header, widths[i])
	}
	fmt.Println(strings.Join(headerParts, "  "))

	// Render separator
	separatorParts := make([]string, len(t.Headers))
	for i := range t.Headers {
		separatorParts[i] = strings.Repeat("-", widths[i])
	}
	fmt.Println(strings.Join(separatorParts, "  "))

	// Render rows
	for _, row := range t.Rows {
		rowParts := make([]string, len(t.Headers))
		for i, cell := range row {
			if i < len(widths) {
				rowParts[i] = padRight(cell, widths[i])
			}
		}
		fmt.Println(strings.Join(rowParts, "  "))
	}
}

// padRight pads a string to the right with spaces.
func padRight(s string, width int) string {
	if len(s) >= width {
		return s
	}
	return s + strings.Repeat(" ", width-len(s))
}
