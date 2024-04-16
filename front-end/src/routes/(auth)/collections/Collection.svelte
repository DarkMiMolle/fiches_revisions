<script lang="ts">
    import type { Collection } from "$lib";
    import { Card, Heading, Hr, List, Li, P, GradientButton, } from "flowbite-svelte";
    import { EditSolid, PlaySolid } from "flowbite-svelte-icons";

    export let collection: Collection
</script>

<Card size={"md"}> 
    <!-- ajouter des tabs: 1: info; 2: fiches; 3: réglage -->
    <Heading tag="h5">{collection.name}</Heading>
    <Hr/>
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
</Card>