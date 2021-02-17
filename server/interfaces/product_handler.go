package interfaces

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Prashant-sharma3012/ecom/tree/main/server/application"
	"github.com/Prashant-sharma3012/ecom/tree/main/server/domain/entity"
	"github.com/gorilla/mux"
)

var productAPIInstance productAPI

type productAPI struct {
	pa application.ProductAppInterface
}

func SetupProductRoutes(r *mux.Router, p application.ProductAppInterface) {
	productAPIInstance = productAPI{
		pa: p,
	}

	r.HandleFunc("/product/all", productAPIInstance.getProducts).Methods("GET")
	r.HandleFunc("/product", productAPIInstance.createProduct).Methods("POST")
	r.HandleFunc("/products/:productId", productAPIInstance.getProductByID).Methods("GET")
	r.HandleFunc("/products/:productId", productAPIInstance.updateProduct).Methods("PUT")
	r.HandleFunc("/product/:productId", productAPIInstance.deleteProduct).Methods("DELETE")
}

func (p *productAPI) getProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting Products")
	products, err := p.pa.GetAllProduct(0, 10)
	if err != nil {
		w.Write([]byte("Error Occured"))
		return
	}

	log.Println("completed")

	log.Println("creating response")
	response, err := json.Marshal(products)
	if err != nil {
		w.Write([]byte("Error Occured"))
		return
	}

	log.Println("sending")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (p *productAPI) createProduct(w http.ResponseWriter, r *http.Request) {
	var product *entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.Write([]byte("Error Occured"))
		return
	}

	savedProduct, err := p.pa.SaveProduct(product)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error Occured"))
		return
	}

	response, err := json.Marshal(savedProduct)
	if err != nil {
		w.Write([]byte("Error Occured"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (p *productAPI) getProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]

	product, err := p.pa.GetProduct(productId)
	if err != nil {
		w.Write([]byte("Error Occured"))
		return
	}

	response, err := json.Marshal(product)
	if err != nil {
		w.Write([]byte("Error Occured"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (p *productAPI) updateProduct(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// productId := vars["productId"]

	var product *entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.Write([]byte("Error Occured"))
		return
	}

	message, err := p.pa.UpdateProduct(product)
	if err != nil {
		w.Write([]byte("Error Occured"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

func (p *productAPI) deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]

	message, err := p.pa.DeleteProduct(productId)
	if err != nil {
		w.Write([]byte("Error Occured"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(message)
}