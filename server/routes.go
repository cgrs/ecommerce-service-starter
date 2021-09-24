package server

import (
	"encoding/json"
	"github.com/cgrs/ecommerce-service-starter/items"
	"io"
	"net/http"
)

func RootHandler(rw http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(rw)
	rw.WriteHeader(http.StatusOK)
	enc.Encode(map[string]interface{}{"status": 200, "message": "OK"})
}

var Mux = http.NewServeMux()

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

func FindItem(rw http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(rw)
	i := items.FindItem(r.URL.Query().Get("name"))
	if i == nil {
		rw.WriteHeader(http.StatusNotFound)
		enc.Encode(map[string]interface{}{"error": "item not found", "status": 404})
		return
	}
	enc.Encode(i)
}

func init() {
	Mux.HandleFunc("/items", AddItem)
	Mux.HandleFunc("/items/get", FindItem)
}