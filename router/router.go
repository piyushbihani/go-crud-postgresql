package router

import (
	"github.com/gorilla/mux"
	"github.com/piyushbihani/go_stocks_crud/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock", middleware.GetAllStocks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/createstock", middleware.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/updatestock/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deletestock/{id}", middleware.DeleteStock).Methods("DELETE", "OPTIONS")
	return router
}
