package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/piyushbihani/go_stocks_crud/models"
)

type response struct {
	ID      int    `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	db, err := sql.Open("postgres", os.Getenv("POSTGRESQL_URL"))

	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	fmt.Println("Successfully connected to database")
	return db
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Unable to decode the request body: %v", err)
	}
	// db := CreateConnection()
	// query, err := fmt.Printf("INSERT INTO stocksdb(name, price, company) VALUES(%s, %f, %s)", stock.Name, stock.Price, stock.Company)
	// if err != nil {
	// 	log.Fatal("Error in query")
	// }
	// db.Exec(string(query))
	insertId, err := insertStock(stock)

	if err != nil {
		log.Fatalf("Unable to insert stock: %v", err)
	}

	res := response{
		ID:      insertId,
		Message: "Stock Inserted Successfully",
	}
	json.NewEncoder(w).Encode(res)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert a string: %v", err)
	}
	stock, err := getStock(id)

	if err != nil {
		log.Fatalf("Unable to get stock: %v", err)
	}

	if stock.StockId == 0 {
		res := response{
			ID:      stock.StockId,
			Message: fmt.Sprintf("StockId: %v not found", id),
		}
		json.NewEncoder(w).Encode(res)
	} else {
		json.NewEncoder(w).Encode(stock)
	}
}

func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := getAllStocks()

	if err != nil {
		log.Fatalf("Unable to get all stocks: %v", err)
	}
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(stocks)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert a string: %v", err)
	}

	var stock models.Stock
	err = json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatalf("Error decoding the request body: %v", err)
	}

	updatedRows, err := updateStock(id, stock)

	if err != nil {
		log.Fatalf("Unable to update the stock: %v", err)
	}

	var msg string
	if updatedRows != 0 {
		msg = fmt.Sprintf("Stock Updated Successfully. Total rows affected: %v", updatedRows)
	} else {
		msg = "Stock not found"
	}

	res := response{ID: id, Message: msg}
	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert a string: %v", err)
	}

	deletedRows, err := deleteStock(id)

	if err != nil {
		log.Fatalf("Unable to delete the stock: %v", err)
	}

	var msg string
	if deletedRows != 0 {
		msg = fmt.Sprintf("Stock Deleted Successfully. Total rows affected: %v", deletedRows)
	} else {
		msg = "Stock not found"
	}

	res := response{ID: id, Message: msg}
	json.NewEncoder(w).Encode(res)
}

func insertStock(stock models.Stock) (int, error) {
	db := createConnection()
	defer db.Close()
	query := `INSERT INTO stocks(name, price, company) VALUES($1, $2, $3) RETURNING stockid`
	var id int
	err := db.QueryRow(query, stock.Name, stock.Price, stock.Company).Scan(&id)
	if err != nil {
		return 0, err
	}
	fmt.Printf("Inserted a single row %v", id)
	return id, nil
}

func getStock(id int) (models.Stock, error) {
	db := createConnection()
	defer db.Close()
	query := `SELECT * FROM stocks WHERE stockid = $1`
	var stock models.Stock
	row := db.QueryRow(query, id)
	err := row.Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
		stock.StockId = 0
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to scan the row: %v", err)
	}
	return stock, err
}

func getAllStocks() ([]models.Stock, error) {
	db := createConnection()
	defer db.Close()
	query := `SELECT * FROM stocks`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Unable to query: %v", err)
		return []models.Stock{}, err
	}
	var stocks []models.Stock
	for rows.Next() {
		var stock models.Stock
		err := rows.Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company)
		if err != nil {
			log.Fatalf("Unable to scan the row: %v", err)
		}

		stocks = append(stocks, stock)
	}
	return stocks, err
}

func updateStock(id int, stock models.Stock) (int, error) {
	db := createConnection()
	defer db.Close()
	query := `UPDATE stocks SET name = $2, price = $3, company = $4 WHERE stockid = $1`
	row, err := db.Exec(query, id, stock.Name, stock.Price, stock.Company)
	if err != nil {
		log.Fatalf("Unable to execute the query: %v", err)
	}
	rowsAffected, err := row.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows: %v", err)
	}
	return int(rowsAffected), err
}

func deleteStock(id int) (int, error) {
	db := createConnection()
	defer db.Close()
	query := `DELETE FROM stocks WHERE stockid=$1`
	row, err := db.Exec(query, id)
	if err != nil {
		log.Fatalf("Unable to execute the query: %v", err)
	}
	rowsAffected, err := row.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows: %v", err)
	}
	return int(rowsAffected), err
}
