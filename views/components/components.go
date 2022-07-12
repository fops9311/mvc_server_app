package components

import (
	"bytes"
	"sync"
	"text/template"

	"github.com/fops9311/mvc_server_app/tassets"
)

type ViewRequest struct {
	Params    map[string]interface{}
	Respounce chan ViewRespounce
}
type ViewRespounce struct {
	Result string
	Error  error
}

var (
	Head               chan ViewRequest = make(chan ViewRequest, 5)
	Header             chan ViewRequest = make(chan ViewRequest, 5)
	Alarm_panel        chan ViewRequest = make(chan ViewRequest, 5)
	Footer             chan ViewRequest = make(chan ViewRequest, 5)
	Layout_concat2     chan ViewRequest = make(chan ViewRequest, 5)
	Layout_fullwh_cent chan ViewRequest = make(chan ViewRequest, 5)
	Layout_grid2x2     chan ViewRequest = make(chan ViewRequest, 5)
	Layout_htmlpage    chan ViewRequest = make(chan ViewRequest, 5)
	Layout_tiles2      chan ViewRequest = make(chan ViewRequest, 5)
	Loginbutton        chan ViewRequest = make(chan ViewRequest, 5)
	Menu               chan ViewRequest = make(chan ViewRequest, 5)
	MessagePanel       chan ViewRequest = make(chan ViewRequest, 5)
	Navbar             chan ViewRequest = make(chan ViewRequest, 5)
	ObjectPanel        chan ViewRequest = make(chan ViewRequest, 5)
	SinginPanel        chan ViewRequest = make(chan ViewRequest, 5)
	SignupPanel        chan ViewRequest = make(chan ViewRequest, 5)
	SummaryPanel       chan ViewRequest = make(chan ViewRequest, 5)
	TrendPanel         chan ViewRequest = make(chan ViewRequest, 5)
	Layout_htmldoc     chan ViewRequest = make(chan ViewRequest, 5)
)

var (
	DIR_Head               string = "templates/components/head/index.html"
	DIR_Header             string = "templates/components/header/index.html"
	DIR_Alarm_panel        string = "templates/components/alarm_panel/index.html"
	DIR_Footer             string = "templates/components/footer/index.html"
	DIR_Layout_concat2     string = "templates/components/layout_concat2/index.html"
	DIR_Layout_fullwh_cent string = "templates/components/layout_fullwh_cent/index.html"
	DIR_Layout_grid2x2     string = "templates/components/layout_grid2x2/index.html"
	DIR_Layout_htmlpage    string = "templates/components/layout_htmlpage/index.html"
	DIR_Layout_tiles2      string = "templates/components/layout_tiles2/index.html"
	DIR_Loginbutton        string = "templates/components/loginbutton/index.html"
	DIR_MessagePanel       string = "templates/components/message_panel/index.html"
	DIR_Navbar             string = "templates/components/navbar/index.html"
	DIR_ObjectPanel        string = "templates/components/object_panel/index.html"
	DIR_SinginPanel        string = "templates/components/singin_panel/index.html"
	DIR_SignupPanel        string = "templates/components/signup_panel/index.html"
	DIR_SummaryPanel       string = "templates/components/summary_panel/index.html"
	DIR_TrendPanel         string = "templates/components/trend_panel/index.html"
	DIR_Layout_htmldoc     string = "templates/components/layout_htmldoc/index.html"
)

func Init() {
	view := func(c chan ViewRequest, dir string) {
		var req ViewRequest
		for {
			req = <-c
			req.Respounce <- renderTemplate(
				req.Params,
				tassets.GetAsset(dir),
				dir,
			)
		}
	}
	go view(
		Head,
		DIR_Head,
	)
	go view(
		Header,
		DIR_Header)
	go view(
		Alarm_panel,
		DIR_Alarm_panel)
	go view(
		Footer,
		DIR_Footer)
	go view(
		Layout_concat2,
		DIR_Layout_concat2)
	go view(
		Layout_fullwh_cent,
		DIR_Layout_fullwh_cent)
	go view(
		Layout_grid2x2,
		DIR_Layout_grid2x2)
	go view(
		Layout_htmlpage,
		DIR_Layout_htmlpage)
	go view(
		Layout_tiles2,
		DIR_Layout_tiles2)
	go view(
		Loginbutton,
		DIR_Loginbutton)
	go view(
		MessagePanel,
		DIR_MessagePanel)
	go view(
		Navbar,
		DIR_Navbar)
	go view(
		ObjectPanel,
		DIR_ObjectPanel)
	go view(
		SinginPanel,
		DIR_SinginPanel)
	go view(
		SignupPanel,
		DIR_SignupPanel)
	go view(
		SummaryPanel,
		DIR_SummaryPanel)
	go view(
		TrendPanel,
		DIR_TrendPanel)
	go view(
		Layout_htmldoc,
		DIR_Layout_htmldoc)
}

var tmap map[string]*template.Template = make(map[string]*template.Template)
var m sync.Mutex

func renderTemplate(params map[string]interface{}, templ string, templateName string) (resp ViewRespounce) {
	m.Lock()
	defer m.Unlock()
	if _, ok := tmap[templateName]; !ok {
		tmap[templateName], resp.Error = template.New(templateName).Parse(templ)
		if resp.Error != nil {
			return resp
		}
	}

	buf := bytes.NewBuffer([]byte{})
	resp.Error = tmap[templateName].Execute(buf, params)
	if resp.Error != nil {
		return resp
	}
	resp.Result = buf.String()
	return resp
}

func Render(comp chan ViewRequest, params map[string]interface{}) (result string) {
	var resp chan ViewRespounce = make(chan ViewRespounce)
	comp <- ViewRequest{
		Params:    params,
		Respounce: resp,
	}
	r := <-resp
	if r.Error != nil {
		return ""
	}
	return r.Result
}
