package controllers

import "github.com/astaxie/beego"

import "github.com/360EntSecGroup-Skylar/excelize"

type MainCtrl struct {
	beego.Controller
}

func (c *MainCtrl) Get() {
	xlsx, err := excelize.OpenFile("./data/data.xlsx")

	if err != nil {
		c.Data["msg"] = err
	}

	rows := xlsx.GetRows("Sheet1")

	var thList = make([]string, 0)

	var trList = make([][]string, 0)

	for i, row := range rows {

		var tdList = make([]string, 0)

		for _, colCell := range row {
			if i > 0 {

				tdList = append(tdList, colCell)
			} else {
				thList = append(thList, colCell)
			}
		}
		if i > 0 {
			trList = append(trList, tdList)
		}

	}

	c.Data["thList"] = thList
	c.Data["trList"] = trList

	c.TplName = "index.html"
}
