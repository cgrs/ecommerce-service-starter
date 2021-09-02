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
	customersCount := customers.Count()
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
	cookie, err := r.Cookie("session")
	if err != nil && err != http.ErrNoCookie {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	if cookie != nil && cart.Find(cookie.Value) != nil {
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

	sessionID := fmt.Sprintf("%x", md5.Sum([]byte(body.Username+strconv.FormatInt(time.Now().UTC().UnixNano(), 10))))

	cart.Create(sessionID, body.Username)

	http.SetCookie(rw, &http.Cookie{
		Name:     "session",
		Value:    sessionID,
		HttpOnly: true,
	})
}

func HandleLogout(rw http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(rw)
	cookie, err := r.Cookie("session")
	if err != nil && err != http.ErrNoCookie {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	if cookie == nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := cart.Delete(cookie.Value); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		enc.Encode(map[string]interface{}{"status": http.StatusInternalServerError, "message": err.Error()})
	}
	cookie.MaxAge = -1
	http.SetCookie(rw, cookie)
}

func GetOrders(rw http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(rw)
	cookie, err := r.Cookie("session")
	if err != nil && err == http.ErrNoCookie {
		rw.WriteHeader(http.StatusBadRequest)
		enc.Encode("missing cookie")
		return
	}
	customer := customers.FindBySession(cookie.Value)
	if customer == nil {
		rw.WriteHeader(http.StatusNotFound)
		enc.Encode("customer not found")
		return
	}
	enc.Encode(customer.Orders)
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

func notAllowed(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusMethodNotAllowed)
	enc := json.NewEncoder(rw)
	enc.Encode(map[string]interface{}{"error": 405, "message": "method not allowed"})
}

func GetCart(rw http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(rw)
	cookie, err := r.Cookie("session")
	if err != nil && err == http.ErrNoCookie {
		rw.WriteHeader(http.StatusBadRequest)
		enc.Encode("missing cookie")
		return
	}
	c := cart.Find(cookie.Value)
	if c == nil {
		rw.WriteHeader(http.StatusNotFound)
		enc.Encode("cart not found")
		return
	}
	enc.Encode(c)
}

func UpdateCart(rw http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(rw)
	cookie, err := r.Cookie("session")
	if err != nil && err == http.ErrNoCookie {
		rw.WriteHeader(http.StatusBadRequest)
		enc.Encode("missing cookie")
		return
	}
	dec := json.NewDecoder(r.Body)
	var body struct {
		Lines []cart.Line `json:"lines"`
	}
	if err := dec.Decode(&body); err != nil && err != io.EOF {
		rw.WriteHeader(http.StatusInternalServerError)
		enc.Encode(err)
		return
	}
	if err := cart.Update(cookie.Value, body.Lines); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		enc.Encode(err)
		return
	}
}

func HandleCart(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetCart(rw, r)
	case "PUT":
		UpdateCart(rw, r)
	default:
		notAllowed(rw, r)
	}
}

func HandleCheckout(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		notAllowed(rw, r)
		return
	}
	enc := json.NewEncoder(rw)
	cookie, err := r.Cookie("session")
	if err != nil && err == http.ErrNoCookie {
		rw.WriteHeader(http.StatusBadRequest)
		enc.Encode("missing cookie")
		return
	}

	if err := customers.DoCheckout(cookie.Value); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		enc.Encode(err)
	}

}

func GetItems(rw http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(rw)
	itms := items.List()
	enc.Encode(itms)
}

func AddItem(rw http.ResponseWriter, r *http.Request) {

}

func HandleItems(rw http.ResponseWriter, r *http.Request) {
	prefix := "/api/items/"
	path := r.URL.Path[len(prefix):]
	if path == "" {
		switch r.Method {
		case "GET":
			GetItems(rw, r)
		case "POST":
			AddItem(rw, r)
		default:
			notAllowed(rw, r)
		}
	} else {
		switch r.Method {
		case "GET":
		case "PUT":
		case "DELETE":
			return
		default:
			notAllowed(rw, r)
		}
	}
}

func init() {
	MainMux.HandleFunc("/stats", GetStats)
	MainMux.HandleFunc("/api/login", HandleLogin)
	MainMux.HandleFunc("/api/logout", HandleLogout)
	MainMux.HandleFunc("/api/register", HandleRegister)
	MainMux.HandleFunc("/api/orders", GetOrders)
	MainMux.HandleFunc("/api/cart", HandleCart)
	MainMux.HandleFunc("/api/cart/checkout", HandleCheckout)
	MainMux.HandleFunc("/api/items/", HandleItems)
}
