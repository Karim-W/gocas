package users

type Entity struct {
	Id      string `json:"id"`
	Email   string `json:"email"`
	Age     int    `json:"age"`
	Country string `json:"country"`
}
