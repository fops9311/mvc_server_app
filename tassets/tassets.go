package tassets

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var templates map[string]string

var init_flag bool = false
var m sync.Mutex

func init() {

}
func InitDir(dir string) {
	if init_flag {
		return
	}
	init_flag = true
	templates = map[string]string{}
	traverceDir(dir)

}
func GetAsset(name string) (templ string) {
	m.Lock()
	defer m.Unlock()
	return templates[name]
}
func traverceDir(dir string) {
	_ = filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())
			if !info.IsDir() {
				b, _ := os.ReadFile(path)
				templates[path] = (string(b))
			}
			return nil
		})
}
