package main

import(
	"linkernetworks.com/mongodb-mesos/scheduler/repo"
	"encoding/json"
	"fmt"
)

func main(){
	repo.InitZK([]string{"192.168.3.223"},"/mongodb-mesos")
	dbnode := repo.DBNode{Name:"tang",Cpu:0.1,Memory:512}
	
	bytes,_ := json.Marshal(&dbnode)
	
	fmt.Println(string(bytes))
	
	var dbnode2 repo.DBNode
	json.Unmarshal(bytes,&dbnode2)
	
	fmt.Println(dbnode2)
	
	
	//repo.SaveStandalone(&dbnode2)
}