<script lang="ts">
	let selectedSegment = $state('');
	let shouldTranslateDesc = $state(false);

	function segmentDescription(selectedSegment: string, shouldTranslate: boolean) {
		switch (selectedSegment) {
			case 'toukei':
				return shouldTranslate
					? 'You can view your study statistics here.'
					: 'まだありません、しばらくお待ちください。';
			case 'renshuu':
				return shouldTranslate
					? 'This is where you Practice. If you want to remember vocabularies, or learn how to read Kanji, you can do it here.'
					: '練習なら、ここです。言葉が覚えたいとか漢字の読み方を勉強したいとか、ここでできます。';
			case 'arimasen':
				return shouldTranslate ? 'Upcoming feature soon' : 'まだありません、しばらくお待ちください';
			default:
				return 'Click on a card segment.';
		}
	}
</script>

<div class="flex h-200 w-full p-12">
	<p class="text-5xl font-semibold">Get Started!</p>
</div>

<div class="relative flex h-207 w-full flex-col">
	<div class="flex h-[50%]">
		<button
			class="flex h-full w-[33%] items-end justify-center bg-[#D5CEBE]"
			onmouseenter={() => {
				selectedSegment = 'toukei';
			}}
		>
			<p class="mb-14 font-noto-sans-jp text-4xl font-semibold">統計</p>
		</button>
		<button
			onmouseenter={() => {
				selectedSegment = 'renshuu';
			}}
			class="flex h-full w-[33%] items-end justify-center bg-[#E6E3D1]"
		>
			<p class="mb-14 font-noto-sans-jp text-4xl font-semibold">練習</p>
		</button>
		<button
			onmouseenter={() => {
				selectedSegment = 'arimasen';
			}}
			class="flex h-full w-[33%] items-end justify-center bg-[#D5CEBE]"
		>
			<p class="mb-14 font-noto-sans-jp text-4xl font-semibold">ありません</p>
		</button>
	</div>
	<button
		class="flex h-[50%] items-center justify-center"
		onmouseenter={() => {
			shouldTranslateDesc = true;
		}}
		onmouseleave={() => {
			shouldTranslateDesc = false;
		}}
	>
		<p class="w-[95%] font-noto-sans text-4xl font-semibold">
			{segmentDescription(selectedSegment, shouldTranslateDesc)}
		</p>
	</button>
</div>

<div class="relative flex h-250 w-full bg-[#E6E3D1]">
	<div class="absolute mt-25 ml-24 rounded-2xl bg-[#D5CEBE] p-10 drop-shadow-md drop-shadow-black">
		<p class="text-3xl font-semibold underline">Personalize your learning!</p>
		<p class="mt-4 ml-4 text-3xl font-semibold">Create your own learning sets</p>
	</div>
	<div class="absolute mt-90 ml-250 rounded-2xl bg-[#D5CEBE] p-10 drop-shadow-md drop-shadow-black">
		<p class="text-3xl font-semibold underline">Track your growth!</p>
		<p class="mt-4 ml-4 text-3xl font-semibold">Look your statistics to gain insights</p>
	</div>
	<div class="absolute mt-175 ml-90 rounded-2xl bg-[#D5CEBE] p-10 drop-shadow-md drop-shadow-black">
		<p class="text-3xl font-semibold underline">Simplicity in itself!</p>
		<p class="mt-4 ml-4 text-3xl font-semibold">Tailor the platform to your liking</p>
	</div>
</div>
