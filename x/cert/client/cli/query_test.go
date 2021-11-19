package cli

import (
	"context"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/testutil/network"

	"github.com/certikfoundation/shentu/v2/x/cert/types"
)

type IntegrationTestSuite struct {
	suite.Suite
	cfg     network.Config
	network *network.Network
}

func (suite *IntegrationTestSuite) SetupSuite() {
	suite.T().Log("Setting up integration test suite")

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
			"abc",
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
