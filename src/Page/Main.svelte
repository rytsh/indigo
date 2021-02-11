<script>
	import Code from "../component/Code.svelte";
	import Title from "../component/Title.svelte";
	let VERSION = LATEST_VERSION;
</script>

<p>Serve any json file with GET, POST, PUT, PATCH or DELETE request data, even most inner object and root path.</p>
<p>Serve folder with SPA, browsable support options.</p>

<Title id="installation">Installation</Title>
<p>The most recent version of <b>indigo</b> is {VERSION}. Downloads are available on <a href="https://github.com/rytsh/indigo/releases/latest" target="_blank">GitHub</a>:</p>

<ul>
	<li>Linux 64-bit: <a href="https://github.com/rytsh/indigo/releases/latest/download/indigo-linux-amd64-{VERSION}.tar.gz">indigo-linux-amd64-{VERSION}.tar.gz</a> | <a href="https://github.com/rytsh/indigo/releases/latest/download/indigo-linux-arm64-{VERSION}.tar.gz">indigo-linux-arm64-{VERSION}.tar.gz</a></li>
	<li>macOS 64-bit: <a href="https://github.com/rytsh/indigo/releases/latest/download/indigo-darwin-amd64-{VERSION}.tar.gz">indigo-darwin-amd64-{VERSION}.tar.gz</a></li>
	<li>Windows 64-bit: <a href="https://github.com/rytsh/indigo/releases/latest/download/indigo-windows-amd64-{VERSION}.zip">indigo-windows-amd64-{VERSION}.zip</a></li>
</ul>

<p>Get in Linux amd64</p>
<Code button code="curl -fsSL https://github.com/rytsh/indigo/releases/latest/download/indigo-linux-amd64-{VERSION}.tar.gz | sudo tar --overwrite -zx -C /usr/local/bin/"/>

<p>Run in <a href="https://hub.docker.com/r/ryts/indigo" target="_blank">docker</a></p>
<Code button code="docker run --rm -it -p 3000:3000 ryts/indigo:latest https://rytsh.github.io/indigo/test/users.json"/>

<Title id="options">Options</Title>

<Code code={
`indigo [OPTIONS] <file_or_URL>
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
`}/>

<p>s + enter will create a snapshot of the db on a new file.</p>

<p>If same URL uses, order is UI &gt; API &gt; FILE</p>

<p>Gzip compress can usable with <span class="color">Accept-Encoding: gzip</span> header set.</p>

<Title id="examples">Examples</Title>

<p>Give any json file to serve, could be a file or a URL</p>
<Code code="$ indigo https://api.punkapi.com/v2/beers/1/"/>
<p>Use REST API to get and change data event most inner object</p>
<Code code={
`$ curl localhost:3000/1/name

Buzz`}/>

<p>Serve api with basic auth</p>
<Code code="$ indigo --auth-basic user:pass db.json"/>

<p>Share just a folder</p>
<Code code="$ indigo --folder ./public --no-api --no-ui"/>

<p>Share just a folder with SPA</p>
<Code code="$ indigo --folder ./public --no-api --no-ui --spa"/>

<p>Share just a folder with just browsable support</p>
<Code code="$ indigo --folder ./public --no-api --no-ui --browsable --no-index"/>

<hr/>

<Title id="examples-req" className="t2">Examples of using requests</Title>
<p>Open a test json file</p>
<Code code="$ indigo https://rytsh.github.io/indigo/test/users.json"></Code>

<Title id="examples-get" className="t2">Get Request</Title>
<Code wrap code={
`$ curl http://localhost:3000

{"users":[{"age":22,"hobies":["traveller","movies","sport","kahveci"],"id":1,"name":"Selin"},{"age":55,"hobies":["games","sci-fi","kamyoncu"],"id":"xx","name":"Eray"},{"age":52,"hobies":["cars","camping","fishing","job changer"],"id":2,"name":"Ali"},{"age":50,"hobies":["photography","scout","cooking","siemens"],"id":3,"name":"Sinem"},{"age":53,"hobies":["theater","costume","star","professional interviewer"],"id":4,"name":"Yasin"},{"age":44,"hobies":["mushrooms","swiming","planes","no hello no mello"],"id":5,"name":"Aysun"},{"age":67,"hobies":["cars","mars","koç koç"],"id":67,"name":"Utku"},{"age":49,"hobies":["gastronomi","literature","ege'nin incisi"],"id":6,"name":"Zeynep"},{"age":50,"hobies":["books","games","fast write"],"id":7,"name":"Cagatay"},{"age":50,"hobies":["painting","presentation","pasta"],"id":8,"name":"Cansu"}]}

$ curl http://localhost:3000/users/1

{"age":50,"hobies":["traveller","movies","sport"],"id":1,"name":"Selin"}

$ curl http://localhost:3000/users/1/name

Selin

$ curl -s -H "Accept-Encoding: gzip" http://localhost:3000/users/5 | zcat

{"age":44,"hobies":["mushrooms","swiming","planes"],"id":5,"name":"Aysun"}
`}/>

<Title id="examples-post" className="t2">Post Request</Title>
<p>Append a new data to field. Post location should be an array.</p>
<Code wrap code={
`$ curl http://localhost:3000/users/1/hobies

["traveller","movies","sport"]

$ curl -d 'running' -X POST http://localhost:3000/users/1/hobies

{"msg":"success"}

$ curl http://localhost:3000/users/1/hobies

["traveller","movies","sport","running"]
`}/>

<Title id="examples-put" className="t2">Put Request</Title>
<Code wrap code={
`$ curl http://localhost:3000/users/7

{"age":50,"hobies":["books","games","fast write"],"id":7,"name":"Cagatay"}

$ curl -d '{"outside": ["running"], "inside": ["movies"]}' -X PUT http://localhost:3000/users/7/hobies

{"msg":"success"}

$ curl http://localhost:3000/users/7

{"age":50,"hobies":{"inside":["movies"],"outside":["running"]},"id":7,"name":"Cagatay"}

$ curl http://localhost:3000/users/7/hobies/inside

["movies"]
`}/>

<Title id="examples-patch" className="t2">Patch Request</Title>
<p>Patch location and given data must be an object.</p>
<Code wrap code={
`$ curl http://localhost:3000/users/xx

{"age":55,"hobies":["games","sci-fi","driver"],"id":"xx","name":"Eray"}

$ curl -d '{"id": 100}' -X PATCH http://localhost:3000/users/xx

{"msg":"success"}

$ curl http://localhost:3000/users/xx

{"err": "Not found!"}

$ curl http://localhost:3000/users/100

{"age":55,"hobies":["games","sci-fi","driver"],"id":100,"name":"Eray"}
`}/>


<Title id="examples-delete" className="t2">Delete Request</Title>
<Code wrap code={
`$ curl http://localhost:3000/users/8

{"age":50,"hobies":["painting","presentation","pasta"],"id":8,"name":"Cansu"}

$ curl -X DELETE http://localhost:3000/users/8/hobies

{"msg":"success"}

$ curl http://localhost:3000/users/8

{"age":50,"id":8,"name":"Cansu"}

$ curl -X DELETE http://localhost:3000

{"msg":"success"}

$ curl http://localhost:3000

null
`}/>

<hr/>

<p class="mail"><a href="https://github.com/rytsh/indigo/blob/master/LICENSE" target="_blank">MIT Licensed</a> | <a href="mailto:rytsh@devusage.com">rytsh@devusage.com</a></p>

<style>
	.mail {
		text-align: right;
	}
</style>
