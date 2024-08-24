package model

type CreateFormDto struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	ItemList []Item `json:"itemList"`
	BlobId   string `json:"blobId"`
}
