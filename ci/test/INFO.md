# Dockerfile

This dockerfile using for testing and building this repository.

Run command in project dir.

```shell
docker build -t ryts/go:1.15 -f ci/test/Dockerfile ci/test
```

Get in docker shell

```shell
docker run -it --rm -v $(pwd):/app ryts/go:1.15
```
