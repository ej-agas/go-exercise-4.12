package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type comic struct {
	Num        int    `json:"num"`
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "no file given\n")
		os.Exit(1)
	}

	fileName := os.Args[1]

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "no search term given\n")
		os.Exit(1)
	}

	var (
		items       []comic
		searchTerms []string
		input       io.ReadCloser
		count       int
		err         error
	)

	if input, err = os.Open(fileName); err != nil {
		fmt.Fprintf(os.Stderr, "could not open file: %s\n", err)
		os.Exit(1)
	}

	if err = json.NewDecoder(input).Decode(&items); err != nil {
		fmt.Fprintf(os.Stderr, "could not decode json: %s", err)
		os.Exit(1)
	}

	fmt.Printf("read %d comics\n", len(items))

	for _, t := range os.Args[2:] {
		searchTerms = append(searchTerms, strings.ToLower(t))
	}

outer:
	for _, item := range items {
		title := strings.ToLower(item.Title)
		transcript := strings.ToLower(item.Transcript)

		for _, searchTerm := range searchTerms {
			if !strings.Contains(title, searchTerm) && !strings.Contains(transcript, searchTerm) {
				continue outer
			}
		}

		fmt.Printf("https://xkcd.com/%d/ %s/%s/%s %q\n", item.Num, item.Month, item.Day, item.Year, item.Title)

		count++
	}

	fmt.Printf("found %d comic(s)\n", count)
}
