package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func main()  {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		utfEncoding := transform.NewReader(resp.Body,
			simplifiedchinese.GBK.NewDecoder())
		all, err := ioutil.ReadAll(utfEncoding)

		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n" ,all)
	}

}