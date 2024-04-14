import { env } from "$env/dynamic/public";
import type { Actions } from "@sveltejs/kit";

export const actions = {
    default: async ({cookies, request}) => {
        const data = await request.formData()
        const tokenResp = await fetch(`${env.PUBLIC_BACKEND}/login`, {
            method: "POST",
            body: JSON.stringify({
                email: data.get("email"),
                password: data.get("password")
            })
        })

        if (!tokenResp.ok) {
            return
        }
        cookies.set("jwt", await tokenResp.json(), {
            path: "/"
        })
    }
} satisfies Actions