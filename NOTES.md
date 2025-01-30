The communication begins with the beacon sending its initial message, which contains metadata. This packet includes details about the machine, a randomly generated session ID linked to the beacon ID, the IP address, and other relevant information.

This request registers the implant with the server.

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

The implant continuously calls back to the server, acting as a heartbeat to signal its status and request new commands. When a command is included, the server's response follows a structured format.

For the following message, the total content length is $x$, while the hash is last $16$ bytes.

```
200 51 81 249 224 202 112 126 4 141 215 97 111 183 87 160 124 120 14 137 38 230 13 113 243 201 189 187 227 36 89 47 | Data
0 172 250 93 167 54 67 144 23 60 224 91 56 90 192 141 | HMAC
```

Splitting the message reveals its components: a timestamp, the length of the payload, the command ID, and the data.

Once the package is decrypted using AES, the extracted sequence is as follows:

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