<template>
	<div id="bigpanel" class="panel panel-default">
		<div class="panel-body">
			<myupload></myupload>
		</div>
	</div>
</template>


<script>
module.exports = {
	data() {
		return {};
	},
	components: {
		myupload: httpVueLoader("/upload.vue"),
	},
};
</script>

<style scoped>
#bigpanel {
	height: auto;
	min-height: 80%;
}
</style>
