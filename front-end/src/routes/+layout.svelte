<script lang="ts">
import "../app.pcss";

import { Navbar, NavBrand, NavLi, NavUl, NavHamburger } from "flowbite-svelte";

export let data: {
    isAuth: boolean
}


</script>

<Navbar rounded color="form" class="px-2 sm:px-4 py-2.5 sticky w-full z-20 top-0 start-0 border-b mb-5">
    <NavBrand href="/">
    <!-- <img src="" class="me-3 h-6 sm:h-9" alt="Flowbite Logo" /> -->
    <span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white">Fiches - révision</span>
    </NavBrand>
    <NavHamburger />
    <NavUl>
        <NavLi href="/" active={true}>Accueil</NavLi>
        <NavLi href="/about">A propos</NavLi>
        <NavLi href="/collections">Collections</NavLi>
        <NavLi href="/">S'entrainer</NavLi>
        <NavLi href={data.isAuth ? "/logout" : "/login"}>{data.isAuth ? "Logout" : "Login"}</NavLi>
    </NavUl>
</Navbar> 
<div class="px-5">
    <slot></slot>
</div>
