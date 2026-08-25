<script lang="ts">
	import type { PageProps } from './$types';
	import type { Studyword } from '../../../../aux/data_interfaces';
	import { fisherYatesShuffler, getRandomNum } from '../../../../aux/helper_functions';

	let { data }: PageProps = $props();
	let proxyItems = $state<Studyword[]>(structuredClone(data.res!.items));
	let selectedAnswer = $state<Studyword>();
	let optionItems = $state<Studyword[]>([]);

	function selectAnswer() {
		if (proxyItems.length < 1) {
			console.log("It's empty blyat");
			proxyItems = structuredClone(data.res!.items);
		}

		let diceNum: number = getRandomNum(0, proxyItems.length - 1);
		// console.log(`Spliced ${proxyItems[diceNum].kanji}`);
		selectedAnswer = proxyItems[diceNum];
		// console.log(`Remaining:`);
		proxyItems.splice(diceNum, 1);
		// console.log(`Remaining items: ${proxyItems.length}`);
		// proxyItems.forEach((e) => {
		// 	console.log(e.kanji);
		// });

		rearrangeStudyword();
	}

	function rearrangeStudyword() {
		if (selectedAnswer == null) {
			console.log('Select an answer first bruh');
			return;
		}

		let diceNum: number = getRandomNum(5, 10);

		let proxyArr: Studyword[] = [];
		proxyArr = [];

		while (proxyArr.length < diceNum) {
			let rNum: number = getRandomNum(0, data.res!.items.length - 1);
			if (!proxyArr.includes(data.res!.items[rNum])) {
				proxyArr.push(data.res!.items[rNum]);
			}
		}

		// console.log(proxyArr);
		let lastIndex: number = proxyArr.length - 1;

		if (!proxyArr.some((e) => e.kanji == selectedAnswer!.kanji)) {
			console.log('IT IS NOT INCLUDED WHAT THE FUCK????');
			proxyArr[lastIndex] = $state.snapshot(selectedAnswer);
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
			selectAnswer();
			rearrangeStudyword();
		} else {
			console.log('Nah, incorrect');
		}
	}
</script>

<div class="flex flex-1">
	{#if selectedAnswer == null}
		<div class="flex-1">
			<button onclick={() => selectAnswer()}><p>Start</p></button>
			<p>Try generating an answer first</p>
		</div>
	{:else}
		<div class="flex-1">
			<p>{proxyItems.length} / {data.res?.items.length}</p>
		</div>
		<div class="flex flex-3 flex-col items-center">
			<p>Selected answer: {selectedAnswer.english}</p>
			<div>
				{#if proxyItems.length > 0}
					{#each optionItems as item, index (index)}
						<button onclick={() => checkClickedAnswer(item.kanji)}>
							<p>{item.kanji}</p>
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
