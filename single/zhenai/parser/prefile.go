package parser

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"log"
	"regexp"
	"reptile/single/engine"
	"reptile/single/model"
)

var re = regexp.MustCompile(`<script>window.__INITIAL_STATE__=(.+);\(function`)

func ParseProfile(contents []byte,name string)engine.ParseResult{
	profile := model.Profile{}
	profile.Name = name

	match := re.FindSubmatch(contents)
	if len(match) >=2{
		json := match[1]
		fmt.Printf("json:%s\n",json)
	}
	return engine.ParseResult{}

	//for _,m := range matches{
	//	//age,err := strconv.Atoi( extractString(m(1) ) )
	//	//if err != nil{
	//	//	profile.Age = age
	//	//}
	//	profile.Marriage = string(m[0])
    //   profile.Xinzuo = string(m[2])
    //   profile.Education = string(m[6])
    //   profile.Car = "ben chi 998"
	//}
	//result := engine.ParseResult{Items:[]interface{}{profile}}
	//return result
}

func extractString(contents []byte)string{
	if len(contents)>=2{
		return string(contents[1])
	}else{
		return ""
	}
}

func parseJson(json []byte)model.Profile{
	res,err := simplejson.NewJson(json)
	if err != nil{
		log.Println("解析json失败")
	}
	infos,err := res.Get("objectInfo").Get("basicInfo").Array()
}