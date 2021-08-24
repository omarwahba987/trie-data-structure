package api

import (
	"encoding/json"
	"net/http"
	"va_test_d/httpHelper"
	"va_test_d/internal"
)

type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

var Trie *internal.Trie

func InitTrie() {
	Trie = internal.InitTrie()
}
func InsertPerson(w http.ResponseWriter, req *http.Request) {
	var person Person
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&person)
	httpHelper := httpHelper.HttpHandlerRepo{}
	if err != nil {
		httpHelper.SendError(w, "there is an error decoding")
		return
	}
	err = Trie.Insert(person.Name, person.Address)
	if err != nil {
		httpHelper.SendResponseMessage(w, err)
	}
	httpHelper.SendResponseMessage(w, "person added success")

}
func Visualize(w http.ResponseWriter, req *http.Request) {

	httpHelper := httpHelper.HttpHandlerRepo{}

	httpHelper.SendResponseMessage(w, Trie)

}
