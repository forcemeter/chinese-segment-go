package main

import (
	"encoding/json"
	"fmt"
	"github.com/forcemeter/chinese-segment-go/segment"
	"net/http"
	"strings"
)

var sgJieba = segment.Jieba{}
var sgGse = segment.Gse{}

func segmentHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": r,
			})
		}
	}()

	w.Header().Set("Content-Type", "application/json")

	var dict = r.URL.Query().Get("dict")
	var text = r.URL.Query().Get("text")
	var use = r.URL.Query().Get("use")

	var sg segment.Segment
	sg = &sgJieba

	if use == "gse" {
		sg = &sgGse
	}

	if dict != "" {
		sg.LoadDictWords(strings.Split(dict, ","))
	}

	var result = segment.Tokens{}

	if text != "" {
		fmt.Println(text)
		result = sg.Cut(string(text))
	}

	json.NewEncoder(w).Encode(result)
}

func main() {
	sgJieba.Init()
	sgGse.Init()

	http.HandleFunc("/", segmentHandler)

	fmt.Println("Server is running on port 38080...")
	err := http.ListenAndServe(":38080", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
