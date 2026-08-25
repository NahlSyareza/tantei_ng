import type { PageLoad } from './$types';
import { devApi } from '../../../aux/route';
import type { Studyset } from '../../../aux/data_interfaces';

export const load: PageLoad = async () => {
	try {
		const res = await devApi.get<Studyset>('/studyset/6a887e885b5a683961b05698');

		console.log(res.data.items);

		return { res: res.data };
	} catch (e) {
		console.error(e);

		return { res: null };
	}
};
