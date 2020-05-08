package main

import (
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		website := req.URL.Query().Get("website")
		if website != "" {
			png, _ := qrcode.Encode(website, qrcode.Medium, 80)
			w.Write(png)
			return
		}

		w.WriteHeader(400)
	})
	fmt.Println(http.ListenAndServe(":8080", nil))
}
