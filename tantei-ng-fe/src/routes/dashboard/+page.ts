import type { PageLoad } from './$types';
import axios from 'axios';

export const load: PageLoad = async ({ params }) => {
	try {
		let data;
		const res = await axios.get('http://localhost:28080/ng_sets');

		console.log(res.data);
		return { res: res.data };
	} catch (error) {
		console.error(error);
		return { res: [] };
	}
};
