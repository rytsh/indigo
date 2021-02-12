<script>
    import Code from "../../component/Code.svelte";
    import Title from "../../component/Title.svelte";
    export let version;
</script>

<p>You can serve your exported React,Svelte any other static site with <b>INDIGO</b></p>
<p>Just mount your folder in anywere in our docker image, this is a scratch image just indigo executable inside.</p>

<Title id="docker-get" className="t2">Run it</Title>

<Code wrap code={
`docker run -d --name myawesomeapp -v $(pwd)/public:/public -p 3000:3000 ryts/indigo:${version} --folder /public --no-api --no-ui --spa`
}/>

<p>Now open your <b>http://localhost:3000</b> and use it</p>
