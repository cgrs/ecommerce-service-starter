package server

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/cgrs/ecommerce-service-starter/cart"
	"github.com/cgrs/ecommerce-service-starter/customers"
	"github.com/cgrs/ecommerce-service-starter/items"
	"github.com/cgrs/ecommerce-service-starter/orders"
	"github.com/gin-gonic/gin"
)

type contextKey string

func RootHandler(rw http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(rw)
	rw.WriteHeader(http.StatusOK)
	enc.Encode(map[string]interface{}{"status": 200, "message": "OK"})
}

func GetStats(ctx *gin.Context) {
	customersCount := customers.Count()
	ordersCount := customers.OrdersCount()
	itemsCount := items.Count()
	cartsCount := cart.Count()
	ctx.JSON(http.StatusOK, gin.H{"stats": gin.H{"customers": customersCount, "orders": ordersCount, "items": itemsCount, "carts": cartsCount}})
}

func HandleLogin(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie("session")
	if err != nil && !errors.Is(err, http.ErrNoCookie) {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if cookie != nil && cart.Find(cookie.Value) != nil {
		return
	}
	var body struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := customers.Login(body.Username, body.Password); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	sessionID := fmt.Sprintf("%x", md5.Sum([]byte(body.Username+strconv.FormatInt(time.Now().UTC().UnixNano(), 10))))

	if err := cart.Create(sessionID, body.Username); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.SetCookie("session", sessionID, 0, "", "", false, true)
}

func HandleLogout(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie("session")
	if err != nil && !errors.Is(err, http.ErrNoCookie) {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if cookie == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := cart.Delete(cookie.Value); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	cookie.MaxAge = -1
	ctx.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
}

func GetOrders(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie("session")
	if err != nil && errors.Is(err, http.ErrNoCookie) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	customer := customers.FindBySession(cookie.Value)
	if customer == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "customer not found"})
	}

	ctx.JSON(http.StatusOK, orders.FindByCustomer(customer.Username))
}

func HandleRegister(ctx *gin.Context) {
	var body struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Admin    bool   `json:"admin"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := customers.Register(body.Username, body.Password, body.Email, body.Admin); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.Status(http.StatusCreated)
}

func GetCart(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie("session")
	if err != nil && errors.Is(err, http.ErrNoCookie) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c := cart.Find(cookie.Value)
	if c == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "cart not found"})
	}
	t := c.GetTotal()
	var result struct {
		*cart.Cart
		Total float64 `json:"total"`
	}
	result.Cart = c
	result.Total = t
	ctx.JSON(http.StatusOK, result)
}

func UpdateCart(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie("session")
	if err != nil && errors.Is(err, http.ErrNoCookie) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var body struct {
		Lines []cart.Line `json:"lines" binding:"dive"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := cart.Update(cookie.Value, body.Lines); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.Status(http.StatusNoContent)
}

func HandleCheckout(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie("session")
	if err != nil && errors.Is(err, http.ErrNoCookie) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	order, err := customers.DoCheckout(cookie.Value)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, order)
}

func GetItems(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, items.List())
}

func AddItem(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie("session")
	if err != nil && errors.Is(err, http.ErrNoCookie) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	customer := customers.FindBySession(cookie.Value)
	if customer == nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "customer not found"})
	}
	if !customer.Admin {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "you do not have permission to add items"})
	}
	var body struct {
		Name        string  `json:"name" binding:"required"`
		Description string  `json:"description" binding:"required"`
		UnitPrice   float64 `json:"price" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	i := &items.Item{
		ID:          items.GetNextID(),
		Name:        body.Name,
		Description: body.Description,
		UnitPrice:   body.UnitPrice,
	}
	if err := items.Add(i); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusCreated, i)
}

func GetSingleItem(ctx *gin.Context) {
	itemId, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	item := items.Find(int(itemId))
	if item == nil {
		ctx.JSON(http.StatusNotFound, item)
	}
	ctx.JSON(http.StatusOK, item)
}

func UpdateItem(ctx *gin.Context) {
	itemId, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	item := items.Find(int(itemId))
	if item == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "item does not exist"})
	}
	cookie, err := ctx.Request.Cookie("session")
	if err != nil && errors.Is(err, http.ErrNoCookie) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	customer := customers.FindBySession(cookie.Value)
	if customer == nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "customer not found"})
	}
	if !customer.Admin {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "you do not have permission to update items"})
	}
	var body struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		UnitPrice   float64 `json:"price"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	i := &items.Item{
		ID:          int(itemId),
		Name:        body.Name,
		Description: body.Description,
		UnitPrice:   body.UnitPrice,
	}
	if err := items.Update(i); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.Status(http.StatusNoContent)
}

func RemoveItem(ctx *gin.Context) {
	itemId, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	item := items.Find(int(itemId))
	if item == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "item does not exist"})
	}
	cookie, err := ctx.Request.Cookie("session")
	if err != nil && errors.Is(err, http.ErrNoCookie) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	customer := customers.FindBySession(cookie.Value)
	if customer == nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "customer not found"})
	}
	if !customer.Admin {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "you do not have permission to delete items"})
	}
	if err := items.Delete(int(itemId)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.Status(http.StatusNoContent)
}

func Router(e *gin.Engine) error {
	e.GET("/stats", GetStats)
	api := e.Group("/api")
	api.POST("/login", HandleLogin)
	api.GET("/logout", HandleLogout)
	api.POST("/register", HandleRegister)
	api.GET("/orders", GetOrders)
	cart := api.Group("/cart")
	cart.GET("/", GetCart)
	cart.PUT("/", UpdateCart)
	cart.POST("/checkout", HandleCheckout)
	items := api.Group("/items")
	items.GET("/", GetItems)
	items.POST("/", AddItem)
	items.GET("/:id", GetSingleItem)
	items.PUT("/:id", UpdateItem)
	items.DELETE("/:id", RemoveItem)
	return nil
}
