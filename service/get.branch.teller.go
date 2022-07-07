package service

import (
	"context"

	"git.pactindo.com/ebanking/common/log"
	"git.pactindo.com/ebanking/common/pg"
	"git.pactindo.com/ebanking/common/trycatch"
	wtproto "git.pactindo.com/ebanking/web-teller/proto"
)

type branchTeller struct {
	BranchCode string `json:"branch_code"`
	BranchName string `json:"branch_name"`
	Address    string `json:"address"`
}

func (h *WebTellerHandler) GetBranchTeller(_ context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {
	defer func() {
		log.Infof("[%s] response: %v", req.Headers["Request-ID"], string(res.Response))
	}()
	defer trycatch.Catch(func(e trycatch.Exception) {
		log.Infof("[%s] error : %v", req.Headers["Request-ID"], e)
		res.Response, _ = json.Marshal(newResponse("99", "Internal Server Error"))
	})

	sql := "select branch_code,branch_name, address from m_branch_teller"
	var result []branchTeller
	rows, err := pg.DB.Query(sql)
	if err != nil {
		return nil
	}
	//log.Infof("Test query:", rows)
	for rows.Next() {
		datas := branchTeller{}
		err := rows.Scan(
			&datas.BranchCode,
			&datas.BranchName,
			&datas.Address,
		)
		if err != nil {
			return nil
		}
		result = append(result, datas)
	}
	res.Response, _ = json.Marshal(newResponseWithData("00", "Sukses", result))
	return nil
}
