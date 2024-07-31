package models

type Account struct {
	Domain   string `json:"domain"`
	Username string `json:"username"`
	Password string `json:"password"`
	Desc     string `json:"Desc"`
}
