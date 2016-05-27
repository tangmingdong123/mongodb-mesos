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
		Param(ws.PathParameter("id","Standalone Mongodb id")))
		
	ws.Route(ws.GET("/{id}").To(r.CreateStandalone).
		// docs
		Doc("getStandalone Info").
		Operation("getStandalone").
		Param(ws.PathParameter("id", "identifier of the standalone").DataType("string")).
		Consumes("text/plain"))
	
	return ws
}