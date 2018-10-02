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
	s.mux.HandleFunc(global.FromSharePath,s.FromShare)
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

	//为什么这样就不行？ 非法的dereference
	//var buf *bytes.Buffer
	//if err := json.NewEncoder(buf).Encode(&resp);err != nil{
	//	log.Errorf("Error when encode error response %v",err)
	//	//never happen
	//}
	//return buf

	//buf 只是一个pointer,他并没有指向合法的内存空间，而NewEncoder需要将编码之后的数据写入
	//所以写入失败
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
