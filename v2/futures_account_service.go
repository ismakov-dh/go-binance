package binance

import (
	"context"
	"net/http"
)

// GetFuturesAccountService get futures account info
type (
	GetFuturesAccountService struct {
		c *Client
	}
	FuturesAccount struct {
		FeeTier                     int                      `json:"feeTier"`
		CanTrade                    bool                     `json:"canTrade"`
		CanDeposit                  bool                     `json:"canDeposit"`
		CanWithdraw                 bool                     `json:"canWithdraw"`
		UpdateTime                  int                      `json:"updateTime"`
		TotalInitialMargin          string                   `json:"totalInitialMargin"`
		TotalMaintMargin            string                   `json:"totalMaintMargin"`
		TotalWalletBalance          string                   `json:"totalWalletBalance"`
		TotalUnrealizedProfit       string                   `json:"totalUnrealizedProfit"`
		TotalMarginBalance          string                   `json:"totalMarginBalance"`
		TotalPositionInitialMargin  string                   `json:"totalPositionInitialMargin"`
		TotalOpenOrderInitialMargin string                   `json:"totalOpenOrderInitialMargin"`
		TotalCrossWalletBalance     string                   `json:"totalCrossWalletBalance"`
		TotalCrossUnPnl             string                   `json:"totalCrossUnPnl"`
		AvailableBalance            string                   `json:"availableBalance"`
		MaxWithdrawAmount           string                   `json:"maxWithdrawAmount"`
		Assets                      []*FutureAccountAsset    `json:"assets"`
		Positions                   []*FutureAccountPosition `json:"positions"`
	}
	FutureAccountAsset struct {
		Asset                  string `json:"asset"`
		WalletBalance          string `json:"walletBalance"`
		UnrealizedProfit       string `json:"unrealizedProfit"`
		MarginBalance          string `json:"marginBalance"`
		MaintMargin            string `json:"maintMargin"`
		InitialMargin          string `json:"initialMargin"`
		PositionInitialMargin  string `json:"positionInitialMargin"`
		OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
		CrossWalletBalance     string `json:"crossWalletBalance"`
		CrossUnPnl             string `json:"crossUnPnl"`
		AvailableBalance       string `json:"availableBalance"`
		MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
		MarginAvailable        bool   `json:"marginAvailable"`
		UpdateTime             int64  `json:"updateTime"`
	}
	FutureAccountPosition struct {
		Symbol                 string `json:"symbol"`
		InitialMargin          string `json:"initialMargin"`
		MaintMargin            string `json:"maintMargin"`
		UnrealizedProfit       string `json:"unrealizedProfit"`
		PositionInitialMargin  string `json:"positionInitialMargin"`
		OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
		Leverage               string `json:"leverage"`
		Isolated               bool   `json:"isolated"`
		EntryPrice             string `json:"entryPrice"`
		MaxNotional            string `json:"maxNotional"`
		BidNotional            string `json:"bidNotional"`
		AskNotional            string `json:"askNotional"`
		PositionSide           string `json:"positionSide"`
		PositionAmt            string `json:"positionAmt"`
		UpdateTime             int    `json:"updateTime"`
	}
)

// Do send request
func (s *GetFuturesAccountService) Do(ctx context.Context, opts ...RequestOption) (res *FuturesAccount, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v2/account",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(FuturesAccount)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
