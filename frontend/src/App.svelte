<script lang="ts">
	import { onMount } from 'svelte';
	import ky from 'ky';
	import type { Log } from './log';

	let logs: Log[] = [];
	let message: string = '';

	const fetchLogs = async () => {
		try {
			const res = await ky.get('/api/log')
			if (res.status < 200 || res.status >= 300) {
				throw 'status: ' + res.statusText
			}
			logs = await res.json()
		} catch (error) {
			alert('error fetching: '+error)
		}
	}

	const postLogs = async () => {
		try {
			const res = await ky.post('/api/log', {body: message})
			if (res.status !== 201) {
				throw 'status: ' + res.statusText
			}
			fetchLogs()
			alert('created')
		} catch (error) {
			alert('error posting: '+error)
		}
	}

	onMount(fetchLogs);
</script>

<h1>
	Logs
</h1>

<button on:click={fetchLogs}>Fetch Logs</button>

<table>
	<tr>
		<th>Timestamp</th>
		<th>Message</th>
	</tr>
	{#each logs as {timestamp, message}}
		<tr>
			<td>{timestamp}</td>
			<td>{message}</td>
		</tr>
	{/each}
</table>

<h2>Create new log</h2>

<label for="loginput">Message: </label>
<input id="loginput" bind:value={message} />
<br />
<button on:click={postLogs}>Post Log</button>
