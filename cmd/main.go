package main

import (
	"go-api/controller"
	db "go-api/infrastructure"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	dbConnection, err := db.ConnectDB()
	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	if err != nil {
		panic(err)
	}
	ProductController := controller.NewProductController(ProductUseCase)
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.GET("/products", ProductController.GetProducts)
	server.GET("/product/:productId", ProductController.GetProductById)
	server.POST("/product", ProductController.CreateProduct)
	server.Run(":8080")
}
