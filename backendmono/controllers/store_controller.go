package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "backendmono/models" // Pastikan ini mengarah ke tempat struct CreateStoreRequest berada
    "backendmono/services" // Pastikan ini mengarah ke service yang tepat
	"strings"
	"log"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"strconv"
)

type StoreController struct {
    StoreService *services.StoreService // Pastikan untuk menggunakan StoreService
}

func NewStoreController(storeService *services.StoreService) *StoreController {
    return &StoreController{StoreService: storeService}
}

// Handler untuk membuat toko baru
func (sc *StoreController) CreateStore(c *gin.Context) {
    var createStoreReq models.CreateStoreRequest

    // Bind JSON request ke struct
    if err := c.ShouldBindJSON(&createStoreReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
        return
    }

    parts := strings.Split(authHeader, "Bearer ")
    if len(parts) != 2 {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
        return
    }
    tokenString := parts[1]

    // Parse token untuk mendapatkan claims
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte("your_jwt_secret_key"), nil // Ganti dengan secret keymu
    })

    if err != nil || !token.Valid {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || claims["user_id"] == nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
        return
    }

    userID := uint(claims["user_id"].(float64)) // Misalnya, jika user_id disimpan sebagai float64

    log.Println("User ID:", userID)

    // Logika untuk membuat toko baru dengan user_id
    store := &models.Store{
        UserID: userID, // Mengaitkan user_id
        Name:   createStoreReq.Name,
    }

    store, err = sc.StoreService.CreateStore(store) // Pastikan method ini sesuai
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create store"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Store created successfully", "store": store})
}

func (sc *StoreController) FindProductsByStoreID(c *gin.Context) {
	vars := mux.Vars(c.Request) // Mengambil parameter dari request
    idStr := vars["id"]         // Dapatkan nilai dari parameter "id"
    storeID, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
        return
    }

    products, err := sc.StoreService.GetProductsByStoreID(uint(storeID))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "No products found for this store"})
        return
    }

    c.JSON(http.StatusOK, products)
}

// Handler untuk membuat produk baru
func (sc *StoreController) CreateProduct(c *gin.Context) {
    var createProductReq models.CreateProductRequest

    // Bind JSON request ke struct
    if err := c.ShouldBindJSON(&createProductReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

	vars := mux.Vars(c.Request) // Mengambil parameter dari request
    idStr := vars["id"]         // Dapatkan nilai dari parameter "id"
    storeID, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
        return
    }

    // Panggil service untuk membuat produk
    if err := sc.StoreService.CreateProduct(uint(storeID), &createProductReq); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product created successfully"})
}
