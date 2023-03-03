package main

import (
	"fmt"
	"net/http"
	"strconv"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		data := ""
		if req.URL.Query().Get("website") != "" {
			data = req.URL.Query().Get("website")
		} else if req.URL.Query().Get("url") != "" {
			data = req.URL.Query().Get("url")
		} else if req.URL.Query().Get("data") != "" {
			data = req.URL.Query().Get("data")
		} else if req.URL.Query().Get("text") != "" {
			data = req.URL.Query().Get("text")
		}

		size := 150
		if req.URL.Query().Get("size") != "" {
			i, err := strconv.Atoi(req.URL.Query().Get("size"))
			if err == nil {
				size = i
			}
		}

		disableBorder := false
		if req.URL.Query().Get("disableBorder") != "" {
			disableBorder = true
		}

		if data != "" {
			code, _ := qrcode.New(data, qrcode.Medium)
			code.DisableBorder = disableBorder
			png, _ := code.PNG(size)
			w.Header().Set("Content-Type", "image/png")
			w.Write(png)
			return
		}

		w.Write([]byte(`<html><head><title>QR Code</title></head><body>
<form action="/" method="GET" style="margin: 1em auto; text-align: center;">
<label>Website: <input type="text" name="website" /></label><br/>
<label>Size: <input type="text" name="size" value="150" /></label><br/>
<label>Disable border: <input type="checkbox" name="disableBorder" /></label><br/>
<input type="submit"/>
</form>
</body></html>`))
	})
	fmt.Println(http.ListenAndServe(":8080", nil))
}
