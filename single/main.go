package main

import "net/http"

func main(){
   resp,err := http.Get("http://www.zhenai.com/zhenghun")
   if err != nil{
   	panic(err)
   }
   defer resp.Body.Close()

}
