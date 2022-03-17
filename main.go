package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

const (
	DEBUG    = true
	LAYOUT   = "2006/01/02"
	SKIPROWS = 2
)

type (
	UserDatum struct {
		Entrydate  time.Time `json:"entrydate"`
		UserID     string    `json:"userid"`
		Name       string    `json:"name"`
		Sex        string    `json:"sex"`
		Age        int       `json:"age"`
		TotalMoney int       `json:"totalmoney"`
		BirthDay   time.Time `json:"birthday"`
	}
	UserData []UserDatum
)

func main() {
	f, err := excelize.OpenFile("sample_data.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Entry Excel data as go struct
	rows, err := f.GetRows("ランダムデータ群")
	if err != nil {
		fmt.Println(err)
		return
	}
	data := UserData{}
	for i, row := range rows {
		if i < 2 {
			continue
		}
		r0, _ := time.Parse(LAYOUT, row[0])
		r6, _ := time.Parse(LAYOUT, row[6])
		r4, _ := strconv.Atoi(row[4])
		r5, _ := strconv.Atoi(row[5])
		datum := UserDatum{r0, row[1], row[2], row[3], r4, r5, r6}
		data = append(data, datum)
	}
	if DEBUG {
		fmt.Printf("%v\n", data)
	}
}
