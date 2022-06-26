package layout

import (
	"fmt"
	"strings"

	"github.com/fops9311/mvc_server_app/tassets"
)

const format string = "<inner_content_%d/>"
const format1 string = "<inner_content_%s/>"

func init() {
	tassets.InitDir("./templates")
}

/*Layout replaces <inner_content_<*>/> with inner_contents args.
<*> is index of argument starting with 1. Also replaces <inner_content_<*>/> from tassets where <*> is map key.*/
func Layout(layout string, inner_contents ...string) (res string) {
	res = layout
	for i, s := range inner_contents {
		old := fmt.Sprintf(format, i+1)
		res = strings.ReplaceAll(res, old, s)
	}
	for _, s := range tassets.GetKeys() {
		old := fmt.Sprintf(format1, s)
		res = strings.ReplaceAll(res, old, tassets.GetAsset(s))
	}
	return res
}
