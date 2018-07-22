package parser

import (
	"regexp"
	"../../engine"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9+])" [^>]*>([^<]+)</a>`

func ParseCity(content []byte) engine.ParseResult{
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	for _, i := range  matches {
		name := string(i[2])
		result.Items = append(result.Items, "user" + string(i[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(i[1]),
				ParserFunc: func(c []byte) engine.ParseResult {
					return  ParseProfile(c, name)
				},
			})
	}
	return result
}