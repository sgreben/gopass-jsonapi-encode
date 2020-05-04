# gopass-jsonapi-encode

Encode JSON as [gopass](https://github.com/gopasspw/gopass) JSONAPI messages.

`gopass-jsonapi-encode` reads JSON on stdin and writes JSONAPI messages on stdout.

## Contents

- [Contents](#contents)
  - [Examples](#examples)
- [Get it](#get-it)
  - [Using `go get`](#using-go-get)
  - [Pre-built binary](#pre-built-binary)

### Examples

```sh
$ jq -n '{"type":"query","query":"github.com"}' | gopass-jsonapi-encode | gopass jsonapi listen

_["github.com/example"]
```

```sh
$ jq -n '{"type":"create","entry_name":"127.0.0.1/admin","login":"admin","password":"hunter2"}' | gopass-jsonapi-encode | gopass jsonapi listen

_{"username":"admin","password":"hunter2"}
```

## Get it

### Using `go get`

```sh
go get -u github.com/sgreben/gopass-jsonapi-encode
```

### Pre-built binary

Or [download a binary](https://github.com/sgreben/gopass-jsonapi-encode/releases/latest) from the releases page, or from the shell:

```sh
# Linux
curl -L https://github.com/sgreben/gopass-jsonapi-encode/releases/download/0.0.3/gopass-jsonapi-encode_0.0.3_linux_x86_64.tar.gz | tar xz

# OS X
curl -L https://github.com/sgreben/gopass-jsonapi-encode/releases/download/0.0.3/gopass-jsonapi-encode_0.0.3_osx_x86_64.tar.gz | tar xz

# Windows
curl -LO https://github.com/sgreben/gopass-jsonapi-encode/releases/download/0.0.3/gopass-jsonapi-encode_0.0.3_windows_x86_64.zip
unzip gopass-jsonapi-encode_0.0.3_windows_x86_64.zip
```
