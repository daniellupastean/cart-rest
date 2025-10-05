package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"cart-rest/internal"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	telemetry := setupTelemetry()
	defer shutdownTelemetry(telemetry)

	cartService := internal.NewCartService()
	cartHandler := internal.NewCartHandler(cartService)

	r := gin.Default()
	r.Use(internal.MetricsMiddleware(telemetry))

	setupRoutes(r, cartHandler)
	setupMonitoring(r)

	log.Println("Starting server on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func setupTelemetry() *internal.Telemetry {
	telemetry, err := internal.InitTelemetry("cart-rest-service")
	if err != nil {
		log.Fatal("Failed to initialize telemetry:", err)
	}
	return telemetry
}

func shutdownTelemetry(telemetry *internal.Telemetry) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := telemetry.Shutdown(ctx); err != nil {
		log.Printf("Error shutting down telemetry: %v", err)
	}
}

func setupMonitoring(r *gin.Engine) {
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "healthy",
			"service":   "cart-rest-service",
			"timestamp": time.Now().Unix(),
		})
	})
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
