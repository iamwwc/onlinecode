package main

import (
	"bytes"
	"chaochaogege.com/onlinecode/global"
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)


const(
	MaxSnippetSize = 10000
)

type ShareRequest struct{
	Data string `json:"data"`
}

type ShareResponse struct{
	ID string `json:"id"`
	Error string `json:"error"'`
}

//{
//	"action":"share",
//	"details":[]ShareDetail
//}

func (s *server)handleShare(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Allow-Control-Allow-Origin","*")
	if r.Method == http.MethodOptions{
		return
	}

	if r.Method != http.MethodPost{
		report(w,"Requires POST")
		return
	}
	var buf bytes.Buffer
	io.Copy(&buf,io.LimitReader(r.Body,MaxSnippetSize + 1))

	if buf.Len() > MaxSnippetSize{
		report(w,"exceed max snippet length")
		return
	}
	snippet := &global.Snippet{
		Src:buf.Bytes(),
		ID:"",
	}

	id, err := s.db.PutSnippet(context.Background(),snippet)
	if err != nil{
		report(w,"error when share")
		return
	}
	res := ShareResponse{
		ID:id,
		Error:"",
	}
	var resBuf bytes.Buffer
	if err := json.NewEncoder(&resBuf).Encode(&res); err != nil{
		report(w,"error when share")
		return
	}
	count, err := io.Copy(w,&resBuf)
	if err != nil{
		logrus.Error(err)
	}
	logrus.Debugf("%d has been written",count)
}

func report(w http.ResponseWriter, res string){
	io.Copy(w,errorResponse(res))
}