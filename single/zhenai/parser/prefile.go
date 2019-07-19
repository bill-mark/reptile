package parser

import (
	"regexp"
	"reptile/single/engine"
	"reptile/single/model"
	"strconv"
)

var ageRe =regexp.MustCompile(`<td><span>年龄</span></td>`)
var marriageRe =regexp.MustCompile(`婚况`)

func ParseProfile(contents []byte,name string)engine.ParseResult{
	profile := model.Profile{}
	profile.Name = name

	age,err := strconv.Atoi(extractString(contents,ageRe))  //atoi 将字符串转换为int类型
	if err != nil{
		profile.Age = age
	}
	profile.Marriage = extractString(contents,marriageRe)

	result := engine.ParseResult{Items:[]interface{}{profile}}
	return result
}

func extractString(contents []byte,re *regexp.Regexp)string{
	match := marriageRe.FindSubmatch(contents)
	if len(match)>=2{
		return string(match[1])
	}else{
		return ""
	}
}