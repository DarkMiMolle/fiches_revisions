import { Sleep } from "$lib/index"
import {env} from "$env/dynamic/public";

type Opt<T> = T | undefined

export class Form {
    constructor(
        public Question: string, 
        public Answer: string,
        public Tips: Opt<string> = undefined
    ) {}
}

export type Collection = {
    name: string
    forms: Form[]
}

const datas: Collection = {
    name: "Test-local",
    forms: [
        {
            Question: "Signification de 耳",
            Answer: "Oreille",
            Tips: "une partie de la tête"
        }
        ]
}

export async function loadData(): Promise<Collection> {
    console.log("HERE")
    const url = `${env.PUBLIC_BACKEND}/collection?collection=${"Japonnais"}`
    console.log(url)
    try {
        const resp = await fetch("/api/")
        console.log({resp})
        console.log(await resp.json())
    } catch (e) {
        console.log({error: e})
    }

    await Sleep(100)
    return datas
}

export async function loadCollection(fromCollection: string): Promise<Collection> {
    const resp = await fetch("/api/")
    return (await resp.json()) as Collection
}