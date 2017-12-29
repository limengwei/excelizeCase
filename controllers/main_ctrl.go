package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"

	"helloexcel/models"
	_ "helloexcel/models"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Xs struct {
	Id     int64
	Time   string
	Amount string
	Fee    string
	Type   string
	Status string
}

var o orm.Ormer

func hello() {
	var err error
	err = orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/helloexcel?charset=utf8", 30)

	if err != nil {
		fmt.Println(err)
		return
	}

	orm.RegisterModel(new(Xs))
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	o = orm.NewOrm()
}

type MainCtrl struct {
	beego.Controller
}

type Gaosu struct {
}

func (c *MainCtrl) Gis() {

	name := c.GetString("name", "")
	if name != "" {
		c.Data["json"] = models.GetByName(name)
		c.ServeJSON()
		return
	}

	c.Data["nameList"] = models.GetNameList()

	c.TplName = "gis.html"
}

func (c *MainCtrl) Get() {
	xlsx, err := excelize.OpenFile("./data/data.xlsx")

	if err != nil {
		fmt.Println(err)
		c.Data["msg"] = err
	}

	rows := xlsx.GetRows("XS")

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

func xls2db(row []string) {
	xs := Xs{}
	xs.Id, _ = strconv.ParseInt(row[0], 10, 64)
	xs.Time = row[1]
	xs.Amount = row[2]
	xs.Fee = row[3]
	xs.Type = row[4]
	xs.Status = row[5]

	id, err := o.Insert(&xs)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)
}
