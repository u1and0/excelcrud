package api

import "testing"

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

func TestData_RetriveQuery(t *testing.T) {
	datum1 := Datum{UserID: "ABCD", Age: 1}
	datum2 := Datum{UserID: "WXYZ", Age: 2}
	datum3 := Datum{UserID: "BCDE", Age: 3}
	data := Data{datum1, datum2, datum3}
	expected := Data{datum1, datum3}
	actual := data.RetriveID("BC")
	for i, datum := range expected {
		if actual[i] != datum {
			t.Fatalf("got: %v want: %v", actual[i].UserID, datum.UserID)
		}
	}
}

func TestData_RetriveID(t *testing.T) {
	data := Data{
		Datum{UserID: "ABCD", Age: 1},
		Datum{UserID: "WXYZ", Age: 2},
		Datum{UserID: "BCDE", Age: 3},
	}
	expected := Data{data[0], data[2]}
	actual := data.RetriveID("BC")
	for i, datum := range expected {
		if actual[i] != datum {
			t.Fatalf("got: %v want: %v", actual[i].UserID, datum.UserID)
		}
	}
}

func TestData_RetriveAge(t *testing.T) {
	data := Data{
		Datum{UserID: "ABCD", Age: 10},
		Datum{UserID: "WXYZ", Age: 20},
		Datum{UserID: "BCDE", Age: 30},
	}
	expected := Data{data[1], data[2]}
	actual := data.RetriveAge(15, 30)
	for i, datum := range expected {
		if actual[i] != datum {
			t.Fatalf("got: %v want: %v", actual[i].UserID, datum.UserID)
		}
	}
}
