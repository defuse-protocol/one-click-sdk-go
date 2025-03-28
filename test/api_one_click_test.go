package oneclick_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	oneclick "github.com/defuse-protocol/one-click-sdk-go"
)

func Test_OneClickAPIService(t *testing.T) {
	configuration := oneclick.NewConfiguration()
	apiClient := oneclick.NewAPIClient(configuration)

	tokens, _, err := apiClient.OneClickAPI.GetTokens(context.Background()).Execute()
	require.Nil(t, err)
	require.NotNil(t, tokens)

	t.Logf("Tokens: %+v", tokens)

	quote, _, err := apiClient.OneClickAPI.GetQuote(context.Background()).Execute()
	require.Nil(t, err)
	require.NotNil(t, quote)

	submitResp, _, err := apiClient.OneClickAPI.SubmitDepositTx(context.Background()).Execute()
	require.Nil(t, err)
	require.NotNil(t, submitResp)

	status, _, err := apiClient.OneClickAPI.GetExecutionStatus(context.Background()).DepositAddress("").Execute()
	require.Nil(t, err)
	require.NotNil(t, status)
}
