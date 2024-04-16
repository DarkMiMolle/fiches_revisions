<script lang="ts">
    import { Button, Input, Label } from 'flowbite-svelte'
</script>

<form method="POST">
    <Label for="Pseudo">Pseudo</Label>
    <Input type="text" id="Pseudo" name="pseudo" placeholder={"TODO"} />
	<Button type="submit" class="mt-10">log in</Button>
</form>
