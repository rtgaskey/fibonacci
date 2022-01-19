package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Push(element int) {
	*s = append(*s, element)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]

		return element, true
	}
}

func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]

		return element, true
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
