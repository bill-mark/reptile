package parser

import (
	"regexp"
	"reptile/single/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents,-1)

	result := engine.ParseResult{}
	limit := 5
	for _,m := range matches{
		result.Items = append(result.Items,"City "+string(m[2]))
		result.Requests = append(
			result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc:ParseCity,
		})
		limit--
		if(limit == 0){break}
	}
	//fmt.Printf("matches found: %d\n",len(matches))
	return result
}