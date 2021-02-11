<script>
  import Code from "../component/Code.svelte";
  import Title from "../component/Title.svelte";
  import Usage from "./Usage.svelte";
  import { getData } from "../helper/version";

  $: assets = {};
  $: version = "";

  (async ()=> {
    const data = await getData();
    version = data.tag_name;
    data.assets.forEach(asset => {
      const name = asset.name.split("-")
      const [osName, arch] = [name[1], name[2]];
      const info = {
        name: asset.name,
        arch: arch,
        browser_download_url: asset.browser_download_url,
      };
      if (! assets[osName]) {
        assets[osName] = [];
      }
      assets[osName].push(info);
    });
  })();
</script>

<p>Serve any json file with GET, POST, PUT, PATCH or DELETE request data, even most inner object and root path.</p>
<p>Serve folder with SPA, browsable support options.</p>

<Title id="installation">Installation</Title>
<p>The most recent version of <b>indigo</b> is {version}. Downloads are available on <a href="https://github.com/rytsh/indigo/releases/latest" target="_blank">GitHub</a>:</p>

<ul>
  {#each Object.entries(assets) as [osName, asset]}
    <li>
      <span class="font-code">
        {osName.padEnd(10,'-')}| 
      </span>
      {#each asset as {browser_download_url, name}, i}
        <a href={browser_download_url}>{name}</a>{(asset.length-1) > i ? " | " : ""}
      {/each}
    </li>
  {/each}
</ul>

{#if assets["linux"]}
  <p>Get in Linux amd64</p>
  <Code button code={`curl -fsSL ${assets["linux"].find(val => val.arch == "amd64").browser_download_url} | sudo tar --overwrite -zx -C /usr/local/bin/`}/>
{/if}

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

<hr/>

<Usage version={version}/>

<hr/>

<p class="mail"><a href="https://github.com/rytsh/indigo/blob/master/LICENSE" target="_blank">MIT Licensed</a> | <a href="mailto:rytsh@devusage.com">rytsh@devusage.com</a></p>

<style>
	.mail {
		text-align: right;
	}
  .font-code {
    font-family: monospace;
    text-transform: capitalize;
  }
  li:nth-child(odd) {
    background-color: #f9f9f9;
  }
  li:nth-child(even) {
    background-color: #fcfcfc;
  }
  li:hover {
    background-color: #ddd;
  }
</style>
