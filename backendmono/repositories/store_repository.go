package repositories

import (
    "backendmono/models"
    "gorm.io/gorm"
)

// Interface StoreRepository
type StoreRepository interface {
    // Tambahkan metode lain yang diperlukan
	CreateStore(user *models.Store) (*models.Store, error)
	FindStoreByID(id uint) (*models.Store, error)
	FindProductsByStoreID(storeID uint) ([]models.Product, error)
	CreateProduct(storeId uint, product *models.Product) error // Gunakan uint
}

type storeRepository struct {
    DB *gorm.DB
}

// NewStoreRepository membuat instance baru dari StoreRepository
func NewStoreRepository(db *gorm.DB) StoreRepository {
    return &storeRepository{DB: db}
}

// Create menyimpan toko baru ke dalam database
func (repo *storeRepository) CreateStore(store *models.Store) (*models.Store, error) {
	if err := repo.DB.Create(store).Error; err != nil {
        return nil, err // Kembalikan error jika terjadi kesalahan
    }
	// Update store_status untuk user berdasarkan user_id
    if err := repo.DB.Exec("UPDATE users SET store_status = ? WHERE id = ?", "true", store.UserID).Error; err != nil {
        return nil, err // Kembalikan error jika update gagal
    }
    return store, nil // Kembalikan store yang berhasil disimpan
}

// FindByID mengambil toko berdasarkan ID
func (repo *storeRepository) FindStoreByID(id uint) (*models.Store, error) {
    var store models.Store
    if err := repo.DB.First(&store, id).Error; err != nil {
        return nil, err
    }
    return &store, nil
}

func (repo *storeRepository) FindProductsByStoreID(storeID uint) ([]models.Product, error) {
    var products []models.Product
    if err := repo.DB.Where("store_id = ?", storeID).Find(&products).Error; err != nil {
        return nil, err
    }
    return products, nil
}

func (repo *storeRepository) CreateProduct(storeId uint, product *models.Product) error {
    product.StoreID = storeId // Pastikan ini uint
    return repo.DB.Create(product).Error
}