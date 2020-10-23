package server

import (
	"context"
	"gRPC/internal/pkg/config"
	"gRPC/internal/pkg/grpc"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	router  http.Handler
	config  *config.Config
	service InterfaceServer
}

func NewServer() *Server {
	return new(Server)
}

func (s *Server) SetConfig() {
	s.config = config.NewConfigFromEnv() //TODO env
}

func (s *Server) SetHandlers() {
	log.Println("Set Handlers")
	s.router = mux.NewRouter()
	s.SetSubRouter()
}

func (s *Server) SetSubRouter() {
	log.Println("Set Router")
	subR, ok := s.router.(*mux.Router)
	if !ok {
		log.Fatal("Expected mux subRouter.")
	}
	subR.HandleFunc("/login", s.Login).Methods(http.MethodPost)                // Login.
	subR.HandleFunc("/signup", s.SignUp).Methods(http.MethodPost)              // Sign Up.
	subR.HandleFunc("/logout", s.Logout).Methods(http.MethodPost)              // Logout.
	subR.HandleFunc("/ads", s.GetAll).Methods(http.MethodGet)                  // Get Ads.
	subR.HandleFunc("/ad", s.Create).Methods(http.MethodPost)                  // Create Ad.
	subR.HandleFunc("/ad/{id:[1-9]\\d*}", s.Get).Methods(http.MethodGet)       // Get Ad.
	subR.HandleFunc("/ad/{id:[1-9]\\d*}", s.Update).Methods(http.MethodPut)    // Update Ad.
	subR.HandleFunc("/ad/{id:[1-9]\\d*}", s.Delete).Methods(http.MethodDelete) // Delete Ad and update ID other Ad.
}

func (s *Server) SetStorage() {
	log.Println("New service")
	service := new(Service)
	service.SetStorage(grpc.NewGrpcClient())
	s.service = service
	log.Println("Set service")
}

func (s *Server) Run() {
	log.Println("Run")
	server := &http.Server{
		Addr:         ":8000", //s.config.Port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		Handler:      s.router,
	}

	//log.Println("HTTP server launching on port " + s.config.Port)
	//log.Fatal(http.ListenAndServe(s.config.Port, s.router))
	go func() {
		log.Fatal(http.ListenAndServe(":8000", s.router))
	}()
	log.Println("Start server on localhost:8000")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)
	defer close(c)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("HTTP server Shutdown: %v", err)
	}
	log.Printf("HTTP server is shutdown...")
}
