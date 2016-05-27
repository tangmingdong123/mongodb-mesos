package repo

import(
	"encoding/json"
)

func IsStandaloneExist(name string)bool{
	_,ok := meta.StandaloneMap[name]
	return ok
}

func AddStandalone(db *DBNode){
	meta.StandaloneMap[db.Name] = db
	
	SaveStandalone(db)
}
func StandaloneJson()([]byte,error){
	return json.Marshal(&meta.StandaloneMap)
}
func FindStandalone(name string)(*DBNode){
	return meta.StandaloneMap[name]
}
func ListStandalone()[]*DBNode{
	arr := make([]*DBNode,len(meta.StandaloneMap))
	
	i := 0
	for name := range meta.StandaloneMap {
		arr[i] = meta.StandaloneMap[name]
		i = i + 1
	}
	return arr
}