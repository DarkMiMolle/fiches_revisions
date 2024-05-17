<script lang="ts">
    import type { Form } from "$lib"
    import { Card, Label, P } from "flowbite-svelte"
    
    const __filename = "content card.svelt"

    export let fiche: Form

    let answerView = false

    $: cardStyle = `transform: rotateY(${answerView ? -180 : 0}deg); transform-style: preserve-3d;`
</script>

<Card class={`my-2 cursor-pointer ${answerView ? "dark:bg-gray-900 bg-gray-50" : ""} duration-1000`} style={cardStyle} on:click={() => answerView = !answerView}>
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
        <div class="back p-4 sm:p-6">
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
        /* transition: transform 4s; */
        transform-style: preserve-3d;
        height: 100%;
        user-select: none;
        background-color: transparent;
    }

    .front, .back {
		display: flex;
		align-items: center;
		justify-content: space-between;
		backface-visibility: hidden;
		border-radius: 2em;
		border: 1px solid var(--fg-2);
		box-sizing: border-box;
		padding: 0;
	}

    .front {
        background-color: transparent;
    }

    .innerCard.answerView > .front {
        height: 0px;
        overflow: hidden;
    }
    .back {
        width: 100%;
        height: 0;
        transform: rotateY(180deg);
        overflow: hidden;
    }

    .innerCard.answerView > .back {
        height: 100%;
        overflow: visible;
    }
</style>