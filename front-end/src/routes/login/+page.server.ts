import { env } from '$env/dynamic/public';
import { redirect } from '@sveltejs/kit';

export async function load({parent, cookies}) {
	await parent()
	if (cookies.get("jwt") != undefined) {
		throw redirect(303, "/logout")
	}
}

export const actions = {
	default: async ({ request, cookies, url }) => {
		const data = await request.formData()
		const backendUrl = env.PUBLIC_BACKEND
		console.log(backendUrl)
		const resp = await fetch(`${backendUrl}/api/login`, {
			method: "POST",
			body: JSON.stringify({
				email: "florent.carrez@yahoo.fr",
				password: 'passw0rd'
			})
		})
		const result = await resp.json()
		console.log(result)
		if (!resp.ok) {
			return
		}
		cookies.set('jwt', result["jwt"], { path: '/' });
		throw redirect(303, url.searchParams.get('redirectTo') ?? '/');
	}
};
