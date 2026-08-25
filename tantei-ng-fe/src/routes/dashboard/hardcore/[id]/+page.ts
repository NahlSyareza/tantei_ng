import type { PageLoad } from './$types';
import { devApi } from '../../../../aux/route';
import type { Studyset } from '../../../../aux/data_interfaces';

export const load: PageLoad = async ({ params }) => {
	const id = params.id;

	try {
		const res = await devApi.get<Studyset>(`/studyset/${id}`);

		// console.log(res.data.items);

		return { res: res.data };
	} catch (e) {
		console.error(e);

		return { res: null };
	}
};
