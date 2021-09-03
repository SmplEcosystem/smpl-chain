<template>
	<div v-if='initialized'>
		<SpWallet ref='wallet' v-on:dropdown-opened='$refs.menu.closeDropdown()' />
		<SpLayout>
			<template v-slot:sidebar>
				<Sidebar />
			</template>
			<template v-slot:content>
				<router-view />
			</template>
		</SpLayout>
	</div>
</template>

<style>
body {
	margin: 0;
}
</style>

<script>
import './scss/app.scss'
import '@starport/vue/lib/starport-vue.css'
import Sidebar from './components/Sidebar'

export default {
	components: {
		Sidebar
	},
	data() {
		return {
			initialized: false
		}
	},
	computed: {
		hasWallet() {
			return this.$store.hasModule(['common', 'wallet'])
		}
	},
	async created() {
		await this.$store.dispatch(
			'common/env/init',
			{
				apiNode: 'http://smpl-test-node1.s56.net:1317',
				// apiNode: 'http://localhost:1317',
				rpcNode: 'http://smpl-test-node1.s56.net:26657',
				// rpcNode: 'http://localhost:26657',
				wsNode: 'ws://smpl-test-node1.s56.net:26657/websocket',
				// wsNode: 'ws://localhost:26657/websocket',
				chainId: 'smpl-test',
				// chainId: 'linchain',
				addrPrefix: 'smpl',
				// addrPrefix: 'lin',
				sdkVersion: 'Stargate',
				getTXApi: 'http://smpl-test-node1.s56.net:26657/tx?hash=0x',
				// getTXApi: 'http://localhost:26657/tx?hash=0x',
				refresh: 10000
			}
		)
		this.initialized = true
	},
	errorCaptured(err) {
		console.log(err)
		return false
	}
}
</script>
