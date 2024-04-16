import { redirect } from '@sveltejs/kit';

export function load({ cookies, url, locals }) {
	if (!cookies.get('jwt')) {
		throw redirect(303, `/login?redirectTo=${url.pathname}`);
	}
	locals.jwt = cookies.get("jwt")!
}