package main

import (
	// "fmt"
	// "html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BankAccount struct {
	Name    string
	Balance float64
}

var account *BankAccount

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create_account.html", nil)
	})

	r.POST("/create", func(c *gin.Context) {
		name := c.PostForm("name")
		balanceStr := c.PostForm("balance")
		balance, _ := strconv.ParseFloat(balanceStr, 64)
		account = &BankAccount{Name: name, Balance: balance}
		c.Redirect(http.StatusSeeOther, "/")
	})

	r.GET("/deposit", func(c *gin.Context) {
		c.HTML(http.StatusOK, "deposit.html", nil)
	})

	r.POST("/deposit", func(c *gin.Context) {
		amountStr := c.PostForm("amount")
		amount, _ := strconv.ParseFloat(amountStr, 64)
		account.Balance += amount
		c.Redirect(http.StatusSeeOther, "/")
	})

	r.GET("/withdraw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "withdraw.html", nil)
	})

	r.POST("/withdraw", func(c *gin.Context) {
		amountStr := c.PostForm("amount")
		amount, _ := strconv.ParseFloat(amountStr, 64)
		if amount > account.Balance {
			c.HTML(http.StatusOK, "withdraw.html", gin.H{
				"InsufficientFunds": true,
			})
			return
		}
		account.Balance -= amount
		c.Redirect(http.StatusSeeOther, "/")
	})

	r.GET("/balance", func(c *gin.Context) {
		c.HTML(http.StatusOK, "balance.html", gin.H{
			"Balance": account.Balance,
		})
	})

	r.Run(":8080")
}
