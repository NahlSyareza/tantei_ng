import type { PageLoad } from './$types';
import { api } from '../../../aux/route';

export const load: PageLoad = async () => {
	try {
		const studyset = await api.get('/studyset/6a86fb7f92b5df6c368ee1fb');
		const radicalList = await api.get('/radical_list/氵');

		return { studyset: studyset.data, radical_list: radicalList.data };
	} catch (error) {
		console.error(error);

		return { data: [] };
	}
};
