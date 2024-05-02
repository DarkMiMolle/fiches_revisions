<script lang="ts">
    import type { Collection } from "$lib";
    import { Card, Heading, Hr, List, Li, P, GradientButton, Tabs, TabItem } from "flowbite-svelte";
    import { AdjustmentsVerticalSolid, EditSolid, FileLinesSolid, InfoCircleSolid, PlaySolid } from "flowbite-svelte-icons";
    import ContentCard from "./ContentCard.svelte"

    export let collection: Collection
</script>

<Card size={"md"}> 
    <!-- ajouter des tabs: 1: info; 2: fiches; 3: réglage -->
    <Heading tag="h5">{collection.name}</Heading>
    <Tabs tabStyle="underline">
        <TabItem open>
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
                    <ContentCard fiche={form}/>
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