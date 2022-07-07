package main

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
	"github.com/gorilla/mux"
	"net/http"
)

func loadData() []byte {
	jsonFile, err := os.Open("products.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)

	return data
}

type Product struct {
	Uuid string `json:"uuid"`
	Product string `json: "product"`
	Price float64 `jason: "price, string"`
}

type Products struct {
	Products []Product

}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products := loadData()
	w.Write([]byte(products))
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := loadData()

	var products Products
	json.Unmarshal(data, &products)

	for _, v := range products.Products {
		if v.Uuid == vars["id"] {
			product, _ := json.Marshal(v)
			w.Write([]byte(product))
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", ListProducts)
	r.HandleFunc("/product/{id}", GetProductById)
	http.ListenAndServe(":8081", r)
}