# Dev Publish Instruction

Commit last changes and add tag with `git tag v0.x.y`

Build tool

```shell
PLATFORMS="windows:amd64,linux:arm64-amd64,darwin:amd64" ./build.sh --build --clean --pack
./build.sh --publish-page
```

Push commit to repo.

Manuel github release binary.

Docker push automatically handle by drone.io

---

```sh
git update-index --assume-unchanged page/.env
```
