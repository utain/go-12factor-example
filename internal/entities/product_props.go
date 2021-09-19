package entities

// ProductProps model
type ProductProps struct {
	Model
	Key        string `json:"key"`
	Value      string `json:"value"`
	ProductRef string `json:"product_ref"`
}
