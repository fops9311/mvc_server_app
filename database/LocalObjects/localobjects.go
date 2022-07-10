package localobjects

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/fops9311/mvc_server_app/app"
)

func DefineComponent() {
	app.Objects.AddObject = AddObject
	app.Objects.GetObjects = GetObjects
	app.Objects.GetObject = GetObject
	app.Objects.AddSample = AddSample
}

func init() {
	fmt.Println("Using LocalObjects db...")
	readUsersFile()
}
func readUsersFile() {
	usersfile, err := os.ReadFile("./objects.json")
	if err != nil {
		return
	}
	json.Unmarshal(usersfile, &Objects)
}
func writeUsersFile() {
	b, err := json.MarshalIndent(Objects, "", "	")
	if err != nil {
		return
	}
	os.WriteFile("objects.json", b, 0666)
}

var Objects map[string]*app.Object = map[string]*app.Object{}
var m sync.Mutex

var GetObject = func(object_id string, starttime time.Time, endtime time.Time, sampleCount int) (object app.Object, err error) {
	m.Lock()
	defer m.Unlock()
	return *Objects[object_id], nil
}
var AddSample = func(object_id string, value float32) (err error) {
	m.Lock()
	defer writeUsersFile()
	defer m.Unlock()
	if _, ok := Objects[object_id]; !ok {
		Objects[object_id] = &app.Object{
			Id:      object_id,
			Samples: make([]app.Sample, 0),
		}
	}
	NewSample := app.Sample{Timestamp: time.Now(), Value: value}
	Objects[object_id].Samples = append(Objects[object_id].Samples, NewSample)
	Objects[object_id].LastSample = NewSample
	return nil
}
var AddObject = func(obj app.Object) (err error) {
	m.Lock()
	defer writeUsersFile()
	defer m.Unlock()

	Objects[obj.Id] = &obj
	return nil
}
var GetObjects = func(object_id string) (result []app.Object) {
	result = make([]app.Object, 0)
	for key, val := range Objects {
		if strings.HasPrefix(key, object_id+"/") {
			result = append(result, *val)
		}
	}
	return
}
