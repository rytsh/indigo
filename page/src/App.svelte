<script>
	import Code from "./component/Code.svelte";
	let name = "INDIGO";
	let VERSION = LATEST_VERSION;
</script>

<main>
	<p class="title">{name}</p>
	<p>Serve any json file with GET, POST, PUT, PATCH or DELETE request data, even most inner object and root path.</p>
	<p>Get in Linux, for other OS <a href="https://github.com/rytsh/indigo/releases/latest" target="_blank">click here</a></p>
	<Code code={`curl -fsSL https://github.com/rytsh/indigo/releases/latest/download/indigo-linux-amd64-${VERSION}.tar.gz | tar -zx -C /usr/local/bin/`}/>

	<p>Give any json file to serve</p>
	<Code code="indigo https://api.punkapi.com/v2/beers/1/"/>
	<p>Use REST API to get and change data event most inner</p>
	<Code code="curl localhost:3000/1/name"/>

	<p>Share a folder</p>
	<Code code="indigo --folder /mnt/usb --no-api"/>

	<hr/>

	<p>Run in docker</p>
	<Code code="docker run --rm -it -p 3000:3000 ryts/indigo:latest https://api.punkapi.com/v2/beers/1/"/>

</main>

<style>
	main {
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	.title {
		text-align: center;
		font-family: 'Caesar Dressing', cursive;
		font-size: 72pt;
		margin: 0;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>