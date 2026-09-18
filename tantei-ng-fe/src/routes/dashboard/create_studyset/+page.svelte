<script lang="ts">
	import { api } from '../../../aux/route';
	import { type Studyword } from '../../../aux/data_interfaces';
	import { invalidateAll } from '$app/navigation';

	let studysetName = $state('');
	let studysetItems = $state<Studyword[]>([{ kanji: '', furigana: '', latin: '', english: '' }]);

	function handleAddNgItem() {
		studysetItems.push({ kanji: '', furigana: '', latin: '', english: '' });
	}

	// function handleRemoveNgItem() {
	// 	if (studysetItems.length > 1) {
	// 		studysetItems.pop();
	// 	}
	// }

	function handleRemoveNgItem(index: number) {
		studysetItems.splice(index, 1);
	}

	async function trySaveDB() {
		try {
			const res = await api.post(
				'/studyset',
				{
					name: studysetName,
					items: studysetItems,
					// Stil hardcoded, not good
					owner: '6a74bac987651dbbb3c2e71f'
				},
				{
					headers: {
						'Content-Type': 'application/json'
					}
				}
			);

			console.log(res.data);

			invalidateAll();
		} catch (e) {
			console.error(e);
		}
	}
</script>

<div class="flex flex-1 flex-col p-3">
	<div class="mb-2">
		<p>Set Name</p>
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
				{#if studysetItems.length > 1}
					<button onclick={() => handleRemoveNgItem(index)}>Remove Item</button>
				{/if}
			</div>
		{/each}
		<button onclick={handleAddNgItem}>Add Item</button>
	</div>
	<div>
		<button onclick={trySaveDB}>Create</button>
	</div>
</div>
