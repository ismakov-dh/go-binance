package binance

import (
	"context"
	"encoding/json"
)

type (
	CreateSubAccountService struct {
		c   *Client
		tag *string
	}
	CreateSubAccountResponse struct {
		SubAccountId string `json:"subaccountId"`
		Email        string `json:"email"`
		Tag          string `json:"tag"`
	}
	CreateSubAccountApiKeyService struct {
		c            *Client
		subAccountId string
		canTrade     bool
		marginTrade  *bool
		futuresTrade *bool
	}
	CreateSubAccountApiKeyResponse struct {
		SubAccountId string `json:"subaccountId"`
		ApiKey       string `json:"apiKey"`
		SecretKey    string `json:"secretKey"`
		CanTrade     bool   `json:"canTrade"`
		MarginTrade  bool   `json:"marginTrade"`
		FuturesTrade bool   `json:"futuresTrade"`
	}
	DeleteSubAccountApiKeyService struct {
		c                *Client
		subAccountId     string
		subAccountApiKey string
	}
	ListSubAccountsService struct {
		c            *Client
		subAccountId *string
		page         *int64
		size         *int64
	}
	SubAccount struct {
		SubAccountId          string `json:"subaccountId"`
		Email                 string `json:"email"`
		Tag                   string `json:"tag"`
		MakerCommission       string `json:"makerCommission"`
		TakerCommission       string `json:"takerCommission"`
		MarginMakerCommission string `json:"marginMakerCommission"`
		MarginTakerCommission string `json:"marginTakerCommission"`
		CreateTime            int64  `json:"createTime"`
	}
	EnableMarginForSubAccountService struct {
		c            *Client
		subAccountId string
		margin       bool
	}
	EnableFuturesForSubAccountService struct {
		c            *Client
		subAccountId string
		futures      bool
	}
	EnableMarginForSubAccountResponse struct {
		SubAccountId string `json:"subaccountId"`
		EnableMargin bool   `json:"enableMargin"`
		UpdateTime   int64  `json:"updateTime"`
	}
	EnableFuturesForSubAccountResponse struct {
		SubAccountId  string `json:"subaccountId"`
		EnableFutures bool   `json:"enableFutures"`
		UpdateTime    int64  `json:"updateTime"`
	}
	AddIPRestrictionForSubAccountService struct {
		c                *Client
		subAccountId     string
		subAccountApiKey string
		ipAddress        string
	}
	AddIPRestrictionForSubAccountResponse struct {
		SubAccountId string `json:"subaccountId"`
		Apikey       string `json:"apikey"`
		Ip           string `json:"ip"`
		UpdateTime   int64  `json:"updateTime"`
	}
	IPRestrictionForSubAccountService struct {
		c                *Client
		subAccountId     string
		subAccountApiKey string
		ipRestrict       bool
	}
	IPRestrictionForSubAccountResponse struct {
		SubAccountId string   `json:"subaccountId"`
		IpRestrict   string   `json:"ipRestrict"`
		Apikey       string   `json:"apikey"`
		IpList       []string `json:"ipList"`
		UpdateTime   int64    `json:"updateTime"`
	}
	UniversalTransferService struct {
		c               *Client
		fromAccountType string
		toAccountType   string
		asset           string
		amount          float64
		fromId          *string
		toId            *string
		clientTranId    *string
	}
	UniversalTransferResponse struct {
		TxnId        int64  `json:"txnId"`
		ClientTranId string `json:"clientTranId"`
	}
	UniversalTransferHistoryService struct {
		c             *Client
		fromId        *string
		toId          *string
		clientTranId  *string
		startTime     *int64
		endTime       *int64
		page          *int32
		limit         *int32
		showAllStatus *bool
	}
	UniversalTransfer struct {
		FromId          string `json:"fromId,omitempty"`
		ToId            string `json:"toId,omitempty"`
		Asset           string `json:"asset"`
		Qty             string `json:"qty"`
		Time            int64  `json:"time"`
		Status          string `json:"status"`
		TxnId           string `json:"txnId"`
		ClientTranId    string `json:"clientTranId"`
		FromAccountType string `json:"fromAccountType"`
		ToAccountType   string `json:"toAccountType"`
	}
)

func (s *CreateSubAccountService) Tag(tag string) *CreateSubAccountService {
	s.tag = &tag
	return s
}

