package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	//"os"
	//"path/filepath"

	//"github.com/Sirupsen/logrus"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
	//"github.com/magiconair/properties"
	"github.com/tangmingdong123/mongodb-mesos/scheduler/rest"
)

/**
mongodb-mesos -mesos-zk=zk://localhost:2181/mesos -zk zk://localhost:2181/ -name mongodb-mesos 
*/
func main() {
	//parse args
	mesosZk := flag.String("mesos-zk","zk://master.mesos:2181/mesos","zk of mesos")
	zk := flag.String("zk","zk://localhost:2181","repo of mongodb-scheduler")
	name := flag.String("name","mongodb-mesos","framework's name")	
	port := flag.Int("port",37017,"framework's http port")
	
	flag.Parse()
	
	fmt.Println("mongodb-mesos scheduler start...")
	fmt.Printf("mongodb-mesos scheduler mesos-zk:%s,zk:%s,name:%s,port:%d\n",*mesosZk,*zk,*name,*port)
	
	//launch HTTP REST service
	launchHTTP(*port)
}

func launchHTTP(port int){
	fmt.Printf("mongodb-mesos framework listen on %d\n",port)
	
	// accept and respond in JSON unless told otherwise
	restful.DefaultRequestContentType(restful.MIME_JSON)
	restful.DefaultResponseContentType(restful.MIME_JSON)
	restful.DefaultContainer.EnableContentEncoding(true)
	restful.DefaultContainer.Router(restful.CurlyRouter{})
	restful.SetCacheReadEntity(false)
	rest.Register(restful.DefaultContainer, false)
	
	config := swagger.Config{
			WebServices:     restful.DefaultContainer.RegisteredWebServices(),
			WebServicesUrl:  "http://127.0.0.1:"+strconv.Itoa(port),
			ApiPath:         "/apidocs.json",
			SwaggerPath:     "/apidocs/",
			SwaggerFilePath: "d:/swagger-ui/dist",
		}
	
	swagger.RegisterSwaggerService(config,restful.DefaultContainer)

	// If swagger is not on `/` redirect to it
	http.HandleFunc("/", index)
		
	server := &http.Server{Addr: ":"+strconv.Itoa(port), Handler: restful.DefaultContainer}
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/apidocs/", http.StatusMovedPermanently)
}