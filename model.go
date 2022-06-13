package model

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Customer struct {
	Id         int    `gorm:"primaryKey"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
}

type Payment struct {
	Id     int     `json:"id"`
	Amount float32 `json:"amount"`
}

func Findcustomer(c *gin.Context) {
	var customers []Models.Customer
	Model.DB.Find(&customers)
	c.JSON(http.StatusOK, gin.H{"data": customers})
}

func Findpayments(c *gin.Context) {
	var Payments []Models.Payment
	Model.DB.Find(&payment)
	c.JSON(http.StatusOK, gin.H{"datsa": Payments})

}
