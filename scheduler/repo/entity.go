package repo

//a mongodb node
type DBNode struct {
	Name string  `json:"name"`
	Cpu    float32 `json:"cpu"`
	Memory int32   `json:"memory"`
}

type RouterNode struct {
	Name string  `json:"name"`
	Cpu    float32 `json:"cpu"`
	Memory int32   `json:"memory"`
}

//a ReplicaSet ,with name and nodes
type ReplicaSet struct {
	RsName string   `json:"rsName"`
	Nodes  []DBNode `json:"nodes"`
}

//a Shard cluster, with name,configReplicaSet,routers,shards
type ShardCluster struct {
	Name     string       `json:"name"`
	Routers  []RouterNode `json:"routers"`
	ConfigRS ReplicaSet   `json:"configRS"`
	Shards   []ReplicaSet `json:"shards"`
}

type Meta struct {
	StandaloneMap   map[string]DBNode
	ReplicatSetMap  map[string]ReplicaSet
	ShardClusterMap map[string]ShardCluster
}
