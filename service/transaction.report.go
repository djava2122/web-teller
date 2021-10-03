package service

import (
	"context"
	"strconv"

	"gitlab.pactindo.com/ebanking/common/log"
	"gitlab.pactindo.com/ebanking/common/trycatch"
	wtproto "gitlab.pactindo.com/ebanking/web-teller/proto"
	"gitlab.pactindo.com/ebanking/web-teller/repo"
)

func (h *WebTellerHandler) TransactionReport(ctx context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {
	defer func() {
		log.Infof("[%s] response: %v", req.Headers["Request-ID"], string(res.Response))
	}()
	defer trycatch.Catch(func(e trycatch.Exception) {
		log.Infof("[%s] error : %v", req.Headers["Request-ID"], e)
		res.Response, _ = json.Marshal(newResponse("99", "Internal Server Error"))
	})

	jsonReq, _ := json.Marshal(req.Params)
	log.Infof("[%s] request: %v", req.Headers["Request-ID"], string(jsonReq))

	//if req.Params["featureCode"] == "" {
	//	res.Response, _ = json.Marshal(newResponse("01", "params featureCode cannot be empty"))
	//	return nil
	//}
	//
	//if req.Params["startDate"] == "" {
	//	res.Response, _ = json.Marshal(newResponse("01", "params startDate cannot be empty"))
	//	return nil
	//}
	//
	//if req.Params["endDate"] == "" {
	//	res.Response, _ = json.Marshal(newResponse("01", "params endDate cannot be empty"))
	//	return nil
	//}

	datas, err := repo.Transaction.Filter(req.Params["teller"])
	if err != nil {
		log.Errorf("error get data transaction: %v", err)
	}
	log.Infof("Request-ID:[%s] Isi Data: %v", req.Headers["Request-ID"], datas)
	if len(datas) != 0 {
		res.Response, _ = json.Marshal(successResp(ConvertStructTransactionToResult(datas)))
	} else {
		res.Response, _ = json.Marshal(newResponse("80", "Data Not Found"))
	}

	return nil
}

func ConvertStructTransactionToResult(transaction []repo.MTransaction) []*repo.TransactionReport {
	var result []*repo.TransactionReport
	for _, val := range transaction {
		data := repo.TransactionReport{}
		data.FeatureName = val.FeatureName
		data.FeatureCode = strconv.Itoa(val.FeatureCode)
		data.FeatureGroupCode = val.FeatureGroupCode
		data.FeatureGroupName = val.FeatureGroupName
		data.TransactionDate = FormattedTime(val.TransactionDate, "2006-01-02 15:04:05")
		data.TransactionAmount = val.TransactionAmount
		data.Fee = val.Fee
		data.TransactionStatus = val.TransactionStatus
		data.ReferenceNumber = val.ReferenceNumber
		data.CustomerReference = val.CustomerReference
		data.CurrencyCode = val.CurrencyCode
		data.CreatedBy = val.CreatedBy
		data.BranchCode = val.BranchCode
		result = append(result, &data)
	}
	return result
}
