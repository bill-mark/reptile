package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	encoding2 "golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main(){
   resp,err := http.Get("http://www.zhenai.com/zhenghun")
   if err != nil{
   	panic(err)
   }
   defer resp.Body.Close()

   if resp.StatusCode != http.StatusOK{
   	  fmt.Println("error:status code",resp.StatusCode)
   	  return
   }

   e := determineEncoding(resp.Body)
   utf8Reader :=  transform.NewReader(resp.Body,e.NewDecoder())

	all,err := ioutil.ReadAll(utf8Reader)
	if err != nil {
	panic(err)
	}
	//fmt.Printf("%s\n",all)
	printCityList(all)
}

func determineEncoding(r io.Reader) encoding2.Encoding {
	bytes,err := bufio.NewReader(r).Peek(1024)
	if err != nil{
		panic(err)
	}
	e,_,_ := charset.DetermineEncoding(bytes,"")
	return e
}

func printCityList(contents []byte){
   re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
   matches := re.FindAllSubmatch(contents,-1)
   for _,m := range matches{
   	//  for _,subMatch := range m{
		//  fmt.Printf("%s",subMatch)
	  //}
	  fmt.Printf("city: %s,url: %s\n",m[2],m[1])
   	  fmt.Println()
   }
   fmt.Printf("matches found: %d\n",len(matches))
}