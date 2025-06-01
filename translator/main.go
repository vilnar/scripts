// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Translate uses the Google translate API from the command line to translate
// its arguments. By default it auto-detects the input language and translates
// to English.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var (
	target = flag.String("to", "uk", "destination language (two-letter code)")
	source = flag.String("from", "auto", "source language (two-letter code); auto-detected by default")
	text = flag.String("text", "", "text for translate")
)

type Response [][]string

func prepareQuery(text string) string {
	return strings.ReplaceAll(text, "\\n", "\n")
}

func main() {
	flag.Parse()
	v := make(url.Values)
	v.Set("tl", *target)
	if *source != "" {
		v.Set("sl", *source)
	}
	v.Set("dt", "t")
	v.Set("q", *text)

	query := prepareQuery(*text)
	v.Set("q", query)


	urlT := fmt.Sprintf("https://translate.googleapis.com/translate_a/t?client=p&%s", v.Encode())
	client := &http.Client{}
        req, err := http.NewRequest("GET", urlT, nil)
        if err != nil {
		log.Printf("Error: http.NewRequest")
                log.Fatalln(err)
        }
        req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36")

        res, err := client.Do(req)
	if err != nil {
		log.Printf("Error: client.Do")
		log.Fatal(err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error: res.Body read")
		log.Fatal(err)
	}
	// fmt.Printf("debug1 %+v\n", string(data))

	var r Response
	if err := json.Unmarshal(data, &r); err != nil {
		log.Printf("Error: response parse")
		log.Fatal(err)
	}
	trText := r[0][0]

	fmt.Printf("%s\n", trText)
}
