package api

import (
	"net/http"
	database "simplebank/database/sqlc"
	"simplebank/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username 	string 	`json:"username" binding:"required,alphanum"`
	Password 	string 	`json:"password" binding:"required,min=5"`
	Fullname 	string 	`json:"fullname" binding:"required"`
	Email 		string 	`json:"email" binding:"required,email"`
}

type createUserResponse struct {
	Username			string		`json:"username"`
	Fullname			string		`json:"fullname"`
	Email				string		`json:"email"`
	PasswordChangedAt	time.Time	`json:"passwordChangedAt"`
	CreatedAt			time.Time	`json:"createdAt"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON( &req ); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON( http.StatusInternalServerError, errorResponse(err) )
		return
	}

	arg := database.CreateUserParams {
		Username: 			req.Username,
		HashedPassword: 	hashedPassword,
		Email: 				req.Email,
		Fullname: 			req.Fullname,
	}

	user, err := server.store.CreateUser( ctx, arg )
	if err != nil {
		if pqErr, ok := err.( *pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_vio;ation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
			
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := createUserResponse{
		Username: 			user.Username,
		Fullname: 			user.Fullname,
		Email:				user.Email,
		PasswordChangedAt:	time.Now(),
		CreatedAt:			user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}
