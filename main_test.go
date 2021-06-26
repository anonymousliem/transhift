package main

import (
	"fmt"
	"github.com/bregydoc/gtranslate"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestName(t *testing.T) {
	router := httprouter.New()
	var url = "/transhift"
	router.GET(url, func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		queryValues := request.URL.Query()
		text := queryValues.Get("q")
		translated, err := gtranslate.TranslateWithParams(
			text,
			gtranslate.TranslationParams{
				From: "id",
				To:   "en",
			},
		)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(writer, translated)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/transhift?q=halo%20dunia", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder,request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t,"Hello World", string(body))
}
