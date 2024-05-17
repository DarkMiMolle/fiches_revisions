import { type Form, type Collection } from "$lib";
import collections from "./test_collections.json"

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const __filename = "test"

export const CollectionsTest: Collection[] = (() => {
    const newColl = new Array<Collection>()
    for (const i in collections) {
        const newForms = new Array<Form>()
        const collection = collections[i]
        for (const j in collection.forms) {
            const {lastTry, nextTry} = collection.forms[j]
            newForms.push({...collection.forms[j], lastTry: new Date(lastTry), nextTry: new Date(nextTry)})
        }
        newColl.push({...collections[i], forms: newForms})
    }
    return newColl
})()
