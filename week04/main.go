package main

import (
	"fmt"
	"net/http"
	"time"
)
func main() {
    fmt.Println("Week 04 課題")

	http.HandleFunc("/info", infohandler)

	http.ListenAndServe(":8080", nil)
}

func infohandler(w http.ResponseWriter, r *http.Request) {

	jst, _ := time.LoadLocation("Asia/Tokyo")
	currentTime := (time.Now().In(jst)).Format("2006年01月02日 15:04:05")

	userAgent := r.Header.Get("User-Agent")

	fmt.Fprintf(w,
		"今の時刻は**%s**で、利用しているブラウザは「%s」ですね。",
		currentTime,
		userAgent,
	)
}
