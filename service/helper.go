package service

import "time"

func FormattedTime(date string, formatted string) string {
	t, _ := time.Parse(time.RFC3339, date)
	return t.Format(formatted)
}

func ParseRoleTeller(data string) string {
	if data == "?FDS.MENU.MAIN.TELLER" || data == "?KALSEL.MENU.MAIN.TELLER" {
		return "T"
	} else if data == "?FDS.MENU.HEAD.TELLER.MENU" || data == "?KALSEL.MENU.HEAD.TELLER.MENU" {
		return "H"
	} else if data == "?FDS.MENU.MAIN.BO.CABANG" || data == "?KALSEL.MENU.MAIN.BO.CABANG" {
		return "B"
	} else {
		return "I"
	}
	return "I"
}

func ParseRC(rc string) string {
	if rc == "07" {
		return "Tanggal Pembayaran Virtual Account telah Berakhir"
	}
	return "Undefined Error"
}