package controllers

import (
    "backendmono/services"  
    "backendmono/models"  
    "github.com/gin-gonic/gin"
    "github.com/gorilla/mux"
	"net/http"
	"strconv"
	"log"
)

type UserController struct {
    UserService *services.UserService 
}

func NewUserController(userService *services.UserService) *UserController {
    return &UserController{UserService: userService}
}

func (uc *UserController) Register(c *gin.Context) {
    // Handler logic untuk register
    var registerReq models.RegisterRequest

    // Bind request JSON ke struct RegisterRequest
    if err := c.ShouldBindJSON(&registerReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Panggil UserService untuk register
    user, err := uc.UserService.Register(registerReq) // Gunakan registerReq di sini
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Registration failed"})
        return
    }

    // Jika registrasi berhasil, kirimkan respons sukses
    c.JSON(http.StatusOK, gin.H{"message": "Registration successful", "user": user})

}

func (uc *UserController) Login(c *gin.Context) {
    // Handler logic untuk login
	var loginReq models.LoginRequest

    // Bind request JSON ke struct LoginRequest
    if err := c.ShouldBindJSON(&loginReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Panggil UserService untuk verifikasi login
    user, err := uc.UserService.Login(loginReq.Email, loginReq.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    // Jika berhasil, misalnya menghasilkan token (JWT)
    token, err := uc.UserService.GenerateToken(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

	// Simpan email dalam context
    c.Set("email", user.Email) // Simpan email setelah login

    // Kirimkan respon berhasil dengan token
    c.JSON(http.StatusOK, gin.H{
        "message": "Login successful",
        "token":   token,
    })

}

func (uc *UserController) GetProfile(c *gin.Context) {
    vars := mux.Vars(c.Request) // Mengambil parameter dari request
    idStr := vars["id"]         // Dapatkan nilai dari parameter "id"

    log.Println("ID from URL:", idStr)
    log.Println("test")
	id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    user, err := uc.UserService.FindByIDWithStore(uint(id)) // Dapatkan user dan error
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, user)
}





func (uc *UserController) ForgotPassword(c *gin.Context) {
    // Handler logic untuk mendapatkan profil user
    // Misalnya: Ambil user ID dari request dan panggil UserService.GetProfile
}

func (uc *UserController) EditProfile(c *gin.Context) {
    log.Println("Retrieving email from context...")
    
    // Buat instance dari EditProfileRequest
    var editProfileReq models.EditProfileRequest
    if err := c.ShouldBindJSON(&editProfileReq); err != nil {
        log.Println("Invalid request:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    // Log email yang diterima dari request
    log.Println("Email from request:", editProfileReq.Email)

    // Buat struct user yang akan diupdate
    updatedUser := models.User{
        Email:   editProfileReq.Email, // Menggunakan email yang diambil dari request
        Name:    editProfileReq.Name,
        Phone:   editProfileReq.Phone,
        Address: editProfileReq.Address,
    }

    // Panggil service untuk update data berdasarkan email
    err := uc.UserService.EditUserProfile(editProfileReq.Email, updatedUser)
    if err != nil {
        log.Println("Failed to update profile:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}


