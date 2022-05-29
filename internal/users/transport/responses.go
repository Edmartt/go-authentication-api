package transport

//json response for login
type LoginResponse struct{
	Token string `json:"token"`
}

type SignupResponse struct{
	Status string `json:"status"`
}
