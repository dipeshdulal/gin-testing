package main

import "encoding/json"

type DoubleRequest struct {
	Number int `json:"number" binding:"required"`
}

func (r DoubleRequest) String() string {
	res, _ := json.Marshal(r)
	return string(res)
}

type ErrorResponse struct {
	Err string `json:"error"`
}

func (r ErrorResponse) String() string {
	res, _ := json.Marshal(r)
	return string(res)
}

type DataResponse struct {
	Data interface{} `json:"data"`
}

func (r DataResponse) String() string {
	res, _ := json.Marshal(r)
	return string(res)
}
