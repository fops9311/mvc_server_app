package clickhouseobjects

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/fops9311/mvc_server_app/app"
)

func DefineComponent() {
	app.Objects.AddObject = AddObject
	app.Objects.GetObjects = GetObjects
	app.Objects.GetObject = GetObject
	app.Objects.AddSample = AddSample
}

var connect *sql.DB

func init() {
	fmt.Println("[debug][clickhouse]Using LocalObjects db...")
	ver, err := version()
	if err == nil {
		log.Printf("[debug][clickhouse][dbConnect]Successfully connected version %s", ver)
	} else {
		fmt.Printf("[debug][clickhouse][error]connection error %s", err)
	}
}
func version() (string, error) {
	var (
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{"127.0.0.1:19000"},
			Auth: clickhouse.Auth{
				Database: "default",
				Username: "default",
				Password: "",
			},
		})
	)
	if err != nil {
		return "", err
	}
	v, err := conn.ServerVersion()
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

/*
func dbConnect() {
	var err error
	connect, err = clickhouse.Open("clickhouse", "tcp://127.0.0.1:19000") //?debug=true
	if err != nil {
		log.Fatal(err)
	}
	if err := connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return
	}
	_, err = connect.Exec(`
		CREATE TABLE IF NOT EXISTS example (
			object_id    String,
			timestamp    Int64,
			value        Float32
		)
		ENGINE = MergeTree()
		ORDER BY (object_id, timestamp)

	`)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[debug][clickhouse][dbConnect]Successfully connected")
}
*/

var GetObject = func(object_id string, starttime time.Time, endtime time.Time, sampleCount int) (object app.Object, err error) {
	var result = app.Object{}
	return result, nil
}
var AddSample = func(object_id string, value float32) (err error) {
	buffer := Buffers.GetBuffer(object_id)
	buffer.AddSample(app.Sample{Timestamp: time.Now().UnixNano(), Value: value})
	return nil
}

func dbWtireSamples(object_id string, samples []app.Sample) (err error) {
	log.Printf("[debug][clickhouse][dbWtireSamples]...%s", object_id)
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:19000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		//Debug:           true,
		DialTimeout:     time.Second,
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: time.Hour,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
	})
	if err != nil {
		return err
	}
	defer conn.Close()
	ctx := clickhouse.Context(context.Background(), clickhouse.WithSettings(clickhouse.Settings{
		"max_block_size": 10,
	}), clickhouse.WithProgress(func(p *clickhouse.Progress) {
		fmt.Println("progress: ", p)
	}), clickhouse.WithProfileInfo(func(p *clickhouse.ProfileInfo) {
		fmt.Println("profile info: ", p)
	}))
	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Catch exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return err
	}

	batch, err := conn.PrepareBatch(ctx, "INSERT INTO example (object_id,timestamp,value)")
	if err != nil {
		return err
	}
	for _, v := range samples {
		if err := batch.Append(object_id, v.Timestamp, v.Value); err != nil {
			return err
		}
	}
	if err := batch.Send(); err != nil {
		return err
	}

	log.Printf("[debug][clickhouse][dbWtireSamples]added %d samples", len(samples))

	return nil
}

var AddObject = func(obj app.Object) (err error) {
	return nil
}
var GetObjects = func(object_id string) (result []app.Object) {
	log.Printf("[debug][clickhouse][dbWtireSamples]...%s", object_id)
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:19000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		//Debug:           true,
		DialTimeout:     time.Second,
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: time.Hour,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
	})
	if err != nil {
		//return err
	}
	defer conn.Close()
	ctx := clickhouse.Context(context.Background(), clickhouse.WithSettings(clickhouse.Settings{
		"max_block_size": 10,
	}), clickhouse.WithProgress(func(p *clickhouse.Progress) {
		fmt.Println("progress: ", p)
	}), clickhouse.WithProfileInfo(func(p *clickhouse.ProfileInfo) {
		fmt.Println("profile info: ", p)
	}))
	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Catch exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		//return err
	}

	type Sample struct {
		Object_id string
		Timestamp int64
		Value     float32
	}

	var m = make(map[string]*app.Object)

	rows, err := conn.Query(ctx,
		"SELECT object_id,timestamp,value FROM example WHERE object_id LIKE ? ORDER BY object_id,timestamp DESC limit 2000 by object_id",
		object_id+"/%")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var timestampstring string
	for rows.Next() {
		var s Sample = Sample{}

		if err := rows.Scan(&s.Object_id, &s.Timestamp, &s.Value); err != nil {
			log.Fatal(err)
		}
		if _, ok := m[s.Object_id]; !ok {
			m[s.Object_id] = &app.Object{
				Id:      s.Object_id,
				Samples: make([]app.Sample, 0),
			}
		}
		//timestampstring = fmt.Sprint(time.Unix(0, s.Timestamp))
		//strs := strings.Split(timestampstring, " +")
		//var newSample = app.Sample{Timestamp: strs[0], Value: s.Value}
		var newSample = app.Sample{Timestamp: s.Timestamp, Value: s.Value}
		m[s.Object_id].Samples = append(m[s.Object_id].Samples, newSample)
		m[s.Object_id].LastSample = newSample
	}
	for _, val := range m {
		result = append(result, *val)
	}
	log.Printf("[debug][clickhouse][GetObjects]Last sample timestamp %s\n", timestampstring)
	return result
}

type objectBuffer struct {
	app.Object
	sync.Mutex
	sampleInDb []bool
	active     bool
}

func (ob *objectBuffer) Init(object_id string) {
	ob.Lock()
	defer ob.Unlock()
	ob.Id = object_id
	ob.Samples = make([]app.Sample, 0)
	ob.sampleInDb = make([]bool, 0)
	dbDumpDemon := func() {
		log.Printf("[debug][(ob *objectBuffer) Init()][dbDumpDemon] started id:%s\n", ob.Id)
		for {
			<-time.NewTimer(20 * time.Second).C
			ob.Lock()
			samples := make([]app.Sample, 0)
			for i, v := range ob.sampleInDb {
				if !v {
					samples = append(samples, ob.Samples[i])
					ob.sampleInDb[i] = true
				}
			}
			ob.Unlock()
			err := dbWtireSamples(ob.Id, samples)
			if err != nil {
				log.Printf("[debug][dbDumpDemon][error] %s", err)
			}
			log.Printf("[debug][(ob *objectBuffer) Init()][dbDumpDemon] dumped %d samples, object_id:%s\n", len(samples), ob.Id)
		}
	}
	go dbDumpDemon()

}
func (ob *objectBuffer) AddSample(sample app.Sample) {
	ob.Lock()
	defer ob.Unlock()
	ob.Samples = append(ob.Samples, sample)
	ob.sampleInDb = append(ob.sampleInDb, false)
	ob.LastSample = sample
}

type objectBuffers struct {
	sync.Mutex
	mp map[string]*objectBuffer
}

var Buffers objectBuffers = objectBuffers{mp: make(map[string]*objectBuffer)}

func (buffs objectBuffers) GetBuffer(buffer_id string) (result *objectBuffer) {
	buffs.Lock()
	defer buffs.Unlock()
	if _, ok := buffs.mp[buffer_id]; !ok {
		newBuf := &objectBuffer{}
		newBuf.Init(buffer_id)
		buffs.mp[buffer_id] = newBuf
		log.Printf("[debug][GetBuffer] New buffer inited %s", newBuf.Id)
		return newBuf
	}
	return buffs.mp[buffer_id]
}
