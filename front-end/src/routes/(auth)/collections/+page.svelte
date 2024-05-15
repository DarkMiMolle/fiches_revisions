<script lang="ts">
    import { goto, invalidateAll } from "$app/navigation";
    import type { Collection } from "$lib";
    import { Button, Dropdown, P, Radio, Search, Spinner } from "flowbite-svelte";
    import { ChevronDownOutline, UserRemoveSolid } from 'flowbite-svelte-icons';
    import{ default as CollectionCard } from "./Collection.svelte";

    export let data: {error?: any, collections: Collection[], done: Promise<void>}

    let hasErr = data.error != undefined

    let currentCollection: number|undefined

    let search: string = ""

    $: collectionNames = (data.collections ?? new Array<Collection>())
                            .map(collection => collection.name)
                            .filter(name => search === "" || name.toLowerCase().startsWith(search.toLowerCase()))
    $: console.log("collection.page", {currentCollection, collectionNames})

    let isDropdownOpen: boolean = false
</script>

{#await data.done}
<div class="dark:bg-gray-800 flex justify-center items-center absolute opacity-85 w-dvw h-dvh m-0 p-0" style="z-index: 10000000">
    <Spinner size=10/>
</div>
{:then} 
    {#if hasErr}
        <P>{ JSON.stringify(data.error)}</P>
        <Button on:click={() => { invalidateAll(); goto("/") } }>Recharge la page</Button>
    {:else}
        <Button>Selectionne une collection <ChevronDownOutline class="w-6 h-6 ms-2 text-white dark:text-white" /></Button>
        <Dropdown bind:open={isDropdownOpen} class="overflow-y-auto px-3 pb-3 text-sm h-44" placement="bottom-start">
        <div slot="header" class="p-3">
            <Search size="md" bind:value={search}/>
        </div>
        {#each collectionNames as collectionName, idx}
        <li class="rounded hover:bg-gray-100 dark:hover:bg-gray-600">
            <Radio name="collection" bind:group={currentCollection} value={idx}>{collectionName}</Radio>
        </li>
        {/each}
        <div slot="footer" class="flex justify-center">
            <Button outline disabled={currentCollection === undefined} on:click={() => {
                currentCollection = undefined
                setTimeout(() => isDropdownOpen=false, 500)
            }}>Deselectionner</Button>
        </div>

        </Dropdown>
    {/if}

    {#if currentCollection !== undefined}
        <div class="flex justify-center mt-5">
            <CollectionCard collection={data.collections[currentCollection]}/>
        </div>
    {/if}
{/await}


<style>
    
</style>