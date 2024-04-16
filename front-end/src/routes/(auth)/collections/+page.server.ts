import { env } from '$env/dynamic/public'
import type { Collection } from '$lib'


export async function load({cookies, parent}) {
    await parent()
    const jwt = cookies.get("jwt")!
    const endpoint = env.PUBLIC_BACKEND
    const resp = await fetch(`${endpoint}/collections`, {
        method: "GET",
        headers: {
            Authorization: `Bearer ${jwt}` 
        }
    })
    const result = await resp.json()

    console.log(result)
    if (!resp.ok) {
        return {
            error: result,
            collection: []
        }
    }

    return {
        collections: result as Collection[]
    }
}