package api

import (
	"fmt"

	db "github.com/AzfarInan/go-masterclass/simplebank/db/sqlc"
	"github.com/AzfarInan/go-masterclass/simplebank/db/util"
	"github.com/AzfarInan/go-masterclass/simplebank/token"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// / This server will serve HTTP requests for our banking service
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// / NewServer creates a new HTTP server and set up routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// / Add Routes to router: Accounts
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)
	authRoutes.PATCH("/accounts/:id", server.updateAccount)
	authRoutes.DELETE("/accounts/:id", server.deleteAccount)

	// / Add Routes to router: Transfer
	authRoutes.POST("/transfers", server.createTransfer)

	// / Add Routes to router: Users
	router.POST("/users", server.createUser)

	// / Add Routes to router: Login
	router.POST("/users/login", server.loginUser)

	server.router = router
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
