<script>
    import Title from "../component/Title.svelte";
    import CLI from "./usages/CLI.svelte";
    import Req from "./usages/Requests.svelte";
    import Docker from "./usages/Docker.svelte";

    // const urlParams = new URLSearchParams(window.location.search);
    let param = window.location.href.split('#').pop().split('-')[0];
    switch (param) {
        case "request":
            param = "Req";
            break;
        case "docker":
            param = "Docker";
            break;
        default:
            param = "CLI";
            break;
    }
    $: select = param;

    export let version;
</script>

<div>
    <div class="button-group">
        <button class={select == "CLI" ? "selected" : undefined} on:click|stopPropagation={() => select = "CLI"}>
            Command Line
        </button>
        <button class={select == "Req" ? "selected" : undefined} on:click|stopPropagation={() => select = "Req"}>
            Requests
        </button>
        <button class={select == "Docker" ? "selected" : undefined} on:click|stopPropagation={() => select = "Docker"}>
            Docker Image Build
        </button>
    </div>
    <div class={select == "CLI" ? undefined : "invisible"}>
        <CLI/>
    </div>
    <div class={select == "Req" ? undefined : "invisible"}>
        <Req/>
    </div>
    <div class={select == "Docker" ? undefined : "invisible"}>
        <Docker version={version}/>
    </div>
</div>

<style>
    .invisible {
        display: none;
    }
    .button-group {
        border-bottom: 5px solid #333;
        margin: 1em 0 0 0;
        display: flex;
        justify-content: space-around;
    }
    .button-group button {
        border: unset;
        border-radius: unset;
        margin: 0;
        width: 100%;
        padding: 0.4em 1em;
        font-weight: 700;
        font-size: x-large;
    }
    .button-group button:hover {
        cursor: pointer;
        text-decoration: underline;
    }
    .button-group button + .button-group button {
        margin: 0;
        padding: 0;
    }
    .selected {
        text-decoration: underline;
        background-color: #333;
        color: #fff;
    }
</style>