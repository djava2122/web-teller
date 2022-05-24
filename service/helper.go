package service

import (
	"strings"
	"time"
)

func FormattedTime(date string, formatted string) string {
	t, _ := time.Parse(time.RFC3339, date)
	return t.Format(formatted)
}

func ParseRoleTeller(data string) string {
	if data == "?FDS.MENU.MAIN.TELLER" || data == "?KALSEL.MENU.MAIN.TELLER" || data == "?KALSEL.SYA.MENU.TEL" || data == "?KALSEL.SYA.MENU.TELSKN" || data == "?KALSEL.SYA.MENU.TELKFS" ||
		data == "?KALSEL.MENU.MAIN.TELLERWEEKEND" {
		return "T"
	} else if data == "?FDS.MENU.HEAD.TELLER" || data == "?KALSEL.MENU.MAIN.HT" || data == "?KALSEL.SYA.HTEL" || data == "?KALSEL.SYA.MENU.HTELSKN" || data == "?KALSEL.MENU.MAIN.KBCS" {
		return "H"
	} else if data == "?FDS.MENU.MAIN.BO.CABANG" || data == "?KALSEL.MENU.MAIN.BOFLA" || data == "?KALSEL.MENU.MAIN.BOFKL" ||
		data == "?KALSEL.MENU.MAIN.KBBOFKL" || data == "?KALSEL.APV.USER.MENU.KBBOF" || data == "?KALSEL.SYA.MENU.BO" || data == "?KALSEL.SYA.BOSKN" ||
		data == "?KALSEL.SYA.MENU.BOSKN" || data == "?KALSEL.SYA.MENU.KTKP" || data == "?KALSEL.MENU.BO.KP.PEJABAT" || data == "?KALSEL.MENU.MAIN.BO.KP.STAF" ||
		data == "?KALSEL.SYA.MENU.TKP" || data == "?KALSEL.SYA.MENU.TELKFS" {
		return "B"
	} else if data == "?KALSEL.SYA.MENU.CS" || data == "?KALSEL.SYA.CS" {
		return "CS"
	} else if data == "" || data == "?KALSEL.MENU.MAIN.DIVISI" {
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

func ParseBranchCode(data string) string {
	branchCode := strings.ReplaceAll(data, " ", "")
	return branchCode[0:9]
}
