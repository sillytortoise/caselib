<template>
	<div class="panel panel-default">
		<div class="panel-body">
			<mytree></mytree>
		</div>
	</div>
</template>

<script>
module.exports = {
	components: {
		mytree: httpVueLoader("tree.vue"),
	},
	data() {
		return {};
	},
};
</script>

<style scoped>
.panel {
	height: 100%;
}
</style>
