package mesos

import(
	//log "github.com/Sirupsen/logrus"
	"github.com/gogo/protobuf/proto"
	mesos "github.com/mesos/mesos-go/mesosproto"
	//sched "github.com/mesos/mesos-go/scheduler"
	util "github.com/mesos/mesos-go/mesosutil"
	//"time"
	//"encoding/json"
	"github.com/tangmingdong123/mongodb-mesos/scheduler/repo"
	//"strconv"
)

func genStandaloneTask(db *repo.DBNode,offer *mesos.Offer)*mesos.TaskInfo{
	taskID := &mesos.TaskID{
				Value: proto.String(PREFIX_TASK_STANDALONE + db.Name),
			}
	taskType := mesos.ContainerInfo_DOCKER
	task := &mesos.TaskInfo{
				Name:    proto.String(PREFIX_TASK_STANDALONE + db.Name),
				TaskId:  taskID,
				SlaveId: offer.SlaveId,
				Container: &mesos.ContainerInfo{
					Type: &taskType,
					Docker: &mesos.ContainerInfo_DockerInfo{
						Image: proto.String("busybox"),
					},
				},
				Command: &mesos.CommandInfo{
					Shell:     proto.Bool(false),
					Arguments: []string{},
				},
				Resources: []*mesos.Resource{
					util.NewScalarResource("cpus", float64(db.Cpu)),
					util.NewScalarResource("mem", float64(db.Memory)),
				},
			}
	return task
}