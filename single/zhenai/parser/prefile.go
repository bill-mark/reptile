package parser

import (
	"regexp"
	"reptile/single/engine"
	"strconv"
)

const ageRe =`<td><span>年龄</span></td>`

func ParseProfile(contents []byte)engine.ParseResult{
	re := regexp.MustCompile(ageRe)
	match := re.FindSubmatch(contents)

	if match != nil{
		age,err := strconv.Atoi(string(match[1]))
		if err != nil{

		}
	}
}