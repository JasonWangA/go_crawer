package engine

import (
	"log"
	"../fetch"
)


type SimpleEngine struct {

}
func (e SimpleEngine)Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parserResult, err:= worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parserResult.Requests...)


		for _, item := range parserResult.Items {
			log.Printf("got item %v", item)
		}
	}
}


func worker(r Request)  (ParseResult, error){
	body, err := fetch.Fetch(r.Url)
	log.Println(r.Url)
	if err != nil {
		log.Printf("fetch: error fetching, url: %s: %v", r.Url, err)
       return  ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}