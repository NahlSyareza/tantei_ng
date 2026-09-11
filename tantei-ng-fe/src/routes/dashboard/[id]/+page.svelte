<script lang="ts">
	import type { PageProps } from './$types';
	import play_button from '$lib/assets/play_button.png';
	import edit_button from '$lib/assets/edit_button.png';
	import { api } from '../../../aux/route';
	import type { Studyword } from '../../../aux/data_interfaces';
	import { resolve } from '$app/paths';
	import { fisherYatesShuffler, getRandomNum } from '../../../aux/helper_functions';

	let { data }: PageProps = $props();
	let selectedAnswer = $state<Studyword | null>(null);
	let availableAnswers: Studyword[] = $state([]);
	let hasStarted: boolean = $state(false);
	// let itemsProxy : Studyword[] = structuredClone(data.res!.items)
	let itemsProxy: Studyword[] = $derived(structuredClone(data.res!.items));

	async function generateNewSequence() {
		// availableAnswers = [];
		if (itemsProxy.length < 1) {
			// console.log('Refill imminent');
			itemsProxy = structuredClone(data.res!.items);

			let payload = {
				set_id: data.id,
				owner_id: '6a74bac987651dbbb3c2e71f'
			};

			const res = await api.post(`/tracker/remove_item`, payload);

			console.log(payload);
			console.log(res.data);
		}

		let randomNumber: number = getRandomNum(0, itemsProxy.length - 1);
		// console.log(`Generated number: ${randomNumber}`);
		selectedAnswer = itemsProxy[randomNumber];
		let selectedAnswerProxy: Studyword = itemsProxy[randomNumber];
		// availableAnswers.push(itemsProxy[randomNumber]);
		itemsProxy.splice(randomNumber, 1);

		console.log(selectedAnswerProxy);
		console.log(selectedAnswer);

		let availableAnswersProxy: Studyword[] = [];
		availableAnswersProxy = [];

		while (availableAnswersProxy.length < 4) {
			// Now the problem is that it may be double for the options
			let randomNumber: number = getRandomNum(0, data.res!.items.length - 1);
			let rolledItem: Studyword = data.res!.items[randomNumber];
			if (!availableAnswersProxy.some((e) => e.kanji == rolledItem.kanji)) {
				availableAnswersProxy.push(rolledItem);
			}
		}

		if (!availableAnswersProxy.some((e) => e.kanji == selectedAnswerProxy.kanji)) {
			let lastIndex: number = availableAnswersProxy.length - 1;
			availableAnswersProxy[lastIndex] = selectedAnswerProxy;
		}

		fisherYatesShuffler<Studyword>(availableAnswersProxy);

		availableAnswers = availableAnswersProxy;
	}

	function handleStartButton() {
		generateNewSequence();
		hasStarted = true;
	}

	function handleAnswerButtonClick(answer: string) {
		if (selectedAnswer!.english === answer) {
			generateNewSequence();
		}
		// console.log(`${answer}`);
	}

	function getStyle(item: Studyword) {
		let someBool: boolean = false;

		if (someBool) {
			if (item.english === selectedAnswer!.english) {
				return 'flex flex-1 text-2xl flex-1 rounded-xl bg-green-500';
			} else {
				return 'flex flex-1 text-2xl flex-1 rounded-xl bg-red-500';
			}
		}

		return 'flex flex-1 text-2xl rounded-xl bg-[#D5CEBE] justify-center items-center';
	}

	let questionDisplayType = 'english';

	function getQuestionDisplayType(item: Studyword) {
		if (questionDisplayType == 'furigana') {
			return item.furigana;
		}

		if (questionDisplayType == 'kanji') {
			return item.kanji;
		}

		return item.english;
	}

	let answerDisplayType = 'kanji';

	function getAnswerDisplayType(item: Studyword) {
		if (answerDisplayType == 'furigana') {
			return item.furigana;
		}

		if (answerDisplayType == 'kanji') {
			return item.kanji;
		}

		return item.english;
	}
</script>

<div class="flex flex-1">
	{#if data.res!.items.length > 0}{:else}{/if}

	{#if !hasStarted}
		<div class="flex flex-1 flex-col">
			<div class="mb-10 flex flex-col rounded-b-2xl bg-[#E6E3D1]">
				<div class="flex flex-col items-center justify-center space-y-4 p-8">
					<p class="text-4xl font-bold">{data.res!.name}</p>
					<p class="text-2xl font-bold">{data.res!.items.length} words</p>
				</div>
				<div class="flex justify-center space-x-8 p-8">
					<button onclick={() => handleStartButton()}>
						<img src={play_button} alt="play_button.svg" class="h-16" />
					</button>

					<img src={edit_button} alt="edit_button.svg" class="h-16" />

					<a
						href={resolve(`/dashboard/hardcore/${data.id}`)}
						class="flex aspect-square h-16 items-center justify-center rounded-4xl bg-red-600"
						><p>Hack</p></a
					>

					<!-- <a href={resolve(`/dashboard/hardcore/${data.id}`)}> Needle in Haystack </a> -->
				</div>
			</div>

			<div class="mb-8 flex flex-wrap justify-center gap-5">
				{#each data.res!.items as item, index (index)}
					<div class="flex h-50 w-50 items-center justify-center rounded-xl bg-[#D5CEBE]">
						<p class="text-2xl font-semibold">{item.kanji}</p>
					</div>
				{/each}
			</div>
		</div>
	{:else}
		<div class="flex-1">
			<p>{itemsProxy.length} / {data.res!.items.length}</p>
		</div>
		<div class="flex flex-3 flex-col">
			<div class="m-4 flex flex-1">
				<div class="flex flex-1 items-center justify-center rounded-xl bg-[#D5CEBE]">
					<p class="text-4xl font-semibold">{getQuestionDisplayType(selectedAnswer!)}</p>
				</div>
			</div>
			<div class="m-4 grid flex-1 grid-cols-2 gap-4">
				{#each availableAnswers as item, index (index)}
					<button onclick={() => handleAnswerButtonClick(item.english)} class={getStyle(item)}
						><p class="text-3xl font-semibold">{getAnswerDisplayType(item)}</p></button
					>
				{/each}
			</div>
		</div>
		<div class="flex-1"></div>
	{/if}
</div>
