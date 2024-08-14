package models

type Ghoul struct {
	FirstName    string `form:"firstName"`
	LastName     string `form:"lastName"`
	Descsription string `form:"description"`
	Img          string `form:"img"`
	Status       string `form:"status"`
	Rank         string `form:"rank"`
}
