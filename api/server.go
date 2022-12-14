package api

import (
	database "simplebank/database/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serve http request for our banking system
type Server struct {
	store database.Store
	router *gin.Engine /* send each API request to the correct handler for processing */
}

// NewServer create a new HTTP server and serup routing.
func NewServer( store database.Store ) *Server {
	server := &Server{ store: store }
	router := gin.Default()

	if v, ok := binding.Validator.Engine().( *validator.Validate ); ok {
		v.RegisterValidation( "currency", validCurrency )
	}

	/* start adding routes to router */

	/*------- user ------*/
	router.POST("/users", server.createUser)

	/*----- account -----*/
	router.POST( "/accounts", server.createAccount )
	router.GET( "accounts/:id", server.getAccount )
	router.GET( "accounts", server.listAccount )

	/*----- transfer -----*/
	router.POST( "/transfers", server.createTransfer )

	/* end adding routes to router */


	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run( address )
}

// handle error in general
func errorResponse(err error) gin.H {
	return gin.H{ "error": err.Error() }
}
