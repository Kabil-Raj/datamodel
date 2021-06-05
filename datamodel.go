package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/scrapData", productScrappedData).Methods("POST")
	log.Fatal(http.ListenAndServe(":10001", myRouter))
}

func productScrappedData(w http.ResponseWriter, req *http.Request) {
	fmt.Println(("checking.."))
	var result map[string]string
	json.NewDecoder(req.Body).Decode(&result)
	saveDataInDatabase(result["ProductName"], result["ProductImageUrl"], result["ProductDescription"], result["ProductPrice"], result["ProductReviews"], time.Now())
}

func main() {
	connectMySql()
	handleRequests()
}

func connectMySql() {

	// Connect MySql
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	fmt.Println(DBURL)
	db, err := sql.Open("mysql", DBURL)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Connected to Sql")

	}

	fmt.Println(db.Ping())

	// Create Database
	_, err = db.Exec("CREATE DATABASE AmazonProductDatabase")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(" Database created successfully ")
	}

	// Choose Database
	_, err = db.Exec("USE AmazonProductDatabase")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(" AmazonProductDatabase selected successfully ")
	}

	// Table Creation Query
	sqlStatement, err := db.Prepare("CREATE Table AmazonProductDetails(id int NOT NULL AUTO_INCREMENT, ProductName varchar(255), ProductImageUrl varchar(255), ProductDescription varchar(10000), ProductPrice varchar(255), ProductReviews varchar(255),CreatedTime DATETIME,PRIMARY KEY (id));")

	if err != nil {
		fmt.Println(err.Error())
	}

	// Table Execution Query
	_, err = sqlStatement.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Table created successfully  ")
	}

	defer db.Close()

}

func saveDataInDatabase(productName string, productImageUrl string, productDescription string, productPrice string, productReviews string, createdTime time.Time) {
	// Connecting to database
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	fmt.Println(DBURL)
	db, err := sql.Open("mysql", DBURL)

	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = db.Exec("USE AmazonProductDatabase")

	if err != nil {
		fmt.Println(err.Error())
	}
	// Insert values into database
	sqlInsertStatement, err := db.Prepare("INSERT INTO AmazonProductDetails (ProductName,ProductImageUrl,ProductDescription,ProductPrice,ProductReviews,CreatedTime) VALUES (?,?,?,?,?,?);")
	if err != nil {
		fmt.Println("error")
		fmt.Println(err.Error())
	}

	fmt.Sprintln(productName, productImageUrl, productDescription, productPrice, productReviews)

	_, err = sqlInsertStatement.Exec(productName, productImageUrl, productDescription, productPrice, productReviews, createdTime)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()
}
