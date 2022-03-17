package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

const (
	DEBUG      = false
	LAYOUT     = "2006/01/02"
	SKIPHEADER = 2
	FILENAME   = "ランダムデータ群"
	PORT       = ":8080"
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

func (d *UserData) parseExcel(rows [][]string) error {
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

var (
	data UserData
)

func init() {
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

	// Entry Excel data as Go struct
	rows, err := f.GetRows(FILENAME)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = data.parseExcel(rows)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	r := gin.Default()

	r.GET("/data", func(c *gin.Context) {
		c.JSON(http.StatusOK, data)
	})
	if DEBUG {
		fmt.Printf("%v\n", data)
	}
	r.Run(PORT)
}
