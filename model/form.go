package model

type CreateFormDto struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	ItemList    []Item `json:"itemList"`
	Type        int    `json:"type"`
	Creator     string `json:"creator"`
	Participant string `json:"participant"`
}

type GetFormDto struct {
	BlobId string `json:"blobId"`
}
