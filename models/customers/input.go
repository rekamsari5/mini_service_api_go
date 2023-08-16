package customers

type RequestCustomer struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type SearchCustomer struct {
	Name    string `json:"name" `
	Address string `json:"address" `
}

type RequestDelete struct {
	ID int `json:"id" binding:"required"`
}

type RequestUpdate struct {
	ID      int    `json:"id" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}
