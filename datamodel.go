package datamodel

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type ProductDetail struct {
	ProductName        string    `json:"ProductName"`
	ProductImageUrl    string    `json:"ProductImageUrl"`
	ProductDescription string    `json:"ProductDescription"`
	ProductPrice       string    `json:"ProductPrice"`
	ProductReviews     string    `json:"ProductReviews"`
	CreatedAt          time.Time `json:"CreatedAt"`
}

func ConnectMySql() {
	fmt.Println("Connecting MySql")
	// Connect MySql
	db, err := sql.Open("mysql", "root:Electronic1702!@tcp(127.0.0.1:3306)/")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Connected to Sql")
	}

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

func SaveData(productName string, productImageUrl string, productDescription string, productPrice string, productReviews string, createdTime time.Time) {
	// Connecting to database
	db, err := sql.Open("mysql", "root:Electronic1702!@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Connected to Sql")
	}

	// select database
	_, err = db.Exec("USE AmazonProductDatabase")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(" AmazonProductDatabase selected successfully ")
	}

	// Insert values into database
	sqlInsertStatement, err := db.Prepare("INSERT INTO AmazonProductDetails (ProductName,ProductImageUrl,ProductDescription,ProductPrice,ProductReviews,CreatedTime) VALUES (?,?,?,?,?,?);")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = sqlInsertStatement.Exec(productName, productImageUrl, productDescription, productPrice, productReviews, createdTime)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()
}
