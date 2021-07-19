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

func ParseRC(rc string) string {
	if rc == "07" {
		return "Tanggal Pembayaran Virtual Account telah Berakhir"
	}
	return "Undefined Error"
}