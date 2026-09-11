<script lang="ts">
	import { api } from '../../../aux/route';
	import {type Studyword} from '../../../aux/data_interfaces'

	let studysetName = $state('');
	let studysetItems = $state<Studyword[]>([
		{ kanji: '', furigana: '', latin: '', english: '', radical: '' }
	]);

	function handleAddNgItem() {
		studysetItems.push({ kanji: '', furigana: '', latin: '', english: '', radical: '' });
	}

	function handleRemoveNgItem() {
		if (studysetItems.length > 1) {
			studysetItems.pop();
		}
	}

	function testPrintToConsole() {
		console.log(
			JSON.stringify({
				name: studysetName,
				items: studysetItems
			})
		);
	}

	async function trySaveDB() {
		try {
			const res = await api.post(
				'/studyset',
				{
					name: studysetName,
					items: studysetItems,
					// Stil hardcoded, not good
					owner: "6a74bac987651dbbb3c2e71f"
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
		<input type="text" bind:value={studysetName} placeholder="Insert name" />
	</div>
	<div class="space-y-2">
		<p>Items</p>
		{#each studysetItems as item, index (index)}
			<div>
				<input type="text" placeholder="Kanji" bind:value={item.kanji} />
				<input type="text" placeholder="Furigana" bind:value={item.furigana} />
				<input type="text" placeholder="Latin" bind:value={item.latin} />
				<input type="text" placeholder="English" bind:value={item.english} />
				<input type="text" placeholder="Indonesian" bind:value={item.radical} />
			</div>
		{/each}
		<button onclick={handleAddNgItem}>Add</button>
		<button onclick={handleRemoveNgItem}>Remove</button>
	</div>
	<div>
		<p>Created set will be:</p>
		<p>{studysetName}</p>
		<!-- <p>{arrayOStrings}</p> -->
		<div class="space-y-2">
			{#each studysetItems as item, index (index)}
				<div>
					<p>{item.kanji}</p>
					<p>{item.furigana}</p>
					<p>{item.latin}</p>
					<p>{item.english}</p>
					<p>{item.radical}</p>
				</div>
			{/each}
		</div>

		<button onclick={testPrintToConsole}>Test Print</button>
		<button onclick={trySaveDB}>Try Save to DB</button>
	</div>
</div>
