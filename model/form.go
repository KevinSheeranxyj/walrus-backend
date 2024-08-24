package model

type CreateFormDto struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description int8   `json:"description"`
	ItemList    []Item `json:"itemList"`
}
