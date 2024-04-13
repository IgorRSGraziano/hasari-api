package models

type Album struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Cover []byte  `json:"cover"`
	Music []Music `json:"music"`
}
