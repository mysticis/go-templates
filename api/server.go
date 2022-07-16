package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/mysticis/golang-templates-app/db/sqlc"
)

//Server serves http requests for the application
type Server struct {
	store  db.Store
	router *gin.Engine
}

//NewServer creates a new http server and setup routing

func NewServer(store db.Store) *Server {

	server := &Server{
		store: store,
	}

	router := gin.Default()

	router.POST("/addtask", server.createTask)
	router.GET("/task/:id", server.getTask)
	router.GET("/tasks", server.listTasks)

	server.router = router
	return server
}

//Start runs the http server on a specfic address

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
