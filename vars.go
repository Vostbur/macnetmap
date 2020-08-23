package main

const (
	// [1] GigabitEthernet0/1 | FastEthernet0/1 | Vlan1
	// [2] mac-address
	// Ex. map[FastEthernet0/1:0005.5ee6.8e01]
	shIntRe = `(?im)(^\w+\d+/*\d*).+\n.+([0-9a-f]{4}\.[0-9a-f]{4}\.[0-9a-f]{4}) `
	// [1] mac-address
	// [2] Fa0/1
	// Ex. map[0005.5ee6.8e01:Fa0/1]
	shMacRe = `([0-9a-f]{4}\.[0-9a-f]{4}\.[0-9a-f]{4}).+ (\w+\d+/*\d*)`
	// [0] sw0# sw0
	// [1] sw0
	// Ex. [sw0# sw0]
	hostRe = `(\w+?)[>|#]`
)

type Device struct {
	Hostname   string
	Interfaces map[string]string
	MacTable   map[string]string
}
