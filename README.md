<img align="left" width="50%" src="doc/assets/logo.png"/>

[![Go Report Card](https://goreportcard.com/badge/github.com/rytsh/gojson?style=flat-square)](https://goreportcard.com/badge/github.com/rytsh/gojson)
[![License](https://img.shields.io/github/license/rytsh/gojson?color=blue&style=flat-square)](https://raw.githubusercontent.com/rytsh/gojson/master/LICENSE)
[![Discord](https://img.shields.io/discord/706631996478324898?style=flat-square)](https://discordapp.com/channels/706631996478324898)

Serve any json file with GET, POST, PUT, PATCH or DELETE request data, even most inner object and root path.

---

## Notes

gojson hold all data in memory and case sensetive like what you see in json file.

Add an `id` field when PUT, POST, PATCH if you working on an array. gojson not put an auto-id. `id` field help us to find data in array.  
PATCH location and data should be an object.  
Root path's GET method reserved for UI but you can request other methods.

## Options

```txt
--port <3000>
  Set port
--host <localhost>
  Set host
```

s + enter will create a snapshot of the db on a new file.

## Build

You can download binary form in release section.  
Run `make` to build for major platforms or just specify `make windows`

---

## Examples

Run the server

```shell
gojson test/ex.json
```

### Get Data

Get whole list or an item.

```shell
curl http://localhost:3000/inner
[{"data":[{"id":11,"name":"11-inner"}],"id":1},{"data":[{"abc":{"value":5},"id":2,"name":"2-inner"}],"id":2}]

curl http://localhost:3000/inner/1
{"data":[{"id":11,"name":"11-inner"}],"id":1}

curl http://localhost:3000/inner/1/data/11
{"id":11,"name":"11-inner"}

curl http://localhost:3000/inner/1/data/11/name
"11-inner"
```

### Post data

Append a new data to field. Post location should be an array.

```shell
curl http://localhost:3000/users/
[{"age":"2","id":2,"name":"selin"},{"age":"5","id":"xx","name":"eray"},{"age":"3","id":4,"name":"ali"},{"age":"2","id":5,"name":"sinem"}]

curl -d '{"name":"ea","age":100}' -X POST http://localhost:3000/users/
{"msg":"success"}

curl http://localhost:3000/users/
[{"age":"2","id":2,"name":"selin"},{"age":"5","id":"xx","name":"eray"},{"age":"3","id":4,"name":"ali"},{"age":"2","id":5,"name":"sinem"},{"age":100,"name":"ea"}]

curl http://localhost:3000/inner
[{"data":[{"id":11,"name":"11-inner"}],"id":1},{"data":[{"abc":{"value":5},"id":2,"name":"2-inner"}],"id":2}]

curl -d '{"value":"Coool"}' -X POST http://localhost:3000/inner/1/data
{"msg":"success"}

curl http://localhost:3000/inner
[{"data":[{"id":11,"name":"11-inner"},{"value":"Coool"}],"id":1},{"data":[{"abc":{"value":5},"id":2,"name":"2-inner"}],"id":2}]
```

### Put data

You can send PUT request everywhere.

```shell
curl http://localhost:3000/inner
[{"data":[{"id":11,"name":"11-inner"}],"id":1},{"data":[{"abc":{"value":5},"id":2,"name":"2-inner"}],"id":2}]

curl -d '{"data": [{"id": 100, "x":"abc"}]}' -X PUT http://localhost:3000/inner/2/data/2/abc
{"msg":"success"}

curl http://localhost:3000/inner
[{"data":[{"id":11,"name":"11-inner"}],"id":1},{"data":[{"abc":{"data":[{"id":100,"x":"abc"}]},"id":2,"name":"2-inner"}],"id":2}]
```

### Patch data

Patch location and given data must be an object.

```shell
curl http://localhost:3000/inner/1/
{"data":[{"id":11,"name":"11-inner"}],"id":1}

curl -d '{"data":"new value"}' -X PATCH http://localhost:3000/inner/1/
{"msg":"success"}

curl http://localhost:3000/inner/1/
{"data":"new value","id":1}
```

### Delete data

If you delete root path, it will set an empty array or object depends what you have.  
In inner paths deleting that.

```shell
curl http://localhost:3000/inner
[{"data":[{"id":11,"name":"11-inner"}],"id":1},{"data":[{"abc":{"value":5},"id":2,"name":"2-inner"}],"id":2}]

curl -X DELETE http://localhost:3000/inner/1/data
{"msg":"success"}

curl http://localhost:3000/inner
[{"id":1},{"data":[{"abc":{"value":5},"id":2,"name":"2-inner"}],"id":2}]

curl -X DELETE http://localhost:3000/inner/2
{"msg":"success"}

curl http://localhost:3000/inner
[{"id":1}]
```

## License

gojson is [MIT licensed](./LICENSE).
