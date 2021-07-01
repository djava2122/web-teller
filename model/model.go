package model

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

//RestResult for api response
type RestResult struct {
	Code    string        `json:"code"`
	Message string        `json:"message"`
	TrxFlag string        `json:"trxFlag"`
	Data    []interface{} `json:"data"`
}

// ResSuccess .
func ResSuccess(resData ...interface{}) RestResult {
	r := RestResult{}
	strs := []string{}
	names := make([]interface{}, len(strs))
	r.Code = "00"
	r.Message = "Berhasil|Success"
	r.TrxFlag = "0"
	if len(resData) == 0 {
		r.Data = names
		return r
	}
	for _, val := range resData {
		if !reflect.ValueOf(val).IsNil() {
			r.Data = append(r.Data, val)
		} else {
			fmt.Print("data null")
		}
	}
	return r
}

// ResSuccess .
func ResSuccess2(resData ...interface{}) RestResult {
	r := RestResult{}
	r.Code = "00"
	r.Message = "Berhasil|Success"
	r.TrxFlag = "0"
	for _, val := range resData {
		data := reflect.ValueOf(val)
		if !reflect.ValueOf(&val).IsNil() {
			if reflect.TypeOf(val).Kind() == reflect.Array || reflect.TypeOf(val).Kind() == reflect.Slice {
				for i := 0; i < data.Len(); i++ {
					r.Data = append(r.Data, data.Index(i).Interface())
				}
			} else {
				r.Data = append(r.Data, val)
			}
		}
	}

	return r
}

// ResError .
func ResError(responseCode string, message string) RestResult {
	str := strings.Split(responseCode, "|")
	r := RestResult{}
	r.Code = str[0]
	r.Message = message
	r.TrxFlag = "3"
	if len(str) > 1 {
		if str[1] != "" {
			r.TrxFlag = str[1]
		}
	}
	strs := []string{}
	names := make([]interface{}, len(strs))
	r.Data = names
	return r
}

func ResErrorData(responseCode string, message string, resData interface{}) RestResult {
	str := strings.Split(responseCode, "|")
	r := RestResult{}
	r.Code = str[0]
	r.Message = message
	r.TrxFlag = "3"
	if len(str) > 1 {
		if str[1] != "" {
			r.TrxFlag = str[1]
		}
	}

	r.Data = append(r.Data, resData)
	return r
}

type BaseSetruct struct {
	Created   sql.NullTime   `json:"created,omitempty"`
	CreatedBy sql.NullString `json:"createdBy,omitempty"`
	Updated   sql.NullTime   `json:"updated,omitempty"`
	UpdateBy  sql.NullString `json:"updatedBy,omitempty"`
}

type ReqValidationTransfer struct {
	CustomerId string `json:"customerId, omitempty" validate:"required"`
	FeatureId  string `json:"featureId, omitempty" validate:"required"`
}

type ReqValidationCreateTransfer struct {
	CustomerId    string `json:"customerId, omitempty" validate:"required"`
	FeatureId     string `json:"featureId, omitempty" validate:"required"`
	DestAccNumber string `json:"destAccNumber, omitempty" validate:"required"`
	SrcAccNumber  string `json:"destAccNumber, omitempty" validate:"required"`
}

type TransferData struct {
	Id                int32          `json:"id,omitempty"`
	CustomerId        int32          `json:"customerId,omitempty"`
	FeatureId         int32          `json:"featureId,omitempty"`
	CustomerReference string         `json:"customerReference,omitempty"`
	RegisterAlias     string         `json:"registerAlias,omitempty"`
	UseCount          int32          `json:"useCount,omitempty"`
	Currency          string         `json:"currency,omitempty"`
	Data1             string         `json:"data1,omitempty"`
	Data2             string         `json:"data2,omitempty"`
	Data3             string         `json:"data3,omitempty"`
	Data4             sql.NullString `json:"data4,omitempty"`
	Data5             sql.NullString `json:"data5,omitempty"`
	*BaseSetruct
	IsFavourite string `json:"isFavourite,omitempty"`
}

type Transfers struct {
	ID          string `json:"id,omitempty"`
	BankName    string `json:"bankName,omitempty"`
	BankCode    string `json:"bankCode,omitempty"`
	AccNumber   string `json:"accNumber,omitempty"`
	AccName     string `json:"accName,omitempty"`
	Alias       string `json:"alias"`
	Fee         string `json:"fee"`
	TxDate      string `json:"txDate"`
	IsFavourite string `json:"isFavourite"`
	Amount      string `json:"amount,omitempty"`
	VaType      string `json:"vaType,omitempty"`
	DueDate     string `json:"dueDate,omitempty"`
}

type TransferReq struct {
	CustomerId     string `json:"customerId,omitempty"`
	DestBankName   string `json:"destBankName,omitempty"`
	DestBankCode   string `json:"destBankCode,omitempty"`
	DestAccNumber  string `json:"destAccNumber,omitempty"`
	DestAccName    string `json:"destAccName,omitempty"`
	SrcAccName     string `json:"srcAccName"`
	SrcAccProdType string `json:"srcAccProdType"`
	TransferAlias  string `json:"transferAlias"`
	FeatureId      string `json:"featureId,omitempty"`
	SrcAccNumber   string `json:"srcAccNumber, omitempty"`
	Fee            string `json:"fee, omitempty"`
}

