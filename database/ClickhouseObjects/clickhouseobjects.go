package clickhouseobjects

import (
	"fmt"
	"log"
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

var GetObject = func(object_id string, starttime time.Time, endtime time.Time, sampleCount int) (object app.Object, err error) {
	var result = app.Object{}
	return result, nil
}
var AddSample = func(object_id string, value float32) (err error) {
	buffer := Buffers.GetBuffer(object_id)
	buffer.AddSample(object_id, app.Sample{Timestamp: time.Now().UnixNano(), Value: value})
	return nil
}

var AddObject = func(obj app.Object) (err error) {
	return nil
}
var GetObjects = func(object_id string) (result []app.Object) {
	log.Printf("[debug][clickhouse][GetObjects]...%s", object_id)
	return Buffers.GetBuffer(object_id).GetObjects()
}
