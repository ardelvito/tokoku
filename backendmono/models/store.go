package models
import "time"


type Store struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    UserID    uint      `json:"user_id"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}


type CreateStoreRequest struct {
    Name string `json:"name" binding:"required"`
}

type Product struct {
    StoreID      uint      `gorm:"column:store_id`  // Referensi ke Store
    ID           uint      `gorm:"primaryKey" json:"id"`                      // Primary key untuk tabel
    ProductName  string    `gorm:"size:255;not null" json:"product_name"`     // Nama produk
    Price        float64   `gorm:"not null" json:"price"`                     // Harga produk
    Description  string    `gorm:"type:text;not null" json:"description"`     // Deskripsi produk
    URLImage     string    `gorm:"type:text;not null" json:"url_image"`       // URL gambar produk
    CreatedAt    time.Time `json:"created_at"`                                // Waktu dibuat
    UpdatedAt    time.Time `json:"updated_at"`                                // Waktu diperbarui
}

type CreateProductRequest struct {
    StoreID     string  `json:"store_id"`
    ProductName string  `json:"product_name"`
    Price       float64 `json:"price"`
    Description string  `json:"description"`
    URLImage    string  `json:"url_image"`
}
