var topologyData = {
	"links": [
		{
			"id": 0,
			"source": 0,
			"srcDevice": "sw3",
			"srcIfName": "FastEthernet0/1",
			"target": 2,
			"tgtDevice": "sw2",
			"tgtIfName": "Fa0/2"
		},
		{
			"id": 1,
			"source": 1,
			"srcDevice": "sw1",
			"srcIfName": "FastEthernet0/1",
			"target": 3,
			"tgtDevice": "sw0",
			"tgtIfName": "Fa0/1"
		},
		{
			"id": 2,
			"source": 1,
			"srcDevice": "sw1",
			"srcIfName": "FastEthernet0/2",
			"target": 2,
			"tgtDevice": "sw2",
			"tgtIfName": "Fa0/1"
		}
	],
	"nodes": [
		{
			"icon": "switch",
			"id": 0,
			"name": "sw3"
		},
		{
			"icon": "switch",
			"id": 1,
			"name": "sw1"
		},
		{
			"icon": "switch",
			"id": 2,
			"name": "sw2"
		},
		{
			"icon": "switch",
			"id": 3,
			"name": "sw0"
		}
	]
};
