package services

import (
    "backendmono/repositories" // Import package userRepositories
    // "gorm.io/gorm" // Pastikan ini digunakan jika perlu
    "backendmono/models" // Import models jika diperlukan
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

)

// Definisikan struktur UserService
type UserService struct {
    Repo repositories.UserRepository
}

// Fungsi untuk membuat instance UserService
func NewUserService(repo repositories.UserRepository) *UserService {
    return &UserService{Repo: repo}
}

func (us *UserService) Register(req models.RegisterRequest) (*models.User, error) {
    // Hash password sebelum menyimpan
    hashedPassword, err := HashPassword(req.Password)
    if err != nil {
        return nil, err // Mengembalikan error jika hashing gagal
    }

    // Buat instance User dari request
    user := &models.User{
        Email:    req.Email,
        Password: hashedPassword,
        Name:     req.Name,
        Phone:    req.Phone,
        Address:  req.Address,
    }

    // Memanggil metode dari repository untuk menyimpan user
    return us.Repo.CreateUser(user)
}

func (us *UserService) Login(email, password string) (*models.User, error) {
    // Cari user berdasarkan email
    user, err := us.Repo.FindByEmail(email)
    if err != nil || !CheckPassword(user.Password, password) {
        return nil, errors.New("invalid credentials")
    }
    return user, nil
}

func (us *UserService) GetProfileByID(id uint) (*models.User, error) {
    // Memanggil repository untuk mencari user berdasarkan ID
    user, err := us.Repo.FindByID(id) // Tangkap kedua nilai
    if err != nil {
        return nil, err // Kembalikan error jika terjadi kesalahan
    }
    return user, nil // Kembalikan user jika berhasil
}

func (us *UserService) EditUserProfile(email string, updatedUser models.User) error {
    // Panggil repository untuk update data user
    err := us.Repo.UpdateUserProfile(email, updatedUser)
    if err != nil {
        return err
    }

    return nil
}


// Contoh fungsi untuk memeriksa password (hash comparison)
func CheckPassword(hashedPassword, password string) bool {
    // Implementasikan logika hash comparison (misal bcrypt.CompareHashAndPassword)
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

var jwtKey = []byte("your_jwt_secret_key")  // Secret key untuk menandatangani token

// Custom claims untuk JWT, kamu bisa menambahkan field tambahan jika diperlukan
type Claims struct {
    UserID uint `json:"user_id"`
    Email  string `json:"email"`
    jwt.StandardClaims
}

// Mengubah fungsi untuk mengisi user dan mengembalikan nilai user dan error
func (us *UserService) FindByIDWithStore(id uint) (*models.User, error) {
    return us.Repo.FindByID(id) // Mengembalikan hasil dari repository
}

// Fungsi untuk generate JWT token
func (us *UserService) GenerateToken(user *models.User) (string, error) {
    // Set expiration time untuk token (misalnya 24 jam)
    expirationTime := time.Now().Add(24 * time.Hour)

    // Buat claims dengan informasi user yang akan disimpan di token
    claims := &Claims{
        UserID: user.ID,
        Email:  user.Email,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    // Buat token dengan method signing HS256 dan claims yang telah dibuat
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Tanda tangani token dengan secret key
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

// Fungsi untuk memvalidasi token JWT
func (us *UserService) ValidateToken(tokenString string) (*Claims, error) {
    claims := &Claims{}

    // Parse token dan verifikasi dengan secret key
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, errors.New("invalid token")
    }

    return claims, nil
}

func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}
