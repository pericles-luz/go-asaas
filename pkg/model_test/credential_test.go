package model_test

import (
	"testing"

	"github.com/pericles-luz/go-asaas/model"
	"github.com/stretchr/testify/require"
)

func TestCredentialShouldUnmarshal(t *testing.T) {
	credential := model.NewCredential()
	data := []byte(`{"access_token":"test_token","link":"http://example.com"}`)
	err := credential.Unmarshal(data)
	require.NoError(t, err, "Credential should unmarshal successfully")
	require.Equal(t, "test_token", credential.AccessToken, "Access token should match")
	require.Equal(t, "http://example.com", credential.Link, "Link should match")
}
