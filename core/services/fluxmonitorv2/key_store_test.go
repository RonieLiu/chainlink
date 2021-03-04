package fluxmonitorv2_test

import (
	"testing"

	"github.com/smartcontractkit/chainlink/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/core/services/fluxmonitorv2"
	"github.com/stretchr/testify/require"
)

func TestKeyStore_Accounts(t *testing.T) {
	t.Parallel()

	s, cleanup := cltest.NewStore(t)
	t.Cleanup(cleanup)

	ks := fluxmonitorv2.NewKeyStore(s)

	s.KeyStore.Unlock(cltest.Password)
	account, err := s.KeyStore.NewAccount()
	require.NoError(t, err)

	accounts := ks.Accounts()
	require.Len(t, accounts, 1)
	require.Equal(t, account, accounts[0])
}

func TestKeyStore_GetRoundRobinAddress(t *testing.T) {
	t.Parallel()

	t.Log("TODO")
}
