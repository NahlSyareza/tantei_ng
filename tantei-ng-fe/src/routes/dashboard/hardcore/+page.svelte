<script lang="ts">
	import type { PageProps } from './$types';
	// import type { Studyset, Studyword } from '../../../aux/data_interfaces';
	import type { Studyword } from '../../../aux/data_interfaces';
	import { getRandomNum } from '../../../aux/helper_functions';

	// let { data }: { data: { res: Studyset } } = $props();
	let { data }: PageProps = $props();
	let items = $state<Studyword[]>([]);
	let selectedAnswer = $state<Studyword>();
	let testItems = $state<Studyword[]>([]);

	console.log('TestABC');
	console.log(data.res!.items);
	items = structuredClone(data.res!.items);

	function generateShit() {
		console.log('Jang Wonyoung');

		if (getRandomNum.length < 1) {
			console.log("It's empty blyat");
			items = structuredClone(data.res!.items);
		} else {
			let diceNum: number = getRandomNum(0, items.length - 1);
			console.log(`Spliced ${items[diceNum].kanji}`);
			selectedAnswer = items[diceNum];
			console.log(`Remaining:`);
			items.forEach((e) => {
				console.log(e.kanji);
			});
			// items.splice(diceNum, 1);
		}
	}

	function tryGenerateBullshit() {
		if (selectedAnswer == null) {
			console.log('Select an answer first bruh');
			return;
		}

		let diceNum: number = getRandomNum(2, 5);

		let proxyArr: Studyword[] = [];
		proxyArr = [];

		while (proxyArr.length < diceNum) {
			let rNum: number = getRandomNum(0, data.res!.items.length - 1);
			if (!proxyArr.includes(data.res!.items[rNum])) {
				proxyArr.push(data.res!.items[rNum]);
			} else {
				console.log('Pass');
			}
		}

		console.log(proxyArr);

		if (!proxyArr.includes(selectedAnswer)) {
			console.log('IT IS NOT INCLUDED WHAT THE FUCK????');
		}

		testItems = proxyArr;
		// Then randomize shit okay
	}

	function checkClickedAnswer(ans: string) {
		if (ans == selectedAnswer?.kanji) {
			console.log('Correct ass hole');
			testItems = [];
		} else {
			console.log('Nah, incorrect');
		}
	}

	function bullshitTest() {
		if (selectedAnswer == null) {
			return;
		} else {
			console.log('Frida gomam');
		}

		let pArr: Studyword[] = [];

		data.res!.items.forEach((e) => {
			pArr.push(e);
		});

		console.log(pArr);
		console.log($state.snapshot(selectedAnswer));

		let pv: Studyword = $state.snapshot(selectedAnswer);

		if (pArr.includes(pv)) {
			console.log('Correct');
		}
	}
</script>

{#if data == null}
	<p>Pizdec</p>
{/if}
<div>
	<p>What's</p>
	<p>After LIKE</p>
</div>
<div>
	{#if items.length > 0}
		{#each testItems as item, index (index)}
			<button onclick={() => checkClickedAnswer(item.kanji)}>
				<p>{item.kanji}</p>
			</button>
		{/each}
	{/if}
</div>
<button onclick={() => generateShit()}><p>Ass cheeks</p></button>
<button onclick={() => tryGenerateBullshit()}><p>Try generate bs</p></button>
<button onclick={() => bullshitTest()}><p>B Test</p></button>
<main></main>
