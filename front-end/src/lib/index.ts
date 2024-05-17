// place files you want to import through the `$lib` alias in this folder.
export const Sleep = (ms: number) => new Promise<void>(resolve => setTimeout(resolve, ms))


export type Form = {
    question: string
    answers: string[]
    tips?: string

    // TODO: implement to backend
    lastTry: Date
    nextTry: Date
    level: number
}

export type Collection = {
    name: string
    forms: Form[]
}

export type ServerError = {
    status: number
    code: number
    message: string
}