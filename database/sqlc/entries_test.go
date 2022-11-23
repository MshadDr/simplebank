package database

import (
	"context"
	"database/sql"
	"simplebank/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomEntry( t *testing.T, account Account ) Entry {
	args := CreateEntryParams{
		AccountID: account.ID,
		Amount: utils.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry( context.Background(), args )
	require.NoError( t, err )
	require.NotEmpty( t, entry )

	require.Equal( t, account.ID, entry.AccountID )
	require.Equal( t, args.Amount, entry.Amount )

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry( t *testing.T ) {

	account := createRandomeAccount( t )
	createRandomEntry( t, account )
}

func TestGetEntry( t *testing.T ) {
	account := createRandomeAccount( t )
	entryFirst := createRandomEntry( t, account ) 
	entrySecong, err := testQueries.GetEntry( context.Background(), entryFirst.ID )

	require.NoError( t, err )
	require.NotEmpty( t,entrySecong )

	require.Equal( t, entryFirst.ID, entrySecong.ID)
	require.Equal( t, entryFirst.AccountID, entrySecong.AccountID)
	require.Equal( t, entryFirst.Amount, entrySecong.Amount)
	
	require.WithinDuration( t, entryFirst.CreatedAt, entrySecong.CreatedAt, time.Second )
}

func TestDeleteEntry(t *testing.T) {
	account := createRandomeAccount( t )
	entryFirst := createRandomEntry( t, account )
	err := testQueries.DeleteEntry( context.Background(), entryFirst.ID )
	require.NoError( t, err )

	entrySecond, err := testQueries.GetEntry( context.Background(), entryFirst.ID )
	require.Error( t, err )
	require.EqualError( t, err, sql.ErrNoRows.Error() )
	require.Empty( t, entrySecond )
}

func TestListEntries( t *testing.T ) {
	for i := 0; i < 10; i++ {
		account := createRandomeAccount( t )
		createRandomEntry( t, account )
	}

	args := ListEntriesParams{
		Limit: 5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries( context.Background(), args )
	require.NoError( t, err )
	require.Len( t, entries, int( args.Limit ) )

	for _, entry := range entries {
		require.NotEmpty( t, entry )
	}
}

func TestUpdateEntry( t *testing.T ) {
	account := createRandomeAccount( t )
	entryFirst := createRandomEntry( t, account )

	args := UpdateEntryParams{
		ID: entryFirst.ID,
		AccountID: account.ID,
		Amount: utils.RandomMoney(),
	}

	entrySecond, err := testQueries.UpdateEntry( context.Background(), args )

	require.NoError( t, err )
	require.NotEmpty( t, entrySecond )

	require.Equal( t, entrySecond.AccountID, args.AccountID)
	require.Equal( t, entrySecond.Amount, args.Amount)
	require.Equal( t, entrySecond.ID, entryFirst.ID)

	require.WithinDuration( t, entrySecond.CreatedAt, entryFirst.CreatedAt, time.Second )
}
