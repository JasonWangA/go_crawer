package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main()  {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		all, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n" ,all)
	}

}