<script>
    export let id;
    export let className = "t1";
</script>

<span class={`header ${className}`} id={id}>
    <a aria-hidden="true" tabindex="-1" class="hash-link" href="#{id}" title="Direct link to heading">
        <span class="title"><slot/></span>
        <span class="hash">#</span>
    </a>
</span>

<style>
    .hash {
        visibility: hidden;;
    }

    .t1 {
        font-size: 26px;
        font-weight: 600;
    }
    .t2 {
        font-size: 18px;
        font-weight: 600;
    }
    .header:hover {
        text-decoration: underline;
    }
    .header:hover .hash {
        visibility: initial;
    }

    .title {
        color: initial;
    }

</style>