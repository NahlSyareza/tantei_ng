<script lang="ts">
	import type { PageProps } from './$types';
	import type { Studyword } from '../../../../aux/data_interfaces';
	import { fisherYatesShuffler, getRandomNum } from '../../../../aux/helper_functions';
	import { api } from '../../../../aux/route';

	let { data }: PageProps = $props();
	let itemsProxy = $derived<Studyword[]>(structuredClone(data.res!.items));
	let selectedAnswer = $state<Studyword>();
	let optionItems = $state<Studyword[]>([]);

	async function generateSequence() {
		if (data.res == null) {
			console.log('We can smash the club make the park go rough');
			return;
		}

		if (itemsProxy.length < 1) {
			console.log("It's empty blyat");
			itemsProxy = structuredClone(data.res!.items);

			let payload = {
				set_id: data.id,
				owner_id: '6a74bac987651dbbb3c2e71f'
			};

			const res = await api.post(`/tracker/remove_item`, payload);

			console.log(payload);
			console.log(res.data);
		}

		let diceNum: number = getRandomNum(0, itemsProxy.length - 1);
		// console.log(`Spliced ${itemsProxy[diceNum].kanji}`);
		selectedAnswer = itemsProxy[diceNum];
		let selectedAnswerProxy: Studyword = itemsProxy[diceNum];
		// console.log(`Remaining:`);
		itemsProxy.splice(diceNum, 1);
		// console.log(`Remaining items: ${itemsProxy.length}`);
		// itemsProxy.forEach((e) => {
		// 	console.log(e.kanji);
		// });

		let maxPoolLength: number = data.res?.items.length;
		diceNum = getRandomNum(Math.ceil(0.25 * maxPoolLength), maxPoolLength);

		let proxyArr: Studyword[] = [];
		proxyArr = [];

		while (proxyArr.length < diceNum) {
			let rNum: number = getRandomNum(0, data.res!.items.length - 1);
			let selectedRandom: Studyword = data.res!.items[rNum];
			if (!proxyArr.some((e) => e.kanji == selectedRandom.kanji)) {
				proxyArr.push(selectedRandom);
			} else {
				continue;
			}
		}

		// console.log(proxyArr);
		let lastIndex: number = proxyArr.length - 1;

		if (!proxyArr.some((e) => e.kanji == selectedAnswerProxy.kanji)) {
			console.log('IT IS NOT INCLUDED WHAT THE FUCK????');
			proxyArr[lastIndex] = selectedAnswerProxy;
		}

		// console.log(proxyArr);
		// Then randomize shit okay

		fisherYatesShuffler<Studyword>(proxyArr);
		// console.log(proxyArr);

		optionItems = proxyArr;
	}

	function checkClickedAnswer(ans: string) {
		if (ans == selectedAnswer?.kanji) {
			console.log('Correct ass hole');
			optionItems = [];
			generateSequence();
		} else {
			console.log('Nah, incorrect');
		}
	}
</script>

<div class="flex flex-1">
	{#if selectedAnswer == null}
		{#if itemsProxy.length > 0}
			<div class="flex-1">
				<button onclick={() => generateSequence()}><p>Start</p></button>
				<p>Try generating an answer first</p>
			</div>
		{/if}
	{:else}
		<div class="flex-1">
			<p>{itemsProxy.length} / {data.res?.items.length}</p>
		</div>
		<div class="flex flex-3 flex-col items-center">
			<p>Selected answer: {selectedAnswer.english}</p>
			<div>
				{#if itemsProxy.length > 0}
					{#each optionItems as item, index (index)}
						<button onclick={() => checkClickedAnswer(item.kanji)}>
							<p class="text-2xl">{item.kanji}</p>
						</button>
					{/each}
				{/if}
			</div>
		</div>
		<div class="flex-1"></div>
	{/if}

	<!-- <button onclick={() => generateShit()}><p>Ass cheeks</p></button>
<button onclick={() => tryGenerateBullshit()}><p>Try generate bs</p></button>
<button onclick={() => bullshitTest()}><p>B Test</p></button> -->
</div>
