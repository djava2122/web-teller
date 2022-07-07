package repo

import (
	"git.pactindo.com/ebanking/common/pg"
)

const FindRole = "SELECT init_role, role FROM m_teller_roles WHERE init_role=$1;"

func (_ transaction) FindRole(custRef string) (string, string, error) {
	var initrole, role string
	stmt, err := pg.DB.Prepare(FindRole)
	if err != nil {
		return "", "", err
	}
	defer stmt.Close()
	err = stmt.QueryRow(custRef).Scan(&initrole, &role)
	if err != nil {
		return "", "", err
	}
	return initrole, role, nil
}
