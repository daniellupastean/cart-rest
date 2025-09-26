package main

import (
	"log"

	"cart-rest/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	cartService := internal.NewCartService()
	cartHandler := internal.NewCartHandler(cartService)

	r := gin.Default()
	setupRoutes(r, cartHandler)

	log.Println("Starting server on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func setupRoutes(r *gin.Engine, cartHandler *internal.CartHandler) {
	cartRoutes := r.Group("/cart")
	{
		cartRoutes.GET("/:id", cartHandler.GetCart)
		cartRoutes.POST("/:id/items", cartHandler.AddItem)
		cartRoutes.PUT("/:id/items/:productId", cartHandler.UpdateItem)
		cartRoutes.DELETE("/:id/items/:productId", cartHandler.RemoveItem)
	}
}
