<script lang="ts">
    import type { Form } from "$lib"
    import { Card, Label, P } from "flowbite-svelte"

    export let fiche: Form

    let answerView = false
</script>

<Card class={`my-2 cursor-pointer ${answerView ? "dark:bg-gray-900 bg-gray-50" : ""}`} on:click={() => answerView = !answerView}>
    <div class="innerCard" class:answerView>
        <div class="front flex-col">
            <div class="flex justify-between w-full">
                <Label>Question:</Label>
                <P>{ fiche.question }</P>
            </div>
            <div class="flex justify-between w-full">
                <Label>Tips:</Label>
                <P>{ fiche.tips ?? "ø"}</P>
            </div>
                
        </div>
        <div class="back p-4 sm:p-6" >
            <Label>Réponses:</Label>
            <div class="flex flex-col justify-between">
                {#each fiche.answers as aswer}
                <P>{ aswer }</P>
                {/each}
            </div>
        </div>
    </div>
</Card>

<style>
    .innerCard {
        position: relative;
        transform: rotateY(0deg);
        transition: transform 0.4s;
        transform-style: preserve-3d;
        height: 100%;
        /* width: 100%; */
        user-select: none;
    }
    .innerCard.answerView {
        transform: rotateY(180deg);
    }
    

    .front, .back {
		display: flex;
		align-items: center;
		justify-content: space-between;
		left: 0;
        top: 0;
		backface-visibility: hidden;
		border-radius: 2em;
		border: 1px solid var(--fg-2);
		box-sizing: border-box;
		padding: 0;
        /* width: 100%; */

/*         
        min-height: 250px;
        max-height: 500px; */
	}

    .front {
        background-color: transparent;
    }
    .back {
        position: absolute;
        width: 100%;
        height: 100%;
        transform: rotateY(180deg);
    }
</style>