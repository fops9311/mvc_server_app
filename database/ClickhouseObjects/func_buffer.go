package clickhouseobjects

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/fops9311/mvc_server_app/app"
)

type objectBuffer struct {
	sync.Mutex
	dbSamples
	sampleInDb []bool
	expired    chan bool
}
type dbSample struct {
	object_id string
	timestamp int64
	value     float32
}
type dbSamples []dbSample

func (ob *objectBuffer) Init(object_id string) {
	var s1 rand.Source = rand.NewSource(time.Now().UnixNano())
	var r1 *rand.Rand = rand.New(s1)
	var timeToLive int = 20000 + r1.Intn(2000)

	ob.dbSamples = make(dbSamples, 0)
	ob.dbSamples = append(ob.dbSamples, readDbSamples(object_id)...)
	ob.sampleInDb = make([]bool, 0)
	for range ob.dbSamples {
		ob.sampleInDb = append(ob.sampleInDb, true)
	}
	ob.expired = make(chan bool)
	ob.Lock()
	defer ob.Unlock()
	dbDumpDemon := func() {
		log.Printf("[debug][*objectBuffer Init][dbDumpDemon] demon started id:%s\n", object_id)
		<-time.NewTimer(time.Duration(timeToLive) * time.Millisecond).C
		ob.Lock()
		log.Printf("[debug][*objectBuffer Init][dbDumpDemon] dump started id:%s\n", object_id)
		samples := make(dbSamples, 0)
		log.Printf("[debug][*objectBuffer Init][dbDumpDemon] ob.sampleInDb len=%d ob.dbSamples=%d\n", len(ob.sampleInDb), len(ob.dbSamples))
		for i, v := range ob.sampleInDb {
			if !v {
				samples = append(samples, ob.dbSamples[i])
				ob.sampleInDb[i] = true
			}
		}

		ob.expired <- true
		err := dbWtireSamples(samples)
		if err != nil {
			log.Printf("[debug][dbDumpDemon][error] %s", err)
		}
		ob.Unlock()
		log.Printf("[debug][*objectBuffer Init][dbDumpDemon] dumped %d samples, object_id:%s\n", len(samples), object_id)
	}

	go dbDumpDemon()

}
func (ob *objectBuffer) AddSample(object_id string, sample app.Sample) {
	ob.Lock()
	defer ob.Unlock()
	ob.dbSamples = append(ob.dbSamples, dbSample{object_id: object_id, timestamp: sample.Timestamp, value: sample.Value})
	ob.sampleInDb = append(ob.sampleInDb, false)
	log.Printf("[debug][*objectBuffer AddSample] Sample added ob.sampleInDb len=%d ob.dbSamples=%d\n", len(ob.sampleInDb), len(ob.dbSamples))
}

func (ob *objectBuffer) GetObjects() (result []app.Object) {
	ob.Lock()
	defer ob.Unlock()

	var m = make(map[string]*app.Object)

	for _, sample := range ob.dbSamples {
		if len(sample.object_id) > 0 {
			sample.object_id = sample.object_id[:len(sample.object_id)-1]
		}
		if _, ok := m[sample.object_id]; !ok {
			m[sample.object_id] = &app.Object{
				Id:      sample.object_id,
				Samples: make([]app.Sample, 0),
			}
		}
		var newSample = app.Sample{Timestamp: sample.timestamp, Value: sample.value}
		m[sample.object_id].Samples = append(m[sample.object_id].Samples, newSample)
		m[sample.object_id].LastSample = newSample
	}
	for _, val := range m {
		result = append(result, *val)
	}
	log.Printf("[debug][clickhouse][GetObjects]result lens %d\n", len(result))

	return result
}

type objectBuffers struct {
	sync.Mutex
	mp map[string]*objectBuffer
}

var Buffers objectBuffers = objectBuffers{mp: make(map[string]*objectBuffer)}

func (buffs *objectBuffers) GetBuffer(buffer_id string) (result *objectBuffer) {
	buffs.Lock()
	defer buffs.Unlock()
	if _, ok := buffs.mp[buffer_id]; !ok {
		newBuf := &objectBuffer{}
		newBuf.Init(buffer_id)
		buffs.mp[buffer_id] = newBuf
		log.Printf("[debug][GetBuffer] New buffer inited %s", buffer_id)
		go func() {
			<-newBuf.expired
			delete(buffs.mp, buffer_id)
			log.Printf("[debug][BufferDeleteDemon] buffer %s deleted", buffer_id)
		}()
		return newBuf
	}
	return buffs.mp[buffer_id]
}
