package Helper

type GenericCountResponse struct {
	Count int
}

type GenericIDResponse struct {
	ID int
}

type GenericCountResponseFloat struct {
	Count float32
}

type GenericResponse struct {
	Message string `json : "message"`
}
