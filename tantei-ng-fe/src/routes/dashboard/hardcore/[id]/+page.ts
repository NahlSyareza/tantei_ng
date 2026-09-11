import type { PageLoad } from './$types';
import { api } from '../../../../aux/route';
import type { Studyset } from '../../../../aux/data_interfaces';

export const load: PageLoad = async ({ params }) => {
	const id = params.id;

	try {
		const res = await api.get<Studyset>(`/studyset/${id}`);

		console.log(res.data.items);

		return {id: id, res: res.data };
	} catch (e) {
		console.error(e);

		return { res: null };
	}
};
