import type { PageLoad } from './$types';
import { api } from '../../../aux/route';

export const load: PageLoad = async ({ params }) => {
	const id = params.id;

	try {
		const res = await api.get(`/studyset/${id}`);

		// console.log(res.data);

		// return {
		// 	res: res.data
		// };
		return { id: id, o: res.data, c: structuredClone(res.data) };
	} catch (error) {
		console.error(error);
		// return {
		// 	res: {}
		// };

		return { id: id, o: {}, c: {} };
	}
};
