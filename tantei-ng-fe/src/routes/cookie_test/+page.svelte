<script lang="ts">
	import { enhance } from '$app/forms';

	let videoAssReferee = $state('');

	function testApp(val: string) {
		if (val.length < 1) {
			return '';
		}

		const jwtHeader = {
			typ: 'at+jwt',
			alg: 'RS256'
		};

		return btoa(JSON.stringify(jwtHeader)) + '.' + btoa(val);
	}
</script>

<main>
	<p>I'm cool with that</p>
	<form method="POST" action="?/radahn" use:enhance>
		<input type="text" bind:value={videoAssReferee} placeholder="Universal Collapse" />
		<p>{testApp(videoAssReferee)}</p>
		<button type="submit"><p>Hold on Tight</p></button>
	</form>
</main>
