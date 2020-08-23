package main

import (
	"fmt"
	"log"
	"os"
)

func run() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: netmap.exe DIR_WITH_OUTPUT_FILES")
		return
	}
	dir := os.Args[1]
	files, err := readDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	d := reader(files)
	for indDevice, valDevice := range d {
		for interfDevice, macDevice := range valDevice.Interfaces {
			for otherIndDevice, otherValDevice := range d {
				if otherIndDevice == indDevice {
					continue
				}
				for macFromTableOtherDevice, intFromTableOtherDevice := range otherValDevice.MacTable {
					if macFromTableOtherDevice == macDevice {
						fmt.Printf("Device '%s' port '%s' is connected to device '%s' port '%s'\n",
							d[indDevice].Hostname, interfDevice, d[otherIndDevice].Hostname, intFromTableOtherDevice)
					}
				}
			}
		}
	}
}

func main() {
	run()
}
