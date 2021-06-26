package main

import (
	"fmt"
	"github.com/bregydoc/gtranslate"
	"github.com/julienschmidt/httprouter"
	_ "github.com/julienschmidt/httprouter"
	"net/http"
	_ "net/url"
)

func main() {
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
	server := http.Server{
		Handler: router,
		Addr: "localhost:3000",
	}

	server.ListenAndServe()

}
