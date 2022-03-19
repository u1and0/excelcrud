package api

import (
	"math"
	"strings"

	query "github.com/u1and0/excelcrud/cmd/query"
)

// Retrive は検索して、部分一致したdataを返す
// Match は検索して、完全一致したdatumを返す

// MatchID : Get a datum by UserID
func (d *Data) MatchID(id string) Datum {
	for _, datum := range *d {
		if datum.UserID == id {
			return datum
		}
	}
	return *New()
}

// RetriveQuery : Get data by Query
func (d *Data) RetriveQuery(q *query.Query) (data Data) {
	// Shallow copy
	data = *d
	if q.UserID != "" {
		data = data.RetriveID(q.UserID)
	}
	if q.AgeGreaterEqual != 0 && q.AgeLessEqual != math.MaxInt {
		data = data.RetriveAge(q.AgeGreaterEqual, q.AgeLessEqual)
	}
	return
}

// RetriveID : filtering Data.ID contains "s"
func (d *Data) RetriveID(s string) (data Data) {
	for _, datum := range *d {
		if strings.Contains(datum.UserID, s) {
			data = append(data, datum)
		}
	}
	return
}

// RetriveAge : filtering Data.Age of greater equal "g", less equal "l"
func (d *Data) RetriveAge(g, l int) (data Data) {
	for _, datum := range *d {
		if datum.Age >= g && datum.Age <= l {
			data = append(data, datum)
		}
	}
	return
}
