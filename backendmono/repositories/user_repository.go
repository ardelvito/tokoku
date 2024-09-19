package repositories

import (
    "gorm.io/gorm"
    "backendmono/models"
	"errors"
)

// Interface UserRepository
type UserRepository interface {
    // Tambahkan metode lain yang diperlukan
	CreateUser(user *models.User) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByID(id uint) (*models.User, error)
	UpdateUserProfile(email string, updatedUser models.User) error

}
// Struktur userRepository
type userRepository struct {
    DB *gorm.DB
}

// Fungsi untuk membuat instance userRepository
func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{DB: db}
}

// Implementasi metode CreateUser
func (repo *userRepository) CreateUser(user *models.User) (*models.User, error) {
    if err := repo.DB.Create(user).Error; err != nil {
        return nil, err // Kembalikan error jika terjadi kesalahan
    }
    return user, nil // Kembalikan user yang berhasil disimpan
}


// FindByEmail: mencari user berdasarkan email
func (repo *userRepository) FindByEmail(email string) (*models.User, error) {
    var user models.User

    // Query ke database untuk mencari user berdasarkan email
    if err := repo.DB.Where("email = ?", email).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("user not found")
        }
        return nil, err
    }

    return &user, nil
}

func (repo *userRepository) FindByID(id uint) (*models.User, error) {
    var user models.User
    if err := repo.DB.Preload("Store").First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}


//UpdateUserProfile: update user berdasarkan email
func (repo *userRepository) UpdateUserProfile(email string, updatedUser models.User) error {
    
    var user models.User
    err := repo.DB.Where("email = ?", email).First(&user).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return errors.New("user not found")
        }
        return err
    }

    // Update informasi user
    user.Name = updatedUser.Name
    user.Phone = updatedUser.Phone
    user.Address = updatedUser.Address

    // Simpan perubahan ke database
    if err := repo.DB.Save(&user).Error; err != nil {
        return err
    }

    return nil
}

