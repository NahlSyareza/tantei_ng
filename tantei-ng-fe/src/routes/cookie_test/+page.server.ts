import type { PageServerLoadEvent } from './$types';

export function load({ cookies }: PageServerLoadEvent) {
	const visited = cookies.get('visited');

	console.log(visited);

	return {
		visited: visited === 'true'
	};
}

export const actions = {
	radahn: async () => {
		console.log('La peace');
	}
};
