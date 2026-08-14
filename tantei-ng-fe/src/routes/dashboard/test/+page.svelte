<script lang="ts">
	let ahhCheeks: string[] = ['opt1', 'opt2', 'opt3', 'opt4'];
</script>

<div class="flex flex-1">
	<div class="flex-1"></div>
	<div class="flex flex-3 flex-col">
		<div class="m-4 flex flex-1">
			<div class="flex flex-1 items-center justify-center bg-[#D5CEBE]"><p>Ass Cheeks</p></div>
		</div>
		<div class="m-4 grid flex-1 grid-cols-2 gap-4">
			{#each ahhCheeks as cheek, index (index)}
				<div class="flex flex-1 items-center justify-center bg-[#D5CEBE]"><p>{cheek}</p></div>
			{/each}
		</div>
	</div>
	<div class="flex-1"></div>
</div>
