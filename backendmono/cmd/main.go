package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
    "github.com/gin-gonic/gin"
    "backendmono/controllers"
    "backendmono/database"
    "backendmono/repositories"
    "backendmono/services"
    "backendmono/middlewares"
)

// Adapter untuk mengkonversi handler Gin menjadi handler HTTP
func ginHandlerToHTTP(ginHandler func(c *gin.Context)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Membuat konteks Gin baru dan panggil handler Gin
        c, _ := gin.CreateTestContext(w) // Membuat konteks Gin untuk testing
        c.Request = r
        ginHandler(c) // Panggil handler Gin
    }
}

func main() {
    // Inisialisasi koneksi database
    db, err := database.InitDB()
    if err != nil {
        panic("Failed to connect to database")
    }

    // Inisialisasi UserRepository dan UserService
    userRepo := repositories.NewUserRepository(db)
    userSvc := services.NewUserService(userRepo)

    // Inisialisasi UserController dengan UserService
    userController := controllers.NewUserController(userSvc)

	// Inisialisasi StoreRepository dan StoreService
	storeRepo := repositories.NewStoreRepository(db)
	storeSvc := services.NewStoreService(storeRepo)

	// Inisialisasi StoreController dengan StoreService
	storeController := controllers.NewStoreController(storeSvc)

    // Buat router menggunakan Gorilla Mux
  	router := mux.NewRouter()

    // Setup routes untuk login dan register tanpa middleware
    router.HandleFunc("/login", ginHandlerToHTTP(userController.Login)).Methods("POST")
    router.HandleFunc("/register", ginHandlerToHTTP(userController.Register)).Methods("POST")

    // Buat subrouter untuk route yang memerlukan middleware
    authRouter := router.PathPrefix("/").Subrouter() // Subrouter ini akan menggunakan middleware
    authRouter.Use(middlewares.JWTAuthMiddleware("your_jwt_secret_key")) // Terapkan middleware

    // Setup routes dengan middleware
    authRouter.HandleFunc("/profile/{id}", ginHandlerToHTTP(userController.GetProfile)).Methods("GET")
    authRouter.HandleFunc("/editprofile", ginHandlerToHTTP(userController.EditProfile)).Methods("POST")


    authRouter.HandleFunc("/createstore", ginHandlerToHTTP(storeController.CreateStore)).Methods("POST")
    authRouter.HandleFunc("/myproducts/{id}", ginHandlerToHTTP(storeController.FindProductsByStoreID)).Methods("GET")
	authRouter.HandleFunc("/stores/{id}/products", ginHandlerToHTTP(storeController.CreateProduct)).Methods("POST")




	// router := gin.Default()
	// router.GET("/profile/:id", userController.GetProfile) // Gunakan Gin tanpa adapter


    // CORS Options
    corsOptions := handlers.AllowedOrigins([]string{"http://localhost:3000"}) // Ganti dengan domain aplikasi React kamu
    corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
    corsHeaders := handlers.AllowedHeaders([]string{"Origin", "Content-Type", "Authorization"})
    allowCredentials := handlers.AllowCredentials()

    // Menjalankan server dengan handler CORS
    http.ListenAndServe(":8080", handlers.CORS(corsOptions, corsMethods, corsHeaders, allowCredentials)(router))
}