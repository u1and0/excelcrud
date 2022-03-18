package api

import (
	"strconv"
	"time"
)

const (
	LAYOUT     = "2006/01/02"
	SKIPHEADER = 2
)

func (d *UserData) ParseExcel(rows [][]string) error {
	for i, row := range rows {
		if i < SKIPHEADER {
			continue
		}
		r0, err := time.Parse(LAYOUT, row[0])
		if err != nil {
			return err
		}
		r6, err := time.Parse(LAYOUT, row[6])
		if err != nil {
			return err
		}
		r4, err := strconv.Atoi(row[4])
		if err != nil {
			return err
		}
		r5, err := strconv.Atoi(row[5])
		if err != nil {
			return err
		}
		datum := UserDatum{r0, row[1], row[2], row[3], r4, r5, r6}
		*d = append(*d, datum)
	}
	return nil
}
