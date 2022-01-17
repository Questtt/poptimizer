package usd

import (
	"fmt"
	"github.com/WLM1ke/gomoex"
	"github.com/WLM1ke/poptimizer/data/internal/domain"
	"github.com/WLM1ke/poptimizer/data/internal/rules/template"
)

func validator(table domain.Table[gomoex.Candle], rows []gomoex.Candle) error {
	prev := rows[0].Begin
	for _, row := range rows[1:] {
		if prev.Before(row.Begin) {
			prev = row.Begin
			continue
		}

		return fmt.Errorf("%w: not increasing dates %+v", template.ErrNewRowsValidation, prev)
	}

	n := len(table.Rows)
	if n == 0 {
		return nil
	}

	lastRow := table.Rows[n-1]
	firstRow := rows[0]
	if lastRow != firstRow {
		return fmt.Errorf("%w: old rows %+v not match new %+v", template.ErrNewRowsValidation, lastRow, firstRow)
	}

	return nil
}