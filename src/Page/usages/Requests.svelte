<script>
	import Code from "../../component/Code.svelte";
    import Title from "../../component/Title.svelte";
</script>

<p>Open a test json file</p>
<Code code="$ indigo https://rytsh.github.io/indigo/test/users.json"></Code>

<Title id="request-get" className="t2">Get Request</Title>
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

<Title id="request-post" className="t2">Post Request</Title>
<p>Append a new data to field. Post location should be an array.</p>
<Code wrap code={
`$ curl http://localhost:3000/users/1/hobies

["traveller","movies","sport"]

$ curl -d 'running' -X POST http://localhost:3000/users/1/hobies

{"msg":"success"}

$ curl http://localhost:3000/users/1/hobies

["traveller","movies","sport","running"]
`}/>

<Title id="request-put" className="t2">Put Request</Title>
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

<Title id="request-patch" className="t2">Patch Request</Title>
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


<Title id="request-delete" className="t2">Delete Request</Title>
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
