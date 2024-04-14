import {env} from "$env/dynamic/public";

export function GET() {
        console.log("HERE (server)")
        const url = `${env.PUBLIC_BACKEND}/collection?collection=${"Japonnais"}`
        console.log(url)
        return fetch(url)
}