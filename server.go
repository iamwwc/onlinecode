package main

import (
	"bytes"
	"chaochaogege.com/onlinecode/global"
	"context"
	"encoding/base64"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	MaxBuildAndRunTimes = time.Second * 10
)

type server struct {
	mux        *http.ServeMux
	controller *global.DockerController
	db *global.DatabaseController
}

func newServer(options ...func(*server) error) (*server, error) {

	s := &server{
		mux: http.NewServeMux(),
		db: global.NewController(),
	}
	for _, option := range options {
		option(s)
	}
	s.init()
	return s, nil
}

func (s *server) init() {
	s.mux.HandleFunc(global.CompilePath, s.commandHandler(compileAndRun))
	s.mux.HandleFunc(global.SharePath,s.handleShare)
	s.mux.HandleFunc("/",s.handleFromIndex)
	staticHandler := http.StripPrefix("/static/",http.FileServer(http.Dir("./static/")))
	s.mux.Handle("/static/", staticHandler)
}

func (s *server) commandHandler(cmd func(*global.Request) (*global.Response, error)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Allow-Control-Allow-Origin", "*")
			return
		}

		var req global.Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			resp := errorResponse("Error when decode request\n")
			log.Errorf("error when decode request %v", err)
			io.Copy(w, resp)
			return
		}

		for _, v := range req.Files {
			raw, err := base64.StdEncoding.DecodeString(v["src"])
			if err != nil {
				res := errorResponse("error when decode request\n")
				io.Copy(w, res)
				return
			}
			v["src"] = string(raw)
		}

		response, err := cmd(&req)
		if err != nil {

		}
		res := encodeResponse(response)
		io.Copy(w, res)
	}
}

func (s *server)handleFromIndex(w http.ResponseWriter, r *http.Request){
	if strings.HasPrefix(r.URL.Path, "/share/"){
		id := r.URL.Path[7:]
		snippet := &global.Snippet{
			Src:nil,
			ID:id,
		}
		res, err :=s.db.GetSnippet(context.Background(),snippet)
		if err != nil{
			log.Error(err)
			report(w,"error when get Snippet")
			return
		}

		if r, ok := res.(string); !ok{
			log.Errorf("error res is not string")
			report(w,"error res is not string")
			return
		}else{
			GenerateShareTemplate(w,r)
			return
		}
	}
	w.Header().Set("Content-Type","text/html")
	GenerateShareTemplate(w,"")
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func encodeResponse(r *global.Response) *bytes.Buffer {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(r); err != nil {
		return errorResponse("Error when encode response\n")
	}
	return &buf
}

func errorResponse(msg string) *bytes.Buffer {
	resp := &global.Response{
		Error:  msg,
		Events: nil,
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(&resp); err != nil {
		log.Errorf("Error when encode error response %v", err)
		//never happen
	}
	return &buf
}

func compileAndRun(r *global.Request) (*global.Response, error) {
	tempDir, err := ioutil.TempDir(global.VolumePath, "sandbox")
	//defer os.RemoveAll(tempDir)
	if err != nil {
		log.Errorf("Error when create temp dir %v", err)
		return &global.Response{
			Error:  "Error when build",
			Events: nil,
		}, nil
	}

	if _, err := FileWritter(tempDir, r.Files); err != nil {
		log.Errorf("Error when write to file %v", err)
		return &global.Response{
			Error:  "Error when build",
			Events: nil,
		}, nil
	}
	c := &global.DockerController{
		WorkingDir: tempDir,
	}
	runner := langFactory(r.Lang)
	ctx, cancel := context.WithTimeout(context.Background(), MaxBuildAndRunTimes)
	defer cancel()
	ctx, imageName, cmd := runner.Run(ctx)
	return c.SubmitContainer(ctx, cmd, imageName), nil
}
