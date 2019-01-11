package xserver

import (
	"errors"
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ServerHttp struct {
	Server *http.Server
	Router *gin.Engine
	conf   Config
}

func NewServerHttp(cfg Config) (s ServerHttp, err error) {
	if len(cfg.Address) == 0 {
		err = errors.New("address must be specified")
		return
	}
	return ServerHttp{conf: cfg}, nil
}

func (s *ServerHttp) Listen() (err error) {
	fmt.Println("server start")
	s.Server = &http.Server{
		Addr:         s.conf.Address,
		Handler:      s.Router,
		ReadTimeout:  time.Duration(s.conf.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(s.conf.WriteTimeout) * time.Second,
	}
	fmt.Println("Start Listening and serving HTTP on " + s.conf.Address)
	if err := gracehttp.Serve(s.Server); err != nil {
		return errors.New("failed during server listen and serve")
	}
	return
}

func (s *ServerHttp) Close() {
	if s.Server != nil {
		s.Server.Close()
	}
	fmt.Println("shutdown server successfully")
}
