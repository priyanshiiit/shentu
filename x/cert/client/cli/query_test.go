package cli

import (
	"context"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/client"
	sdksimapp "github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/certikfoundation/shentu/v2/simapp"
	"github.com/certikfoundation/shentu/v2/x/cert/keeper"
	"github.com/certikfoundation/shentu/v2/x/cert/types"
)

var (
	acc1 = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc2 = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc3 = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc4 = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
)

type IntegrationTestSuite struct {
	suite.Suite
	app     *simapp.SimApp
	ctx     sdk.Context
	keeper  keeper.Keeper
	address []sdk.AccAddress
}

func (suite *IntegrationTestSuite) SetupSuite() {
	suite.T().Log("Setting up integration test suite")
	suite.app = simapp.Setup(false)
	suite.ctx = suite.app.BaseApp.NewContext(false, tmproto.Header{})
	suite.keeper = suite.app.CertKeeper

	for _, acc := range []sdk.AccAddress{acc1, acc2, acc3, acc4} {
		err := sdksimapp.FundAccount(
			suite.app.BankKeeper,
			suite.ctx,
			acc,
			sdk.NewCoins(
				sdk.NewCoin("uctk", sdk.NewInt(10000000000)), // 1,000 CTK
			),
		)
		if err != nil {
			panic(err)
		}
	}

	suite.address = []sdk.AccAddress{acc1, acc2, acc3, acc4}
	suite.keeper.SetCertifier(suite.ctx, types.NewCertifier(suite.address[0], "", suite.address[0], ""))
}

func (suite *IntegrationTestSuite) TestCertifierQueryCmd() {

	type errArgs struct {
		contains   string
		shouldPass bool
	}

	testcases := []struct {
		name    string
		arg     string
		errArgs errArgs
	}{
		{
			"valid query",
			suite.address[0].String(),
			errArgs{
				shouldPass: true,
				contains:   "",
			},
		},

		{
			"invalid address",
			"pertik18w5txca7skklxe7y54nsxz22tqgtvrkzffrem8",
			errArgs{
				shouldPass: false,
				contains:   "decoding bech32 failed",
			},
		},
		{
			"empty address",
			"",
			errArgs{
				shouldPass: false,
				contains:   "empty address",
			},
		},
	}

	for _, tc := range testcases {
		tc := tc
		suite.Run(tc.name, func() {
			want := tc.arg
			cliClient, _ := client.NewClientFromNode("//127.0.0.1:26657")

			queryClient := types.NewQueryClient(client.Context{Client: cliClient})
			var req = types.QueryCertifierRequest{
				Alias:   viper.GetString(FlagAlias),
				Address: tc.arg,
			}
			res, err := queryClient.Certifier(
				context.Background(),
				&req,
			)
			if tc.errArgs.shouldPass {
				suite.Require().NoError(err)
				got := res.Certifier.Address
				suite.Require().Equal(got, want)
			} else {
				suite.Require().Error(err, tc.name)
				suite.Require().True(strings.Contains(err.Error(), tc.errArgs.contains))
			}

		})

	}
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
