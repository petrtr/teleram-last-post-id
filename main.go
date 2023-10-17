package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"time"
)

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "pass the channel name in the argument\n")
		os.Exit(1)
	}

	channel := args[0]
	requestURL := fmt.Sprintf("https://t.me/s/%s", channel)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Get(requestURL)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, fmt.Sprintf("error making http request: %s\n", err))
		os.Exit(1)
	}

	if res.StatusCode != http.StatusOK {
		_, _ = fmt.Fprintf(os.Stderr, fmt.Sprintf("invalid status code: %d\n", res.StatusCode))
		os.Exit(1)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, fmt.Sprintf("could not read response body: %s\n", err))
		os.Exit(1)
	}

	re, _ := regexp.Compile(fmt.Sprintf(`data-post="%s/(\d+)"`, channel))
	matched := re.FindAllStringSubmatch(string(resBody), -1)
	if len(matched) == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "not found post id\n")
		os.Exit(1)
	}
	last := matched[len(matched)-1]

	fmt.Print(last[1])
}
