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
00 00       | Build
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


For the following message the total content length is $x$, the hash is $x-16$ from the right side.
```
200 51 81 249 224 202 112 126 4 141 215 97 111 183 87 160 124 120 14 137 38 230 13 113 243 201 189 187 227 36 89 47 | Data
0 172 250 93 167 54 67 144 23 60 224 91 56 90 192 141 | HMAC
```


Decoding the package with AES the following sequence is taken.
```
103 153 129 205 | timestamp
0 0 0 19 | length
0 0 0 53 | Command (ls in this case)
0 0 0 11 255 255 255 254 0 0 0 3 46 92 42 65 65 65 65 65 | data
```


References
- https://github.com/b1tg/cobaltstrike-beacon-rust/blob/main/src/main.rs
- https://github.com/darkr4y/geacon
- https://github.com/kyxiaxiang/Beacon_Source
- https://unit42.paloaltonetworks.com/cobalt-strike-metadata-encryption-decryption/