// Code generated by sqlc. DO NOT EDIT.
// source: account.sql

package db

import (
	"context"
	"testing"
	"time"

	"github.com/simplebank/util"
	"github.com/stretchr/testify/require"
)


func createRandomAccount(t *testing.T) Account{
	want := CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	got, err := testQueries.CreateAccount(context.Background(), want)
	require.NoError(t, err)
	require.NotEmpty(t, got)

	require.Equal(t, want.Owner, got.Owner)
	require.Equal(t, want.Balance, got.Balance)
	require.Equal(t, want.Currency, got.Currency)

	require.NotZero(t, got.ID)
	require.NotZero(t, got.CreatedAt)
	return got
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}


func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Balance,account2.Balance)
	require.Equal(t, account1.Currency,account2.Currency)
	require.Equal(t, account1.Owner,account2.Owner)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	arg := UpdateAccountParams {
		ID: account1.ID,
		Balance: util.RandomMoney(),
	}
	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Owner,account2.Owner)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	// require.Equal(t, ErrRecordNotFound.Error(), err.Error())
	require.Contains(t, err.Error(), ErrRecordNotFound.Error())
	require.Empty(t, account2)

}