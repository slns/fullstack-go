package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/nsf/jsondiff"
	"github.com/slns/fullstack/api/responses"
)

	type Result struct {
		Result    string
		Message []string
	}

    
	type ResultJson struct {
		Value1 []string
		Value2 []string
	}

// ComparePost , method to compare two jsons
func (server *Server) ComparePost(w http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		rj := ResultJson{}
		err = json.Unmarshal(body, &rj)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		jsValue1, err := json.Marshal(rj.Value1)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		jsValue2, err := json.Marshal(rj.Value2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		opts := jsondiff.DefaultHTMLOptions()
		opts.PrintTypes = true

		diff, text := jsondiff.Compare([]byte(jsValue1), []byte(jsValue2), &opts)

		compare := Result{diff.String(), []string{text}}
		
		w.Header().Set("Content-Type", "application/json")
		responses.JSON(w, http.StatusOK, compare)
}