<script lang="ts">
	import './layout.css';
	import logo from '$lib/assets/meitanteisan.svg';
	import { resolve } from '$app/paths';
	// import favicon from '$lib/assets/favicon.svg';

	let { children } = $props();

	let meitanteisan = '名探偵さん';
</script>

<svelte:head><link rel="icon" href={logo} /></svelte:head>
<main class="relative flex min-h-screen w-full flex-col bg-[#FFFFEF] font-noto-sans">
	<div class="relative flex h-24 w-full bg-[#B94047]">
		<a href={resolve('/')} class="relative flex h-full w-[50%] justify-start">
			<img src={logo} alt="meitanteisan.svg" class="mx-2 my-auto h-18" />
			<p class="mx-2 my-auto font-noto-sans-jp text-4xl text-white">
				{meitanteisan}
			</p>
		</a>
		<div class="relative flex h-full w-[50%] justify-end">
			<a
				href={resolve('/dashboard')}
				class="mx-4 my-auto flex h-14 cursor-pointer items-center justify-center rounded-4xl bg-[#8a2e32] px-4"
			>
				<p class="font-noto-sans text-white">Dashboard</p>
			</a>
			<a
				href={resolve('/register')}
				class="mx-4 my-auto flex h-14 cursor-pointer items-center justify-center rounded-4xl bg-[#8a2e32] px-4"
			>
				<p class="font-noto-sans text-white">Register</p>
			</a>
			<a
				href={resolve('/sign-in')}
				class="mx-4 my-auto flex h-14 cursor-pointer items-center justify-center rounded-4xl bg-[#8a2e32] px-4"
				><p class="font-noto-sans text-white">Sign In</p>
			</a>
		</div>
	</div>

	<div class="relative flex flex-1 flex-col">
		{@render children()}
	</div>

	<div class="flex h-20 items-center justify-center bg-[#264348]">
		<p class="font-noto-sans text-sm text-white">© Eszaray 2026</p>
	</div>
</main>
