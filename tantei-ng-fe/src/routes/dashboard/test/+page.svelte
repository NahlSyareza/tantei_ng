<script lang="ts">
	import { type Studyset, type Studyword, type RadicalList } from '../../../aux/data_interfaces';

	// let ahhCheeks: string[] = ['opt1', 'opt2', 'opt3', 'opt4'];

	let { data }: { data: { studyset: Studyset; radical_list: RadicalList } } = $props();

	let answer = $state<Studyword | undefined>(undefined);
	let options = $state<string[]>([]);

	function getRandomNum(min: number, max: number): number {
		return Math.floor(Math.random() * (max - min + 1)) + min;
	}

	function generateOptions(min: number, max: number): string[] {
		const arr: number[] = [];

		while (arr.length < 4) {
			const currentNum = getRandomNum(min, max);
			if (arr.includes(currentNum)) {
				continue;
			}
			arr.push(currentNum);
		}

		const outputArr: string[] = [];

		arr.forEach((item) => {
			// outputArr.add(data.radical_list.radical_variants[item]);
			outputArr.push(data.radical_list.radical_variants[item]);
		});

		return outputArr;
	}

	// I CAN do it myself, but it's too boring
	function shuffleArray<T>(arr: T[]): T[] {
		for (let i = arr.length - 1; i > 0; i--) {
			const j = Math.floor(Math.random() * (i + 1));
			[arr[i], arr[j]] = [arr[j], arr[i]];
		}
		return arr;
	}

	async function pickAnswer(arr: Studyword[]) {
		if (data.studyset.items.length < 1 || data.radical_list.radical_variants.length < 1) {
			console.error('Ey, one of them is empty blin!');
			return;
		}

		const max: number = arr.length - 1;
		const number: number = getRandomNum(0, max);

		answer = arr[number];

		const max2: number = data.radical_list.radical_variants.length - 1;
		options = generateOptions(0, max2);

		console.log('Initial Options');
		console.log($state.snapshot(options));
		console.log($state.snapshot(answer));

		if (!options.includes(answer.kanji)) {
			options[3] = answer.kanji;
		} else {
			console.log('Luck you, it has already been included!');
		}

		console.log('Options with the answer included');
		console.log($state.snapshot(options));

		shuffleArray(options);

		console.log('Options shuffled');
		console.log($state.snapshot(options));
	}

	function handleClickAnswer(ans: string) {
		if (ans == answer?.kanji) {
			console.log('Correct');
			pickAnswer(data.studyset.items)
		} else {
			console.log('Incorrect');
		}
	}
</script>

<div class="flex flex-1">
	<div class="flex-1"></div>
	<div class="flex flex-3 flex-col">
		<div class="m-4 flex flex-1">
			<button
				onclick={() => pickAnswer(data.studyset.items)}
				class="flex flex-1 items-center justify-center bg-[#D5CEBE]"><p>{answer?.english}</p></button
			>
		</div>
		<div class="m-4 grid flex-1 grid-cols-2 gap-4">
			{#each options as option, index (index)}
				<button
					onclick={() => handleClickAnswer(option)}
					class="flex flex-1 items-center justify-center bg-[#D5CEBE]"><p>{option}</p></button
				>
			{/each}
		</div>
	</div>
	<div class="flex-1"></div>
</div>
