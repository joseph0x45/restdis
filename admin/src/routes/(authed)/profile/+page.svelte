<script lang="ts">
	let username = 'username';
	let password = '*************';
	let saving = false;
</script>

<form action="" class="flex flex-col gap-2 w-full">
	<h1 class=" text-2xl">General</h1>
	<hr />
	<h1>Username</h1>
	<input
		type="text"
		bind:value={username}
		placeholder="Your username"
		class="p-2 rounded-md border"
	/>
	<h1>Password</h1>
	<input
		type="password"
		bind:value={password}
		placeholder="Your password"
		class="p-2 rounded-md border"
	/>
	<div>
		<button class="shadow-lg p-2 border w-32 rounded-md">
			{#if !saving}
				<h1>Save</h1>
			{:else}
				<h1>Saving...</h1>
			{/if}
		</button>
	</div>
</form>
