package restful

import (
	"encoding/json"
	"fmt"
	"github.com/xiaomy1024/huobi_futures/sdk/linearswap"
	"github.com/xiaomy1024/huobi_futures/sdk/linearswap/restful/response/account"
	"github.com/xiaomy1024/huobi_futures/sdk/log"
	"github.com/xiaomy1024/huobi_futures/sdk/reqbuilder"
)

type AccountClient struct {
	PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (ac *AccountClient) Init(accessKey string, secretKey string, host string) *AccountClient {
	if host == "" {
		host = linearswap.LINEAR_SWAP_DEFAULT_HOST
	}
	ac.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return ac
}

func (ac *AccountClient) GetBalanceValuationAsync(valuationAsset string) (*account.GetBalanceValuationResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_balance_valuation", nil)

	// content
	content := ""
	if valuationAsset != "" {
		content = fmt.Sprintf(",\"valuation_asset\": \"%s\"", valuationAsset)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.GetBalanceValuationResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetBalanceValuationResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) IsolatedGetAccountInfoAsync(contractCode string, subUid int64) (*account.GetAccountInfoResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_account_info", nil)
	if subUid != 0 {
		url = ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_sub_account_info", nil)
	}

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if subUid != 0 {
		content += fmt.Sprintf(",\"sub_uid\": %d", subUid)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.GetAccountInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAccountInfoResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) CrossGetAccountInfoAsync(marginAccount string, subUid int64) (*account.GetAccountInfoResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_account_info", nil)
	if subUid != 0 {
		url = ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_sub_account_info", nil)
	}

	// content
	content := ""
	if marginAccount != "" {
		content = fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	}
	if subUid != 0 {
		content += fmt.Sprintf(",\"sub_uid\": %d", subUid)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.GetAccountInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAccountInfoResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) IsolatedGetAccountPositionAsync(contractCode string, subUid int64) (*account.GetAccountPositionResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_position_info", nil)
	if subUid != 0 {
		url = ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_sub_account_info", nil)
	}

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if subUid != 0 {
		content += fmt.Sprintf(",\"sub_uid\": %d", subUid)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.GetAccountPositionResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAccountPositionResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) CrossGetAccountPositionAsync(contractCode string, subUid int64) (*account.GetAccountPositionResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_position_info", nil)
	if subUid != 0 {
		url = ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_sub_position_info", nil)
	}

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if subUid != 0 {
		content += fmt.Sprintf(",\"sub_uid\": %d", subUid)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.GetAccountPositionResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAccountPositionResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) IsolatedGetAssetsPositionAsync(contractCode string) (*account.GetAssetsPositionResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_account_position_info", nil)

	// content
	content := fmt.Sprintf("{\"contract_code\": \"%s\"}", contractCode)

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.GetAssetsPositionResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAssetsPositionResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) CrossGetAssetsPositionAsync(marginAccount string) (*account.GetAssetsPositionResponseSingle, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_account_position_info", nil)

	// content
	content := fmt.Sprintf("{\"margin_account\": \"%s\"}", marginAccount)

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.GetAssetsPositionResponseSingle{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAssetsPositionResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) SetSubAuthAsync(subUid string, subAuth int) (*account.SetSubAuthResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_sub_auth", nil)

	// content
	content := fmt.Sprintf("{\"sub_uid\": \"%s\", \"sub_auth\": %d}", subUid, subAuth)

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.SetSubAuthResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SetSubAuthResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) IsolatedGetSubAccountListResponseAsync(contractCode string) (*account.GetSubAccountListResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_sub_account_list", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.GetSubAccountListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAllSubAssetsResponse error: %s", getErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) CrossGetSubAccountListAsync(marginAccount string) (*account.GetSubAccountListResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_sub_account_list", nil)

	// content
	content := ""
	if marginAccount != "" {
		content = fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.GetSubAccountListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSubAccountListResponse error: %s", getErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) IsolatedGetSubAccountInfoListAsync(
	contractCode string, pageIndex int, pageSize int) (*account.GetSubAccountInfoListResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_sub_account_info_list", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if pageIndex != 0 {
		content = fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content = fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.GetSubAccountInfoListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSubAccountInfoListResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) CrossGetSubAccountInfoListAsync(
	marginAccount string, pageIndex int, pageSize int) (*account.GetSubAccountInfoListResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_sub_account_info_list", nil)

	// content
	content := ""
	if marginAccount != "" {
		content = fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	}
	if pageIndex != 0 {
		content = fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content = fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.GetSubAccountInfoListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSubAccountInfoListResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) AccountTransferAsync(asset string, fromMarginAccount string, toMarginAccount string, amount float64,
	subUid int64, fcType string) (*account.AccountTransferResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_master_sub_transfer", nil)
	if subUid == 0 {
		url = ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_transfer_inner", nil)
	}

	// content
	content := fmt.Sprintf(",\"asset\":\"%s\", \"from_margin_account\":\"%s\", \"to_margin_account\":\"%s\", \"amount\":%f",
		asset, fromMarginAccount, toMarginAccount, amount)
	if subUid != 0 {
		content += fmt.Sprintf(",\"sub_uid\": %d,\"type\": \"%s\"", subUid, fcType)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.AccountTransferResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to AccountTransferResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) GetAccountTransHisAsync(marginAccount string,
	beMasterSub bool, fcType string, createDate int, pageIndex int, pageSize int) (*account.GetAccountTransHisResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_financial_record", nil)
	if beMasterSub {
		url = ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_master_sub_transfer_record", nil)
	}

	// content
	content := fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	if fcType != "" {
		content += fmt.Sprintf(",\"type\": \"%s\"", fcType)
		if beMasterSub {
			content += fmt.Sprintf(",\"transfer_type\": \"%s\"", fcType)
		}
	}
	if createDate != 0 {
		content += fmt.Sprintf(",\"create_date\": %d", createDate)
	} else {
		if beMasterSub {
			createDate = 7
			content += fmt.Sprintf(",\"create_date\": %d", createDate)
		}
	}
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.GetAccountTransHisResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAccountTransHisResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) GetFinancialRecordExactAsync(marginAccount string,
	contractCode string, fcType string, startTime int64, endTime int64, fromId int64, size int, direct string) (*account.GetFinancialRecordExactResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_financial_record_exact", nil)

	// content
	content := fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if fcType != "" {
		content += fmt.Sprintf(",\"type\": \"%s\"", fcType)
	}
	if startTime != 0 {
		content += fmt.Sprintf(",\"start_time\": %d", startTime)
	}
	if endTime != 0 {
		content += fmt.Sprintf(",\"end_time\": %d", endTime)
	}
	if fromId != 0 {
		content += fmt.Sprintf(",\"from_id\": %d", fromId)
	}
	if size != 0 {
		content += fmt.Sprintf(",\"size\": %d", size)
	}
	if direct != "" {
		content += fmt.Sprintf(",\"direct\": \"%s\"", direct)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.GetFinancialRecordExactResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetFinancialRecordExactResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) IsolatedGetSettlementRecordsAsync(contractCode string,
	startTime int64, endTime int64, pageIndex int, pageSize int) (*account.IsolatedGetSettlementRecordsResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_user_settlement_records", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if startTime != 0 {
		content += fmt.Sprintf(",\"start_time\": %d", startTime)
	}
	if endTime != 0 {
		content += fmt.Sprintf(",\"end_time\": %d", endTime)
	}
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.IsolatedGetSettlementRecordsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to IsolatedGetSettlementRecordsResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) CrossGetSettlementRecordsAsync(marginAccount string,
	startTime int64, endTime int64, pageIndex int, pageSize int) (*account.CrossGetSettlementRecordsResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_user_settlement_records", nil)

	// content
	content := fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	if startTime != 0 {
		content += fmt.Sprintf(",\"start_time\": %d", startTime)
	}
	if endTime != 0 {
		content += fmt.Sprintf(",\"end_time\": %d", endTime)
	}
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.CrossGetSettlementRecordsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to CrossGetSettlementRecordsResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) IsolatedGetValidLeverRateAsync(contractCode string) (*account.GetValidLeverRateResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_available_level_rate", nil)

	// content
	content := fmt.Sprintf("{ \"contract_code\": \"%s\" }", contractCode)
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.GetValidLeverRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetValidLeverRateResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) CrossGetValidLeverRateAsync(contractCode string) (*account.GetValidLeverRateResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_available_level_rate", nil)

	// content
	content := fmt.Sprintf("{ \"contract_code\": \"%s\" }", contractCode)
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.GetValidLeverRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetValidLeverRateResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) GetOrderLimitAsync(orderPriceType string, contractCode string) (*account.GetOrderLimitResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_order_limit", nil)

	// content
	content := fmt.Sprintf(",\"order_price_type\":\"%s\"", orderPriceType)
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.GetOrderLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetOrderLimitResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) GetFeeAsync(contractCode string) (*account.GetFeeResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_fee", nil)

	// content
	content := fmt.Sprintf("{ \"contract_code\": \"%s\" }", contractCode)
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.GetFeeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetFeeResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) IsolatedGetTransferLimitAsync(contractCode string) (*account.GetTransferLimitResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_transfer_limit", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.GetTransferLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetTransferLimitResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) CrossGetTransferLimitAsync(marginAccount string) (*account.GetTransferLimitResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_transfer_limit", nil)

	// content
	content := ""
	if marginAccount != "" {
		content += fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.GetTransferLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetTransferLimitResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) IsolatedGetPositionLimitAsync(contractCode string) (*account.GetPositionLimitResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_position_limit", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.GetPositionLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetPositionLimitResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) CrossGetPositionLimitAsync(contractCode string) (*account.GetPositionLimitResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_position_limit", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.GetPositionLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetPositionLimitResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}

func (ac *AccountClient) GetApiTradingStatusAsync(contractCode string) (*account.GetApiTradingStatusResponse, error) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/linear-swap-api/v1/swap_api_trading_status", nil)

	// content is nil
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
		return nil, getErr
	}
	result := account.GetApiTradingStatusResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetApiTradingStatusResponse error: %s", jsonErr)
		return nil, jsonErr
	}
	return &result, nil
}
