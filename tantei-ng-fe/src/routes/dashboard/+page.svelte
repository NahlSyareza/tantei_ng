<script lang="ts">
	import axios from 'axios';
	import type { PageProps } from './$types';
	import { resolve } from '$app/paths';

	interface NgSet {
		_id: string;
		name: string;
		createdAt: string;
		updatedAt: string;
	}

	let { data }: { data: { res: NgSet[] } } = $props();

	let sampleReading = [
		{
			name: 'R1'
		},
		{
			name: 'R2'
		},
		{
			name: 'R3'
		}
	];

	let samplePatternRecognition = [
		{
			name: 'PR1'
		},
		{
			name: 'PR2'
		},
		{
			name: 'PR3'
		},
		{
			name: 'PR4'
		}
	];

	// let data = $state(null);
	//
	// $effect(() => {
	// axios.get('http://localhost:28080/ng_sets').then((res) => (data = res.data));
	// console.log(data);
	//
	// return () => {};
	// });

	// export let data;
</script>

<div class="flex flex-1 flex-col p-6">
	<p class="text-2xl font-semibold">Reading</p>
	<div class="mt-8 flex flex-wrap gap-5">
		{#each data.res as { name, _id }}
			<a
				href={`/dashboard/${_id}`}
				class="flex h-50 w-50 items-center justify-center rounded-xl bg-[#D5CEBE]"
			>
				<p class="text-2xl font-semibold">{name}</p>
			</a>
		{/each}
		<!-- {#each sampleReading as { name } (name)}
			<div class="flex h-50 w-50 items-center justify-center rounded-xl bg-[#D5CEBE]">
				<p class="text-2xl font-semibold">{name}</p>
			</div>
		{/each} -->
		<a
			href={resolve('/dashboard/create_new_set')}
			class="flex h-50 w-50 items-center justify-center rounded-xl bg-[#E6E3D1]"
		>
			<p class="text-4xl font-semibold">+</p>
		</a>
	</div>
</div>
<div class="flex flex-1 flex-col p-6">
	<p class="text-2xl font-semibold">Pattern Recognition</p>
	<div class="mt-8 flex flex-wrap gap-5">
		{#each samplePatternRecognition as { name } (name)}
			<div class="flex h-50 w-50 items-center justify-center rounded-xl bg-[#D5CEBE]">
				<p class="text-2xl font-semibold">{name}</p>
			</div>
		{/each}
		<div class="flex h-50 w-50 items-center justify-center rounded-xl bg-[#E6E3D1]">
			<p class="text-4xl font-semibold">+</p>
		</div>
	</div>
</div>
