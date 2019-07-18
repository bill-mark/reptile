package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//contents,err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents,err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	//ParseCityList(contents)
	//fmt.Printf("%s\n",conte nts)

	result := ParseCityList(contents)
	//expectedUrls := []string{
	//	"","","",
	//}
	//expectedCities := []string{
	//	"","","",
	//}

	const resultSize = 470
	if len(result.Requests) != resultSize {
        t.Errorf("result should have %d"+"request;bot had %d",resultSize,len(result.Request))
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d"+"request;bot had %d",resultSize,len(result.Request))
	}
}