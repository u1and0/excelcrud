package api

import (
	"strconv"
	"time"
)

const (
	LAYOUT     = "2006/01/02"
	SKIPHEADER = 2
)

func (d *Data) ParseExcel(rows [][]string) error {
	for i, datum := range rows {
		if i < SKIPHEADER {
			continue
		}
		r0, err := time.Parse(LAYOUT, datum[0])
		if err != nil {
			return err
		}
		r6, err := time.Parse(LAYOUT, datum[6])
		if err != nil {
			return err
		}
		r4, err := strconv.Atoi(datum[4])
		if err != nil {
			return err
		}
		r5, err := strconv.Atoi(datum[5])
		if err != nil {
			return err
		}
		datum := Datum{r0, datum[1], datum[2], datum[3], r4, r5, r6}
		*d = append(*d, datum)
	}
	return nil
}
