package dtos

type Response struct {
	Success string      `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
