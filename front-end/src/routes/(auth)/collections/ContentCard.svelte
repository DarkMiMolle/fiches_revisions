<script lang="ts">
    import type { Form } from "$lib"
    import Separator from "$lib/Separator.svelte"
    import { Card, Indicator, Label, P, Tooltip } from "flowbite-svelte"
    import { createEventDispatcher } from "svelte"
   
    const __filename = "content card.svelt"

    export let cardClass: string = ""
    export let fiche: Form

    let initialFicheJSON = JSON.stringify(fiche)
    $: saved = initialFicheJSON == JSON.stringify(fiche)
    let prevState = true
    const dispatcher = createEventDispatcher()
    $: {
        if (saved != prevState) {
            console.log(__filename, {synchro: saved})
            dispatcher("synchro", saved)
            prevState = saved
        }
    }
    
    export let inputSave = false
    $: {
        if (inputSave) {
            initialFicheJSON = JSON.stringify(fiche)
            setTimeout(() => inputSave = false, 1)
        }
    }

    let answerView = false

    $: cardStyle = `transform: rotateY(${answerView ? -180 : 0}deg); transform-style: preserve-3d;`

    function editable(node: HTMLElement) { // , update: (newVal: (data: string) => string) => void
        function editionClick(ev: MouseEvent): any {
            ev.stopPropagation()
            node.classList.add('editing')
            node.addEventListener('keypress', function edit (ev: KeyboardEvent): any {
                node.removeEventListener('click', editionClick)
                if (ev.key == 'Enter' && !ev.shiftKey) {
                    node.blur()
                    node.removeEventListener('keypress', edit)
                    return
                }
            })
            node.addEventListener('blur', function onBlur(ev) {
                node.classList.remove('editing')
                node.addEventListener('click', editionClick)
                node.removeEventListener('blur', onBlur)
            })
        }
        node.addEventListener('click', editionClick)
        return {
            destroy() {
                node.removeEventListener('click', editionClick)
            }
        }
    }
    
    let tips = fiche.tips ?? 'ø'
    $: fiche.tips = tips == 'ø' ? undefined : tips
    $: {
        if (tips == '') {
            tips = 'ø'
        }
    }

    export function Save() {
        initialFicheJSON = JSON.stringify(fiche)
    }

    export function IsSaved(): boolean {
        return saved
    }


</script>
<Card class={`my-2 cursor-pointer ${answerView ? "dark:bg-gray-900 bg-gray-50" : ""} duration-1000 ${cardClass}`} style={cardStyle} on:click={() => answerView = !answerView}>
    {#if !saved}
    <Indicator color="orange" placement={answerView ? "top-left" : "top-right"}/>
    <Tooltip placement="right">Il faut sauvegarder les changements</Tooltip>
    {/if}
    <div class="innerCard" class:answerView>
        <div class="front flex-col">
            <div class="flex justify-between w-full">
                <Label>Question:</Label>
                <P><div class="cursor-context-menu px-[5px]" use:editable contenteditable="true" bind:textContent={fiche.question}></div></P>
            </div>
            <div class="flex justify-between w-full">
                <Label>Tips:</Label>
                <P><div class="cursor-context-menu px-[5px]" use:editable contenteditable="true" bind:textContent={tips} on:focus|stopPropagation={() => { setTimeout(() => {if (tips == 'ø') tips = '\0'}, 100)}} on:blur={() => {if (tips == '\0') tips = ''}}> {tips}</div></P>
            </div>
            <Separator darkColor="gray-700"/>
            <div class="flex justify-between w-full">
                <Label>Réviser le</Label>
                <P>{fiche.lastTry.toLocaleDateString()}</P>
            </div>
            <div class="flex justify-between w-full">
                <Label>Prochaine révision le</Label>
                <P>{fiche.nextTry.toLocaleDateString()}</P>
            </div>
            <div class="flex justify-between w-full">
                <Label>Niveau de maitrise</Label>
                <P>{fiche.level}</P>
            </div>
        </div>
        <div class="back p-4 sm:p-6">
            <Label>Réponses:</Label>
            <div class="flex flex-col justify-between">
                {#each fiche.answers as answer}
                <P><div class="cursor-context-menu px-[5px]" use:editable contenteditable="true" bind:textContent={answer}></div></P>
                {/each}
            </div>
        </div>
    </div>
</Card>

<style>

    .editing {
        border: black;
        border-width: 1px;
        padding: 2px;
    }

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
		border: 1px solid transparent;
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
        /* max-height: 70px;
        scrollbar-color: rgb(46, 62, 123) rgb(35, 29, 82) */
    }

    .innerCard.answerView > .back {
        height: 100%;
        overflow: visible;
        /* overflow-y: scroll; */
    }
</style>