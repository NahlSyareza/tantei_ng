import type { PageLoad } from './$types';
import {api} from '../../aux/route'

export const load: PageLoad = async () => {
	try {
		// let data;
		const res = await api.get('/studysets/6a74bac987651dbbb3c2e71f');

		// console.log(res.data);
		return { res: res.data };
	} catch (error) {
		console.error(error);
		return { res: [] };
	}
};
