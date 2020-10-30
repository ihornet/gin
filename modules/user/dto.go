package user

type dto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type body struct {
	Id   int    `binding:"required"`
	Name string `binding:"required"`
}
