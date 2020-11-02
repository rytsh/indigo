# Dev Publish Instruction

Commit last changes and add tag with `git tag v0.x.y`

Build tool, docker image and publish page

```shell
./build.sh --build --clean --publish-page --build-docker
```

Push commit to repo.

Manuel push docker image and github release binary.
