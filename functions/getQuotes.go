package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rhyuen/types"
)

type Payload struct {
	Path string        `json:"path"`
	Data []types.Quote `json:"quotes"`
}

//Handler ... Exported Handler REQ, RES
func Handler(w http.ResponseWriter, r *http.Request) {
	latest := types.Quote{"Mahatma Ghandi", "Be the change you wish to see in the world."}
	second := types.Quote{"Batman", "vengeance is the night."}
	list := make([]types.Quote, 0)
	list = append(list, latest, second)
	send := Payload{"path is here", list}
	json.NewEncoder(w).Encode(send)

}
