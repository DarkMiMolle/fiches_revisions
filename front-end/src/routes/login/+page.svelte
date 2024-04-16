<script lang="ts">
    import { Button, Input, Label } from 'flowbite-svelte'
</script>

<form method="POST">
    <Label for="Pseudo">Pseudo</Label>
    <Input type="text" id="Pseudo" name="pseudo"/>
	<Button type="submit">log in</Button>
</form>
