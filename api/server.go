package api

import (
	database "simplebank/database/sqlc"
	"github.com/gin-gonic/gin"
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

	/* start adding routes to router */
	router.POST( "/accounts", server.createAccount )
	router.GET( "accounts/:id", server.getAccount )
	router.GET( "accounts", server.listAccount )
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
