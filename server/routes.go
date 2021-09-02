package server

import (
	"context"
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
	"github.com/cgrs/ecommerce-service-starter/orders"
)

type contextKey string

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
	if r.Method != http.MethodPost {
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
	orders.FindByCustomer(customer.Username)
	enc.Encode(customer.Orders)
}

func HandleRegister(rw http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(rw)
	if r.Method != http.MethodPost {
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
	t := c.GetTotal()
	var result struct {
		*cart.Cart
		Total float64 `json:"total"`
	}
	result.Cart = c
	result.Total = t
	enc.Encode(result)
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
	case http.MethodGet:
		GetCart(rw, r)
	case http.MethodPut:
		UpdateCart(rw, r)
	default:
		notAllowed(rw, r)
	}
}

func HandleCheckout(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
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

	order, err := customers.DoCheckout(cookie.Value)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		enc.Encode(err)
	}
	enc.Encode(order)
}

func GetItems(rw http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(rw)
	itms := items.List()
	enc.Encode(itms)
}

func AddItem(rw http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(rw)
	cookie, err := r.Cookie("session")
	if err != nil && err == http.ErrNoCookie {
		rw.WriteHeader(http.StatusBadRequest)
		enc.Encode("missing cookie")
		return
	}
	customer := customers.FindBySession(cookie.Value)
	if customer == nil {
		rw.WriteHeader(http.StatusInternalServerError)
		enc.Encode("customer not found")
		return
	}
	if !customer.Admin {
		rw.WriteHeader(http.StatusForbidden)
		enc.Encode("you do not have permission to add items")
		return
	}
	dec := json.NewDecoder(r.Body)
	var body struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		UnitPrice   float64 `json:"price"`
	}
	if err := dec.Decode(&body); err != nil && err != io.EOF {
		rw.WriteHeader(http.StatusInternalServerError)
		enc.Encode(err)
		return
	}
	i := &items.Item{
		ID:          items.GetNextID(),
		Name:        body.Name,
		Description: body.Description,
		UnitPrice:   body.UnitPrice,
	}
	if err := items.Add(i); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		enc.Encode(err)
		return
	}
	rw.WriteHeader(http.StatusCreated)
	enc.Encode(i)
}

func HandleItems(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetItems(rw, r)
	case http.MethodPost:
		AddItem(rw, r)
	default:
		notAllowed(rw, r)
	}
}

func GetSingleItem(rw http.ResponseWriter, r *http.Request) {
	itemId := r.Context().Value(contextKey("path")).(int)
	item := items.Find(itemId)
	if item == nil {
		rw.WriteHeader(http.StatusNotFound)
	}
	json.NewEncoder(rw).Encode(item)
}

func UpdateItem(rw http.ResponseWriter, r *http.Request) {
	itemId := r.Context().Value(contextKey("path")).(int)
	enc := json.NewEncoder(rw)
	item := items.Find(itemId)
	if item == nil {
		rw.WriteHeader(http.StatusNotFound)
		enc.Encode("item does not exist")
		return
	}
	cookie, err := r.Cookie("session")
	if err != nil && err == http.ErrNoCookie {
		rw.WriteHeader(http.StatusBadRequest)
		enc.Encode("missing cookie")
		return
	}
	customer := customers.FindBySession(cookie.Value)
	if customer == nil {
		rw.WriteHeader(http.StatusInternalServerError)
		enc.Encode("customer not found")
		return
	}
	if !customer.Admin {
		rw.WriteHeader(http.StatusForbidden)
		enc.Encode("you do not have permission to update items")
		return
	}
	dec := json.NewDecoder(r.Body)
	var body struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		UnitPrice   float64 `json:"price"`
	}
	if err := dec.Decode(&body); err != nil && err != io.EOF {
		rw.WriteHeader(http.StatusInternalServerError)
		enc.Encode(err)
		return
	}
	i := &items.Item{
		ID:          itemId,
		Name:        body.Name,
		Description: body.Description,
		UnitPrice:   body.UnitPrice,
	}
	if err := items.Update(i); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		enc.Encode(err)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}

func RemoveItem(rw http.ResponseWriter, r *http.Request) {
	itemId := r.Context().Value(contextKey("path")).(int)
	enc := json.NewEncoder(rw)
	item := items.Find(itemId)
	if item == nil {
		rw.WriteHeader(http.StatusNotFound)
		enc.Encode("item does not exist")
		return
	}
	cookie, err := r.Cookie("session")
	if err != nil && err == http.ErrNoCookie {
		rw.WriteHeader(http.StatusBadRequest)
		enc.Encode("missing cookie")
		return
	}
	customer := customers.FindBySession(cookie.Value)
	if customer == nil {
		rw.WriteHeader(http.StatusInternalServerError)
		enc.Encode("customer not found")
		return
	}
	if !customer.Admin {
		rw.WriteHeader(http.StatusForbidden)
		enc.Encode("you do not have permission to delete items")
		return
	}
	if err := items.Delete(itemId); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		enc.Encode(err)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}

func HandleSingleItems(rw http.ResponseWriter, r *http.Request) {
	prefix := "/api/items/"
	pathString := r.URL.Path[len(prefix):]
	path, err := strconv.ParseInt(pathString, 10, 0)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(`"invalid item id"`))
		return
	}
	r = r.WithContext(context.WithValue(r.Context(), contextKey("path"), int(path)))
	switch r.Method {
	case http.MethodGet:
		GetSingleItem(rw, r)
	case http.MethodPut:
		UpdateItem(rw, r)
	case http.MethodDelete:
		RemoveItem(rw, r)
	default:
		notAllowed(rw, r)
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
	MainMux.HandleFunc("/api/items", HandleItems)
	MainMux.HandleFunc("/api/items/", HandleSingleItems)
}
