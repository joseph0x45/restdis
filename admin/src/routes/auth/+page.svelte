<script lang="ts">
	let username = '';
	let password = '';
	let err_message = '';
	let err = false;
	let loading = false;
	async function login() {
		loading = true;
	}
</script>

<div class=" flex flex-col items-center justify-center h-screen gap-5">
	<h1>Log into your RestDis instance administration page</h1>
	<form class="flex flex-col gap-5" on:submit|preventDefault={login}>
		<input
			required
			bind:value={username}
			class="p-2 rounded-md w-72 border"
			type="text"
			placeholder="Username"
		/>
		<input
			required
			bind:value={password}
			class="p-2 rounded-md w-72 border"
			type="password"
			placeholder="Password"
		/>
		<button
			disabled={loading}
			class={`${loading ? 'bg-blue-500' : ' bg-blue-600'} p-2 rounded-md text-white`}
		>
			{#if loading}
				<h1>...</h1>
			{:else}
				<h1>Login</h1>
			{/if}
		</button>
		{#if err}
			<h1 class="text-center text-xs text-red-500">{err_message}</h1>
		{/if}
	</form>
</div>
