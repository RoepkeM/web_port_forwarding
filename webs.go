package main


type Webs struct {
	Webs []YourWeb "json:'webs'"
}

type YourWeb struct {
	Name string "json:'name'"
	Ip int "json:'ip'"
	Port int "json:'port'"
}