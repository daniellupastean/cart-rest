package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartService *CartService
}

func NewCartHandler(cartService *CartService) *CartHandler {
	return &CartHandler{
		cartService: cartService,
	}
}

type AddItemRequest struct {
	ProductID string `json:"productId" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
}

type UpdateItemRequest struct {
	Quantity int `json:"quantity" binding:"required,min=0"`
}

func (h *CartHandler) GetCart(c *gin.Context) {
	cartID := c.Param("id")

	cart, err := h.cartService.GetCart(cartID)
	if err != nil {
		if err == ErrCartNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "cart not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, cart)
}

func (h *CartHandler) AddItem(c *gin.Context) {
	cartID := c.Param("id")

	var req AddItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cart, err := h.cartService.AddItemToCart(cartID, req.ProductID, req.Quantity)
	if err != nil {
		if err == ErrInvalidQuantity {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid quantity"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, cart)
}

func (h *CartHandler) UpdateItem(c *gin.Context) {
	cartID := c.Param("id")
	productID := c.Param("productId")

	var req UpdateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cart, err := h.cartService.UpdateItemInCart(cartID, productID, req.Quantity)
	if err != nil {
		if err == ErrCartNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "cart not found"})
			return
		}
		if err == ErrInvalidQuantity {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid quantity"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, cart)
}

func (h *CartHandler) RemoveItem(c *gin.Context) {
	cartID := c.Param("id")
	productID := c.Param("productId")

	cart, err := h.cartService.RemoveItemFromCart(cartID, productID)
	if err != nil {
		if err == ErrCartNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "cart not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, cart)
}
