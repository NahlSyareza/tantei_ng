<script lang="ts">
	import axios from 'axios';

	interface NgSetItem {
		kanji: string;
		furigana: string;
		latin: string;
		english: string;
		indonesian: string;
	}

	let ngSetName = $state('');
	let ngSetItems = $state<NgSetItem[]>([
		{ kanji: '', furigana: '', latin: '', english: '', indonesian: '' }
	]);

	function handleAddNgItem() {
		ngSetItems.push({ kanji: '', furigana: '', latin: '', english: '', indonesian: '' });
	}

	function handleRemoveNgItem() {
		if (ngSetItems.length > 1) {
			ngSetItems.pop();
		}
	}

	function testPrintToConsole() {
		console.log(
			JSON.stringify({
				name: ngSetName,
				items: ngSetItems
			})
		);
	}

	async function trySaveDB() {
		try {
			const res = await axios.post(
				'http://localhost:28080/ng_set',
				{
					name: ngSetName,
					items: ngSetItems
				},
				{
					headers: {
						'Content-Type': 'application/json'
					}
				}
			);

			console.log(res.data);
		} catch (e) {
			console.error(e);
		}
	}
</script>

<div class="flex flex-1 flex-col">
	<div>
		<p>New Set Name</p>
		<input type="text" bind:value={ngSetName} placeholder="Insert name" />
	</div>
	<div class="space-y-2">
		<p>Items</p>
		{#each ngSetItems as item}
			<div>
				<input type="text" placeholder="Kanji" bind:value={item.kanji} />
				<input type="text" placeholder="Furigana" bind:value={item.furigana} />
				<input type="text" placeholder="Latin" bind:value={item.latin} />
				<input type="text" placeholder="English" bind:value={item.english} />
				<input type="text" placeholder="Indonesian" bind:value={item.indonesian} />
			</div>
		{/each}
		<button onclick={handleAddNgItem}>Add</button>
		<button onclick={handleRemoveNgItem}>Remove</button>
	</div>
	<div>
		<p>Created set will be:</p>
		<p>{ngSetName}</p>
		<!-- <p>{arrayOStrings}</p> -->
		<div class="space-y-2">
			{#each ngSetItems as item}
				<div>
					<p>{item.kanji}</p>
					<p>{item.furigana}</p>
					<p>{item.latin}</p>
					<p>{item.english}</p>
					<p>{item.indonesian}</p>
				</div>
			{/each}
		</div>

		<button onclick={testPrintToConsole}>Test Print</button>
		<button onclick={trySaveDB}>Try Save to DB</button>
	</div>
</div>
