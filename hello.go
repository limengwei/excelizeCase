package main

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

//编译命令
//rsrc -manifest hello.manifest -o rsrc.syso
//go build -ldflags="-H windowsgui"
func main() {
	var wv *walk.WebView

	beego.Router("/", &MainCtrl{})
	go beego.Run()

	MainWindow{
		Icon:    Bind("'office.ico'"),
		Title:   "FuckExcel",
		MinSize: Size{800, 600},
		Layout:  VBox{MarginsZero: true},
		Children: []Widget{
			WebView{
				AssignTo: &wv,
				Name:     "wv",
				URL:      "http://localhost:8080/",
			},
		},
		Functions: map[string]func(args ...interface{}) (interface{}, error){
			"icon": func(args ...interface{}) (interface{}, error) {
				if strings.HasPrefix(args[0].(string), "https") {
					return "check", nil
				}

				return "stop", nil
			},
		},
	}.Run()
}

type MainCtrl struct {
	beego.Controller
}

func (c *MainCtrl) Get() {
	c.TplName = "index.html"
}