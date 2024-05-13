import { env } from '$env/dynamic/public'
import type { Collection } from '$lib'
import { CollectionsTest } from '$lib/test.js'


export async function load({cookies, parent}) {
    await parent()
    const jwt = cookies.get("jwt")!
    const endpoint = env.PUBLIC_BACKEND
    const resp = await fetch(`${endpoint}/api/collections`, {
        method: "GET",
        headers: {
            'Content-type': 'application/json',
            Authorization: `Bearer ${jwt}`,
        }
    })
    console.log({resp})
    const result = await resp.json()
    console.log({"valid resp": resp.ok, result})
    return {collections: CollectionsTest}
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