package db

import (
	"context"
	"testing"

	"github.com/jerosanchez/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)

	gotAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gotAccount)

	require.Equal(t, gotAccount.ID, account.ID)
	require.Equal(t, gotAccount.Owner, account.Owner)
	require.Equal(t, gotAccount.Balance, account.Balance)
	require.Equal(t, gotAccount.Currency, account.Currency)
	require.Equal(t, gotAccount.CreatedAt, account.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, updatedAccount.ID, account.ID)
	require.Equal(t, updatedAccount.Balance, arg.Balance)
	require.Equal(t, updatedAccount.Currency, account.Currency)
	require.Equal(t, updatedAccount.CreatedAt, account.CreatedAt)
}

// Helpers

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, account.Owner, arg.Owner)
	require.Equal(t, account.Balance, arg.Balance)
	require.Equal(t, account.Currency, arg.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}
