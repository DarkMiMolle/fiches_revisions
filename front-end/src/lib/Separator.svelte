<script lang="ts">
    export let width: number = 2
    export let lightColor: string = 'gray-100'
    export let darkColor: string = 'gray-900'

    export let classNames: string = ''
    classNames = `border-[${width}px] border-${lightColor} dark:border-${darkColor} ` + classNames
</script>

<div class={`separator ${classNames}`}></div>

<style>
    .separator {
        min-width: 40px;
        border-width: 1px;
        border-radius: 5px;
        width: 100%;
        height: 1px;
        margin-top: 2px;
        margin-bottom: 2px;
    }
</style>