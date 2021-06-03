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

	db, err := sql.Open("mysql", "root:Electronic1702!@tcp(127.0.0.1:3306)/")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Connected to Sql")
	}

	_, err = db.Exec("CREATE DATABASE AmazonProductDetailsDatabase")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(" Database created successfully ")
	}

	sqlStatement, err := db.Prepare("CREATE Table AmazonProductDetails(id int NOT NULL AUTO_INCREMENT, ProductName, ProductImageUrl, ProductDescription, ProductPrice, ProductReviews,CreatedTime,PRIMARY KEY (id));")

	if err != nil {
		fmt.Println(err.Error())
	}

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
	_, err = db.Exec("USE AmazonProductDetailsDatabase")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(" testDB selected successfully ")
	}

	// Insert values into database
	sqlInsertStatement, err := db.Prepare("INSERT INTO employee (ProductName, ProductImageUrl, ProductDescription, ProductPrice, ProductReviews,CreatedTime) VALUES (productName,ProductImageUrl,productDescription,productPrice, productReviews,createdTime);")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = sqlInsertStatement.Exec()
	if err != nil {
		fmt.Println(err.Error())
	}
}
