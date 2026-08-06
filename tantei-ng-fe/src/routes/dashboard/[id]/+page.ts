// import type { PageLoad } from './$types';
import type { PageLoad } from '../$types';
import axios from 'axios';

export const load: PageLoad = async ({ params }) => {
	const id = (params as any).id;

	try {
		const res = await axios.get(`http://localhost:28080/ng_set/${id}`);

		// console.log(res.data);

		// return {
		// 	res: res.data
		// };
		return { o: res.data, c: structuredClone(res.data) };
	} catch (error) {
		console.error(error);
		// return {
		// 	res: {}
		// };

		return { o: {}, c: {} };
	}
};
