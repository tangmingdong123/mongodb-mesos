package mesos

import(
	log "github.com/Sirupsen/logrus"
	"github.com/gogo/protobuf/proto"
	sched "github.com/mesos/mesos-go/scheduler"
	mesos "github.com/mesos/mesos-go/mesosproto"
	//util "github.com/mesos/mesos-go/mesosutil"
	//"time"
)

type MongodbScheduler struct{
	
}

func Start(master *string){
	log.Infoln("startScheduler master:",*master)
	
	fwinfo := &mesos.FrameworkInfo{
		User: proto.String(""),
		Name: proto.String("mongodb-mesos"),
	}
	
	log.Infoln("startScheduler makeFrameworkInfo ok:")
	
	config := sched.DriverConfig{
		Scheduler:  newMongodbScheduler(),
		Framework:  fwinfo,
		Master:     *master,
	}
	
	log.Infoln("startScheduler makeDriverConfig ok:")
	
	driver, err := sched.NewMesosSchedulerDriver(config)
	if err != nil {
		log.Errorln("Unable to create a SchedulerDriver ", err.Error())
	}
	
	log.Infof("startScheduler makeDriver ok:%v",driver)
	
	stat, err := driver.Run()
	if err != nil {
		log.Infof("Framework stopped with status %s and error: %s", stat.String(), err.Error())
	}
	
	log.Infoln("startScheduler driver run ok:")
	log.Infof("stat:%v",stat)
}


func newMongodbScheduler()*MongodbScheduler{
	return &MongodbScheduler{}
}


func (sched *MongodbScheduler) Registered(driver sched.SchedulerDriver, frameworkId *mesos.FrameworkID, masterInfo *mesos.MasterInfo) {
	log.Infoln("Framework Registered with Master ", masterInfo)
}

func (sched *MongodbScheduler) Reregistered(driver sched.SchedulerDriver, masterInfo *mesos.MasterInfo) {
	log.Infoln("Framework Re-Registered with Master ", masterInfo)
}

func (sched *MongodbScheduler) Disconnected(sched.SchedulerDriver) {
	log.Warningf("disconnected from master")
}

func (sched *MongodbScheduler) ResourceOffers(driver sched.SchedulerDriver, offers []*mesos.Offer) {
	log.Warningf("Framework resourceOffer")
}

func (sched *MongodbScheduler) StatusUpdate(driver sched.SchedulerDriver, status *mesos.TaskStatus) {
	log.Infoln("Status update: task", status.TaskId.GetValue(), " is in state ", status.State.Enum().String())
}

func (sched *MongodbScheduler) OfferRescinded(_ sched.SchedulerDriver, oid *mesos.OfferID) {
	log.Errorf("offer rescinded: %v", oid)
}
func (sched *MongodbScheduler) FrameworkMessage(_ sched.SchedulerDriver, eid *mesos.ExecutorID, sid *mesos.SlaveID, msg string) {
	log.Errorf("framework message from executor %q slave %q: %q", eid, sid, msg)
}
func (sched *MongodbScheduler) SlaveLost(_ sched.SchedulerDriver, sid *mesos.SlaveID) {
	log.Errorf("slave lost: %v", sid)
}
func (sched *MongodbScheduler) ExecutorLost(_ sched.SchedulerDriver, eid *mesos.ExecutorID, sid *mesos.SlaveID, code int) {
	log.Errorf("executor %q lost on slave %q code %d", eid, sid, code)
}
func (sched *MongodbScheduler) Error(_ sched.SchedulerDriver, err string) {
	log.Errorf("Scheduler received error: %v", err)
}