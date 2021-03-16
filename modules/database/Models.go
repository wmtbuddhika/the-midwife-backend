package database

const (
	UsersSchema = "Users"
)

type Login struct {
	Id       int     `json:"id"`
	NIC      string  `json:"nic"`
	Email      string `json:"email"`
	UserName string `json:"user_name"`
	Password string  `json:"password"`
	Token    string  `json:"token"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	UserTypeId   int `json:"user_type_id"`
}

type Response struct {
	Status bool   `json:"status"`
	StatusCode int   `json:"status_code"`
	Data   interface{} `json:"data"`
	Token  string `json:"token"`
}
