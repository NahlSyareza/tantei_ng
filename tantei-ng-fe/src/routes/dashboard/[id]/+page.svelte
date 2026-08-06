<script lang="ts">
	import play_button from '$lib/assets/play_button.png';
	import edit_button from '$lib/assets/edit_button.png';

	interface NgSetItem {
		kanji: string;
		furigana: string;
		latin: string;
		english: string;
		indonesian: string;
	}

	interface NgSet {
		// _id: string;
		name: string;
		// createdAt: string;
		// updatedAt: string;
		items: NgSetItem[];
	}

	let {
		data
	}: {
		data: {
			o: NgSet;
			c: NgSet;
		};
	} = $props();
	// let dataCopy: NgSetItem[] = [...data.o];
	let selectedQuestion = $state<NgSetItem | null>(null);
	let availableAnswers: NgSetItem[] = $state([]);
	let hasStarted: boolean = $state(false);
	let remainingCounter: number = $state(0);

	function generateNewSequence() {
		availableAnswers = [];

		if (data.c.items.length < 1) {
			// console.log('Refill imminent');
			data.c.items = structuredClone(data.o.items);
			remainingCounter = data.o.items.length;
		}

		while (availableAnswers.length < 4) {
			if (availableAnswers.length == 0) {
				let randomNumber: number = Math.floor(Math.random() * data.c.items.length);
				// console.log(`Generated number: ${randomNumber}`);
				selectedQuestion = data.c.items[randomNumber];
				availableAnswers.push(data.c.items[randomNumber]);
				data.c.items.splice(randomNumber, 1);
				remainingCounter -= 1;
				continue;
			}

			// Now the problem is that it may be double for the options
			let randomNumber: number = Math.floor(Math.random() * data.o.items.length);
			while (data.o.items[randomNumber].kanji === availableAnswers[0].kanji) {
				randomNumber = Math.floor(Math.random() * data.o.items.length);
			}

			availableAnswers.push(data.o.items[randomNumber]);
		}

		// console.log(`Before: ${JSON.stringify(availableAnswers)}`);

		availableAnswers = shuffle(availableAnswers);

		// console.log(`After: ${JSON.stringify(availableAnswers)}`);

		// console.log(selectedItems);
		// console.log(data.o.items.length);
		// console.log(data.c.items.length);
	}

	// Too lazy to make, so I get AI to make... THIS FUNCTION ONLY
	function shuffle<T>(array: T[]): T[] {
		const copy = [...array]; // 1. Create a shallow copy
		for (let i = copy.length - 1; i > 0; i--) {
			const j = Math.floor(Math.random() * (i + 1));
			[copy[i], copy[j]] = [copy[j], copy[i]];
		}
		return copy; // 2. Return the new shuffled array
	}

	function handleStartButton() {
		generateNewSequence();
		remainingCounter = data.o.items.length;
		hasStarted = true;
	}

	function handleAnswerButtonClick(answer: string) {
		if (selectedQuestion!.english === answer) {
			generateNewSequence();
		}
		console.log(`${answer}`);
	}

	function getStyle(item: NgSetItem) {
		let someBool: boolean = false;

		if (someBool) {
			if (item.english === selectedQuestion!.english) {
				return 'text-2xl h-16 w-48 flex-1 rounded-xl bg-green-500';
			} else {
				return 'text-2xl h-16 w-48 flex-1 rounded-xl bg-red-500';
			}
		}

		return 'text-2xl h-32 w-150 rounded-xl bg-[#D5CEBE]';
	}

	let questionDisplayType = 'english';

	function getQuestionDisplayType(item: NgSetItem) {
		if (questionDisplayType == 'furigana') {
			return item.furigana;
		}

		if (questionDisplayType == 'indonesian') {
			return item.indonesian;
		}

		if (questionDisplayType == 'kanji') {
			return item.kanji;
		}

		return item.english;
	}

	let answerDisplayType = 'kanji';

	function getAnswerDisplayType(item: NgSetItem) {
		if (answerDisplayType == 'furigana') {
			return item.furigana;
		}

		if (answerDisplayType == 'indonesian') {
			return item.indonesian;
		}

		if (answerDisplayType == 'kanji') {
			return item.kanji;
		}

		return item.english;
	}
</script>

<div class="flex flex-1 flex-col">
	{#if !hasStarted}
		<div class="mb-10 flex flex-col rounded-b-2xl bg-[#E6E3D1]">
			<div class="flex flex-col items-center justify-center space-y-4 p-8">
				<p class="text-4xl font-bold">{data.o.name}</p>
				<p class="text-2xl font-bold">{data.o.items.length} words</p>
			</div>
			<div class="flex justify-center space-x-8 p-8">
				<button onclick={handleStartButton}>
					<img src={play_button} alt="play_button.svg" class="h-16" />
				</button>

				<img src={edit_button} alt="edit_button.svg" class="h-16" />
			</div>
		</div>

		<div class="mb-8 flex flex-wrap justify-center gap-5">
			{#each data.o.items as item, index (index)}
				<div class="flex h-50 w-50 items-center justify-center rounded-xl bg-[#D5CEBE]">
					<p class="text-2xl font-semibold">{item.kanji}</p>
				</div>
			{/each}
		</div>
	{:else}
		<p>{remainingCounter} / {data.o.items.length}</p>
		<div class="flex flex-4 p-4 font-semibold">
			<p class="flex flex-1 items-center justify-center rounded-xl bg-[#E6E3D1] text-6xl">
				<!-- {selectedQuestion!.kanji} -->
				{getQuestionDisplayType(selectedQuestion!)}
			</p>
		</div>
		<div
			class="flex flex-1 flex-wrap items-center justify-center gap-5 space-x-2 px-8 font-semibold"
		>
			{#each availableAnswers as item, index (index)}
				<button onclick={() => handleAnswerButtonClick(item.english)} class={getStyle(item)}
					>{getAnswerDisplayType(item)}</button
				>
			{/each}
		</div>
	{/if}
</div>
