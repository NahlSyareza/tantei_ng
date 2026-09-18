<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import type { Studyword } from '../../../../aux/data_interfaces';
	import { api } from '../../../../aux/route';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();

	let items: Studyword[] = $state(data.res!.items);
	let name: string = $state(data.res!.name);

	function addItem() {
		items.push({ kanji: '', furigana: '', latin: '', english: '' });
	}

	function removeItem(index: number) {
		items.splice(index, 1);
	}

	async function updateStudyset() {
		try {
			const res = await api.put(
				`/studyset/${data.id}`,
				{
					name: name,
					items: items
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

<div class="flex flex-1 flex-col">
	<div class="mb-2">
		<p>Name</p>
		<input placeholder="name" bind:value={name} />
	</div>
	<div class="space-y-2">
		<p>Items</p>
		{#each items as item, index (index)}
			<!-- <p>{item.kanji}</p> -->
			<div class="flex space-x-2">
				<input placeholder="kanji" bind:value={item.kanji} />
				<input placeholder="furigana" bind:value={item.furigana} />
				<input placeholder="latin" bind:value={item.latin} />
				<input placeholder="english" bind:value={item.english} />
                <button onclick={() => removeItem(index)}>Remove</button>
			</div>
		{/each}
		<button onclick={addItem}>Add Item</button>
		<button onclick={updateStudyset}>Update</button>
	</div>
</div>
