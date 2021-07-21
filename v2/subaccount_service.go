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
	DeleteSubAccountApiKeyResponse struct{}
	ListSubAccountsService         struct {
		c            *Client
		subAccountId *string
		page         *int64
		size         *int64
	}
	SubAccount struct {
		SubAccountId          string `json:"subaccountId"`
		Email                 string `json:"email"`
		Tag                   string `json:"tag"`
		MakerCommission       int    `json:"makerCommission"`
		TakerCommission       int    `json:"takerCommission"`
		MarginMakerCommission int    `json:"marginMakerCommission"`
		MarginTakerCommission int    `json:"marginTakerCommission"`
		CreateTime            int64  `json:"createTime"`
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

func (s *DeleteSubAccountApiKeyService) Do(ctx context.Context, opts ...RequestOption) (res *CreateSubAccountApiKeyResponse, err error) {
	data, err := s.deleteSubAccountApiKey(ctx, "/sapi/v1/broker/subAccountApi", opts...)
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
	r.setParam("subAccountId", s.subAccountId)
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