func (s *CreateSubAccountService) createSubAccount(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{}
	if s.tag != nil {
		m["tag"] = *s.tag
	}
	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func (s *CreateSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *CreateSubAccountResponse, err error) {
	data, err := s.createSubAccount(ctx, "/sapi/v1/broker/subAccount", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateSubAccountResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CreateSubAccountApiKeyService) SubAccountId(subAccountId string) *CreateSubAccountApiKeyService {
	s.subAccountId = subAccountId
	return s
}

func (s *CreateSubAccountApiKeyService) CanTrade(canTrade bool) *CreateSubAccountApiKeyService {
	s.canTrade = canTrade
	return s
}

func (s *CreateSubAccountApiKeyService) MarginTrade(marginTrade bool) *CreateSubAccountApiKeyService {
	s.marginTrade = &marginTrade
	return s
}

func (s *CreateSubAccountApiKeyService) FuturesTrade(futuresTrade bool) *CreateSubAccountApiKeyService {
	s.futuresTrade = &futuresTrade
	return s
}

func (s *CreateSubAccountApiKeyService) createSubAccountApiKey(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId": s.subAccountId,
		"canTrade":     s.canTrade,
	}
	if s.marginTrade != nil {
		m["marginTrade"] = *s.marginTrade
	}
	if s.futuresTrade != nil {
		m["futuresTrade"] = *s.futuresTrade
	}
	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func (s *CreateSubAccountApiKeyService) Do(ctx context.Context, opts ...RequestOption) (res *CreateSubAccountApiKeyResponse, err error) {
	data, err := s.createSubAccountApiKey(ctx, "/sapi/v1/broker/subAccountApi", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateSubAccountApiKeyResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *DeleteSubAccountApiKeyService) SubAccountId(subAccountId string) *DeleteSubAccountApiKeyService {
	s.subAccountId = subAccountId
	return s
}

func (s *DeleteSubAccountApiKeyService) SubAccountApiKey(subAccountApiKey string) *DeleteSubAccountApiKeyService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *DeleteSubAccountApiKeyService) deleteSubAccountApiKey(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "DELETE",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId":     s.subAccountId,
		"subAccountApiKey": s.subAccountApiKey,
	}
	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func (s *DeleteSubAccountApiKeyService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	_, err = s.deleteSubAccountApiKey(ctx, "/sapi/v1/broker/subAccountApi", opts...)
	return err
}

func (s *ListSubAccountsService) SubAccountId(subAccountId string) *ListSubAccountsService {
	s.subAccountId = &subAccountId
	return s
}

func (s *ListSubAccountsService) Page(page int64) *ListSubAccountsService {
	s.page = &page
	return s
}

func (s *ListSubAccountsService) Size(size int64) *ListSubAccountsService {
	s.size = &size
	return s
}

func (s *ListSubAccountsService) Do(ctx context.Context, opts ...RequestOption) (res []*SubAccount, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/broker/subAccount",
		secType:  secTypeSigned,
	}
	if s.subAccountId != nil {
		r.setParam("subAccountId", *s.subAccountId)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*SubAccount{}, err
	}
	res = make([]*SubAccount, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*SubAccount{}, err
	}
	return res, nil
}

func (s *EnableMarginForSubAccountService) SubAccountId(subAccountId string) *EnableMarginForSubAccountService {
	s.subAccountId = subAccountId
	return s
}

func (s *EnableMarginForSubAccountService) Margin(margin bool) *EnableMarginForSubAccountService {
	s.margin = margin
	return s
}

func (s *EnableMarginForSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *EnableMarginForSubAccountResponse, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/subAccount/margin",
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId": s.subAccountId,
		"margin":       s.margin,
	}
	r.setFormParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(EnableMarginForSubAccountResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *EnableFuturesForSubAccountService) SubAccountId(subAccountId string) *EnableFuturesForSubAccountService {
	s.subAccountId = subAccountId
	return s
}

func (s *EnableFuturesForSubAccountService) Futures(futures bool) *EnableFuturesForSubAccountService {
	s.futures = futures
	return s
}

func (s *EnableFuturesForSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *EnableFuturesForSubAccountResponse, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/subAccount/futures",
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId": s.subAccountId,
		"futures":      s.futures,
	}
	r.setFormParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(EnableFuturesForSubAccountResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AddIPRestrictionForSubAccountService) SubAccountId(subAccountId string) *AddIPRestrictionForSubAccountService {
	s.subAccountId = subAccountId
	return s
}

func (s *AddIPRestrictionForSubAccountService) SubAccountApiKey(subAccountApiKey string) *AddIPRestrictionForSubAccountService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *AddIPRestrictionForSubAccountService) IPAddress(ipAddress string) *AddIPRestrictionForSubAccountService {
	s.ipAddress = ipAddress
	return s
}

