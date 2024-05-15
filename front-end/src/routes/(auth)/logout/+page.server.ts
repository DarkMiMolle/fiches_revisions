import { redirect } from '@sveltejs/kit';

export const actions = {
	default: ({ cookies }) => {
		cookies.delete('jwt', { path: '/' });
		throw redirect(303, '/');
	}
};