type ConfirmIBFTRes struct {
	CustomerId    string `json:"customerId,omitempty"`
	DestBankName  string `json:"destBankName,omitempty"`
	DestBankCode  string `json:"destBankCode,omitempty"`
	AccNumber     string `json:"accNumber,omitempty"`
	AccName       string `json:"accName,omitempty"`
	TransferAlias string `json:"transferAlias"`
	FeatureId     string `json:"featureId,omitempty"`
	SrcAccNumber  string `json:"srcAccNumber, omitempty"`
	Fee           string `json:"fee, omitempty"`
	TxDate        string `json:"txDate"`
}

type TransferConfirmReq struct {
	CustomerId    string `json:"customerId,omitempty"`
	DeviceId      string `json:"destBankName,omitempty"`
	TrfListId     string `json:"trfListId,omitempty"`
	SrcAccNumber  string `json:"srcAccNumber,omitempty"`
	DestAccNumber string `json:"destAccNumber,omitempty"`
	TxReffNumber  string `json:"txReffNumber,omitempty"`
	TxAmount      string `json:"txAmount,omitempty"`
	TxDescription string `json:"txDescription"`
}

type CustomerInfoCore struct {
	Address      string `json:"address"`
	BirtDate     string `json:"birtDate"`
	BirthPlace   string `json:"birthPlace"`
	BranchCode   string `json:"branchCode"`
	Cif          string `json:"cif"`
	CustomerName string `json:"customerName"`
	Email        string `json:"email"`
	Gender       string `json:"gender"`
	LegalDocExp  string `json:"legalDocExp"`
	LegalDocName string `json:"legalDocName"`
	LegalID      string `json:"legalID"`
	MotherName   string `json:"motherName"`
	PhoneNumber  string `json:"phoneNumber"`
}

type RequestPosting struct {
	Cif        string `json:"cif"`
	DeviceID   string `json:"deviceId"`
	CustomerID string `json:"customerId"`
	Username   string `json:"username"`
	TrxToken   string `json:"trxToken"`
	JSON       string `json:"json"`
}

type JsonTransfer struct {
	CustomerId              string `json:"customerId"`
	Cif                     string `json:"cif"`
	Username                string `json:"username"`
	TxFee                   string `json:"txFee"`
	FeatureID               string `json:"featureId"`
	FeatureName             string `json:"featureName"`
	FeatureCode             string `json:"featureCode"`
	TrfListID               string `json:"trfListId"`
	SrcAccNumber            string `json:"srcAccNumber"`
	SrcAccName              string `json:"srcAccName"`
	SrcAccProdType          string `json:"srcAccProdType"`
	DestAccNumber           string `json:"destAccNumber"`
	DestAccName             string `json:"destAccName"`
	CurrencyCode            string `json:"currencyCode"`
	TxAmount                string `json:"txAmount"`
	TxDescription           string `json:"txDescription"`
	TxDate                  string `json:"txDate"`
	DestBankCode            string `json:"destBankCode"`
	DestBankName            string `json:"destBankName"`
	CustomerReferenceNumber string `json:"customerReferenceNumber"`
	TxExecutionType         string `json:"txExecutionType"`
	PeriodType              string `json:"periodType"`
	PeriodValue             string `json:"periodValue"`
	EffectiveDate           string `json:"effectiveDate"`
	ExpireDate              string `json:"expireDate"`
	Content                 string `json:"content"`
}

type AutoGenerated struct {
	FeatureID       string `json:"featureId"`
	FeatureCode     string `json:"featureCode"`
	FeatureName     string `json:"featureName"`
	SrcAccNumber    string `json:"srcAccNumber"`
	SrcAccProdType  string `json:"srcAccProdType"`
	TxExecutionType string `json:"txExecutionType"`
	PeriodType      string `json:"periodType"`
	PeriodValue     string `json:"periodValue"`
	Content         []struct {
		FeatureID     string `json:"featureId"`
		FeatureCode   string `json:"featureCode"`
		FeatureName   string `json:"featureName"`
		DestAccNumber string `json:"destAccNumber"`
		DestAccName   string `json:"destAccName"`
		CurrencyCode  string `json:"currencyCode"`
		TxAmount      string `json:"txAmount"`
		TxDescription string `json:"txDescription"`
	} `json:"content"`
}

type TransactionHistory struct {
	Id                string `json:"id"`
	FeatureId         string `json:"featureId"`
	FeatureName       string `json:"featureName"`
	CustomerReference string `json:"customerReference"`
	BillerName        string `json:"billerName"`
	TransactionName   string `json:"transactionName"`
	TransactionAlias  string `json:"transactionAlias"`
	TransactionDate   string `json:"transactionDate"`
}

type InqTransferIBFT struct {
	SrcAccNumber      string `json:"srcAccNumber"`
	SrcAccName        string `json:"srcAccName"`
	SrcAccType        string `json:"srcAccType"`
	DestBankCode      string `json:"destBankCode"`
	DestBankName      string `json:"destBankName"`
	DestAccNumber     string `json:"destAccNumber"`
	DestAccType       string `json:"destAccType"`
	TxAmount          string `json:"txAmount"`
	Currency          string `json:"currency"`
	TxReferenceNumber string `json:"txReferenceNumber"`
}

type ResInqIBFT struct {
	DestBankCode      string `json:"destBankCode"`
	DestBankName      string `json:"destBankName"`
	DestAccNumber     string `json:"destAccNumber"`
	DestAccName       string `json:"destAccName"`
	TxAmount          string `json:"txAmount,omitempty"`
	TxFee             string `json:"txFee,omitempty"`
	TxReferenceNumber string `json:"txReferenceNumber"`
}
