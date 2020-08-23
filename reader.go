package main

import (
	"fmt"
	"sync"
)

func reader(files []string) []Device {
	rawChan := make(chan string, len(files))
	wg := &sync.WaitGroup{}
	for _, f := range files {
		wg.Add(1)
		go func(f string, r chan string, wg *sync.WaitGroup) {
			defer wg.Done()
			tmp, err := readFile(f)
			if err != nil {
				fmt.Println(err.Error())
			}
			r <- tmp
		}(f, rawChan, wg)
	}
	wg.Wait()
	close(rawChan)

	devChan := make(chan Device, len(rawChan))
	for raw := range rawChan {
		wg.Add(1)

		go func(r string, c chan Device, wg *sync.WaitGroup) {
			defer wg.Done()
			wgInner := &sync.WaitGroup{}
			var tmp Device
			var err error

			wgInner.Add(1)
			go func(r string, tmp *Device, w *sync.WaitGroup) {
				defer w.Done()
				tmp.Hostname, err = parseHostname(r, hostRe)
				if err != nil {
					fmt.Println(err.Error())
				}
			}(r, &tmp, wgInner)

			wgInner.Add(1)
			go func(r string, tmp *Device, w *sync.WaitGroup) {
				defer w.Done()
				tmp.Interfaces, err = parseShow(r, shIntRe)
				if err != nil {
					fmt.Println(err.Error())
				}
			}(r, &tmp, wgInner)

			wgInner.Add(1)
			go func(r string, tmp *Device, w *sync.WaitGroup) {
				defer w.Done()
				tmp.MacTable, err = parseShow(r, shMacRe)
				if err != nil {
					fmt.Println(err.Error())
				}
			}(r, &tmp, wgInner)

			wgInner.Wait()
			c <- tmp
		}(raw, devChan, wg)
	}
	wg.Wait()
	close(devChan)

	var devices []Device = make([]Device, len(devChan))
	for i := range devices {
		devices[i] = <-devChan
	}
	return devices
}
