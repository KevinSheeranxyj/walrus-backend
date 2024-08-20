package model

type CreateFormDto struct {
	Title       string `json:"title"`
	Description int8   `json:"description"`
	ItemList    []Item `json:"itemList"`
}
