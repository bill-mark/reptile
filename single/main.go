package main

import (
	"reptile/single/engine"
	"reptile/single/zhenai/parser"
)

func main(){
   engine.Run(engine.Request{
   	  Url:"http://www.zhenai.com/zhenghun",
   	  ParserFunc:parser.ParseCityList,
   })
}
