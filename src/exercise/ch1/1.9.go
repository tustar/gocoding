package main

import (
	"os"
	"net/http"
	"fmt"
	"io"
	"strings"
	"io/ioutil"
)

func main() {
	for _, url := range os.Args[1:] {
		s := "http://"
		if !strings.HasPrefix(url, s) {
			url = s + url
		}
		resp, err := http.Get(url)
		fmt.Printf("fetch: status code: %d\n", resp.StatusCode)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		n, err := io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", n)
	}
}