package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/farizziezhi/layered/internal/handler"
	authMiddleware "github.com/farizziezhi/layered/internal/middleware"
	"github.com/farizziezhi/layered/internal/repository"
	"github.com/farizziezhi/layered/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() (*sql.DB, error) {
    db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/product_go?parseTime=true")
    if err != nil {
        return nil, err
    }
    db.SetMaxIdleConns(25)
    db.SetMaxOpenConns(100)
    db.SetConnMaxIdleTime(5 * time.Minute)
    db.SetConnMaxLifetime(60 * time.Minute)
    return db, nil
}

func main() {
	db, err := GetConnection()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Gagal terhubung ke database: ", err)
	}
	log.Println("Berhasil terhubung ke database")

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

    r := chi.NewRouter()
    r.Use(middleware.Logger)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/register", userHandler.Register)
		r.Post("/login", userHandler.Login)

		r.Group(func(r chi.Router) {
			r.Use(authMiddleware.AuthMiddleware)
			r.Get("/products", productHandler.GetProducts)
			r.Post("/products", productHandler.CreateProduct) 
			r.Get("/products/{id}", productHandler.GetProduct) 
			r.Put("/products/{id}", productHandler.UpdateProduct)
			r.Delete("/products/{id}", productHandler.DeleteProduct)
		})
	})
	log.Println("Server berjalan di port :3000")
    http.ListenAndServe(":3000", r)
}
