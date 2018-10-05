package main

import (
	"html/template"
	"io"
)

func GenerateShareTemplate(w io.Writer, config string){
	c :=struct{
		CodesValue string `json:"data"`
	}{
		CodesValue:config,
	}


	//all index.html will be render by template engine
	//
	temp:= template.Must(template.ParseFiles("./index.html"))
	temp.Execute(w,c)
}
