<script lang="ts">
    import { onMount, tick } from "svelte";
    import type { Form } from "$lib/data";

    export let data: Form

    let seeAnswer = false
    const turnCard = async () => {
        seeAnswer = !seeAnswer
        await tick()
    }
    let button: HTMLButtonElement,
        front: HTMLDivElement, back: HTMLDivElement

    let userAnswer: string = ""
    $: isGoodAnswer = userAnswer.toLowerCase() == data.Answer.toLowerCase()
    $: result = isGoodAnswer ? 'Bonne réponse' : 'Mauvaise réponse'

    onMount(() => {
        let frontHeight = front.getBoundingClientRect().height
        let backHeight = back.getBoundingClientRect().height
        let height = Math.trunc(frontHeight > backHeight ? frontHeight : backHeight) 
                        + 10 + "px"
        button.style.height = height 
        front.style.height = height 
        back.style.height = height 
    })
</script>

<button 
    on:click={turnCard}
    on:keypress={({key}) => {key == "Enter" ? turnCard() : null}}
    class="card"
    bind:this={button}
>
    <div 
        class="innerCard"
        class:seeAnswer
    >
        <div class="front" bind:this={front}>
            <p>{data.Question}</p>
            <input type="text" bind:value={userAnswer} on:click|stopPropagation={()=>{}}>
        </div>
        <div class="back" bind:this={back} class:isGoodAnswer>
            <p>{result}</p>
            {#if !isGoodAnswer && data.Tips != undefined}
                <p>Tips: {data.Tips}</p>
            {/if}
        </div>
    </div>
    
</button>

<style>
    .card {
        display: flex;
        flex-direction: column;
        width: 100%;
        height: 100%;
        min-width: 200px;
        max-width: 500px;
        background-color: blue;
        border-color: transparent;
        cursor: pointer;
        padding: 0px;
        font-family: "Roboto, Futura; Verdana; Arial";
        font-size: 24px;
    }
    .innerCard {
        position: relative;
        transform: rotateY(0deg);
        transition: transform 0.4s;
        transform-style: preserve-3d;
        height: 100%;
        user-select: none;
    }

    
    .innerCard.seeAnswer {
        transform: rotateY(180deg);
    }
    

    .front, .back {
		display: flex;
		align-items: center;
		justify-content: space-around;
        flex-direction: column;
		position: absolute;
		left: 0;
		top: -2px;
		backface-visibility: hidden;
		border-radius: 2em;
		border: 1px solid var(--fg-2);
		box-sizing: border-box;
		padding: 2em;
        width: 100%;

        
        min-height: 250px;
        max-height: 500px;
	}

    .front {
        background-color: cyan;
    }
    .back {
        background-color: rgb(230, 104, 104);
        transform: rotateY(180deg);
    }
    .isGoodAnswer {
        background-color: rgb(88, 231, 136);
    }
    input {
        width: 70%;
        height: 10%;
        
    }
</style>