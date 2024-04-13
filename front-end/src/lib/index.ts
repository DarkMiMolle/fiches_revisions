// place files you want to import through the `$lib` alias in this folder.
export const Sleep = async (ms: number) => new Promise((resolver) => setTimeout(resolver, ms))