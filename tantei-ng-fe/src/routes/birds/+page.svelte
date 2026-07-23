<script>
	import axios from 'axios';

	// let { data } = $props();
	// $inspect(data);

	let japanese = $state('');

	async function runAxios() {
		await axios
			.get('http://100.117.82.121:3001/word/get/68396284412fb927646fe56b')
			.then(function (response) {
				const res = response.data;
				console.log(res);
			});
	}
</script>

<main>
	<p>Refreshin'</p>
	<input bind:value={japanese} />
	<button onclick={runAxios}>Mack</button>
	<p>{japanese}</p>
</main>
