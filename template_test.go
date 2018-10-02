package main

import (
	"chaochaogege.com/onlinecode/global"
	"context"
	"html/template"
	"log"
	"os"
	"testing"
)

func Must(i interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return i
}

func TestTemplate(t *testing.T) {
	controller := global.NewController()
	hashKey := "b91bb55eff930a3ebd6454da0f11ec3a"
	res, _ := controller.GetSnippet(context.Background(), &global.Snippet{
		ID:  hashKey,
		Src: nil,
	})
	if r, ok := res.(string); !ok {

	} else {
		Config := &struct {
			CodesValue string `json:"codesvalue"`
		}{r}

		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		file := Must(os.Create("templates/result.html"))
		if f, ok := file.(*os.File); !ok {
			log.Println("err file must")
		} else {
			tmpl.Execute(f, Config)
		}

	}

}
