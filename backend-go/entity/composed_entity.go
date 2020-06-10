package entity

type CartProduct struct {
	Cart
	Product
}

type PromoteSection struct {
	Title    string    `json:"title"`
	Category int       `json:"category"`
	Products []Product `json:"products"`
}
