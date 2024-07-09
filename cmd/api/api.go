package api

import (
	"encoding/json"
	"net/http"

	"github.com/MuhAndriansyah/go-redis-crud/internal/book"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	r := mux.NewRouter()
	v1 := r.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		info := map[string]string{"status": "ok"}

		res, _ := json.Marshal(info)

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(200)
		w.Write(res)
	})

	v1.HandleFunc("/books", book.ListBook)

	logrus.Info("starting up server...")

	return http.ListenAndServe(s.addr, r)
}
