<script lang="ts">
	import { location } from '$lib/stores';
	const links = [
		{
			title: 'Profile',
			href: '/profile'
		},
		{
			title: 'Access tokens',
			href: '/version'
		},
		{
			title: 'Users',
			href: '/version'
		},
		{
			title: 'Configuration',
			href: '/config'
		}
	];
</script>

<header class="p-2 w-full border-b">
	<div class="flex justify-between items-center w-full">
		<h1 class="text-xl">Restdis v0.1</h1>
		<h1>Logged in as username</h1>
	</div>
</header>
<div class=" flex w-full gap-5 p-2 border-b">
	{#each links as link}
		<a class={` ${link.href === $location ? 'underline' : ''} hover:underline `} href={`/console/${link.href}`}
			>{link.title}</a
		>
	{/each}
</div>
<div class="m-auto flex justify-center w-[50%]">
	<slot />
</div>
