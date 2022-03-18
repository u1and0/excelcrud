package main

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

const (
	DEBUG      = true
	LAYOUT     = "2006/01/02"
	SKIPHEADER = 2
	FILENAME   = "ランダムデータ群"
	PORT       = ":8080"
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
	Query    struct {
		// EntryDateGreaterEqual  [8]int `form:"edge"`
		// EntryDateLessEqual  [8]int `form:"edle"`
		UserID string `form:"userid"`
		// Name       string    `form:"name"`
		// Sex        string    `form:"sex"`
		AgeGreaterEqual int `form:"agege"`
		AgeLessEqual    int `form:"agele"`
		// TotalMoney int       `form:"totalmoney"`
		// BirthDay   time.Time `form:"birthday"`
	}
)

// New : Query constructor
// Default value Logging: ture <= always log search query
//									if ommited URL request &logging
// Default value Limit: -1 <= dump all result
//									if ommited URL request &limit
func New(c *gin.Context) (*Query, error) {
	// query := Query{Logging: true, Limit: -1}
	var q Query
	err := c.ShouldBind(&q)
	return &q, err
}

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
		q, err := New(c)
		if DEBUG {
			fmt.Printf("%#v\n", q)
		}
		if err != nil {
			fmt.Println(err)
		}
		if *q == (Query{}) {
			// if not query parameter
			// Return all data
			c.JSON(http.StatusOK, data)
			return
		}
		// if query parameter
		// Traverse alldata
		traversedata, err := data.TraverseQuery(q)
		if err != nil {
			c.JSON(404, UserData{})
			fmt.Println(err)
			return
		}
		c.JSON(http.StatusOK, traversedata)
	})
	// One datum from UserID
	// curl localhost:8080/OD77412
	r.GET("/data/:userid", func(c *gin.Context) {
		id := c.Param("userid")
		datum, err := data.TraverseID(id)
		if err != nil {
			c.JSON(404, UserDatum{})
			return
		}
		c.JSON(http.StatusOK, datum)
	})
	// curl localhost:8080/age?gt=10&lt=30
	r.GET("/data/age", func(c *gin.Context) {
		gt := intQuery(c, "gt") // age?gt=10 => gt==10
		lt := intQuery(c, "lt") // age?lt=100 => lt==100
		if lt == 0 {
			lt = math.MaxInt64
		}
		ageData, err := data.TraverseAge(gt, lt)
		if DEBUG {
			fmt.Println("gt, lt", gt, lt)
			fmt.Println("ageData", ageData)
			fmt.Println(err)
		}
		if err != nil {
			c.JSON(404, UserData{})
			return
		}
		c.JSON(http.StatusOK, ageData)
	})
	r.Run(PORT)
}

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
func (d *UserData) TraverseQuery(q *Query) (data UserData, err error) {
	data = d.MatchID(q.UserID)
	data = data.MatchAge(q.AgeGreaterEqual, q.AgeLessEqual)
	if len(data) == 0 {
		err = errors.New("no data")
	}
	return
}

func (d *UserData) MatchID(s string) (data UserData) {
	for _, datum := range *d {
		if strings.Contains(datum.UserID, s) {
			data = append(data, datum)
		}
	}
	return
}

func (d *UserData) MatchAge(g, l int) (data UserData) {
	for _, datum := range *d {
		if datum.Age >= g && datum.Age <= l {
			data = append(data, datum)
		}
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
		err = errors.New("no data")
	}
	return
}

// intQuery parse query as int
func intQuery(c *gin.Context, q string) int {
	s, ok := c.GetQuery(q)
	if !ok {
		return 0
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}
