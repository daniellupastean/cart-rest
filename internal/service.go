package internal

import (
	"errors"
	"sync"
)

var (
	ErrCartNotFound    = errors.New("cart not found")
	ErrInvalidQuantity = errors.New("invalid quantity")
)

type CartService struct {
	carts map[string]*Cart
	mutex sync.RWMutex
}

func NewCartService() *CartService {
	return &CartService{
		carts: make(map[string]*Cart),
	}
}

func (s *CartService) GetCart(id string) (*Cart, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	cart, exists := s.carts[id]
	if !exists {
		return nil, ErrCartNotFound
	}
	return cart, nil
}

func (s *CartService) CreateCart(id string) *Cart {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if cart, exists := s.carts[id]; exists {
		return cart
	}

	cart := NewCart(id)
	s.carts[id] = cart
	return cart
}

func (s *CartService) AddItemToCart(cartID, productID string, quantity int) (*Cart, error) {
	if quantity <= 0 {
		return nil, ErrInvalidQuantity
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	cart, exists := s.carts[cartID]
	if !exists {
		cart = NewCart(cartID)
		s.carts[cartID] = cart
	}

	cart.AddItem(productID, quantity)
	return cart, nil
}

func (s *CartService) UpdateItemInCart(cartID, productID string, quantity int) (*Cart, error) {
	if quantity < 0 {
		return nil, ErrInvalidQuantity
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	cart, exists := s.carts[cartID]
	if !exists {
		return nil, ErrCartNotFound
	}

	if quantity == 0 {
		cart.RemoveItem(productID)
	} else {
		cart.UpdateItem(productID, quantity)
	}

	return cart, nil
}

func (s *CartService) RemoveItemFromCart(cartID, productID string) (*Cart, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	cart, exists := s.carts[cartID]
	if !exists {
		return nil, ErrCartNotFound
	}

	cart.RemoveItem(productID)
	return cart, nil
}
