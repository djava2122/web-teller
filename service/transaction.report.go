package service

import (
	"context"
	"fmt"
	"gitlab.pactindo.com/ebanking/common/log"
	"gitlab.pactindo.com/ebanking/common/micro"
	eb_response "gitlab.pactindo.com/ebanking/common/response"
	"gitlab.pactindo.com/ebanking/common/trycatch"
	wtproto "gitlab.pactindo.com/ebanking/web-teller/proto"
	"gitlab.pactindo.com/ebanking/web-teller/repo"
	"io/ioutil"
	"os"
	"strconv"
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

	params := req.Params
	if params == nil {
		params = make(map[string]string)
	}

	page := 1
	pageSize := 50
	if v, ok := params["page"]; ok {
		i, err := strconv.Atoi(v)
		if err != nil {
			res.Response = eb_response.Error("01", "invalid input for page")
			return nil
		}

		if i <= 0 {
			res.Response = eb_response.Error("01", "page must be greater than 0")
			return nil
		}
		page = i
	}

	if v, ok := params["pageSize"]; ok {
		i, err := strconv.Atoi(v)
		if err != nil {
			res.Response = eb_response.Error("01", "invalid input for pageSize")
			return nil
		}

		if i <= 0 {
			res.Response = eb_response.Error("01", "pageSize must be greater than 0")
			return nil
		}
		pageSize = i
	}

	jsonReqParams, _ := json.Marshal(req.Params)
	filter := Filter{}
	_ = json.Unmarshal(jsonReqParams, &filter)

	filterSlice := parseParamsToMapString(&filter)
	filterString := parseSliceToString(filterSlice, true)

	//cabang := ""
	//if req.Params["teller"] == "All" {
	//	cabang = req.Params["cabang"]
	//} else {
	//	cabang = ""
	//}
	data, dataCount, err := repo.Transaction.FilterTransaction(filterString, page, pageSize)
	if err != nil {
		log.Errorf("error get data transaction: %v", err)
	}

	if len(data) != 0 {
		res.Response, _ = json.Marshal(SuccessWithPadding(dataCount, page, pageSize, ConvertStructTransactionToResult(data)))
	} else {
		res.Response, _ = json.Marshal(newResponse("80", "Data Not Found"))
	}

	return nil
}

func (h *WebTellerHandler) GenerateFileReport(ctx context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {
	defer func() {
		log.Infof("[%s] response: %v", req.Headers["Request-ID"], string(res.Response))
	}()
	defer trycatch.Catch(func(e trycatch.Exception) {
		log.Infof("[%s] error : %v", req.Headers["Request-ID"], e)
		res.Response, _ = json.Marshal(newResponse("99", "Internal Server Error"))
	})

	jsonReq, _ := json.Marshal(req.Params)
	log.Infof("[%s] request: %v", req.Headers["Request-ID"], string(jsonReq))

	jsonReqParams, _ := json.Marshal(req.Params)
	filter := Filter{}
	_ = json.Unmarshal(jsonReqParams, &filter)

	filterSlice := parseParamsToMapString(&filter)
	filterString := parseSliceToString(filterSlice, true)

	datas, err := repo.Transaction.TransactionReport(filterString)
	if err != nil {
		panic(err)
	}

	if datas == nil {
		res.Response, _ = json.Marshal(newResponse("", ""))
		return nil
	}

	status, file, err := ExportTransaction(datas)
	if err != nil {
		panic(err)
	}

	if !status {
		res.Response, _ = json.Marshal(newResponse("99", "Failed generate file"))
		return nil
	}

	res.Response, _ = json.Marshal(newResponseWithData("00", "Success", file))
	return nil
}

func (h *WebTellerHandler) DownloadFileReport(ctx context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {
	defer func() {
		log.Infof("[%s] response: %v", req.Headers["Request-ID"], "image")
	}()
	defer trycatch.Catch(func(e trycatch.Exception) {
		log.Infof("[%s] error : %v", req.Headers["Request-ID"], fmt.Sprintf("%v", e))
		res.Response, _ = json.Marshal(newResponse("99", "Internal Server Error"))
	})

	conf := micro.GetConfig()
	file, err := os.Open(conf["FILE_LOCATION"] + req.Params["filename"])
	if file != nil {
		defer file.Close()
	}

	if err != nil {
		panic(err)
	}

	defer file.Close()

	header := map[string]string{}
	header["Content-Disposition"] = "attachment; filename= " + req.Params["filename"]
	//data := make([]byte, 1024)
	bFile, _:= ioutil.ReadFile(conf["FILE_LOCATION"] + req.Params["filename"])
	res.Response = bFile
	res.Headers = header
	return nil
}

func ConvertStructTransactionToResult(transaction []repo.TransactionReport) []*repo.TransactionReport {
	var result []*repo.TransactionReport
	for _, val := range transaction {
		data := repo.TransactionReport{}
		data.FeatureName = val.FeatureName
		data.FeatureCode = val.FeatureCode
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
		data.BranchName = val.BranchName
		data.ResponseCode = val.ResponseCode
		data.SrcAccount = val.SrcAccount
		data.TransactionType = val.TransactionType
		result = append(result, &data)
	}
	return result
}

func parseParamsToMapString(filter *Filter) []string {
	var filterString []string
	if len(filter.Teller) > 0 {
		filterString = append(filterString, fmt.Sprintf("createdby = '%v'", filter.Teller))
	}
	if len(filter.FeatureCode) > 0 {
		filterString = append(filterString, fmt.Sprintf("feature_code = '%v'", filter.FeatureCode))
	}
	if len(filter.Branch) > 0 {
		filterString = append(filterString, fmt.Sprintf("branch_code = '%v'", filter.Branch))
	}
	if len(filter.StartDate) > 0 && len(filter.EndDate) > 0 {
		filterString = append(filterString, fmt.Sprintf("transaction_date between '%s 00:00:00' and '%s 23:59:59'", filter.StartDate, filter.EndDate))
	}
	if len(filter.Status) > 0 {
		filterString = append(filterString, fmt.Sprintf("transaction_status = '%v'", filter.Status))
	}
	return filterString
}

func parseSliceToString(filter []string, count bool) string {
	var filterStr string
	if len(filter) == 0 {
		return " ORDER BY ce.created DESC"
	}
	for i := 0; i < len(filter); i++ {
		if i == 0 {
			filterStr = fmt.Sprintf("WHERE %v ", filter[i])
			if !count {
				if len(filter) == 1 {
					filterStr = fmt.Sprintf("%v ORDER BY id DESC", filterStr)
				}
			}
		} else if (len(filter) - 1) == i {
			if !count {
				filterStr = fmt.Sprintf("%v AND %v ORDER BY id DESC", filterStr, filter[i])
			} else {
				filterStr = fmt.Sprintf("%v AND %v ", filterStr, filter[i])
			}
		} else {
			filterStr = fmt.Sprintf("%v AND %v ", filterStr, filter[i])
		}
	}
	return filterStr
}
