package database

import (
	"context"
	"database/sql"
	"simplebank/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomeAccount( t *testing.T) Account {

	arg := CreateAccountParams{
		Owner: utils.RandomOwner(),
		Balance: utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount( context.Background(), arg )

	require.NoError( t, err )
	require.NotEmpty( t, account )

	require.Equal( t, arg.Owner, account.Owner )
	require.Equal( t, arg.Balance, account.Balance )
	require.Equal( t, arg.Currency, account.Currency )
	
	require.NotZero( t, account.ID )
	require.NotZero( t, account.CreatedAt )

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomeAccount( t )
}

func TestGetAccount(t *testing.T) {
	accountFirst := createRandomeAccount( t )
	accountSecond, err := testQueries.GetAccount( context.Background(), accountFirst.ID )

	require.NoError( t, err )
	require.NotEmpty( t, accountSecond )

	require.Equal( t, accountFirst.ID, accountSecond.ID )
	require.Equal( t, accountFirst.Owner, accountSecond.Owner )
	require.Equal( t, accountFirst.Balance, accountSecond.Balance )
	require.Equal( t, accountFirst.Currency, accountSecond.Currency )
	
	require.WithinDuration( t, accountFirst.CreatedAt, accountSecond.CreatedAt, time.Second )
}

func TestUpdateAccount(t *testing.T) {
	accountFirst := createRandomeAccount( t )

	args := UpdateAccountParams {
		ID: accountFirst.ID,
		Balance: utils.RandomMoney(),
	}

	accountSecond, err := testQueries.UpdateAccount( context.Background(), args )
	
	require.NoError( t, err )
	require.NotEmpty( t, accountSecond )

	require.Equal( t, accountFirst.ID, accountSecond.ID )
	require.Equal( t, args.Balance, accountSecond.Balance )
	require.Equal( t, accountFirst.Currency, accountSecond.Currency )
	require.Equal( t, accountFirst.Owner, accountSecond.Owner )

	require.WithinDuration( t, accountFirst.CreatedAt, accountSecond.CreatedAt, time.Second)
}

func TestDeleteAccount( t *testing.T ) {
	accountFirst := createRandomeAccount( t )
	err := testQueries.DeleteAccount( context.Background(), accountFirst.ID )
	require.NoError( t, err )

	accountSecond, err := testQueries.GetAccount( context.Background(), accountFirst.ID )
	require.Error( t, err )
	require.EqualError( t, err, sql.ErrNoRows.Error() )
	require.Empty( t, accountSecond )
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomeAccount( t )
	}

	args := ListAccountsParams {
		Limit: 5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts( context.Background(), args)
	require.NoError( t, err )
	require.Len( t, accounts, 5 )

	for _, account := range accounts {
		require.NotEmpty( t, account )
	}
}
