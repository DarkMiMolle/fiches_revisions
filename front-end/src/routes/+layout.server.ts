
export function load({cookies}) {
    return {
        isAuth: cookies.get("jwt") != undefined
    }
}