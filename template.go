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

	temp:= template.Must(template.ParseFiles("./templates/index.html"))
	temp.Execute(w,c)
}
