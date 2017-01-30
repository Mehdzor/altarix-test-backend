package web

import (
	"gopkg.in/gin-gonic/gin.v1"
	"altarix_test/services"
	"altarix_test/model"
	"altarix_test/socket"
	"net/http"
	"strconv"
)

func Run() {
	r := gin.Default()
	r.GET("/list", handleList)
	r.POST("/sell", handleSell)
	r.POST("/buy/:id", handleBuy)
	r.GET("/ws", handleSocket)
	r.Run(":9104")
}

func handleList(c *gin.Context) {
	c.JSON(http.StatusOK, services.ListProducts())
}

func handleSell(c *gin.Context) {
	var p model.Product
	c.Bind(&p)
	services.AddProduct(&p)
	c.JSON(http.StatusCreated, p)
}

func handleBuy(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := services.DeleteProduct(id); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.Writer.WriteHeader(http.StatusNoContent)
	}

}

func handleSocket(c *gin.Context) {
	socket.ActiveSocket.Serve(c.Writer, c.Request)
}

