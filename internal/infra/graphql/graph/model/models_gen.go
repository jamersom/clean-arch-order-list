package model

type DeleteOrderInput struct {
	ID string `json:"id"`
}

type Order struct {
	ID         string  `json:"id"`
	Price      float64 `json:"Price"`
	Tax        float64 `json:"Tax"`
	FinalPrice float64 `json:"FinalPrice"`
}

type OrderInput struct {
	Price float64 `json:"Price"`
	Tax   float64 `json:"Tax"`
}

type UpdateOrderInput struct {
	ID    string  `json:"id"`
	Price float64 `json:"Price"`
	Tax   float64 `json:"Tax"`
}
