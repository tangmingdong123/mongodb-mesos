package rest

import (
	"github.com/emicklei/go-restful"
	"fmt"
)

func (r *Router) CreateStandalone(req *restful.Request, resp *restful.Response){
	fmt.Println("invoke me")
	
	resp.AddHeader("Content-Type","text/plain")
	
	resp.Write([]byte("createStandalone."))
	//resp.WriteEntity("createStandalone.")
}