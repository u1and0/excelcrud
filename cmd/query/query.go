package query

import "github.com/gin-gonic/gin"

type (
	Query struct {
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
