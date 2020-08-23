# macnetmap
Output network map from mac address table

Program show devices connectivity from output command "sh int", "sh mac-address-table" (for switches) or "sh ip arp" (for routers). Based only on mac-address information. For better result ping devices before run program. It update mac address tables.

Argument is folder with files. For example there is *output* folder with several files.

Output:

```
PS D:\projects-go\src\macnetmap> .\macnetmap.exe output
Device 'sw1' port 'FastEthernet0/2' is connected to device 'sw2' port 'Fa0/1'
Device 'sw1' port 'FastEthernet0/1' is connected to device 'sw0' port 'Fa0/1'
Device 'sw2' port 'FastEthernet0/1' is connected to device 'sw1' port 'Fa0/2'
Device 'sw2' port 'FastEthernet0/2' is connected to device 'sw3' port 'Fa0/1'
Device 'sw0' port 'FastEthernet0/1' is connected to device 'sw1' port 'Fa0/1'
Device 'sw3' port 'FastEthernet0/1' is connected to device 'sw2' port 'Fa0/2'
```
