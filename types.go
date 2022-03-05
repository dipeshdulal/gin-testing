package main

import "encoding/json"

type DoubleRequest struct {
	Number int `json:"number" binding:"required"`
}

func (r DoubleRequest) JSON() string {
	res, _ := json.Marshal(r)
	return string(res)
}

type ErrorResponse struct {
	Err string `json:"error"`
}

func (r ErrorResponse) JSON() string {
	res, _ := json.Marshal(r)
	return string(res)
}

type DataResponse struct {
	Data interface{} `json:"data"`
}

func (r DataResponse) JSON() string {
	res, _ := json.Marshal(r)
	return string(res)
}
