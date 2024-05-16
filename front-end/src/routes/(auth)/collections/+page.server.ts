import { env } from '$env/dynamic/public'
import type { Collection, ServerError } from '$lib'
import { HttpCode } from '$lib/statusCode.js'
import { CollectionsTest } from '$lib/test.js'
import { redirect } from '@sveltejs/kit'

const __filename = "collection.server"

export async function load({cookies, parent, url}) {
    await parent()
    const jwt = cookies.get("jwt")!
    const endpoint = env.PUBLIC_BACKEND

    console.log(__filename, {jwt})

    const fetchResp = fetch(`${endpoint}/api/collections`, {
        method: "GET",
        headers: {
            'Content-type': 'application/json',
            Authorization: `Bearer ${jwt}`,
        }
    })

    return {
        done: new Promise<void>((resolver) => {
            fetchResp.then(() => resolver())
        }),
        ...await (async () => {
            const resp = await fetchResp
            const result = await resp.json()
            if (!resp.ok) {
                if ((result as ServerError).status === HttpCode.UNAUTHORIZED) {
                    console.log(__filename, {result})
                    cookies.delete("jwt", {path: '/'})
                    throw redirect(HttpCode.TEMPORARY_REDIRECT, `/login?redirectTo=${url.pathname}`)
                }
                return {
                    error: result,
                    collections: [] as Collection[]
                }
            }
            return {
                collections: CollectionsTest
            }
        })()
    }
}