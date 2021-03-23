package main
/*
Port Forward provides a server program that redirects the client HTTP request for a web to the port of your desire,
so you can serve multiple HTTP request incoming to one port.

Note: This program does not work as a proxy. It tells the clients web browser to request your website on another port.

This code works with a json file in which you must specify the URL name of your website, the IP of your server and
the desire redirect port. Example:
{
  "webs": [
    {
      "name":"example.com",
      "name1":"www.example.com",
      "name2":"www.example.com",
      "ip": "127.0.0.1",
      "port": "9000"
    }
   ]
}

And as an extra this program generates a log file in which are recorded all the request to your websites, it stores the
IP and the time of the request.
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//This is the structure of the json wen decode.
//Webs is array of the substructure YourWeb, and web has as information the IP the port and the different names that your
//web could be search, like whit or without www.
type Webs struct {
	Webs []YourWeb "json:'webs'"
}

type YourWeb struct {
	Name string "json:'name'"
	Name2 string "json:'name2'"
	Name3 string "json:'name3'"
	Ip string "json:'ip'"
	Port string "json:'port'"
}

//The main function receives all the requests in the port :80
func main() {
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

//This function handles the request to and uses the checkWeb and logs functions
func handler(w http.ResponseWriter,r *http.Request){
	webName := r.Host
	urlGet:= r.URL.String()
	newUrl:=checkWeb(webName, urlGet)
	w.Header().Set("Location",newUrl)
	w.WriteHeader(http.StatusSeeOther)
	logs(r)
}

//checkWeb parses the webs.json file and searches the require website to match the needed port
func checkWeb(webName string,urlGet string) string {

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
	newIpAndPort:="http://"+wName+":"+newPort+urlGet
	defer jsonFile.Close()
	return newIpAndPort
}

//this function creates and updates a log file whit some basic information of the http requests
func logs(r *http.Request){
	fmt.Println("en logs2")
	f,err:=os.OpenFile("webLogs.log",os.O_CREATE|  os.O_APPEND|os.O_RDWR , 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	tmp:="\n-->Request for:"+r.Host+"from ip:"+r.RemoteAddr+" time:"+time.Now().Format(time.RFC3339)
	f.WriteString(tmp)
}


