package fetch

import (
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
	"log"
	"time"
)

var rateLimiter = time.Tick(100 * time.Microsecond)
func Fetch(url string) ([]byte, error) {
	<- rateLimiter
	resp, err := http.Get(url)
	if err != nil {
		return  nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return  nil,
		   fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determingRncoding(bodyReader)
	utfEncoding := transform.NewReader(resp.Body,
		e.NewDecoder())
	return  ioutil.ReadAll(utfEncoding)
}

func determingRncoding (r *bufio.Reader) encoding.Encoding{
	bytes, err:= r.Peek(1024)
	if err != nil {
		log.Printf("fetch error %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")

	return  e
}
