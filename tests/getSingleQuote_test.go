package test

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/rhyuen/golang_mongo_faas/types"
)

func TestGetSingleQuote(t *testing.T) {
	url := "http://localhost:3000/api/getSingleQuote/5e08bda676b8cee331461933"
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var payload types.Payload
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		log.Fatal(err)
	}
	checkResponseCode(t, "/getSingleQuoteUpdated", payload.Path)

}

func TestGetAllQuotes(t *testing.T) {

}

func TestMakeNewQuote(t *testing.T) {

}

func TestDeleteQuote(t *testing.T) {

}

func checkResponseCode(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected: %s, Received: %s", expected, actual)
	}
}
