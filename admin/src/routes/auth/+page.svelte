<script lang="ts">
	import { goto } from '$app/navigation';

	let username = '';
	let password = '';
	let err_message =
		'Something unexpected happened. Open an issue on github if the problem persists';
	let err = false;
	let loading = false;
	let api_url = 'http://127.0.0.1:8080';
	async function login() {
		try {
			loading = true;
			const res = await fetch(`${api_url}/auth/login`, {
				method: 'POST',
				body: JSON.stringify({
					username,
					password
				})
			});
			loading = false;
			switch (res.status) {
				case 500:
					console.log('came here');
					throw new Error();
				case 400:
					err = true;
					err_message = 'Invalid credentials';
					break;
				case 200:
					const { token } = (await res.json()) as { token: string };
					localStorage.setItem('token', token);
					goto('/');
				default:
					break;
			}
		} catch (error) {
			console.log(error);
			err = true;
			err_message =
				'Something unexpected happened. Open an issue on github if the problem persists';
			loading = false;
		}
	}
</script>

<div class=" flex flex-col items-center justify-center h-screen gap-5">
	<div class=" flex flex-col items-center justify-center w-96 ">
		<h1 class="text-md pb-5">Log into RestDis</h1>
		<form class="flex flex-col gap-5 w-full" on:submit|preventDefault={login}>
			<input
				required
				bind:value={username}
				class="p-2 rounded-md w-72 border w-full"
				type="text"
				placeholder="Username"
			/>
			<input
				required
				bind:value={password}
				class="p-2 rounded-md w-72 border w-full"
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
				<h1 class="text-center text-sm text-red-500 w-full">{err_message}</h1>
			{/if}
		</form>
	</div>
</div>
