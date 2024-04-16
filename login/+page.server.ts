import { env } from "$env/dynamic/public";
import type { Actions } from "@sveltejs/kit";

export const actions = {
    caca: async ({cookies, request, fetch}) => {
        const data = await request.formData()
        console.log(data)
        const tokenResp = await fetch(`${env.PUBLIC_BACKEND}/login`, {
            method: "POST",
            body: JSON.stringify({
                email: data.get("email"),
                password: "passw0rd"
            })
        })

        if (!tokenResp.ok) {
            console.log(await tokenResp.json())
            return
        }
        const resp = await tokenResp.json()
        cookies.set("jwt", resp["jwt"], {
            path: "/"
        })
        console.log(cookies.get("jwt"))
    }
} satisfies Actions