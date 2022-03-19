package api

import (
	"testing"
	"time"

	query "github.com/u1and0/excelcrud/cmd/query"
)

func TestData_MatchID(t *testing.T) {
	data := Data{
		Datum{UserID: "ABCD", Age: 1},
		Datum{UserID: "WXYZ", Age: 2},
		Datum{UserID: "BCDE", Age: 3},
	}
	expected := data[2]
	actual := data.MatchID("BCDE")
	if actual != expected {
		t.Fatalf("got: %v want: %v", actual, expected)
	}
}

func TestData_FilterByQuery(t *testing.T) {
	datum1 := Datum{UserID: "ABCD", Age: 1}
	datum2 := Datum{UserID: "WXYZ", Age: 2}
	datum3 := Datum{UserID: "BCDE", Age: 3}
	data := Data{datum1, datum2, datum3}
	expected := Data{datum1, datum3}
	actual := data.FilterByID("BC")
	for i, datum := range expected {
		if actual[i] != datum {
			t.Fatalf("got: %v want: %v", actual[i].UserID, datum.UserID)
		}
	}
}

func TestData_FilterByID(t *testing.T) {
	data := Data{
		Datum{UserID: "ABCD", Age: 1},
		Datum{UserID: "WXYZ", Age: 2},
		Datum{UserID: "BCDE", Age: 3},
	}
	expected := Data{data[0], data[2]}
	actual := data.FilterByID("BC")
	for i, datum := range expected {
		if actual[i] != datum {
			t.Fatalf("got: %v want: %v", actual[i].UserID, datum.UserID)
		}
	}
}

func TestData_FilterByAge(t *testing.T) {
	data := Data{
		Datum{UserID: "ABCD", Age: 10},
		Datum{UserID: "WXYZ", Age: 20},
		Datum{UserID: "BCDE", Age: 30},
	}
	expected := Data{data[1], data[2]}
	actual := data.FilterByAge(15, 30)
	for i, datum := range expected {
		if actual[i] != datum {
			t.Fatalf("got: %v want: %v", actual[i].UserID, datum.UserID)
		}
	}
}

func TestData_FilterByEntryDate(t *testing.T) {
	data := Data{
		Datum{UserID: "ABCD",
			EntryDate: time.Date(1999, 1, 2, 0, 0, 0, 0, query.Loc)},
		Datum{UserID: "WXYZ",
			EntryDate: time.Date(2010, 1, 1, 0, 0, 0, 0, query.Loc)},
		Datum{UserID: "BCDE",
			EntryDate: time.Date(2999, 1, 2, 0, 0, 0, 0, query.Loc)},
	}
	expected := Data{data[0], data[1]}
	g := time.Date(1999, 1, 2, 0, 0, 0, 0, query.Loc)
	l := time.Date(2010, 1, 1, 0, 0, 0, 0, query.Loc)
	actual := data.FilterByEntryDate(g, l)
	for i, datum := range actual {
		if expected[i] != datum {
			t.Fatalf("got: %#v want: %#v", actual[i].EntryDate, datum.EntryDate)
		}
	}
}
