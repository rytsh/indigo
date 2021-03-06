![indigo logo](doc/assets/logo.png)

[![Go Report Card](https://goreportcard.com/badge/github.com/rytsh/indigo?style=flat-square)](https://goreportcard.com/report/github.com/rytsh/indigo)
[![License](https://img.shields.io/github/license/rytsh/indigo?color=blue&style=flat-square)](https://raw.githubusercontent.com/rytsh/indigo/master/LICENSE)
[![Discord](https://img.shields.io/discord/706631996478324898?style=flat-square)](https://discordapp.com/channels/706631996478324898)
[![Docker](https://img.shields.io/badge/dockerHub-indigo-blue?style=flat-square&logo=docker)](https://hub.docker.com/r/ryts/indigo)
[![Web](https://img.shields.io/badge/web-gh--pages-blueviolet?style=flat-square)](https://rytsh.github.io/indigo/)
[![Coveralls github branch](https://img.shields.io/coveralls/github/rytsh/indigo/master?style=flat-square)](https://coveralls.io/github/rytsh/indigo?branch=master)
![Drone (cloud)](https://img.shields.io/drone/build/rytsh/indigo?style=flat-square)

Serve any json file with GET, POST, PUT, PATCH or DELETE request data, even most inner object and root path.  
Serve folder with SPA, browsable support options.

---

## Notes

Indigo hold all data in memory and case sensetive like what you see in json file.

Add an `id` field when PUT, POST, PATCH if you working on an array. indigo not put an auto-id. `id` field help us to find data in array.  
If same URL uses, order is UI > API > FILE  
I will add more useful stuff in it and write test cases.  
If you see any error or wants to support something write me.

## Options

```txt
indigo [OPTIONS] <file_or_URL>
Generate RestAPI with JSON file and serve folder

Options:
  --port, -P <3000>
    Set port, tool default is '3000'
  --host, -H <localhost>
    Set host, tool default is 'localhost'

  --location <./data.json>
    Change save location
  --api-path <api_url_path>
    Set API path prefix
  --ui-path <ui_url_path>
    Set UI path default '/indigo'

  --folder <./public>
    Serve folder
  --folder-path <folder_path>
    Set Folder path, works with folder option
  --browsable
    Enable folder browsable
  --spa
    Enable SPA mode
  --no-index
    Stop redirect to index

  --no-api
    Close API server, use just serve folder
  --no-ui
    Close UI server

  --auth-basic <username:password>
    Enable basic authentication with username and password

  --no-color
    Disable color output

  -v, --version
    Show version number
  -h, --help
    Show help

  Examples:
    indigo --api-path api/v1 --auth-basic user:pass test/users.json
```

## Build

You can download binary form in releases.

Or you can build with show PLATFORMS, each platform seperated by `,` and each arch by `-`.

```shell
PLATFORMS="windows:amd64,linux:arm64-amd64,darwin:amd64" ./build.sh --build --clean
```

Check list of platforms and archs here  
<https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63>

---

## License

Indigo is [MIT licensed](./LICENSE).

### Libraries Used

[github.com/fatih/color](https://github.com/fatih/color)
