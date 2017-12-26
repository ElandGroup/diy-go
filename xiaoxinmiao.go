package main

import "encoding/json"

type XXM struct {
}

func (XXM) StructToJson(book *Book) (result string, err error) {
	b, err := json.Marshal(book)
	if err != nil {
		return
	}
	result = string(b)
	return
}
