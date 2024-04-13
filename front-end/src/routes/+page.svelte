<script>
    import Card from "$lib/Card.svelte"
    import {loadData, loadCollection} from "$lib/data";

    const data = loadData()

    
</script>
{#await loadCollection("Japonnais")}
    <h1>loading...</h1>
{:then data}
    <div class="container">
        <Card data={data.forms[0]}/>
    </div>
{:catch e}
    <p>{{e}}</p>
{/await}
yolo

<style>
    .container {
        display: flex;
        justify-content: space-around;
    }
</style>