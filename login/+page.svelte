<script lang="ts">
    import { Input, Label, Helper, Button, Checkbox, A } from 'flowbite-svelte';
</script>
  
<form method="POST">
    <div class="mb-6">
        <Label for="email" class="mb-2">Email address</Label>
        <Input name="email" type="email" id="email" placeholder="john.doe@company.com" required />
      </div>
      <div class="mb-6">
        <Label for="password" class="mb-2">Password</Label>
        <Input name="password" type="password" id="password" placeholder="•••••••••" required />
      </div>
      <Button type="submit">Submit</Button>
</form>