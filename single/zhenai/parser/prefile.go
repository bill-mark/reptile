package parser

import (
	"regexp"
	"reptile/single/engine"
	"reptile/single/model"
)

const basic = `basicInfo"/["(*)","(*)","(*)","(*)","(*)","(*)","(*)"\]`

func ParseProfile(contents []byte,name string)engine.ParseResult{
	profile := model.Profile{}
	profile.Name = name

	re := regexp.MustCompile(basic)
	matches := re.FindAllSubmatch(contents,-1)

	for _,m := range matches{
		//age,err := strconv.Atoi( extractString(m(1) ) )
		//if err != nil{
		//	profile.Age = age
		//}

		profile.Marriage = string(m[0])
        profile.Xinzuo = string(m[2])
        profile.Education = string(m[6])
	}
	result := engine.ParseResult{Items:[]interface{}{profile}}
	return result
}

func extractString(contents []byte)string{
	if len(contents)>=2{
		return string(contents[1])
	}else{
		return ""
	}
}