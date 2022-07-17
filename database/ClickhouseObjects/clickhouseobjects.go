package clickhouseobjects

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go"
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
	fmt.Println("Using LocalObjects db...")
	dbConnect()
}
func dbConnect() {
	var err error
	connect, err = sql.Open("clickhouse", "tcp://127.0.0.1:19000?debug=true")
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
	type Sample struct {
		Object_id string
		Timestamp string
		Value     float32
	}
	_, err = connect.Exec(`
		CREATE TABLE IF NOT EXISTS example (
			object_id    String,
			timestamp    String,
			value        Float32
		)
		ENGINE = MergeTree()
		ORDER BY (object_id, timestamp)
		
	`)

	if err != nil {
		log.Fatal(err)
	}
}

var GetObject = func(object_id string, starttime time.Time, endtime time.Time, sampleCount int) (object app.Object, err error) {
	var result = app.Object{}
	return result, nil
}
var AddSample = func(object_id string, value float32) (err error) {
	var (
		tx, _   = connect.Begin()
		stmt, _ = tx.Prepare("INSERT INTO example (object_id,timestamp,value) VALUES (?, ?, ?)")
	)
	if _, err := stmt.Exec(
		object_id,
		fmt.Sprint(time.Now()),
		value,
	); err != nil {
		return err
	}
	defer stmt.Close()

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
var AddObject = func(obj app.Object) (err error) {
	return nil
}
var GetObjects = func(object_id string) (result []app.Object) {

	type Sample struct {
		Object_id string
		Timestamp string
		Value     float32
	}

	var m = make(map[string]*app.Object)

	rows, err := connect.Query("SELECT object_id,timestamp,value FROM example WHERE object_id LIKE '" + object_id + "/%'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

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
		stime, _ := time.Parse("2006-01-02 15:04:05.9999", s.Timestamp)
		var newSample = app.Sample{Timestamp: stime, Value: s.Value}
		m[s.Object_id].Samples = append(m[s.Object_id].Samples, newSample)
		m[s.Object_id].LastSample = newSample
		log.Printf("object_id: %s,timestamp: %s,value: %.2f,", s.Object_id, s.Timestamp, s.Value)
	}
	for _, val := range m {
		result = append(result, *val)
	}
	return result
}
