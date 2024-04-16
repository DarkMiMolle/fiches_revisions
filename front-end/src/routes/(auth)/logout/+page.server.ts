import { redirect } from '@sveltejs/kit';

export const actions = {
	default: ({ cookies }) => {
		console.log("here")
		cookies.delete('jwt', { path: '/' });
		throw redirect(303, '/');
	}
};
