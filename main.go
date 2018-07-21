package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/net/html/charset"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
)

func main()  {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		e := determingRncoding(resp.Body)
		utfEncoding := transform.NewReader(resp.Body,
			e.NewDecoder())
		all, err := ioutil.ReadAll(utfEncoding)

		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n" ,all)
	}

}

func determingRncoding (r io.ReadCloser) encoding.Encoding{
   bytes, err:= bufio.NewReader(r).Peek(1024)
   if err != nil {
   	panic(err)
   }
   e, _, _ := charset.DetermineEncoding(bytes, "")

   return  e
}