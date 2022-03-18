package api

import "strings"

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
