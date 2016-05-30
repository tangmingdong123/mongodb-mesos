package repo

const STATE_INIT = 0
const STATE_DEPLOYING = 1
const STATE_RUNNING = 2
const STATE_END = 3

//a mongodb node
type DBNode struct {
	Name string  `json:"name"`
	Cpu    float32 `json:"cpu"`
	Memory int32   `json:"memory"`
	State int32 `json:"state"`
}

type RouterNode struct {
	Name string  `json:"name"`
	Cpu    float32 `json:"cpu"`
	Memory int32   `json:"memory"`
	State int32 `json:"state"`
}

//a ReplicaSet ,with name and nodes
type ReplicaSet struct {
	RsName string   `json:"rsName"`
	Nodes  []DBNode `json:"nodes"`
	State int32 `json:"state"`
}

//a Shard cluster, with name,configReplicaSet,routers,shards
type ShardCluster struct {
	Name     string       `json:"name"`
	Routers  []RouterNode `json:"routers"`
	ConfigRS ReplicaSet   `json:"configRS"`
	Shards   []ReplicaSet `json:"shards"`
	State int32 `json:"state"`
}

type Meta struct {
	StandaloneMap   map[string]*DBNode
	ReplicatSetMap  map[string]*ReplicaSet
	ShardClusterMap map[string]*ShardCluster
}
