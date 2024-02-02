package model

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Status    int    `json:"status"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
