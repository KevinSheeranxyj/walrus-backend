package model

type CreateFormDto struct {
	BlobId   string `json:"blobId"`
	Title    string `json:"title"`
	ItemList []Item `json:"itemList"`
}
