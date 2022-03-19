package query

import (
	"math"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	Loc         = time.UTC
	MinDate     = time.Date(1, 1, 1, 0, 0, 0, 0, Loc)
	MaxDate     = time.Date(9999, 12, 31, 0, 0, 0, 0, Loc)
	defaultdate = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
)

type (
	Query struct {
		EntryDateGreaterEqual time.Time `form:"edge" time_format:"2006-01-02" time_utc:"1"`
		EntryDateLessEqual    time.Time `form:"edle" time_format:"2006-01-02" time_utc:"1"`
		UserID                string    `form:"userid"`
		// Name       string    `form:"name"`
		// Sex        string    `form:"sex"`
		AgeGreaterEqual int `form:"agege"`
		AgeLessEqual    int `form:"agele"`
		// TotalMoney int       `form:"totalmoney"`
		// BirthDay   time.Time `form:"birthday"`
	}
)

// New : Query constructor
func New(c *gin.Context) (*Query, error) {
	// query := Query{Logging: true, Limit: -1}
	q := Query{
		EntryDateGreaterEqual: MinDate,
		EntryDateLessEqual:    MaxDate,
	}
	err := c.ShouldBind(&q)
	// Default values
	if q.AgeLessEqual < 1 {
		q.AgeLessEqual = math.MaxInt
	}
	return &q, err
}
