package Helper

type User struct {
	FirstName string `json : "firstname"`
	LastName  string `json : "lastname"`
	Email     string `json : "email"`
	UserName  string `json : "username"`
	Password  string `json : "password"`
}

type Credentials struct {
	UserName string `json : "username"`
	Password string `json : "password"`
}

type GenericCountResponse struct {
	Count int
}

type GenericResponse struct {
	Message string `json : "message"`
}
