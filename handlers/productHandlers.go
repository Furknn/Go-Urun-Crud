package handlers

import (
	. "../helpers"
	. "../models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

var productStore = make(map[string]Product)
var id int = 0

//HTTP Get - /api/products
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []Product
	for _, product := range productStore {
		products = append(products, product)
	}
	data, err := json.Marshal(products)
	CheckError(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

//HTTP Post - /api/product
func PostProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	CheckError(err)
	product.CreatedAt = time.Now()
	product.ID = id
	productStore[strconv.Itoa(id)] = product

	data, err := json.Marshal(product)
	CheckError(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

//HTTP Get - /api/product/{id}
func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	for _, prd := range productStore {
		if prd.ID == key {
			product = prd
		}
	}

	data, err := json.Marshal(product)
	CheckError(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

//HTTP Delete - /api/product/{id}
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	if _, ok := productStore[key]; ok {
		delete(productStore, key)
	} else {
		log.Printf(key, " Dğeri bulunamadı")
	}
	w.WriteHeader(http.StatusOK)
}

//HTTP Put - /api/product/{id}
func PutProductHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	key := vars["id"]

	var updateProduct Product
	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	CheckError(err)

	if _, ok := productStore[key]; ok {
		updateProduct.ID, _ = strconv.Atoi(key)
		updateProduct.ChangedAt = time.Now()
		delete(productStore, key)
		productStore[key] = updateProduct
	} else {
		log.Printf(key, " Dğeri bulunamadı")
	}
	w.WriteHeader(http.StatusOK)
}
