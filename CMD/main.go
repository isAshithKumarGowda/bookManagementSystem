package main

import (
	"github.com/gorilla/mux"
	"github.com/isAshithKumarGowda/bookManagementSystem/PKG/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9010", r))
}
