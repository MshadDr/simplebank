package api

import (
	"bytes"
	// "database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	mockdb "simplebank/database/mock"
	database "simplebank/database/sqlc"
	"simplebank/utils"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	// "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestUserAPI(t *testing.T) {
	user, password := randomUser(t)

	testCases := []struct {
        name          string
        body          gin.H
        buildStubs    func(store *mockdb.MockStore)
        checkResponse func(recorder *httptest.ResponseRecorder)
    }{
        {
            name: "OK",
            body: gin.H{
                "username":  user.Username,
                "password":  password,
                "fullname": user.Fullname,
                "email":     user.Email,
            },
            buildStubs: func(store *mockdb.MockStore) {
				
                store.EXPECT().
                    CreateUser(gomock.Any(), gomock.Any()).
                    Times(1).
                    Return(user, nil)
            },
            checkResponse: func(recorder *httptest.ResponseRecorder) {
			
                require.Equal(t, http.StatusOK, recorder.Code)
                requireBodyMatchUser(t, recorder.Body, user)  /* recorder.Body is Actual and user is Expected */
            },
        },
		// {
		// 	name: "InternalError",
		// 	body: gin.H{
		// 		"username": user.Username,
		// 		"password": password,
		// 		"fullname": user.Fullname,
		// 		"email": user.Email,
		// 	},
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().
		// 		CreateUser(gomock.Any(), gomock.Any()).
		// 		Times(1).
		// 		Return(database.User{}, sql.ErrConnDone)
		// 	},
		// 	checkResponse: func(recorder *httptest.ResponseRecorder) {
		// 		require.Equal( t, http.StatusInternalServerError, recorder.Code )
		// 	},
		// },
		// {
		// 	name: "DuplicateUsername",
		// 	body: gin.H{
		// 		"username": user.Username,
		// 		"password": password,
		// 		"fullname": user.Fullname,
		// 		"email": user.Email,
		// 	},
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().
		// 		CreateUser(gomock.Any(), gomock.Any()).
		// 		Times(1).
		// 		Return(database.User{}, &pq.Error{Code: "23505"})
		// 	},
		// 	checkResponse: func(recorder *httptest.ResponseRecorder) {
		// 		require.Equal( t, http.StatusForbidden, recorder.Code )
		// 	},
		// },
		// {
		// 	name: "InvalidUsername",
		// 	body: gin.H{
		// 		"username": "invalid-username#1",
		// 		"password": password,
		// 		"fullname": user.Fullname,
		// 		"email": user.Email,
		// 	},
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().
		// 		CreateUser(gomock .Any(), gomock.Any()).
		// 		Times(0)		
		// 	},
		// 	checkResponse: func(recorder *httptest.ResponseRecorder) {
		// 		require.Equal( t, http.StatusBadRequest, recorder.Code )
		// 	},
		// },
		// {
		// 	name: "InvalidEmail",
		// 	body: gin.H{
		// 		"username": user.Username,
		// 		"password": password,
		// 		"fullname": user.Fullname,
		// 		"email": "invallid-email",
		// 	},
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().
		// 		CreateUser(gomock.Any(), gomock.Any()).
		// 		Times(0)
		// 	},
		// 	checkResponse: func(recorder *httptest.ResponseRecorder) {
		// 		require.Equal( t, http.StatusBadRequest, recorder.Code )
		// 	},
		// },
		// {
		// 	name: "ToShortPassword",
		// 	body: gin.H{
		// 		"username": user.Username,
		// 		"password": "pas",
		// 		"fullname": user.Fullname,
		// 		"email": user.Email,
		// 	},
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().
		// 		CreateUser(gomock.Any(), gomock.Any()).
		// 		Times(0)
		// 	},
		// 	checkResponse: func(recorder *httptest.ResponseRecorder) {
		// 		require.Equal( t, http.StatusBadRequest, recorder.Code )
		// 	},
		// },
	}

	for i := range testCases {
        tc := testCases[i]

        t.Run(tc.name, func(t *testing.T) {
            ctrl := gomock.NewController(t)
            defer ctrl.Finish()

            store := mockdb.NewMockStore(ctrl)
            tc.buildStubs(store)

            server := NewServer(store)
            recorder := httptest.NewRecorder()

            // Marshal body data to JSON
            data, err := json.Marshal(tc.body)
            require.NoError(t, err)

            url := "/users"
            request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
            require.NoError(t, err)

            server.router.ServeHTTP(recorder, request)
            tc.checkResponse(recorder)
		} )
	}
	
}

// check responses
func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user database.User) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotUser database.User
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)
	require.Equal(t, user.Username, gotUser.Username)
	require.Equal(t, user.Fullname, gotUser.Fullname)
	require.Equal(t, user.Email, gotUser.Email)
	require.Empty(t, gotUser.HashedPassword)
}

// create fake user and password
func randomUser(t *testing.T) ( database.User, string ) {

	password, err := utils.HashPassword( utils.RandomString( 6 ) )
	require.NoError(t, err)

	user := database.User{
		Username: utils.RandomOwner(),
		Fullname: utils.RandomOwner(),
		HashedPassword: password,
		Email: utils.RandomEmail(),
	}

	return user, password 
} 


