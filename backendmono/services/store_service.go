package services

import (
    "backendmono/repositories" // Import package userRepositories
    // // "gorm.io/gorm" // Pastikan ini digunakan jika perlu
    "backendmono/models" // Import models jika diperlukan
	// "errors"
	// "time"
	// "github.com/golang-jwt/jwt/v4"
	// "golang.org/x/crypto/bcrypt"

)

type StoreService struct {
    // Dependency injections, e.g., repositories
	Repo repositories.StoreRepository
}
// Fungsi untuk membuat instance StoreService
func NewStoreService(repo repositories.StoreRepository) *StoreService {
    return &StoreService{Repo: repo}
}

// Method untuk membuat toko baru
// StoreService
func (ss *StoreService) CreateStore(store *models.Store) (*models.Store, error) {
    return ss.Repo.CreateStore(store) // Pastikan repo memiliki method Create yang benar
}

func (ss *StoreService) GetProductsByStoreID(storeID uint) ([]models.Product, error) {
    return ss.Repo.FindProductsByStoreID(storeID)
}

// Implementasi CreateProduct
func (ss *StoreService) CreateProduct(storeId uint, product *models.CreateProductRequest) error {
    newProduct := models.Product{
        StoreID:     storeId, // Pastikan ini uint
        ProductName: product.ProductName,
        Price:       product.Price,
        Description: product.Description,
        URLImage:    product.URLImage,
    }
    return ss.Repo.CreateProduct(storeId, &newProduct)
}


