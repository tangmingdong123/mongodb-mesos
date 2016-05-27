package mesos

import(
	mesos "github.com/mesos/mesos-go/mesosproto"
)

type Summary struct {
	Cpu   float64
	Mem   float64
	PortRanges []*mesos.Value_Range
}

type Used struct{
	Cpu   float64
	Mem   float64
	Ports []uint64
}
