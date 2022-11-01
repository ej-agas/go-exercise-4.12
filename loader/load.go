package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	var (
		output    io.WriteCloser = os.Stdout
		err       error
		count     int
		failCount int
		data      []byte
	)

	if len(os.Args) > 0 {
		output = createOutputFile(os.Args[1])
		defer output.Close()
	}

	fmt.Fprint(output, "[")
	defer fmt.Fprint(output, "]")

	for i := 1; failCount < 2; i++ {
		if data = getOneComic(i); data == nil {
			failCount++
			continue
		}

		if count > 0 {
			fmt.Fprint(output, ",")
		}

		if _, err = io.Copy(output, bytes.NewBuffer(data)); err != nil {
			fmt.Fprintf(os.Stderr, "stopped: %s", err)
			os.Exit(1)
		}

		failCount = 0
		count++
	}

	fmt.Fprintf(os.Stderr, "read %d comics\n", count)
}

func createOutputFile(name string) *os.File {
	file, err := os.Create(name)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating file: %s", err)
		os.Exit(1)
	}

	return file
}

func getOneComic(i int) []byte {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)

	res, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error connecting to xkcd.com: %s\n", err)
		os.Exit(1)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "skipping Comic #%d, got HTTP code: %d\n", i, res.StatusCode)
		return nil
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "could not read xkcd.com response body: %s", err)
		os.Exit(1)
	}

	return body
}
