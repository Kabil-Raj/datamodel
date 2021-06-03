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

	db, err := sql.Open("mysql", "root:Electronic1702!@tcp(127.0.0.1:3306)/test")

	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()
}

func Sample(productName string, productImageUrl string, productDescription string, productPrice string, productReviews string, createdTime time.Time) {
	fmt.Println(" ", productName, " ", productImageUrl, " ", createdTime)
}
