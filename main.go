package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
)
type Product struct{
	ID string
	Name string 
	Quantity float64
	Price float64
}

var Products []Product

func homepage(w http.ResponseWriter, r *http.Request){
	log.Println("Started the server")
	fmt.Fprintf(w,"Welcome to homepage")
}

func returnProducts(w http.ResponseWriter, r *http.Request){
	log.Println("Products hit......")
	json.NewEncoder(w).Encode(Products)
}
func getProducts(w http.ResponseWriter,r *http.Request){
	log.Println("Requested hit.....")
	vars := mux.Vars(r)
	key := vars["id"]

	for _,product := range Products{
		if product.ID == key{
			json.NewEncoder(w).Encode(product)
		}
	}
	
}
func HandleFunc(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",homepage)
	myRouter.HandleFunc("/products",returnProducts)
	myRouter.HandleFunc("/products/{id}",getProducts)
	http.ListenAndServe(":7070",myRouter)
}

func main(){
	Products = []Product{
		{ID:"1",Name:"Iphone",Quantity: 400.00,Price: 125000.00},
		{ID:"2",Name: "Ipad",Quantity: 250.00,Price: 65000.00},
	}
	HandleFunc()
}

