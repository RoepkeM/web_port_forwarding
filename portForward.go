package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func handler(w http.ResponseWriter,r *http.Request){
	webName := r.Host
	//webIpClient := r.RemoteAddr
	newUrl:=checkWeb(webName)
	w.Header().Set("Location",newUrl)
	w.WriteHeader(http.StatusSeeOther)
}

func checkWeb(webName string) string {

	jsonFile, err := os.OpenFile("webs.json",os.O_RDONLY,0755)
	if err != nil{
		fmt.Println(err)
	}


	byteValue, _ := ioutil.ReadAll(jsonFile)
	var webs Webs
	var newPort string
	var wName string


	json.Unmarshal(byteValue,&webs)
	for  _,v :=range webs.Webs{
			if webName == v.Name || webName == v.Name2 || webName == v.Name3 {
				wName = v.Name
				newPort = v.Port
			}
	}
	newIpAndPort:="http://"+wName+":"+newPort
	defer jsonFile.Close()
	return newIpAndPort
}



