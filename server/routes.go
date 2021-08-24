package server

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/cgrs/ecommerce-service-starter/cart"
	"github.com/cgrs/ecommerce-service-starter/customers"
	"github.com/cgrs/ecommerce-service-starter/items"
)

func RootHandler(rw http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(rw)
	rw.WriteHeader(http.StatusOK)
	enc.Encode(map[string]interface{}{"status": 200, "message": "OK"})
}

func GetStats(rw http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(rw)
	customersCount := customers.CustomersCount()
	ordersCount := customers.OrdersCount()
	itemsCount := items.Count()
	cartsCount := cart.Count()
	rw.WriteHeader(http.StatusOK)
	enc.Encode(map[string]interface{}{"stats": map[string]int{"customers": customersCount, "orders": ordersCount, "items": itemsCount, "carts": cartsCount}})
}

var MainMux = http.NewServeMux()

func HandleLogin(rw http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(rw)
	if r.Method != "POST" {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		enc.Encode(map[string]interface{}{
			"status": 405, "message": "method not allowed",
		})
		return
	}
	dec := json.NewDecoder(r.Body)
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := dec.Decode(&body); err != nil && err != io.EOF {
		rw.WriteHeader(http.StatusInternalServerError)
		enc.Encode(err)
		return
	}
	if err := customers.Login(body.Username, body.Password); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		enc.Encode(err)
		return
	}

	http.SetCookie(rw, &http.Cookie{
		Name:     "session",
		Value:    fmt.Sprintf("%x", md5.Sum([]byte(body.Username+strconv.FormatInt(time.Now().UTC().UnixNano(), 10)))),
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
	})
}

func HandleRegister(rw http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(rw)
	if r.Method != "POST" {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		enc.Encode(map[string]interface{}{
			"status": 405, "message": "method not allowed",
		})
		return
	}
	dec := json.NewDecoder(r.Body)
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Admin    bool   `json:"admin"`
	}
	if err := dec.Decode(&body); err != nil && err != io.EOF {
		rw.WriteHeader(http.StatusInternalServerError)
		enc.Encode(err)
		return
	}
	if body.Username == "" || body.Password == "" || body.Email == "" {
		rw.WriteHeader(http.StatusBadRequest)
		enc.Encode(map[string]interface{}{"error": "missing required params: username, password and/or email"})
		return
	}
	if err := customers.Register(body.Username, body.Password, body.Email, body.Admin); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		enc.Encode(err)
		return
	}
}

func init() {
	MainMux.HandleFunc("/stats", GetStats)
	MainMux.HandleFunc("/api/login", HandleLogin)
	MainMux.HandleFunc("/api/register", HandleRegister)
}
