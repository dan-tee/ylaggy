package main

import "encoding/json"

type JsonMap map[string]interface{}

func (r JsonMap) String() (s string, err error) {
	var b []byte
	b, err = json.Marshal(r)
	if err != nil {
		s = ""
		return
	}
	s = string(b)
	return
}
