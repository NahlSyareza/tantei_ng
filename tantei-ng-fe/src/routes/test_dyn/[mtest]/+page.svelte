<script>
	import { page } from '$app/state';

	let id = $derived(page.params.mtest);
</script>

<main>
	<p>{id}</p>
</main>
