package main

import (
	"golang.org/x/net/context"
	"github.com/guregu/kami"
	"net/http"
	"io/ioutil"
)

func main() {
	kami.Context = context.Background()

	kami.Get("/url", func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")

		if url == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		} else {
			resp, err := http.Get(url)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Write(body)
		}
	})

	kami.Serve()
}
