<script lang="ts">
    import { page } from '$app/stores'
    import type { ServerError } from '$lib'
    import { Button, ButtonGroup, Input, InputAddon, Label, AccordionItem, Accordion, Radio, Helper, P, Spinner, Tooltip } from 'flowbite-svelte'
    import { UserCircleSolid, LockSolid, EnvelopeSolid, EyeOutline, EyeSlashOutline, UserSolid } from 'flowbite-svelte-icons'

    const __filename = "login.page"

    let notifEnable = String(false)
    $: notification = notifEnable == String(true)

    let pseudo = ""

    let pseudoHelper: boolean

    $: regexTestPseudo = /^[a-zA-Z_]+[a-zA-Z_0-9\-]*$/.test(pseudo)

    $: pseudoHelper = !regexTestPseudo && pseudo != ""

    $: validPseudo = regexTestPseudo && pseudo.length >= 6


    let signup = false

    let showPwd = false

    export let formResult: ServerError|undefined = undefined

    let htmlForm: HTMLFormElement

    let serverWorking = false

    $: {
        if (formResult != undefined) {
            serverWorking = false
        }
    }
</script>

{#if serverWorking}
<div class="dark:bg-gray-800 flex justify-center items-center absolute opacity-85 w-dvw h-dvh m-0 p-0" style="z-index: 10000000">
    <Spinner size=10/>
</div>
{/if}

<!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
<form bind:this={htmlForm} method="POST" action={`?/${signup ? "signup" : "login"}`} on:submit={() => {
    showPwd = false
    serverWorking = true
    formResult = undefined
}} on:keypress|stopPropagation={({key}) => {
    if (key != "Enter") return
    htmlForm.submit()
}}>
    <input type="text" name="redirection" style="display: none" value={$page.url.searchParams.get("redirectTo")}>
    <Label for="pseudo-login" class="block mb-2 mt-4">Pseudo <span class="text-primary-600">*</span></Label>
    <ButtonGroup class="w-full">
        <InputAddon >
            <UserCircleSolid class="w-4 h-4 text-gray-500 dark:text-gray-400"/>
        </InputAddon>
        <Input id="pseudo-login" name="pseudo" bind:value={pseudo} required outline color={pseudoHelper ? 'red':'base'}/>
    </ButtonGroup>
    {#if pseudoHelper || formResult != undefined}
    <Helper color='red'>
        Un pseudo doit commencer par une lettre et doit avoir au moins 6 caractères parmis: des lettres de base, des chiffres, - ou _
    </Helper>
    {/if}
    <Label for="pwd-login" class="block mb-2 mt-4">Mot de passe <span class="text-primary-600">*</span></Label>
    <ButtonGroup class="w-full">
        <InputAddon>
            <LockSolid class="w-4 h-4 text-gray-500 dark:text-gray-400"/>
        </InputAddon>
        <Input id="pwd-login" class="first:rounded-s-none first:border-s-0" name="password" type={showPwd ? "text" : "password"} required placeholder={showPwd ? "" : "••••••••"}>
            <!-- svelte-ignore a11y-click-events-have-key-events -->
            <!-- svelte-ignore a11y-no-static-element-interactions -->
            <div slot="right" on:click={() => {
                // Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc felis magna, tempor eu purus vel
                // L.olreevm  siuprsuupm  udeo lro...rp mseitt  ,aamnegta,m  csoinlseefc tcentuuNr  .atdiilpei sgcni
                showPwd = !showPwd
            }}>
                {#if showPwd}
                    <EyeOutline class="w-6 h-6" />
                {:else}
                    <EyeSlashOutline class="w-6 h-6" />
                {/if}
        </div>
        </Input>
    </ButtonGroup>
    <Accordion class="mt-4">
        <AccordionItem bind:open={signup}>
            <span slot="header">formulaire d'inscription</span>
            <Label for="pwd-singup" class="block mb-2 mt-4">confirmation du mot de passe <span class="text-primary-600">*</span></Label>
            <ButtonGroup class="w-full">
                <InputAddon>
                    <LockSolid class="w-4 h-4 text-gray-500 dark:text-gray-400"/>
                </InputAddon>
                <Input id="pwd-signup" class="first:rounded-s-none first:border-s-0" type={showPwd ? "text" : "password"} required={signup} placeholder={showPwd ? "" : "••••••••"}>
                    <button slot="right" on:click={() => showPwd = !showPwd}>
                        {#if showPwd}
                            <EyeOutline class="w-6 h-6" />
                        {:else}
                            <EyeSlashOutline class="w-6 h-6" />
                        {/if}
                    </button>
                </Input>
            </ButtonGroup>
            <Label for="name" class="block mb-2 mt-4">Nom</Label>
            <ButtonGroup class="w-full">
                <InputAddon>
                    <UserSolid class="w-4 h-4 text-gray-500 dark:text-gray-400"/>
                </InputAddon>
                <Input name="name"/>
            </ButtonGroup>
            <Label for="email" class="block mb-2 mt-4">email</Label>
            <ButtonGroup class="w-full">
                <InputAddon>
                    <EnvelopeSolid class="w-4 h-4 text-gray-500 dark:text-gray-400"/>
                </InputAddon>
                <Input id="email" name="email" placeholder={`${pseudo}@gmail.com`}/>
            </ButtonGroup>
            <Label for="notif" class="block mt-4">recevoir des notifications par email ?</Label>
            <div class="flex gap-3 w-min" id="notification-radio">
                <Radio group={notifEnable} value="true" name="notif">Oui</Radio>
                <Radio group={notifEnable} value="false" name="notif">Non</Radio>
            </div>
            <Tooltip  class="ml-5" triggeredBy="#notification-radio" placement="right">cette fonctionnalité n'est pas encore disponible</Tooltip>
        </AccordionItem>
    </Accordion>
    <P class="mt-4 mb-4" size="xs"><span class="text-primary-600">*</span> champs requis</P>
	<Button type="submit" class="mt-10">{signup ? "Sign up" : "Log in"}</Button>
</form>
