<script lang="ts">
	interface NgSetItem {
		kanji: string;
		furigana: string;
		latin: string;
		english: string;
		indonesian: string;
	}

	interface NgSet {
		_id: string;
		name: string;
		createdAt: string;
		updatedAt: string;
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

		if (data.c.items.length <= 1) {
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
		// console.log(`${answer}`);
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

		return 'text-2xl h-16 w-48 flex-1 rounded-xl bg-[#D5CEBE]';
	}
</script>

<div class="flex flex-1 flex-col">
	<p>{remainingCounter} / {data.o.items.length}</p>
	{#if !hasStarted}
		<div class="flex flex-1 items-center justify-center">
			<button class="rounded-xl bg-[#D5CEBE] p-8 text-2xl text-black" onclick={handleStartButton}
				>Start</button
			>
		</div>
	{:else}
		<div class="flex flex-1 p-8">
			<p class="flex flex-1 items-center justify-center rounded-xl bg-[#E6E3D1] text-4xl">
				{selectedQuestion!.kanji}
			</p>
		</div>
		<div class="flex flex-1 items-center space-x-2 p-8">
			{#each availableAnswers as items}
				<button onclick={() => handleAnswerButtonClick(items.english)} class={getStyle(items)}
					>{items.english}</button
				>
			{/each}
		</div>
	{/if}
</div>
