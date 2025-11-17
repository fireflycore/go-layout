package model

type UpdateDemo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Status      uint32 `json:"status"`
	Sort        uint32 `json:"sort"`
}
