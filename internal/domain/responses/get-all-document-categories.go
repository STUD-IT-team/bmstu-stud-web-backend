package responses

type GetAllDocumentCategories struct {
	Categories []Category `json:"categories"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
