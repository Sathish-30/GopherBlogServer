package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sathish-30/GopherBlogServer/internal/controller"
	"github.com/sathish-30/GopherBlogServer/internal/database"
	"github.com/sathish-30/GopherBlogServer/internal/initializers"
	"github.com/sathish-30/GopherBlogServer/internal/middleware"
)

// The init function is used to load env variables and to connect to database
func init() {
	initializers.LoadEnv()
	database.ConnectToDb()
	initializers.MigrateModel()
}

type Server struct {
	listenAddr string
}

func createServer(addr string) *Server {
	return &Server{
		listenAddr: addr,
	}
}

func (s *Server) run() {
	router := mux.NewRouter()
	router.HandleFunc("/register", middleware.MakeHandlerFunc(controller.Register)).Methods("POST")
	log.Println("Server is running on the port -> ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func main() {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	server := createServer(port)
	server.run()
}
