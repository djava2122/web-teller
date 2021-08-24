package service

import (
	"context"
	// "encoding/json"
	"strconv"
	"strings"
	"time"

	"gitlab.pactindo.com/ebanking/common/constant"
	"gitlab.pactindo.com/ebanking/common/response"
	cutil "gitlab.pactindo.com/ebanking/common/util"
	feature_and_product "gitlab.pactindo.com/ebanking/proto-common/feature-and-product"
	"gitlab.pactindo.com/ebanking/proto-common/fee"
	feesvc "gitlab.pactindo.com/ebanking/proto-common/fee"
	csvc "gitlab.pactindo.com/ebanking/proto-ibmb/customer"
	msg "gitlab.pactindo.com/ebanking/proto-ibmb/eb-message"
	sysparam "gitlab.pactindo.com/ebanking/proto-ibmb/system-parameter"
	trxsvc "gitlab.pactindo.com/ebanking/proto-ibmb/transaction"

	"gitlab.pactindo.com/ebanking/common/log"
	"gitlab.pactindo.com/ebanking/common/transport"
	"gitlab.pactindo.com/ebanking/common/trycatch"
	"gitlab.pactindo.com/ebanking/transfer/model"
	"gitlab.pactindo.com/ebanking/transfer/util"

	wtproto "gitlab.pactindo.com/ebanking/web-teller/proto"
)

type transferHandler struct {
	custSrv        csvc.CustomerService
	feeSrv         feesvc.FeeService
	systemSrv      sysparam.SystemParameterService
	transactionSrv trxsvc.TransactionService
	ftrSvc         feature_and_product.FeatureAndProductService
}

func (c *transferHandler) getFeature(featureCode, reqId string) *feature_and_product.DataFeature {
	reqFtr := feature_and_product.DataFeature{
		RequestId:   reqId,
		FeatureCode: featureCode,
	}
	res, err := c.ftrSvc.GetFeature(context.TODO(), &reqFtr)
	if err != nil || res.ResponseCode != constant.Success {
		log.Infof("feature not found")
		return nil
	}
	return res
}

