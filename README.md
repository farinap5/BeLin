```
00 00 BE EF | magic number
00 00 00 00 | data size
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
00 00       | ANSI
00 00       | OEM
00 00 00 00 | Beacon ID
00 00 00 00 | PID
00 00       | port
00          | Flag (x32 x64)
00 00       | Version
00 00       |
00 00 00 00 | Prefix
00 00 00 00 | Module Handler A ptr
00 00 00 00 | Get Proc Address ptr
00 00 00 00 | Ip address
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 | data separated by 0x09 (\t)
```


```
00 00 be ef 
00 00 00 4b
19 ea 3f da 57 10 8a 24 15 b3 05 93 b6 5b f5 a6 
03 a8
03 a8
00 04 fa 04
00 02 71 ad
00 00
04
06 01
00 00
00 00 00 00
00 00 00 00
00 00 00 00
8a 0f a8 c0
64 65 62 78 20 28 4c 69 6e 75 78 29 09 61 61 61 61 61 61 09 6d 61 69 6e
```