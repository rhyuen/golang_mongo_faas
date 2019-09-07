package handler

import (
	"encoding/json"
	"net/http"
)

type Payload struct {
	Path string  `json:"path"`
	Data []Quote `json:"quotes"`
}

type Quote struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}

//Handler ... Exported Handler REQ, RES
func Handler(w http.ResponseWriter, r *http.Request) {
	//dburl := os.Getenv("go_mongo_db")
	latest := Quote{"Mahatma Ghandi", "Be the change you wish to see in the world."}
	second := Quote{"Batman", "vengeance is the night."}
	list := make([]Quote, 0)
	list = append(list, latest, second)
	send := Payload{"path is here", list}
	json.NewEncoder(w).Encode(send)

}
