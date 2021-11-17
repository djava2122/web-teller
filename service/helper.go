package service

import "time"

func FormattedTime(date string, formatted string) string {
	t, _ := time.Parse(time.RFC3339, date)
	return t.Format(formatted)
}

func ParseRoleTeller(data string) string {
	if data == "?FDS.MENU.MAIN.TELLER" || data == "?KALSEL.MENU.MAIN.TELLER" || data == "?KALSEL.SYA.MENU.TEL" || data == "?KALSEL.SYA.TELSKN" || data == "?KALSEL.SYA.MENU.TELSKN" ||
		data == "?KALSEL.SYA.MENU.TELKFS" {
		return "T"
	} else if data == "?FDS.MENU.HEAD.TELLER" || data == "?KALSEL.MENU.MAIN.HT" || data == "?KALSEL.SYA.HTEL" || data == "?KALSEL.SYA.HTELSKN" || data == "?KALSEL.MENU.MAIN.KBCS" {
		return "H"
	} else if data == "?FDS.MENU.MAIN.BO.CABANG" || data == "?KALSEL.MENU.MAIN.BOFLA" || data == "?KALSEL.MENU.MAIN.BOFKL" ||
		data == "?KALSEL.MENU.MAIN.KBBOFKL" || data == "?KALSEL.APV.USER.MENU.KBBOF" || data == "?KALSEL.SYA.BO" || data == "?KALSEL.SYA.BOSKN" ||
		data == "?KALSEL.SYA.MENU.BOSKN" || data == "?KALSEL.SYA.MENU.KTKP" || data == "?KALSEL.MENU.BO.KP.PEJABAT" || data == "?KALSEL.MENU.MAIN.BO.KO.STAF" ||
		data == "?KALSEL.SYA.MENU.TKP" {
		return "B"
	} else if data == "?KALSEL.SYA.CS" {
		return "CS"
	} else if data == " "{
		return "I"
	} else {
		return "A"
	}
}

func ParseRC(rc string) string {
	if rc == "07" {
		return "Tanggal Pembayaran Virtual Account telah Berakhir"
	}
	return "Undefined Error"
}
