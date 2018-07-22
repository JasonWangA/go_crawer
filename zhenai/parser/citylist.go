package parser

import (
	"regexp"
	"../../engine"
)

const citylist = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(citylist)
	//matches := re.FindAll(contents, -1)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}

	for _, i := range  matches {
		result.Items = append(result.Items, "city" + string(i[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(i[1]),
				ParserFunc: ParseCity,
			})
	}
	return   result

}