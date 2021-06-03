package datamodel

import (
	"fmt"
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

func Sample(productName string, productImageUrl string, productDescription string, productPrice string, productReviews string, createdTime time.Time)  {
	fmt.Println(" ", productName, " ", productImageUrl, " ", createdTime)
}
