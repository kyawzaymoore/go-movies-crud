package model

type ResponseModelList[T any] struct {
	Code    int    `json:"code:`
	Message string `json:"message"`
	Data    []T    `json:"data"`
}

type ResponseModel[T any] struct {
	Code    int    `json:"code:`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:title`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:firstname`
	Lastname  string `json:lastname`
}
