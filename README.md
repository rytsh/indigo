![indigo logo](doc/assets/logo.png)

[![Go Report Card](https://goreportcard.com/badge/github.com/rytsh/indigo?style=flat-square)](https://goreportcard.com/badge/github.com/rytsh/indigo)
[![License](https://img.shields.io/github/license/rytsh/indigo?color=blue&style=flat-square)](https://raw.githubusercontent.com/rytsh/indigo/master/LICENSE)
[![Discord](https://img.shields.io/discord/706631996478324898?style=flat-square)](https://discordapp.com/channels/706631996478324898)

Serve any json file with GET, POST, PUT, PATCH or DELETE request data, even most inner object and root path.

---

## Notes

Indigo hold all data in memory and case sensetive like what you see in json file.

Add an `id` field when PUT, POST, PATCH if you working on an array. indigo not put an auto-id. `id` field help us to find data in array.  
PATCH location and data should be an object.  
Root path's GET method reserved for UI but you can request other methods.
I will add more useful stuff in it and write test cases. If you see any error or wants to support something write me.

## Options

```txt
  --port, -P <3000>
    Set port, tool default is '3000'
  --host, -H <localhost>
    Set host, tool default is 'localhost'

  --empty <data.json>
    If not given any json file, this is the save location
    Tool default is 'data.json'
  --api <api>
    Start API url with this string
  --folder <./public>
    Serve folder
    if api option is empty, auto set 'api'

  --auth-basic <username:password>
    Enable basic authentication with username and password

  --no-color
    Disable color output

  -v, --version
    Show version number
  -h, --help
    Show help
```

s + enter will create a snapshot of the db on a new file.

## Build

You can download binary form in release section.  
Run `make` to build for major platforms or just specify `PLATFORM=windows make`

---

## Examples

Run the server, if you not give any json file it will start with null content.

```shell
# indigo --api api/v1 --folder /server/public --auth-basic user:pass db.json
# You can use above but I just start with json
indigo test/ex.json
```

<details><summary>Show example requests</summary>

Gzip compress can usable with `Accept-Encoding: gzip` header set.

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

</details>

## License

Indigo is [MIT licensed](./LICENSE).

### Libraries Used

[github.com/fatih/color](https://github.com/fatih/color)
