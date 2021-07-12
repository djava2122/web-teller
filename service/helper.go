package service

import "time"

func FormattedTime(date string, formatted string) string {
	t, _ := time.Parse(time.RFC3339, date)
	return t.Format(formatted)
}

func ParseRoleTeller(data string) string {
	if data == "?FDS.MENU.HEAD.TELLER.MENU" {
		return "00"
	}
	return "01"
}
