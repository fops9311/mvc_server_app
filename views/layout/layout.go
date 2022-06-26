package layout

import (
	"fmt"
	"strings"
)

const format string = "<inner_content_%d/>"

//Layout replaces <inner_content_*/> with inner_contents args.
// * is index of argument starting with 1
func Layout(layout string, inner_contents ...string) (res string) {
	res = layout
	for i, s := range inner_contents {
		old := fmt.Sprintf(format, i+1)
		res = strings.ReplaceAll(res, old, s)
	}
	return res
}
