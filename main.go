package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Starting hello-world server...")
	http.HandleFunc("/", currentTime)
	http.HandleFunc("/local", localTime)
	http.HandleFunc("/utc", utcTime)
	http.HandleFunc("/info", info)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "simple timeserver")
}
func currentTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Local "+yyyy_mm_dd_hh_mm_ss_fff(time.Now()))
	fmt.Fprint(w, "\n")
	fmt.Fprint(w, "UTC   "+yyyy_mm_dd_hh_mm_ss_fff(time.Now().UTC()))
}
func utcTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "UTC   "+yyyy_mm_dd_hh_mm_ss_fff(time.Now().UTC()))
}
func localTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Local "+yyyy_mm_dd_hh_mm_ss_fff(time.Now()))
}

func yyyy_mm_dd_hh_mm_ss_fff(t time.Time) string {
	return fmt.Sprintf("%4d-%02d-%02d %02d:%02d:%02d.%03d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond()/1000)
}
