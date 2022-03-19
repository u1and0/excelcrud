package api

import (
	"math"
	"strings"
	"time"

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
	ge := q.EntryDateGreaterEqual != query.MinDate
	le := q.EntryDateLessEqual != query.MaxDate
	if ge && le {
		data = data.FilterByEntryDate(q.EntryDateGreaterEqual, q.EntryDateLessEqual)
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

// FilterByEntryDate : Filtering data greater equal "g", less equal "l" in Datum.Age
func (d *Data) FilterByEntryDate(g, l time.Time) (data Data) {
	for _, datum := range *d {
		// Beforeの否定で以上
		// After の否定で以下
		if !datum.EntryDate.Before(g) && !datum.EntryDate.After(l) {
			data = append(data, datum)
		}
	}
	return
}
