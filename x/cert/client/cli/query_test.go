package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/certikfoundation/shentu/v2/x/cert/types"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/testutil"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/cosmos/cosmos-sdk/x/bank/client/cli"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("Setting up integration test suite")
	s.cfg = network.DefaultConfig()
	s.network = network.New(s.T(), s.cfg)
}

func ExecTestCLICmd(clientCtx client.Context, cmd *cobra.Command, extraArgs []string) (testutil.BufferWriter, error) {
	cmd.SetArgs(extraArgs)
	fmt.Println(extraArgs)

	_, out := testutil.ApplyMockIO(cmd)
	clientCtx = clientCtx.WithOutput(out)

	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)
	if err := cmd.ExecuteContext(ctx); err != nil {
		fmt.Println("Oh NO", err)
		return out, err
	}
	return out, nil
}
func (s *IntegrationTestSuite) TestCertifierQueryCmd() {
	val := s.network.Validators[0]

	testcases := []struct {
		name      string
		args      []string
		expectErr bool
		// TODO type
		respType types.QueryCertifierResponse
		expected types.QueryCertifierResponse
	}{
		{
			name:      "valid query",
			args:      []string{"certik14rhps44azz6qr2sqqdx4m86rm7272zzhnzpr9g"},
			expectErr: false,
			respType:  types.QueryCertifierResponse{},
			expected: types.QueryCertifierResponse{
				Certifier: types.Certifier{
					Address:     "certik14rhps44azz6qr2sqqdx4m86rm7272zzhnzpr9g",
					Alias:       "",
					Proposer:    "",
					Description: "",
				},
			},
		},
	}

	for _, tc := range testcases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := GetCmdCertifier()
			clientCtx := val.ClientCtx
			// clientCtx, _ := client.NewClientFromNode("//127.0.0.1:26657")
			// out, err := ExecTestCLICmd(client.Context{Client: clientCtx}, cmd, tc.args)
			r, _ := ExecTestCLICmd(clientCtx, cli.GetCmdQueryTotalSupply(), []string{})
			fmt.Println("GetCmdQueryTotalSupply: ", r)
			out, err := ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				fmt.Println(out)
				s.Require().NoError(err)
				s.Require().NoError(json.Unmarshal(out.Bytes(), &tc.respType))
				s.Require().Equal(tc.expected, tc.respType)
			}
		})

	}
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