func (h *WebTellerHandler) ConfirmTransferOverbook(ctx context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {
	defer func() {
		log.Infof("[%s] response: %v", req.Headers["Request-ID"], string(res.Response))
	}()
	defer trycatch.Catch(func(e trycatch.Exception) {
		log.Infof("[%s] Exception %v", req.Headers["Request-ID"], e)
		res.Response = response.InternalServerError()
	})
	jsonReq, _ := json.Marshal(req)
	log.Infof("[%s] request: %v", req.Headers["Request-ID"], string(jsonReq))

	if strings.EqualFold(req.Params["srcAccNumber"], req.Params["destAccNumber"]) {
		res.Response, _ = json.Marshal(model.ResError("41", "Nomor rekening tujuan tidak bisa di daftarkan"))
		return nil
	}

	result := transport.SendToGate("gate.shared", "06", map[string]string{
		"termId":      req.Headers["Term-Id"],
		"srcAccount":  req.Params["srcAccNumber"],
		"destAccount": req.Params["destAccNumber"],
	})

	if result.ResponseCode != "00" {
		res.Response, _ = json.Marshal(model.ResError(result.ResponseCode, result.Description))
		return nil
	}

	//var reqGetFee = fee.ReqFee{FeatureCode: req.Params["featureCode"]}
	//rF, err := c.feeSrv.GetFeatureFee(context.TODO(), &reqGetFee)
	//if err != nil {
	//	panic(err)
	//}
	//if rF.Rc != constant.Success {
	//	res.Response = response.InvalidFee()
	//	return nil
	//}

	var amount, vatype, feeVa, dueDate string
	if v, ok := result.Data["amount"].(string); ok {
		amount = v
	}
	if v, ok := result.Data["vaType"].(string); ok {
		vatype = v
	}
	if v, ok := result.Data["dueDate"].(string); ok {
		dueDate = v
	}
	if v, ok := result.Data["fee"].(string); ok {
		feeVa = v
	} else {
		//feeVa = strconv.Itoa(int(rF.Fee.Charge))
	}
	var dataResTransfer = model.Transfers{
		AccNumber: req.Params["destAccNumber"],
		AccName:   result.Data["receiverName"].(string),
		Fee:       feeVa,
		TxDate:    time.Now().Format(util.DefaultFormatDate),
		Amount:    amount,
		VaType:    vatype,
		DueDate:   dueDate,
	}
	res.Response, _ = json.Marshal(model.ResSuccess(&dataResTransfer))
	return nil
}

func (c *transferHandler) PostingTransferOverbook(ctx context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {
	var receipt = make(map[string]interface{})
	receipt["txRefference"] = "-"
	receipt["txDate"] = time.Now().Format(util.DefaultFormatDate)
	receipt["txStatus"] = constant.FAILED
	defer func() {
		log.Infof("[%s] response: %v", req.Headers["Request-ID"], string(res.Response))
	}()
	defer trycatch.Catch(func(e trycatch.Exception) {
		log.Infof("[%s] Exception %v", req.Headers["Request-ID"], e)
		res.Response = response.SuccessMsg(constant.ResInternalError, receipt)
	})

	jsonReq, _ := json.Marshal(req)
	jsonReqParam, _ := json.Marshal(req.Params)
	log.Infof("[%s] request: %v", req.Headers["Request-ID"], string(jsonReq))

	reqPosting := model.RequestPosting{}
	err := json.Unmarshal(jsonReqParam, &reqPosting)
	if err != nil {
		panic("error parsing requestData")
	}

	reqDataPosting := model.JsonTransfer{}
	err = json.Unmarshal([]byte(reqPosting.JSON), &reqDataPosting)
	if err != nil {
		panic("error unmarshal json transaction")
	}
	ftr := c.getFeature(reqDataPosting.FeatureCode, req.Headers["Request-ID"])
	if ftr != nil {
		reqDataPosting.FeatureName = ftr.FeatureName
	}
	reqDataPosting.Cif = reqPosting.Cif
	reqDataPosting.CustomerId = reqPosting.CustomerID
	reqDataPosting.Username = reqPosting.Username

	if strings.EqualFold(reqDataPosting.SrcAccNumber, reqDataPosting.DestAccNumber) {
		res.Response = response.SuccessMsg("Invalid Destination Account", receipt)
		return nil
	}

	var txfee = 0.0
	if reqDataPosting.TxFee != "" {
		txfee, err = strconv.ParseFloat(reqDataPosting.TxFee, 64)
		if err != nil {
			log.InfoS(err.Error())
			res.Response = response.SuccessMsg(constant.ResInvalidFee, receipt)
			return nil
		}

	}

	var reqGetFee = fee.ReqFee{FeatureCode: reqDataPosting.FeatureCode,
		RequestId: req.Headers["Request-ID"]}
	rF, err := c.feeSrv.GetFeatureFee(ctx, &reqGetFee)
	if err != nil {
		log.InfoS(err.Error())
		res.Response = response.SuccessMsg(constant.ResInvalidFee, receipt)
		return nil
	}
	if rF.Rc != constant.Success {
		res.Response = response.SuccessMsg(constant.ResInvalidFee, receipt)
		return nil
	}

	if txfee != float64(rF.Fee.Charge) && reqDataPosting.FeatureCode != "103" {
		res.Response = response.SuccessMsg(constant.ResInvalidFee, receipt)
		return nil
	}

	amount, err := strconv.ParseFloat(reqDataPosting.TxAmount, 64)
	if err != nil {
		log.InfoS("error parsing ammount")
		res.Response = response.SuccessMsg(constant.ResInvalidAmount, receipt)
		return nil
	}
	customerId, _ := strconv.Atoi(reqPosting.CustomerID)

	mapFields := util.ConvertMap(cutil.StructToMap(&reqDataPosting))
	mapFields["freeData2"] = reqDataPosting.DestAccName
	mapFields["destBankCode"] = "122"
	mapFields["destBankName"] = "BPD Kalsel"
	mapFields["freeData4"] = "BPD Kalsel"
	mapFields["termType"] = req.Headers["Channel"]
	mapFields["termId"] = req.Headers["Term-Id"]

	reqTransaction := msg.TransactionMsg{
		Id:                req.Headers["Request-ID"],
		TransactionType:   "01",
		FeatureId:         reqDataPosting.FeatureID,
		FeatureName:       reqDataPosting.FeatureName,
		CustomerReference: reqDataPosting.DestAccNumber,
		BillerName:        "BPD Kalsel",
		CustomerId:        int32(customerId),
		TrxToken:          reqPosting.TrxToken,
		AccountNumber:     reqDataPosting.SrcAccNumber,
		Amount:            amount,
		Fee:               txfee,
		Description:       reqDataPosting.TxDescription,
		ToAccountNumber:   reqDataPosting.DestAccNumber,
		DeviceId:          reqPosting.DeviceID,
		CurrencyCode:      reqDataPosting.CurrencyCode,
		TransactionDate:   time.Now().Unix(),
		SrcAccProdType:    reqDataPosting.SrcAccProdType,
		FeatureCode:       reqDataPosting.FeatureCode,
		Fields:            mapFields,
	}

	result, err := c.transactionSrv.Post(ctx, &reqTransaction)
	if err != nil {
		log.InfoS(err.Error())
		res.Response = response.SuccessMsg(constant.ResInternalError, receipt)
		return nil
	}
	trxDate := util.ParseIntToDate(result.TransactionDate)
	receipt["txRefference"] = result.ReferenceNumber
	receipt["txDate"] = trxDate.Format(util.DefaultFormatDate)
	receipt["txStatus"] = result.TransactionStatus
	res.Response = response.SuccessMsg(result.ResponseDescription, receipt)
	return nil
}
