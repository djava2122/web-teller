package service

import "time"

func FormattedTime(date string, formatted string) string {
	t, _ := time.Parse(time.RFC3339, date)
	return t.Format(formatted)
}

func ParseRoleTeller(data string) string {
	if data == "?FDS.MENU.MAIN.TELLER" || data == "?KALSEL.MENU.MAIN.TELLER" || data == "KALSEL.SYA.MENU.TEL" || data == "KALSEL.SYA.TELSKN" {
		return "T"
	} else if data == "?FDS.MENU.HEAD.TELLER" || data == "?KALSEL.MENU.MAIN.HT" || data == "KALSEL.SYA.HTEL" || data == "KALSEL.SYA.HTELSKN" || data == "KALSEL.MENU.MAIN.KBCS" {
		return "H"
	} else if data == "?FDS.MENU.MAIN.BO.CABANG" || data == "?KALSEL.MENU.MAIN.BOFLA" || data == "?KALSEL.MENU.MAIN.BOFKL" ||
		data == "KALSEL.MENU.MAIN.KBBOFKL" || data == "?KALSEL.APV.USER.MENU.KBBOF" || data == "?KALSEL.SYA.BO" || data == "KALSEL.SYA.BOSKN" {
		return "B"
	} else {
		return "I"
	}
}

func ParseRC(rc string) string {
	if rc == "07" {
		return "Tanggal Pembayaran Virtual Account telah Berakhir"
	}
	return "Undefined Error"
}
