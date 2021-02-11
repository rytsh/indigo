<script>
    import Code from "../../component/Code.svelte";
</script>

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
