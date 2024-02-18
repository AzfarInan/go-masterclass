package api

import (
	db "github.com/AzfarInan/go-masterclass/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// / This server will serve HTTP requests for our banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// / NewServer creates a new HTTP server and set up routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	/// Add Routes to router: Accounts
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.PATCH("/accounts/:id", server.updateAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)

	/// Add Routes to router: Transfer
	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

// / Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// / Error Response
func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
