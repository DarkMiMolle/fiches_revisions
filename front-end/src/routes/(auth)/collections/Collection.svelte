<script lang="ts">
    import type { Collection } from "$lib";
    import { Card, Heading, Hr, List, Li, P, } from "flowbite-svelte";

    export let collection: Collection
</script>

<Card> 
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
        <Li class="pb-3 sm:pb-4">

        </Li>
    </List>
</Card>