package handler

import (
	"encoding/json"
	"net/http"
)

type payload struct {
	Name string `json:"name"`
	Data []int  `json:"data"`
}

//Handler ... Exported Handler REQ, RES
func Handler(w http.ResponseWriter, r *http.Request) {
	numbers := make([]int, 5)
	numbers = append(numbers, 55, 12, 13, 66, 99)
	latest := payload{"mypayload", numbers}
	json.NewEncoder(w).Encode(latest)

}
