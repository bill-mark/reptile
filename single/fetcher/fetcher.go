package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	encoding2 "golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte,error){
	request, _:=http.NewRequest(http.MethodGet,url,nil)
	request.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	resp,_:=http.DefaultClient.Do(request)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return nil,
		fmt.Errorf("wrong status code: %d",resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader :=  transform.NewReader(bodyReader,e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding2.Encoding {  //获得网站编码
	bytes,err := r.Peek(1024)
	if err != nil{
		//panic(err)
		log.Printf("Fetcher error 9999 : %v",err)
		return unicode.UTF8
	}
	e,_,_ := charset.DetermineEncoding(bytes,"")
	return e
}