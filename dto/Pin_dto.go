package dto

type PIN struct {
	Pin    string `json:"pin"`
	NewPIN string `json:"new_pin"`
}

type Password struct {
	Password    string `json:"password"`
	Newpassword string `json:"new_password"`
}
