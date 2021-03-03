package main


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