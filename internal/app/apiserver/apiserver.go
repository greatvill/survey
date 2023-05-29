package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/greatvill/survey.git/internal/app/store"
	"io"
	"log"
	"net/http"
)

type APIServer struct {
	config *Config
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	s.configureRouter()
	if err := s.configureStore(); err != nil {
		log.Fatal(err)
	}
	log.Print("starting api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, err := io.WriteString(writer, "Hello")
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}
