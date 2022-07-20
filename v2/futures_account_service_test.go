package binance

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type futuresAccountServiceTestSuite struct {
	baseTestSuite
}

func TestFuturesAccountService(t *testing.T) {
	suite.Run(t, new(futuresAccountServiceTestSuite))
}

func (s *futuresAccountServiceTestSuite) TestSingleAssetFuturesAccountService() {
	data := []byte(`{   
    		"feeTier": 0,
    		"canTrade": true,
    		"canDeposit": true,
			"canWithdraw": true,
			"updateTime": 0, 
			"totalInitialMargin": "0.00000000",
			"totalMaintMargin": "0.00000000",
			"totalWalletBalance": "23.72469206",
			"totalUnrealizedProfit": "0.00000000",
			"totalMarginBalance": "23.72469206",
			"totalPositionInitialMargin": "0.00000000",
			"totalOpenOrderInitialMargin": "0.00000000",
			"totalCrossWalletBalance": "23.72469206",
			"totalCrossUnPnl": "0.00000000",
			"availableBalance": "23.72469206",
			"maxWithdrawAmount": "23.72469206",
			"assets": [
				{
					"asset": "USDT",
					"walletBalance": "23.72469206",
					"unrealizedProfit": "0.00000000",
					"marginBalance": "23.72469206",
					"maintMargin": "0.00000000",
					"initialMargin": "0.00000000", 
					"positionInitialMargin": "0.00000000",
					"openOrderInitialMargin": "0.00000000",
					"crossWalletBalance": "23.72469206",
					"crossUnPnl": "0.00000000",
					"availableBalance": "23.72469206",
					"maxWithdrawAmount": "23.72469206",
					"marginAvailable": true,
					"updateTime": 1625474304765 
				},
				{
					"asset": "BUSD",
					"walletBalance": "103.12345678",
					"unrealizedProfit": "0.00000000",
					"marginBalance": "103.12345678",
					"maintMargin": "0.00000000",
					"initialMargin": "0.00000000", 
					"positionInitialMargin": "0.00000000",
					"openOrderInitialMargin": "0.00000000",
					"crossWalletBalance": "103.12345678",
					"crossUnPnl": "0.00000000",
					"availableBalance": "103.12345678",
					"maxWithdrawAmount": "103.12345678",
					"marginAvailable": true,
					"updateTime": 1625474304765
				}
			],
			"positions": [
				{
					"symbol": "BTCUSDT",
					"initialMargin": "0", 
					"maintMargin": "0",
					"unrealizedProfit": "0.00000000",
					"positionInitialMargin": "0",
					"openOrderInitialMargin": "0",
					"leverage": "100",
					"isolated": true,
					"entryPrice": "0.00000",
					"maxNotional": "250000",
					"bidNotional": "0",
					"askNotional": "0",
					"positionSide": "BOTH",
					"positionAmt": "0",
					"updateTime": 0
				}
			]
  }`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewGetFuturesAccountService().Do(newContext())
	s.r().NoError(err)
	e := &FuturesAccount{
		FeeTier:                     0,
		CanTrade:                    true,
		CanDeposit:                  true,
		CanWithdraw:                 true,
		UpdateTime:                  0,
		TotalInitialMargin:          "0.00000000",
		TotalMaintMargin:            "0.00000000",
		TotalWalletBalance:          "23.72469206",
		TotalUnrealizedProfit:       "0.00000000",
		TotalMarginBalance:          "23.72469206",
		TotalPositionInitialMargin:  "0.00000000",
		TotalOpenOrderInitialMargin: "0.00000000",
		TotalCrossWalletBalance:     "23.72469206",
		TotalCrossUnPnl:             "0.00000000",
		AvailableBalance:            "23.72469206",
		MaxWithdrawAmount:           "23.72469206",
		Assets: []*FutureAccountAsset{
			{
				Asset:                  "USDT",
				WalletBalance:          "23.72469206",
				UnrealizedProfit:       "0.00000000",
				MarginBalance:          "23.72469206",
				MaintMargin:            "0.00000000",
				InitialMargin:          "0.00000000",
				PositionInitialMargin:  "0.00000000",
				OpenOrderInitialMargin: "0.00000000",
				CrossWalletBalance:     "23.72469206",
				CrossUnPnl:             "0.00000000",
				AvailableBalance:       "23.72469206",
				MaxWithdrawAmount:      "23.72469206",
				MarginAvailable:        true,
				UpdateTime:             1625474304765,
			},
			{
				Asset:                  "BUSD",
				WalletBalance:          "103.12345678",
				UnrealizedProfit:       "0.00000000",
				MarginBalance:          "103.12345678",
				MaintMargin:            "0.00000000",
				InitialMargin:          "0.00000000",
				PositionInitialMargin:  "0.00000000",
				OpenOrderInitialMargin: "0.00000000",
				CrossWalletBalance:     "103.12345678",
				CrossUnPnl:             "0.00000000",
				AvailableBalance:       "103.12345678",
				MaxWithdrawAmount:      "103.12345678",
				MarginAvailable:        true,
				UpdateTime:             1625474304765,
			},
		},
		Positions: []*FutureAccountPosition{
			{
				Symbol:                 "BTCUSDT",
				InitialMargin:          "0",
				MaintMargin:            "0",
				UnrealizedProfit:       "0.00000000",
				PositionInitialMargin:  "0",
				OpenOrderInitialMargin: "0",
				Leverage:               "100",
				Isolated:               true,
				EntryPrice:             "0.00000",
				MaxNotional:            "250000",
				BidNotional:            "0",
				AskNotional:            "0",
				PositionSide:           "BOTH",
				PositionAmt:            "0",
				UpdateTime:             0,
			},
		},
	}
	s.assertFuturesAccountEqual(e, res)
}

func (s *futuresAccountServiceTestSuite) TestMultiAssetsFuturesAccountService() {
	data := []byte(`{   
			"feeTier": 0,
			"canTrade": true,
			"canDeposit": true,
			"canWithdraw": true,
			"updateTime": 0,
			"totalInitialMargin": "0.00000000",
			"totalMaintMargin": "0.00000000",
			"totalWalletBalance": "126.72469206",
			"totalUnrealizedProfit": "0.00000000",
			"totalMarginBalance": "126.72469206",
			"totalPositionInitialMargin": "0.00000000",
			"totalOpenOrderInitialMargin": "0.00000000",
			"totalCrossWalletBalance": "126.72469206",
			"totalCrossUnPnl": "0.00000000",
			"availableBalance": "126.72469206",
			"maxWithdrawAmount": "126.72469206",
			"assets": [
				{
					"asset": "USDT",
					"walletBalance": "23.72469206",
					"unrealizedProfit": "0.00000000",
					"marginBalance": "23.72469206",
					"maintMargin": "0.00000000",
					"initialMargin": "0.00000000",
					"positionInitialMargin": "0.00000000",
					"openOrderInitialMargin": "0.00000000",
					"crossWalletBalance": "23.72469206",
					"crossUnPnl": "0.00000000",
					"availableBalance": "23.72469206",
					"maxWithdrawAmount": "23.72469206",
					"marginAvailable": true,
					"updateTime": 1625474304765
				},
				{
					"asset": "BUSD",
					"walletBalance": "103.12345678",
					"unrealizedProfit": "0.00000000",
					"marginBalance": "103.12345678",
					"maintMargin": "0.00000000",
					"initialMargin": "0.00000000",
					"positionInitialMargin": "0.00000000",
					"openOrderInitialMargin": "0.00000000",
					"crossWalletBalance": "103.12345678",
					"crossUnPnl": "0.00000000",
					"availableBalance": "103.12345678",
					"maxWithdrawAmount": "103.12345678",
					"marginAvailable": true,
					"updateTime": 1625474304765
				}
			],
			"positions": [
				{
					"symbol": "BTCUSDT",
					"initialMargin": "0",
					"maintMargin": "0",
					"unrealizedProfit": "0.00000000",
					"positionInitialMargin": "0",
					"openOrderInitialMargin": "0",
					"leverage": "100",
					"isolated": true,
					"entryPrice": "0.00000",
					"maxNotional": "250000",
					"bidNotional": "0",
					"askNotional": "0",
					"positionSide": "BOTH",
					"positionAmt": "0",
					"updateTime": 0
				}
			]
  }`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewGetFuturesAccountService().Do(newContext())
	s.r().NoError(err)
	e := &FuturesAccount{
		FeeTier:                     0,
		CanTrade:                    true,
		CanDeposit:                  true,
		CanWithdraw:                 true,
		UpdateTime:                  0,
		TotalInitialMargin:          "0.00000000",
		TotalMaintMargin:            "0.00000000",
		TotalWalletBalance:          "126.72469206",
		TotalUnrealizedProfit:       "0.00000000",
		TotalMarginBalance:          "126.72469206",
		TotalPositionInitialMargin:  "0.00000000",
		TotalOpenOrderInitialMargin: "0.00000000",
		TotalCrossWalletBalance:     "126.72469206",
		TotalCrossUnPnl:             "0.00000000",
		AvailableBalance:            "126.72469206",
		MaxWithdrawAmount:           "126.72469206",
		Assets: []*FutureAccountAsset{
			{
				Asset:                  "USDT",
				WalletBalance:          "23.72469206",
				UnrealizedProfit:       "0.00000000",
				MarginBalance:          "23.72469206",
				MaintMargin:            "0.00000000",
				InitialMargin:          "0.00000000",
				PositionInitialMargin:  "0.00000000",
				OpenOrderInitialMargin: "0.00000000",
				CrossWalletBalance:     "23.72469206",
				CrossUnPnl:             "0.00000000",
				AvailableBalance:       "23.72469206",
				MaxWithdrawAmount:      "23.72469206",
				MarginAvailable:        true,
				UpdateTime:             1625474304765,
			},
			{
				Asset:                  "BUSD",
				WalletBalance:          "103.12345678",
				UnrealizedProfit:       "0.00000000",
				MarginBalance:          "103.12345678",
				MaintMargin:            "0.00000000",
				InitialMargin:          "0.00000000",
				PositionInitialMargin:  "0.00000000",
				OpenOrderInitialMargin: "0.00000000",
				CrossWalletBalance:     "103.12345678",
				CrossUnPnl:             "0.00000000",
				AvailableBalance:       "103.12345678",
				MaxWithdrawAmount:      "103.12345678",
				MarginAvailable:        true,
				UpdateTime:             1625474304765,
			},
		},
		Positions: []*FutureAccountPosition{
			{
				Symbol:                 "BTCUSDT",
				InitialMargin:          "0",
				MaintMargin:            "0",
				UnrealizedProfit:       "0.00000000",
				PositionInitialMargin:  "0",
				OpenOrderInitialMargin: "0",
				Leverage:               "100",
				Isolated:               true,
				EntryPrice:             "0.00000",
				MaxNotional:            "250000",
				BidNotional:            "0",
				AskNotional:            "0",
				PositionSide:           "BOTH",
				PositionAmt:            "0",
				UpdateTime:             0,
			},
		},
	}
	s.assertFuturesAccountEqual(e, res)
}

func (s *futuresAccountServiceTestSuite) assertFuturesAccountEqual(e, a *FuturesAccount) {
	r := s.r()
	r.Equal(e.FeeTier, a.FeeTier, "FeeTier")
	r.Equal(e.CanTrade, a.CanTrade, "CanTrade")
	r.Equal(e.CanDeposit, a.CanDeposit, "CanDeposit")
	r.Equal(e.CanWithdraw, a.CanWithdraw, "CanWithdraw")
	r.Equal(e.UpdateTime, a.UpdateTime, "UpdateTime")
	r.Equal(e.TotalInitialMargin, a.TotalInitialMargin, "TotalInitialMargin")
	r.Equal(e.TotalMaintMargin, a.TotalMaintMargin, "TotalMaintMargin")
	r.Equal(e.TotalWalletBalance, a.TotalWalletBalance, "TotalWalletBalance")
	r.Equal(e.TotalUnrealizedProfit, a.TotalUnrealizedProfit, "TotalUnrealizedProfit")
	r.Equal(e.TotalMarginBalance, a.TotalMarginBalance, "TotalMarginBalance")
	r.Equal(e.TotalPositionInitialMargin, a.TotalPositionInitialMargin, "TotalPositionInitialMargin")
	r.Equal(e.TotalOpenOrderInitialMargin, a.TotalOpenOrderInitialMargin, "TotalOpenOrderInitialMargin")
	r.Equal(e.TotalCrossWalletBalance, a.TotalCrossWalletBalance, "TotalCrossWalletBalance")
	r.Equal(e.TotalCrossUnPnl, a.TotalCrossUnPnl, "TotalCrossUnPnl")
	r.Equal(e.AvailableBalance, a.AvailableBalance, "AvailableBalance")
	r.Equal(e.MaxWithdrawAmount, a.MaxWithdrawAmount, "MaxWithdrawAmount")
	r.Len(a.Assets, len(e.Assets))
	for i := 0; i < len(a.Assets); i++ {
		r.Equal(e.Assets[i].Asset, a.Assets[i].Asset, "Asset")
		r.Equal(e.Assets[i].WalletBalance, a.Assets[i].WalletBalance, "WalletBalance")
		r.Equal(e.Assets[i].UnrealizedProfit, a.Assets[i].UnrealizedProfit, "UnrealizedProfit")
		r.Equal(e.Assets[i].MarginBalance, a.Assets[i].MarginBalance, "MarginBalance")
		r.Equal(e.Assets[i].MaintMargin, a.Assets[i].MaintMargin, "MaintMargin")
		r.Equal(e.Assets[i].InitialMargin, a.Assets[i].InitialMargin, "InitialMargin")
		r.Equal(e.Assets[i].PositionInitialMargin, a.Assets[i].PositionInitialMargin, "PositionInitialMargin")
		r.Equal(e.Assets[i].OpenOrderInitialMargin, a.Assets[i].OpenOrderInitialMargin, "OpenOrderInitialMargin")
		r.Equal(e.Assets[i].CrossWalletBalance, a.Assets[i].CrossWalletBalance, "CrossWalletBalance")
		r.Equal(e.Assets[i].CrossUnPnl, a.Assets[i].CrossUnPnl, "CrossUnPnl")
		r.Equal(e.Assets[i].AvailableBalance, a.Assets[i].AvailableBalance, "AvailableBalance")
		r.Equal(e.Assets[i].MaxWithdrawAmount, a.Assets[i].MaxWithdrawAmount, "MaxWithdrawAmount")
		r.Equal(e.Assets[i].MarginAvailable, a.Assets[i].MarginAvailable, "MarginAvailable")
		r.Equal(e.Assets[i].UpdateTime, a.Assets[i].UpdateTime, "UpdateTime")
	}
	r.Len(a.Positions, len(e.Positions))
	for i := 0; i < len(a.Positions); i++ {
		r.Equal(e.Positions[i].Symbol, a.Positions[i].Symbol, "Symbol")
		r.Equal(e.Positions[i].InitialMargin, a.Positions[i].InitialMargin, "InitialMargin")
		r.Equal(e.Positions[i].MaintMargin, a.Positions[i].MaintMargin, "MaintMargin")
		r.Equal(e.Positions[i].UnrealizedProfit, a.Positions[i].UnrealizedProfit, "UnrealizedProfit")
		r.Equal(e.Positions[i].PositionInitialMargin, a.Positions[i].PositionInitialMargin, "PositionInitialMargin")
		r.Equal(e.Positions[i].OpenOrderInitialMargin, a.Positions[i].OpenOrderInitialMargin, "OpenOrderInitialMargin")
		r.Equal(e.Positions[i].Leverage, a.Positions[i].Leverage, "Leverage")
		r.Equal(e.Positions[i].Isolated, a.Positions[i].Isolated, "Isolated")
		r.Equal(e.Positions[i].EntryPrice, a.Positions[i].EntryPrice, "EntryPrice")
		r.Equal(e.Positions[i].MaxNotional, a.Positions[i].MaxNotional, "MaxNotional")
		r.Equal(e.Positions[i].BidNotional, a.Positions[i].BidNotional, "BidNotional")
		r.Equal(e.Positions[i].AskNotional, a.Positions[i].AskNotional, "AskNotional")
		r.Equal(e.Positions[i].PositionSide, a.Positions[i].PositionSide, "PositionSide")
		r.Equal(e.Positions[i].PositionAmt, a.Positions[i].PositionAmt, "PositionAmt")
		r.Equal(e.Positions[i].UpdateTime, a.Positions[i].UpdateTime, "UpdateTime")
	}
}
