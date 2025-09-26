package internal

type Cart struct {
	ID    string         `json:"id"`
	Items map[string]int `json:"items"`
}

func NewCart(id string) *Cart {
	return &Cart{
		ID:    id,
		Items: make(map[string]int),
	}
}

func (c *Cart) AddItem(productID string, quantity int) {
	if c.Items == nil {
		c.Items = make(map[string]int)
	}
	c.Items[productID] += quantity
}

func (c *Cart) UpdateItem(productID string, quantity int) {
	if c.Items == nil {
		c.Items = make(map[string]int)
	}
	c.Items[productID] = quantity
}

func (c *Cart) RemoveItem(productID string) {
	if c.Items != nil {
		delete(c.Items, productID)
	}
}

func (c *Cart) HasItem(productID string) bool {
	if c.Items == nil {
		return false
	}
	_, exists := c.Items[productID]
	return exists
}

func (c *Cart) GetItemQuantity(productID string) int {
	if c.Items == nil {
		return 0
	}
	return c.Items[productID]
}
