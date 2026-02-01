package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/database"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var products = []Product{
	{ID: 1, Name: "Indomie Godog", Price: 3500, Stock: 10},
	{ID: 2, Name: "Vit 1000ml", Price: 3000, Stock: 40},
}

var categories = []Category{
	{ID: 1, Name: "Makanan", Description: "Makanan ringan"},
	{ID: 1, Name: "Minuman", Description: "Minuman"},
}

// MARK: CATEGORY
func getCategoryById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for _, p := range categories {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	var newCategory Category
	err = json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for i := range categories {
		if categories[i].ID == id {
			newCategory.ID = id
			categories[i] = newCategory

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(categories)
			return
		}
	}
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for i, p := range categories {
		if p.ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "success",
			})
			return
		}
	}
}

func getCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func createCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	newCategory.ID = len(categories) + 1
	categories = append(categories, newCategory)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(newCategory)
}

// MARK: PRODUK
func getProdukByid(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for _, p := range products {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Produk tidak ditemukan", http.StatusNotFound)
}

func updateProduk(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	var newProduct Product
	err = json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for i := range products {
		if products[i].ID == id {
			newProduct.ID = id
			products[i] = newProduct

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products)
			return
		}
	}
}

func deleteProduk(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "success",
			})
			return
		}
	}
}

func main() {
	// GET localhost:8080/api/categories/{id}
	// PUT localhost:8080/api/categories/{id}
	// DELETE localhost:8080/api/categories/{id}
	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getCategoryById(w, r)
		case "PUT":
			updateCategory(w, r)
		case "DELETE":
			deleteCategory(w, r)
		}
	})

	// GET localhost:8080/api/categories
	// POST localhost:8080/api/categories
	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getCategories(w, r)
		case "POST":
			createCategory(w, r)
		}
	})

	// GET localhost:8080/api/produk/{id}
	// PUT localhost:8080/api/produk/{id}
	// DELETE localhost:8080/api/produk/{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getProdukByid(w, r)
		case "PUT":
			updateProduk(w, r)
		case "DELETE":
			deleteProduk(w, r)
		}
	})

	// GET localhost:8080/api/produk
	// POST localhost:8080/api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products)
		case "POST":
			// baca dari request
			var newProduct Product
			err := json.NewDecoder(r.Body).Decode(&newProduct)
			if err != nil {
				http.Error(w, "invalid request", http.StatusBadRequest)
				return
			}

			// masukkan data ke dalam variable produk
			newProduct.ID = len(products) + 1
			products = append(products, newProduct)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated) // 201
			json.NewEncoder(w).Encode(newProduct)
		}
	})

	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API running",
		})
	})

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port: viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	addr := "localhost:" + config.Port
	fmt.Println("server running di " + addr)

	errServer := http.ListenAndServe(addr, nil)
	if errServer != nil {
		fmt.Println("gagal running server")
	}
}
