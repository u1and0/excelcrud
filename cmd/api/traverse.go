package api

import (
	"errors"
	"math"

	query "github.com/u1and0/excelcrud/cmd/query"
)

// TraverseID : get a row by UserID
func (d *Data) TraverseID(id string) (Row, error) {
	for _, row := range *d {
		if row.UserID == id {
			return row, nil
		}
	}
	return Row{}, errors.New("no data")
}

// TraverseQuery : get rows by Query
func (d *Data) TraverseQuery(q *query.Query) (data Data, err error) {
	if q.UserID != "" {
		data = d.MatchID(q.UserID)
	}
	if q.AgeGreaterEqual != 0 && q.AgeLessEqual != math.MaxInt {
		data = data.MatchAge(q.AgeGreaterEqual, q.AgeLessEqual)
	}
	if len(data) == 0 {
		err = errors.New("no match")
	}
	return
}

// TraverseQuery : get a row by Age
func (d *Data) TraverseAge(gt, lt int) (data Data, err error) {
	for _, row := range *d {
		if row.Age >= gt && row.Age <= lt {
			data = append(data, row)
		}
	}
	if len(data) == 0 {
		err = errors.New("no match")
	}
	return
}
