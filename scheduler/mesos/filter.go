package mesos

import (
	mesos "github.com/mesos/mesos-go/mesosproto"
	util "github.com/mesos/mesos-go/mesosutil"
	"github.com/tangmingdong123/mongodb-mesos/scheduler/repo"
)

func IsRunning(status *mesos.TaskStatus) bool {
	if status.GetState() == mesos.TaskState_TASK_LOST ||
		status.GetState() == mesos.TaskState_TASK_KILLED ||
		status.GetState() == mesos.TaskState_TASK_FINISHED ||
		status.GetState() == mesos.TaskState_TASK_ERROR ||
		status.GetState() == mesos.TaskState_TASK_FAILED {
		return false
	} else {
		return true
	}
}

func isMatch(db *repo.DBNode, offers []*mesos.Offer, usedMap map[*mesos.Offer]*Used) *mesos.Offer {
	for _, offer := range offers {
		summary := sum(offer)
		merge(summary,usedMap[offer])
		
		if float64(db.Cpu) <=summary.Cpu && float64(db.Memory) <= summary.Mem {
			return offer
		}
	}

	return nil
}

func merge(summary *Summary,used *Used){
	summary.Cpu = summary.Cpu - used.Cpu
	summary.Mem = summary.Mem - used.Mem
}

func sum(offer *mesos.Offer) *Summary {
	//mem
	memResources := util.FilterResources(offer.Resources, func(res *mesos.Resource) bool {
		return res.GetName() == "mem"
	})
	mems := 0.0
	for _, res := range memResources {
		mems += res.GetScalar().GetValue()
	}

	//cpu
	cpuResources := util.FilterResources(offer.Resources, func(res *mesos.Resource) bool {
		return res.GetName() == "cpus"
	})
	cpus := 0.0
	for _, res := range cpuResources {
		cpus += res.GetScalar().GetValue()
	}

	//ports
	portsResources := util.FilterResources(offer.Resources, func(res *mesos.Resource) bool {
		return res.GetName() == "ports"
	})
	//var ports uint64
	var portArr []*mesos.Value_Range
	for _, res := range portsResources {
		portRanges := res.GetRanges().GetRange()
		for _,rg := range portRanges {
			portArr = append(portArr,rg)
		}
	}
	
	return &Summary{Cpu:cpus,Mem:mems,PortRanges:portArr}
}