func (s *AddIPRestrictionForSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *AddIPRestrictionForSubAccountResponse, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/subAccountApi/ipRestriction/ipList",
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId":     s.subAccountId,
		"subAccountApiKey": s.subAccountApiKey,
		"ipAddress":        s.ipAddress,
	}
	r.setFormParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(AddIPRestrictionForSubAccountResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *IPRestrictionForSubAccountService) SubAccountId(subAccountId string) *IPRestrictionForSubAccountService {
	s.subAccountId = subAccountId
	return s
}

func (s *IPRestrictionForSubAccountService) SubAccountApiKey(subAccountApiKey string) *IPRestrictionForSubAccountService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *IPRestrictionForSubAccountService) IPRestrict(ipRestrict bool) *IPRestrictionForSubAccountService {
	s.ipRestrict = ipRestrict
	return s
}

func (s *IPRestrictionForSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *IPRestrictionForSubAccountResponse, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/subAccountApi/ipRestriction",
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId":     s.subAccountId,
		"subAccountApiKey": s.subAccountApiKey,
		"ipRestrict":       s.ipRestrict,
	}
	r.setFormParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(IPRestrictionForSubAccountResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UniversalTransferService) FromAccountType(fromAccountType string) *UniversalTransferService {
	s.fromAccountType = fromAccountType
	return s
}
func (s *UniversalTransferService) ToAccountType(toAccountType string) *UniversalTransferService {
	s.toAccountType = toAccountType
	return s
}
func (s *UniversalTransferService) Asset(asset string) *UniversalTransferService {
	s.asset = asset
	return s
}
func (s *UniversalTransferService) Amount(amount float64) *UniversalTransferService {
	s.amount = amount
	return s
}
func (s *UniversalTransferService) FromId(fromId string) *UniversalTransferService {
	s.fromId = &fromId
	return s
}
func (s *UniversalTransferService) ToId(toId string) *UniversalTransferService {
	s.toId = &toId
	return s
}
func (s *UniversalTransferService) ClientTranId(clientTranId string) *UniversalTransferService {
	s.clientTranId = &clientTranId
	return s
}

func (s *UniversalTransferService) Do(ctx context.Context, opts ...RequestOption) (res *UniversalTransferResponse, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/universalTransfer",
		secType:  secTypeSigned,
	}
	m := params{
		"fromAccountType": s.fromAccountType,
		"toAccountType":   s.toAccountType,
		"asset":           s.asset,
		"amount":          s.amount,
	}
	r.setFormParams(m)
	if s.clientTranId != nil {
		r.setParam("clientTranId", *s.clientTranId)
	}
	if s.fromId != nil {
		r.setParam("fromId", *s.fromId)
	}
	if s.toId != nil {
		r.setParam("toId", *s.toId)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(UniversalTransferResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UniversalTransferHistoryService) FromId(fromId string) *UniversalTransferHistoryService {
	s.fromId = &fromId
	return s
}
func (s *UniversalTransferHistoryService) ToId(toId string) *UniversalTransferHistoryService {
	s.toId = &toId
	return s
}
func (s *UniversalTransferHistoryService) ClientTranId(clientTranId string) *UniversalTransferHistoryService {
	s.clientTranId = &clientTranId
	return s
}
func (s *UniversalTransferHistoryService) StartTime(startTime int64) *UniversalTransferHistoryService {
	s.startTime = &startTime
	return s
}
func (s *UniversalTransferHistoryService) EndTime(endTime int64) *UniversalTransferHistoryService {
	s.endTime = &endTime
	return s
}
func (s *UniversalTransferHistoryService) Page(page int32) *UniversalTransferHistoryService {
	s.page = &page
	return s
}
func (s *UniversalTransferHistoryService) Limit(limit int32) *UniversalTransferHistoryService {
	s.limit = &limit
	return s
}
func (s *UniversalTransferHistoryService) ShowAllStatus(showAllStatus bool) *UniversalTransferHistoryService {
	s.showAllStatus = &showAllStatus
	return s
}

func (s *UniversalTransferHistoryService) Do(ctx context.Context, opts ...RequestOption) (res []*UniversalTransfer, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/broker/universalTransfer",
		secType:  secTypeSigned,
	}
	if s.fromId != nil {
		r.setParam("fromId", *s.fromId)
	}
	if s.toId != nil {
		r.setParam("toId", *s.toId)
	}
	if s.clientTranId != nil {
		r.setParam("clientTranId", *s.clientTranId)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.showAllStatus != nil {
		r.setParam("showAllStatus", *s.showAllStatus)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*UniversalTransfer{}, err
	}
	res = make([]*UniversalTransfer, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*UniversalTransfer{}, err
	}
	return res, nil
}
