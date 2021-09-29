package server

import (
	"encoding/json"
	"github.com/cgrs/ecommerce-service-starter/items"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func RootHandler(rw http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(rw)
	rw.WriteHeader(http.StatusOK)
	enc.Encode(map[string]interface{}{"status": 200, "message": "OK"})
}

var Mux = mux.NewRouter()

func AddItem(rw http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(rw)
	dec := json.NewDecoder(r.Body)
	var body items.Item
	if err := dec.Decode(&body); err != nil && err != io.EOF {
		rw.WriteHeader(http.StatusBadRequest)
		enc.Encode(map[string]interface{}{"error": err, "status": 400})
		return
	}
	i, err := items.AddItem(&body)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		enc.Encode(map[string]interface{}{"error": err, "status": 500})
		return
	}
	rw.WriteHeader(http.StatusCreated)
	enc.Encode(i)
}

func ListItem(rw http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(rw)
	enc.Encode(items.ListItems())
}

func init() {
	Mux.HandleFunc("/items", AddItem).Methods(http.MethodPost)
	Mux.HandleFunc("/items", ListItem).Methods(http.MethodGet)
}