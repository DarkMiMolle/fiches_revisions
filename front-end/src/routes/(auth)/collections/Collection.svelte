<script lang="ts">
    import type { Collection } from "$lib";
    import { Card, Heading, Hr, List, Li, P, GradientButton, Tabs, TabItem, Indicator, Button } from "flowbite-svelte";
    import { AdjustmentsVerticalSolid, CloudArrowUpSolid, EditSolid, FileLinesSolid, InfoCircleSolid, PlaySolid } from "flowbite-svelte-icons";
    import ContentCard from "./ContentCard.svelte"
    import { fade } from "svelte/transition"
    import Separator from "$lib/Separator.svelte"

    const __filename = "collection.svelte"

    export let collection: Collection
    const initialCollection = JSON.parse(JSON.stringify(collection)) as Collection

    let cardElements: ContentCard[] = Array<ContentCard>(collection.forms.length)

    let formChange = 0
    $: collectionFormsUpdated = formChange > 0 // initialCollectionForms != JSON.stringify(collection.forms)

    function dring(node: SVGSVGElement, {duration = 500}: {duration?: number}) {
        let prevTime = 0
        let direction = 1
        return {
            css: (t: number): string => {
                let time = t * duration
                if (Math.abs(time - prevTime) > 70) {
                    direction = -direction
                    prevTime = time
                }
                return `
                transform: rotateZ(${direction * (20 - Math.random()*10)}deg);
                scale: ${1 + Math.sin(3*t)}
                `
            },
        }
    }

    function slideTop(node: HTMLDivElement, {duration = 1000}: {duration?: number}) {
        const parent = node.parentElement
        
        return {
            css: (t: number) => {
                return `position: relative; top: -${(1-t)*100}%`
            },
            duration,
        }
    }

    let edit = false
    type TabName = 'info'|'fiche'|'reglage'
    let currentTab: TabName = 'info'

    $: infoOpen = currentTab === 'info'
    $: {
        if (!infoOpen) {
            edit = false
        }
    }

    let editedForms: boolean[] = collection.forms.map(() => false)
    let toSave = collection.forms.map(() => false)
    
    function updateSynchro(index: number): ({detail}: {detail: boolean}) => void {
        return ({detail}) => {
            if (detail) {
                formChange > 0 ? formChange-- : null
            } else {
                formChange++
            }
            editedForms[index] = !detail
        }
    }

</script>

<div class="flex flex-col justify-around w-full items-center overflow-hidden">
    <Card size={"md"} style="margin-bottom: 2%; z-index: 2">
        <div class="flex justify-space">
            <Heading tag="h5">{collection.name}</Heading>
            {#if collectionFormsUpdated}
            <svg on:click={() => {
                console.log(__filename, "call to backend to save data")
                for(let index in editedForms) {
                    if (editedForms[index]) {
                        toSave[index] = true
                        editedForms[index] = false
                    }
                }
                formChange = 0
            }} in:dring out:fade class="cursor-pointer w-[26px] h-[26px] text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24">
                <path d="M13.383 4.076a6.5 6.5 0 0 0-6.887 3.95A5 5 0 0 0 7 18h3v-4a2 2 0 0 1-1.414-3.414l2-2a2 2 0 0 1 2.828 0l2 2A2 2 0 0 1 14 14v4h4a4 4 0 0 0 .988-7.876 6.5 6.5 0 0 0-5.605-6.048Z"/>
                <path d="M12.707 9.293a1 1 0 0 0-1.414 0l-2 2a1 1 0 1 0 1.414 1.414l.293-.293V19a1 1 0 1 0 2 0v-6.586l.293.293a1 1 0 0 0 1.414-1.414l-2-2Z"/>
            </svg>
    
            {/if}
        </div>
        <Tabs tabStyle="underline">
            <TabItem bind:open={infoOpen} disabled={formChange > 0} on:click={() => currentTab = 'info'}>
                <span slot="title" class="flex items-center gap-2">
                    <InfoCircleSolid size="md" />
                    Info
                </span>
                <List tag="ul" list="none">
                    <Li class="pb-3 sm:pb-4">
                        <div class="flex items-center justify-between space-x-4 rtl:space-x-reverse">
                            <P>Nombre de question</P>
                            <P class="inline-flex items-center">{collection.forms.length}</P>
                        </div>
                    </Li>
                    <Li class="pb-3 sm:pb-4">
                        <div class="flex items-center justify-between space-x-4 rtl:space-x-reverse">
                            <P>tout un tas d'information</P>
                            <P class="inline-flex items-center">{collection.forms.length}</P>
                        </div>
                    </Li>
                    <Li class="pb-3 max-md:hidden flex justify-around ">
                        <GradientButton color="purpleToPink" on:click={() => {edit = true}}>Editer la collection</GradientButton>
                        <GradientButton color="tealToLime">S'entrainer avec cette collection</GradientButton>
                    </Li>
                    <Li class="pb-4 md:hidden flex justify-around ">
                        <GradientButton outline color="purpleToPink" class="duration-100" on:click={() => {edit = true}}><EditSolid color="Pink" size="lg"/></GradientButton>
                        <GradientButton outline color="tealToLime"><PlaySolid color="Green" size="lg"/></GradientButton>
                        
                    </Li>
                </List>
            </TabItem>
            <TabItem disabled={formChange > 0} on:click={() => currentTab = 'fiche'}>
                <span slot="title" class="flex items-center gap-2">
                    <FileLinesSolid size="md" />
                    Fiches
                </span>
                <div class="flex flex-col justify-between items-center">
                    {#each collection.forms as form, index}
                        <ContentCard bind:fiche={form} bind:inputSave={toSave[index]} on:synchro={updateSynchro(index)} />
                    {/each}
                </div>
            </TabItem>
            <TabItem disabled={formChange > 0} on:click={() => currentTab = 'reglage'}>
                <span slot="title" class="flex items-center gap-2">
                    <AdjustmentsVerticalSolid size="md" />
                    Réglages
                </span>
                <div>
                    mes réglages
                </div>
            </TabItem>
        </Tabs>
    </Card>
    
    {#if edit}
        <div in:slideTop out:slideTop={{duration: 250}} style="z-index: 1;" class="w-full flex justify-center">
            <Card size="md" class="flex flex-col">
                <P>Filtre, recherche</P>
                <Separator/>
                <GradientButton style="margin-bottom: 5%;">Nouvelle fiche</GradientButton>
                {#each collection.forms as form, index }
                    <ContentCard bind:fiche={form} bind:inputSave={toSave[index]} cardClass="self-center" on:synchro={updateSynchro(index)}/>
                    <div class="self-center flex w-full justify-around" style="margin-bottom: 2%; padding: 0 5%;">
                        <Button>Supprimer</Button>
                        {#if editedForms[index]}
                            <Button>Annuler</Button>
                            <Button>Sauvgarder</Button>
                        {/if}
                    </div>
                    
                {/each}
            </Card>
        </div>
    {/if}
</div>

