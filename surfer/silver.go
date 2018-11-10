package surfer

import (
	"encoding/json"
	"sync"
)

type AddReplicaEvent struct {
	collection string            `json:"collection"`
	tags       []string          `json:"tags"`
	location   Machine           `json:"location"`
	metrics    map[string]Metric `json:"metrics"`
}

type Metric struct {
	key   string  `json:"key"`
	value float32 `json:"value"`
	info  string  `json:"info"`
}

type Machine struct {
	hostName     string `json:"hostName"`
	instanceType string `json:"instanceType"`
	url          string `json:"url"`
}

type SSurfer struct {
	data map[Machine]MachineInfo
	lock sync.Mutex
}

type MachineInfo struct {
	replicas map[string]ReplicaInfo `json:"collections"`
}

type ReplicaInfo struct {
	docCount   int         `json:"docCount"`
	sizeFactor int         `json:"sizeFactor"`
	fields     []FieldInfo `json:"fields"`
}

type FieldInfo struct {
	tokenFactor int
}

const (
	addReplica string = "addReplica"
)

type GalactusEvent struct {
	eventType string `json:"eventType"`
	PayLoad   string `json:"payLoad"`
}

func New() SSurfer {
	return SSurfer{
		data: map[Machine]MachineInfo{},
		lock: sync.Mutex{},
	}
}
func (ss SSurfer) OnEvent(msg []byte) error {
	event := GalactusEvent{}
	err := json.Unmarshal(msg, &event)
	if err != nil {
		return err
	}
	switch event.eventType {
	case addReplica:
		ss.AddReplica(event)
	}
	return nil
}
func (ss SSurfer) AddReplica(event GalactusEvent) error {
	add := AddReplicaEvent{}
	err := json.Unmarshal([]byte(event.PayLoad), add)
	if err != nil {
		return err
	}
	rinfo := mapAddEvent(add)
	ss.lock.Lock()
	if _, ok := ss.data[add.location]; ok {
		ss.data[add.location].replicas[add.collection] = rinfo
	} else {
		rep := make(map[string]ReplicaInfo)
		rep[add.collection] = rinfo
		minfo := MachineInfo{
			replicas: rep,
		}
		ss.data[add.location] = minfo
	}
	ss.lock.Unlock()
	return nil
}

func mapAddEvent(event AddReplicaEvent) ReplicaInfo {
	//todo generalize this
	return ReplicaInfo{
		docCount:   int(event.metrics["docCount"].value),
		sizeFactor: int(event.metrics["fieldsUsage"].value),
	}
}
