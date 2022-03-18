package api

import (
	"errors"
	"math"
	"time"

	query "github.com/u1and0/excelcrud/cmd/query"
)

type (
	UserDatum struct {
		EntryDate  time.Time `json:"entrydate"`
		UserID     string    `json:"userid"`
		Name       string    `json:"name"`
		Sex        string    `json:"sex"`
		Age        int       `json:"age"`
		TotalMoney int       `json:"totalmoney"`
		BirthDay   time.Time `json:"birthday"`
	}
	UserData []UserDatum
)

// TraverseID : get a row by UserID
func (d *UserData) TraverseID(id string) (UserDatum, error) {
	for _, datum := range *d {
		if datum.UserID == id {
			return datum, nil
		}
	}
	return UserDatum{}, errors.New("no data")
}

// TraverseQuery : get rows by Query
func (d *UserData) TraverseQuery(q *query.Query) (data UserData, err error) {
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
func (d *UserData) TraverseAge(gt, lt int) (data UserData, err error) {
	for _, datum := range *d {
		if datum.Age >= gt && datum.Age <= lt {
			data = append(data, datum)
		}
	}
	if len(data) == 0 {
		err = errors.New("no match")
	}
	return
}
