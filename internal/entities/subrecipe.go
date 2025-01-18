package entities

type Subrecipe struct {
	ID         int64      `json:"id"`
	Quantity   float64    `json:"quantity"`
	Note       string     `json:"note"`
	Ingredient Ingredient `json:"ingredient"`
	Unit       Unit       `json:"unit"`
}
