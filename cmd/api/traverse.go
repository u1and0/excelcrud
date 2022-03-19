package api

import (
	"math"
	"strings"

	query "github.com/u1and0/excelcrud/cmd/query"
)

// FilterBy は検索して、部分一致したdataを返す
// Match は検索して、完全一致したdatumを返す

// MatchID : Match a datum by UserID
func (d *Data) MatchID(id string) Datum {
	for _, datum := range *d {
		if datum.UserID == id {
			return datum
		}
	}
	return *New()
}

// FilterByQuery : Filtering data by Query
func (d *Data) FilterByQuery(q *query.Query) (data Data) {
	// Shallow copy
	data = *d
	if q.UserID != "" {
		data = data.FilterByID(q.UserID)
	}
	if q.AgeGreaterEqual != 0 && q.AgeLessEqual != math.MaxInt {
		data = data.FilterByAge(q.AgeGreaterEqual, q.AgeLessEqual)
	}
	return
}

// FilterByID : Filtering data contains "s" in Datum.ID
func (d *Data) FilterByID(s string) (data Data) {
	for _, datum := range *d {
		if strings.Contains(datum.UserID, s) {
			data = append(data, datum)
		}
	}
	return
}

// FilterByAge : Filtering data greater equal "g", less equal "l" in Datum.Age
func (d *Data) FilterByAge(g, l int) (data Data) {
	for _, datum := range *d {
		if datum.Age >= g && datum.Age <= l {
			data = append(data, datum)
		}
	}
	return
}
