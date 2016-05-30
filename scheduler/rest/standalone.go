package rest

import (
	"github.com/emicklei/go-restful"
	"github.com/tangmingdong123/mongodb-mesos/scheduler/repo"
	"strconv"
	log "github.com/Sirupsen/logrus"
)

func (r *Router) CreateStandalone(req *restful.Request, resp *restful.Response){
	resp.AddHeader("Content-Type","application/json")
	
	
	id := req.PathParameter("id")
	cpu,_ := strconv.ParseFloat(req.QueryParameter("cpu"),32)
	mem,_ := strconv.Atoi(req.QueryParameter("mem"))
	
	
	if(repo.IsStandaloneExist(id)){
		log.Errorf("standalone mongodb '%s' already exist.",id)
		r := &Response{Code:CODE_ALREADY_EXIST,Desc:"already exist"}
		bs,_ := r.Byte()
		resp.Write(bs)
		return
	}else{
		db := &repo.DBNode{Name:id,Cpu:float32(cpu),
			Memory:int32(mem),
			State:repo.STATE_INIT}
		
		repo.AddStandalone(db)
		
		bs,_ := repo.DBNodeJson(db)
		resp.Write(bs)
	}
}

func (r *Router) ListStandalone(req *restful.Request, resp *restful.Response){
	resp.AddHeader("Content-Type","application/json")
	
	bs,_ := repo.StandaloneJson()
	resp.Write(bs)
}