import { env } from '$env/dynamic/public';
import { fail, redirect } from '@sveltejs/kit';

const backendUrl = env.PUBLIC_BACKEND

export async function load({parent, cookies}) {
	await parent()
	if (cookies.get("jwt") != undefined) {
		throw redirect(303, "/logout")
	}
}

export const actions = {
	login: async ({ request, cookies, url }) => {
		const data = await request.formData()
		const pseudo = data.get("pseudo") as string, password = data.get("password") as string
		const body = {pseudo, password}
		// const errors = []

		// if (!/^([a-zA-Z]+[a-zA-Z0-9_-]*)$/.test(pseudo) || pseudo.length < 6) {
		// 	errors.push({pseudo: `${pseudo} n'est pas valide; un pseudo doit commencer par une lettre, doit avoir au moins 6 caractère parmis: des lettres, des chiffres, '-' ou '_'`})
		// }

		// if (password.length < 7) {
		// 	errors.push({password: "un mot de passe doit avoir au moins 7 caractères"})
		// }
		
		console.log(backendUrl)
		const resp = await fetch(`${backendUrl}/api/login`, {
			method: "POST",
			body: JSON.stringify(body)
		})
		const result = await resp.json()
		console.log(result)
		if (!resp.ok) {
			return fail(resp.status, result)
		}
		cookies.set('jwt', result.jwt, { path: '/' });
		return redirect(303, url.searchParams.get('redirectTo') ?? '/');
	},

	signup: async ({request, cookies, url}) => {
		const data = await request.formData()
		const pseudo = data.get("pseudo") as string, password = data.get("password") as string,
			email =	data.get("email") as string, notif = data.get("notif") as string == "true"

		const body = {pseudo, password, email, notif}
		const resp = await fetch(`${backendUrl}/api/signup`, {
			method: "POST",
			body: JSON.stringify(body)
		})
		const result = await resp.json()
		console.log(result)
		if (!resp.ok) {
			return fail(resp.status, result)
		}
		cookies.set("jwt", result.jwt, {path: "/"})
		console.log(url.origin)
		// return redirect(303, url.searchParams.get('redirectTo') ?? '/');
		return fail(501, {
			msg: "TODO"
		})

	}
};
