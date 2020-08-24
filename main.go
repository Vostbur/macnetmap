package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type linksStruct struct {
	Id        int    `json:"id"`
	Source    int    `json:"source"`
	SrcDevice string `json:"srcDevice"`
	SrcIfName string `json:"srcIfName"`
	Target    int    `json:"target"`
	TgtDevice string `json:"tgtDevice"`
	TgtIfName string `json:"tgtIfName"`
}
type nodesStruct struct {
	Icon string `json:"icon"`
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type linksStructJson struct {
	Links []linksStruct `json:"links"`
}
type nodesStructJson struct {
	Nodes []nodesStruct `json:"nodes"`
}

func run() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: macnetmap.exe DIR_WITH_OUTPUT_FILES")
		return
	}
	dir := os.Args[1]
	files, err := readDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	d := reader(files)
	i := 0
	var ls linksStructJson
	var ns nodesStructJson
	for indDevice, valDevice := range d {
		ns.Nodes = append(ns.Nodes, nodesStruct{
			Icon: "switch",
			Id:   indDevice,
			Name: d[indDevice].Hostname,
		})
		for interfDevice, macDevice := range valDevice.Interfaces {
			for otherIndDevice, otherValDevice := range d {
				if otherIndDevice == indDevice {
					continue
				}
			LOUT:
				for macFromTableOtherDevice, intFromTableOtherDevice := range otherValDevice.MacTable {
					if macFromTableOtherDevice == macDevice {
						// Проверяем на дублирование линков
						if i == 0 {
							goto L1
						}
						for j := 0; j < i; j++ {
							if ls.Links[j].Source == otherIndDevice && ls.Links[j].Target == indDevice {
								continue LOUT
							}
						}
					L1:
						ls.Links = append(ls.Links, linksStruct{
							Id:        i,
							Source:    indDevice,
							SrcDevice: d[indDevice].Hostname,
							SrcIfName: interfDevice,
							Target:    otherIndDevice,
							TgtDevice: d[otherIndDevice].Hostname,
							TgtIfName: intFromTableOtherDevice,
						})
						i += 1
					}
				}
			}
		}
	}
	lsJson, err := json.MarshalIndent(ls, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	nsJson, err := json.MarshalIndent(ns, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	strLsJon := string(lsJson)
	strLsJon = "var topologyData = " + strLsJon[:len(strLsJon)-2] + ","
	strNsJson := string(nsJson)
	strNsJson = strNsJson[1:] + ";\n"
	resultJson := strLsJon + strNsJson
	if err = ioutil.WriteFile("topology.js", []byte(resultJson), 0644); err != nil {
		log.Fatal(err)
	}
}

func main() {
	run()
}
