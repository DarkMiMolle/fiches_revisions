import { env } from '$env/dynamic/public';
import { Sleep, type ServerError } from '$lib'
import { fail, redirect } from '@sveltejs/kit';

const backendUrl = env.PUBLIC_BACKEND

export async function load({parent, cookies}) {
	await parent()
	if (cookies.get("jwt") != undefined) {
		throw redirect(303, "/logout")
	}
}

async function testFail() {
	await Sleep(1000)
	return fail(401, {
		code: 1,
		status: 401,
		message: "pseudo non valide - un pseudo doit commencer par une lettre, doit avoir au moins 6 caractère parmis: des lettres, des chiffres, '-' ou '_'"
	} as ServerError)
}

export const actions = {
	login: async ({ request, cookies, url}) => {
		// return await testFail()
		const data = await request.formData()
		const pseudo = data.get("pseudo") as string, password = data.get("password") as string
		const body = {pseudo, password}
		
		const resp = await fetch(`${backendUrl}/api/login`, {
			method: "POST",
			body: JSON.stringify(body),
		})
		const result = await resp.json()

		console.log('login.server', {result})
		if (!resp.ok) {
			return fail(resp.status, result)
		}
		cookies.set('jwt', result.jwt, { path: '/' });
		return redirect(303, data.get("redirection") as string|undefined ?? '/');
	},

	signup: async ({request, cookies, url}) => {
		return testFail()
		const data = await request.formData()
		const pseudo = data.get("pseudo") as string, password = data.get("password") as string,
			email =	data.get("email") as string, notif = data.get("notif") as string == "true"

		const body = {pseudo, password, email, notif}
		const resp = await fetch(`${backendUrl}/api/signup`, {
			method: "POST",
			body: JSON.stringify(body)
		})
		const result = await resp.json()
		if (!resp.ok) {
			return fail(resp.status, result)
		}
		cookies.set("jwt", result.jwt, {path: "/"})
		console.log(url.origin)
		// return redirect(303, url.searchParams.get('redirectTo') ?? '/');
		return fail(501, )

	}
};
