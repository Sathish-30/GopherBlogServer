package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sathish-30/GopherBlogServer/internal/database"
	"github.com/sathish-30/GopherBlogServer/internal/initializers"
)

// The init function is used to load env variables and to connect to database
func init() {
	initializers.LoadEnv()
	database.ConnectToDb()
	log.Println("Connected to database")
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
	log.Println("Server is running on the port -> ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, nil)

}

func main() {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	server := createServer(port)
	server.run()
}
