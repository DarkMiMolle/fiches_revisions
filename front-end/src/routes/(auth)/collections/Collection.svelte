<script lang="ts">
    import type { Collection } from "$lib";
    import { Card, Heading, Hr, List, Li, P, GradientButton, Tabs, TabItem, Indicator } from "flowbite-svelte";
    import { AdjustmentsVerticalSolid, CloudArrowUpSolid, EditSolid, FileLinesSolid, InfoCircleSolid, PlaySolid } from "flowbite-svelte-icons";
    import ContentCard from "./ContentCard.svelte"
    import { fade } from "svelte/transition"

    const __filename = "collection.svelte"

    export let collection: Collection
    const initialCollectionForms = JSON.stringify(collection.forms)

    $: collectionFormsUpdated = initialCollectionForms != JSON.stringify(collection.forms)

    function dring(node: SVGSVGElement, {duration = 500}) {
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
            }
        }
    }

</script>

<Card size={"md"}>
    <div class="flex justify-space">
        <Heading tag="h5">{collection.name}</Heading>
        {#if collectionFormsUpdated}
        <svg in:dring={{}} out:fade class="w-[26px] h-[26px] text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24">
            <path d="M13.383 4.076a6.5 6.5 0 0 0-6.887 3.95A5 5 0 0 0 7 18h3v-4a2 2 0 0 1-1.414-3.414l2-2a2 2 0 0 1 2.828 0l2 2A2 2 0 0 1 14 14v4h4a4 4 0 0 0 .988-7.876 6.5 6.5 0 0 0-5.605-6.048Z"/>
            <path d="M12.707 9.293a1 1 0 0 0-1.414 0l-2 2a1 1 0 1 0 1.414 1.414l.293-.293V19a1 1 0 1 0 2 0v-6.586l.293.293a1 1 0 0 0 1.414-1.414l-2-2Z"/>
        </svg>

        {/if}
    </div>
    <Tabs tabStyle="underline">
        <TabItem open>
            <span slot="title" class="flex items-center gap-2 p-[2px] md:p-5">
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
                    <GradientButton color="purpleToPink">Editer la collection</GradientButton>
                    <GradientButton color="tealToLime">S'entrainer avec cette collection</GradientButton>
                </Li>
                <Li class="pb-4 md:hidden flex justify-around ">
                    <GradientButton outline color="purpleToPink"><EditSolid color="Pink" size="lg"/></GradientButton>
                    <GradientButton outline color="tealToLime"><PlaySolid color="Green" size="lg"/></GradientButton>
                    
                </Li>
            </List>
        </TabItem>
        <TabItem>
            <span slot="title" class="flex items-center gap-2">
                <FileLinesSolid size="md" />
                Fiches
            </span>
            <div class="flex flex-col justify-between items-center">
                {#each collection.forms as form}
                    <ContentCard bind:fiche={form}/>
                {/each}
            </div>
        </TabItem>
        <TabItem>
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