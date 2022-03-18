package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	api "github.com/u1and0/excelcrud/cmd/api"
	query "github.com/u1and0/excelcrud/cmd/query"
	"github.com/xuri/excelize/v2"
)

const (
	DEBUG    = true
	FILENAME = "ランダムデータ群"
	PORT     = ":8080"
)

var (
	data api.Data
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
	err = data.ParseExcel(rows)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	r := gin.Default()
	r.GET("/data", func(c *gin.Context) {
		q, err := query.New(c)
		if DEBUG {
			fmt.Printf("%#v\n", q)
		}
		if err != nil {
			fmt.Println(err)
		}
		if *q == (query.Query{}) {
			// if not query parameter
			// Return all data
			c.JSON(http.StatusOK, data)
			return
		}
		// if query parameter
		// Traverse alldata
		traversedata, err := data.TraverseQuery(q)
		if err != nil {
			c.JSON(404, api.Data{})
			fmt.Println(err)
			return
		}
		c.JSON(http.StatusOK, traversedata)
	})
	// One row from UserID
	// curl localhost:8080/OD77412
	r.GET("/data/:userid", func(c *gin.Context) {
		id := c.Param("userid")
		row, err := data.TraverseID(id)
		if err != nil {
			c.JSON(404, api.Row{})
			return
		}
		c.JSON(http.StatusOK, row)
	})

	r.Run(PORT)
}
