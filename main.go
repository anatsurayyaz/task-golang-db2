package main

import (
	"log"
	"os"
	"task-golang-db/handlers"
	"task-golang-db/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// Database
	db := NewDatabase()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get DB from GORM:", err)
	}
	defer sqlDB.Close()

	// secret-key
	signingKey := os.Getenv("SIGNING_KEY")

	r := gin.Default()

	// grouping route with /auth
	authhandlers := handlers.NewAuth(db, []byte(signingKey))
	authRoute := r.Group("/auth")
	authRoute.POST("/login", authhandlers.Login)
	authRoute.POST("/upsert", authhandlers.Upsert)

	// grouping route with /account
	accounthandlers := handlers.NewAccount(db)
	accountRoutes := r.Group("/account")
	accountRoutes.POST("/create", accounthandlers.Create)
	accountRoutes.GET("/read/:id", accounthandlers.Read)
	accountRoutes.PATCH("/update/:id", accounthandlers.Update)
	accountRoutes.DELETE("/delete/:id", accounthandlers.Delete)
	accountRoutes.GET("/list", accounthandlers.List)
	accountRoutes.POST("/topup", accounthandlers.TopUp)

	// middleware := middleware.AuthMiddleware(signingKey)
	accountRoutes.GET("/my", middleware.AuthMiddleware(signingKey), accounthandlers.My)
	accountRoutes.GET("/balance", middleware.AuthMiddleware(signingKey), accounthandlers.Balance)
	accountRoutes.POST("/transfer", middleware.AuthMiddleware(signingKey), accounthandlers.Transfer)
	accountRoutes.GET("/mutation", middleware.AuthMiddleware(signingKey), accounthandlers.Mutation)

	// grouping route with /transaction-category
	transaction_categoryhandlers := handlers.NewTransCat(db)
	transaction_categoryRoutes := r.Group("/transaction-category")
	transaction_categoryRoutes.POST("/create", transaction_categoryhandlers.Create)
	transaction_categoryRoutes.GET("/read/:id", transaction_categoryhandlers.Read)
	transaction_categoryRoutes.PATCH("/update/:id", transaction_categoryhandlers.Update)
	transaction_categoryRoutes.DELETE("/delete/:id", transaction_categoryhandlers.Delete)
	transaction_categoryRoutes.GET("/list", transaction_categoryhandlers.List)

	transaction_categoryRoutes.GET("/my", middleware.AuthMiddleware(signingKey), transaction_categoryhandlers.My)

	transactionhandlers := handlers.NewTrans(db)
	transactionRoutes := r.Group("/transaction")
	transactionRoutes.POST("/new", transactionhandlers.NewTransaction)
	transactionRoutes.GET("/list", transactionhandlers.TransactionList)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func NewDatabase() *gorm.DB {
	// dsn := "host=localhost port=5432 user=postgres dbname=digi sslmode=disable TimeZone=Asia/Jakarta"
	dsn := os.Getenv("DATABASE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get DB object: %v", err)
	}

	var currentDB string
	err = sqlDB.QueryRow("SELECT current_database()").Scan(&currentDB)
	if err != nil {
		log.Fatalf("failed to query current database: %v", err)
	}

	log.Printf("Current Database: %s\n", currentDB)

	return db
}
