# Dev Publish Instruction

Commit last changes and add tag with `git tag v0.x.y`

Build tool

```shell
PLATFORMS="windows:amd64,linux:arm64-amd64,darwin:amd64" ./build.sh --build --clean --pack
```

Push commit to repo. `git push origin v0.x.y`

Manuel github release binary.

Docker push automatically handle by drone.io

Manual docker build and push, run build command before this

```sh
./build.sh --build-context --build-docker
```

---

```sh
git update-index --assume-unchanged page/.env
```
