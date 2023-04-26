package models

type Query struct{
	offset int `json:"offset", default:"0"`
	limit int `json:"limit", default:"10"`
	search string `json:"search"`
}

type DefaultError struct{
	Message string `json:"message"`
}

type ErrorResponse struct{
	Message string `json:"message"`
	Code int `json:"code"`
}

type SuccessResponse struct{
	Message string `json:"message"`
	Date interface{} `json:"date"`
}
