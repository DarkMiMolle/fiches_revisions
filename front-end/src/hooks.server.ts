

export function handleFetch({request, fetch}) {
    console.log(request.headers)
    return fetch(request)
}