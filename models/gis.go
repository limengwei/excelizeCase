package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/bitly/go-simplejson"
)

type GisData struct {
	Type     string
	Features []Feature
}

type Feature struct { //特征
	Type       string
	Geometry   Geometry
	Properties Properties
}

type Geometry struct { //坐标
	Type        string
	Coordinates [][]float64
}

type Properties struct { //路段属性
	Z____ID int
	NAME    string
	KIND    string
}

var data GisData

func init() {
	fmt.Println("--ReadDson--")

	bs, err := ioutil.ReadFile("./data/gaosu.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = simplejson.NewJson(bs)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(bs, &data)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(len(data.Features))
}

type Line struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}

var NameList = make([]string, 0)

func GetNameList() []string {
	if len(NameList) > 0 {
		return NameList
	}

	features := data.Features

	nameMap := make(map[string]string)
	for _, f := range features {

		n := f.Properties.NAME
		if n == "" {
			continue
		}
		nameMap[n] = n
	}

	for _, v := range nameMap {
		NameList = append(NameList, v)
	}

	return NameList

}

func GetByName(name string) []Line {

	var lines = make([]Line, 0)

	features := data.Features
	for _, f := range features {
		if f.Properties.NAME == name {
			line := new(Line)
			line.Type = "LineString"
			line.Coordinates = f.Geometry.Coordinates
			lines = append(lines, *line)
		}
	}

	return lines
}
