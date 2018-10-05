package global

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"testing"
)

var db , _ = sql.Open("mysql","root:wxlwuweichao@/onlinecode")


var controller = &DatabaseController{
	db:db,
}


func TestStart(t *testing.T){

	snippet := &Snippet{
		Src:[]byte("iamsrc"),
		ID:"",
	}
	ctx := context.Background()
	if _, err := controller.PutSnippet(ctx,snippet); err != nil{
		logrus.Error(err)
	}

	res, err := controller.GetSnippet(ctx,snippet)
	if err != nil{
		logrus.Errorf("error when get Snippet %v",err)
	}

	if v, ok:= res.(string); !ok{
		t.Error("fetch result failed")
	}else{
		if v != string(snippet.Src){
			t.Error("result not equal")
		}
	}
}

func Must(s string, err error)string{
	if err != nil{
		panic(err)
	}

	return s
}

func TestController(t *testing.T){
	_ = filepath.Dir(Must(os.Getwd()))
}
