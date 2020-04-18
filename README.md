<p align="center"><img src="doc/assets/logo.png" width="50%"/></p>

Serve given json file with first child URL seperation, this is a very basic json-server but can run on any json file.  
gojson hold all data in memory.

## Useful options

```txt
--port <3000>
  Set port
--host <localhost>
  Set host
```

## Build

Run `make` to build listed platforms or just specify `make windows`

---

## Examples

This server is a case insensitive.

Run the server

```shell
gojson test/ex.json
```

### Get Data

Get whole list or an item with id field

```shell
curl http://localhost:3000/UsErS
[{"age":"2","id":2,"name":"selin"},{"age":"5","id":"xx","name":"eray"},{"age":"3","id":4,"name":"ali"},{"age":"2","id":5,"name":"sinem"}]

 curl http://localhost:3000/users/2
{"age":"2","id":2,"name":"selin"}
```

### Post data

Append a new data to field

```shell
curl -d '{"name":"ea","age":100}' -X POST -H 'Content-Type: application/json' http://localhost:3000/users/

curl http://localhost:3000/users
[{"age":"2","id":2,"name":"selin"},{"age":"5","id":"xx","name":"eray"},{"age":"3","id":4,"name":"ali"},{"age":"2","id":5,"name":"sinem"},{"age":100,"name":"ea"}]

curl -d '[{"name":"XYZ","age":10000}]' -X POST -H 'Content-Type: application/json' http://localhost:3000/userS/

curl http://localhost:3000/users
[{"age":"2","id":2,"name":"selin"},{"age":"5","id":"xx","name":"eray"},{"age":"3","id":4,"name":"ali"},{"age":"2","id":5,"name":"sinem"},{"age":100,"name":"ea"},[{"age":10000,"name":"XYZ"}]]

```

### Put data

You can not delete it or put again this content beacuse you didn't pass an id field! Because auto id algorithm doesn't exist.

```shell
curl http://localhost:3000/users
[{"age":"2","id":2,"name":"selin"},{"age":"5","id":"xx","name":"eray"},{"age":"3","id":4,"name":"ali"},{"age":"2","id":5,"name":"sinem"}]

curl -d '{"id":100,"name":"selin","age":100}' -X PUT -H 'Content-Type: application/json' http://localhost:3000/userS/2

curl http://localhost:3000/users
[{"age":100,"id":100,"name":"selin"},{"age":"5","id":"xx","name":"eray"},{"age":"3","id":4,"name":"ali"},{"age":"2","id":5,"name":"sinem"}
```

### Delete data

You can delete with id field, warning without it it will flush data. (is it good?)

```shell
curl http://localhost:3000/test
[{"date":"Wed Apr 15 17:04:14 +03 2020","id":1},{"age":107,"id":2,"name":"AAAA"}]

curl -X DELETE http://localhost:3000/test/2
{}

curl http://localhost:3000/test
[{"date":"Wed Apr 15 17:04:14 +03 2020","id":1}]

curl -X DELETE http://localhost:3000/test
{}

curl http://localhost:3000/test
[]
```

## License

gojson is [MIT licensed](./LICENSE).
