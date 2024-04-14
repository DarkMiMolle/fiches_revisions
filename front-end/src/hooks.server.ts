import { redirect, type Handle } from "@sveltejs/kit";

export const handle: Handle = async ({event, resolve}) => {
    if (event.url.pathname.startsWith("/login")) {
        return resolve(event)
    }
    event.locals.token = event.cookies.get("jwt")
    console.log("hook -", event.locals.token)
    if (event.locals.token === undefined) {
        console.log("redirecting")
        throw redirect(303, "/login")
    }
    
    return resolve(event)
}