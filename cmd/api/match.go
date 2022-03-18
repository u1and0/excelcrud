package api

import "strings"

func (d *Data) MatchID(s string) (data Data) {
	for _, row := range *d {
		if strings.Contains(row.UserID, s) {
			data = append(data, row)
		}
	}
	return
}

func (d *Data) MatchAge(g, l int) (data Data) {
	for _, row := range *d {
		if row.Age >= g && row.Age <= l {
			data = append(data, row)
		}
	}
	return
}
