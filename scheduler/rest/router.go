package rest

import (
	//"encoding/json"
	//"errors"

	"github.com/emicklei/go-restful"
	//"github.com/Sirupsen/logrus"
	//"github.com/compose/mejson"
	//"gopkg.in/mgo.v2/bson"
)




func (r *Router)StandaloneService()*restful.WebService{
	ws := new(restful.WebService)
	ws.Path("/standalone")
	ws.Consumes("*/*")
	ws.Produces(restful.MIME_JSON)
	
	ws.Route(ws.POST("/{id}").
		To(r.CreateStandalone).
		Doc("createStandalone").
		Operation("createStandalone").
		Param(ws.PathParameter("id","Standalone Mongodb id")).
		Param(ws.FormParameter("cpu","mongodb's cpu").DataType("float")).
		Param(ws.FormParameter("mem","mongodb's mem in MB").DataType("int32")))
		
	ws.Route(ws.GET("/{id}").To(r.CreateStandalone).
		// docs
		Doc("getStandalone Info").
		Operation("getStandalone").
		Param(ws.PathParameter("id", "identifier of the standalone").DataType("string")).
		Consumes("text/plain"))
	
	ws.Route(ws.GET("/list").To(r.ListStandalone).
		// docs
		Doc("getStandalone list").
		Operation("listStandalone").
		Consumes("application/json"))
	
	return ws
}