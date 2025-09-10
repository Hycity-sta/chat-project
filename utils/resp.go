package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H struct {
	Code  int
	Msg   string
	Data  any
	Rows  any
	Total any
}

func Resp(w http.ResponseWriter, code int, data any, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	h := H{
		Code: code,
		Data: data,
		Msg:  msg,
	}

	ret, err := json.Marshal(h)

	if err != nil {
		fmt.Println(err)
	}

	w.Write(ret)
}

func RespList(w http.ResponseWriter, code int, data any, total any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	h := H{
		Code:  code,
		Rows:  data,
		Total: total,
	}

	ret, err := json.Marshal(h)

	if err != nil {
		fmt.Println(err)
	}

	w.Write(ret)
}

func RespFail(w http.ResponseWriter, msg string) {
	Resp(w, -1, nil, msg)
}

func RespOK(w http.ResponseWriter, data interface{}, msg string) {
	Resp(w, 0, data, msg)
}

func RespOKList(w http.ResponseWriter, data interface{}, total interface{}) {
	RespList(w, 0, data, total)
}
