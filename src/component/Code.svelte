<script>
    export let code;
    export let button = false;
    export let wrap = false;
    let copied = false;
    const copy = () => {
        navigator.clipboard.writeText(code).then(() => {
            copied = true;
            setTimeout(() => copied = false, 500);
        }, () => {
            console.warn("failed copy");
        });
    }
    $: if (!button) {
        let lines = code.replace(/&/g, "&amp;").replace(/>/g, "&gt;").replace(/</g, "&lt;").replace(/"/g, "&quot;").split('\n');
        code = "";
        for(var line = 0; line < lines.length; line++){
            if (lines[line][0] === '$') {
                code += '<span class="color">' + lines[line] + '</span>\n';
            } else {
                code += lines[line] + "\n";
            }
        }
    }
</script>

<pre class:wrap={wrap}>
    {#if button}
    <button on:click={copy}>
        {#if copied}
        <svg version="1.1" viewBox="0 0 12 16" width="12" height="16" aria-hidden="true" class="octicon octicon-check"><path fill-rule="evenodd" d="M12 5l-8 8-4-4 1.5-1.5L4 10l6.5-6.5L12 5z"></path></svg>
        {:else}
        <svg aria-hidden="true" height="16" version="1.1" viewBox="0 0 14 16" width="14" class="octicon octicon-clippy"><path d="M2 13h4v1H2v-1zm5-6H2v1h5V7zm2 3V8l-3 3 3 3v-2h5v-2H9zM4.5 9H2v1h2.5V9zM2 12h2.5v-1H2v1zm9 1h1v2c-.02.28-.11.52-.3.7-.19.18-.42.28-.7.3H1c-.55 0-1-.45-1-1V4c0-.55.45-1 1-1h3c0-1.11.89-2 2-2 1.11 0 2 .89 2 2h3c.55 0 1 .45 1 1v5h-1V6H1v9h10v-2zM2 5h8c0-.55-.45-1-1-1H8c-.55 0-1-.45-1-1s-.45-1-1-1-1 .45-1 1-.45 1-1 1H3c-.55 0-1 .45-1 1z" fill-rule="evenodd"></path></svg>
        {/if}
    </button>
    {/if}
    <code>{@html code}</code>
</pre>

<style>
    .wrap {
        white-space: break-spaces;
        word-break: break-all;
    }
    svg {
        width: 16px;
    }

    button {
        margin: unset;
        padding: 8px;
        border: unset;
        border-radius: unset;
        border-right: solid 1px #333;
        background-color: #efefef;
        cursor: pointer;
    }

    button:hover {
        background-color: #ddd;
    }

    code {
        padding: 16px;
        overflow: auto;
    }

    pre {
        background-color: #f6f8fa;
        text-align: left;
        display: flex;
        font-size: 12px;
    }
</style>
