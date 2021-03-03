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
	http.Redirect(w, r, newUrl, http.StatusSeeOther)
}

func checkWeb(webName string) string {

	jsonFile, err := os.OpenFile("webs.json",os.O_RDONLY,0755)
	if err != nil{
		fmt.Println(err)
	}


	byteValue, _ := ioutil.ReadAll(jsonFile)
	var webs Webs
	var newPort int


	json.Unmarshal(byteValue,&webs)
	for  _,v :=range webs.Webs{
			if webName == v.Name{
				newPort = v.Port
			}
	}
	newIpAndPort:=webName+":"+string(newPort)
	defer jsonFile.Close()
	return newIpAndPort
}



