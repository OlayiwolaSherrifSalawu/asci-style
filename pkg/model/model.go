package model

type User struct {
	UserId       string `json:"userid"`
	UserName     string `json:"username"`
	EmailAddress string `json:"emailAddress"`
	HashPassword string `json:"hashPassword"`
}

type AsciiStyle struct {
	UserId    string `json:"userid"`
	AsciiText string `json:"asciiText"`
	PlainText string `json:"plainText"`
	Banner    string `json:"banner"`
}
