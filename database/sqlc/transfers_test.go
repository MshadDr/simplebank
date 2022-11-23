package database

import (
	"context"
	"simplebank/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer( t *testing.T, accountFrom, accountTo Account ) Transfer {
	args := CreateTransferParams{
		FromAccountID: accountFrom.ID,
		ToAccountID: accountTo.ID,
		Amount: utils.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer( context.Background(), args )
	require.NoError( t, err )
	require.NotEmpty( t, transfer )

	require.Equal( t, transfer.FromAccountID, accountFrom.ID )
	require.Equal( t, transfer.ToAccountID, accountTo.ID )
	require.Equal( t, transfer.Amount, args.Amount )

	require.NotZero( t, transfer.ID )
	require.NotZero( t, transfer.CreatedAt )

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	accountFrom := createRandomeAccount( t )
	accountTo := createRandomeAccount( t )
	createRandomTransfer( t, accountFrom, accountTo )
}

func TestGetTransfer(t *testing.T) {
	accountFrom := createRandomeAccount( t )
	accountTo := createRandomeAccount( t )
	transferFirst := createRandomTransfer( t, accountFrom, accountTo )
	
	transferSecond, err := testQueries.GetTransfer( context.Background(), transferFirst.ID )
	require.NoError( t, err )
	require.NotEmpty( t, transferSecond )

	require.Equal( t, transferFirst.ID, transferSecond.ID )
	require.Equal( t, transferFirst.Amount, transferSecond.Amount )
	require.Equal( t, transferFirst.FromAccountID, transferSecond.FromAccountID )
	require.Equal( t, transferFirst.ToAccountID, transferSecond.ToAccountID )

	require.WithinDuration( t, transferFirst.CreatedAt, transferSecond.CreatedAt, time.Second )
}

func TestListTransfer(t *testing.T) {
	for i := 0; i < 10; i++ {
		accountFrom := createRandomeAccount( t )
		accountTo := createRandomeAccount( t )
		createRandomTransfer( t, accountFrom, accountTo )
	}

	args := ListTransfersParams{
		Limit: 5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers( context.Background(), args )	
	require.NoError( t, err )
	require.Len( t, transfers, int(args.Limit) )

	for _, transfer := range transfers {
		require.NotEmpty( t, transfer )
	}
}
