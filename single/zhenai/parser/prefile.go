package parser

import (
	"github.com/bitly/go-simplejson"
	"log"
	"regexp"
	"reptile/single/engine"
	"reptile/single/model"
)

var re = regexp.MustCompile(`<script>window.__INITIAL_STATE__=(.+);\(function`)

func ParseProfile(contents []byte,name string)engine.ParseResult{
	match := re.FindSubmatch(contents)
	result := engine.ParseResult{}

	if len(match) >=2{
		json := match[1]
		profile := parseJson(json)
		profile.Name = name
		result.Items = append(result.Items,profile)
	}
	return result
}

func parseJson(json []byte)model.Profile{
	res,err := simplejson.NewJson(json)
	if err != nil{
		log.Println("解析json失败")
	}
	infos,err := res.Get("objectInfo").Get("basicInfo").Array()
	//infos是一个切片,类似是interface

	var profile  model.Profile
	for k,v := range infos {
        //fmt.Printf("k:%v,%T\n",k,k)
		//fmt.Printf("v:%v,%T\n", v, v)

        if e,ok := v.(string);ok{
        	switch k {
			case 0:
				profile.Marriage = e
			case 1:
			//年龄:47岁，我们可以设置int类型，所以可以通过另一个json字段来获取
			   profile.Age = e
			case 2:
			  profile.Xingzuo = e
			case 3:
			  profile.Height = e
			case 4:
			   profile.Weight = e
			case 6:
			   profile.Income = e
			case 7:
			   profile.Occupation = e
			case 8:
			   profile.Education = e
			}
		}
	}
    return profile
}