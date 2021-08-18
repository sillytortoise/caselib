<template>
	<div id="bigpanel" class="panel panel-default">
		<div class="page-header">
			<myupload></myupload>
		</div>
		<div class="panel-default" style="width: 95%; height: 80%">
			<div class="panel-body">这是一个panel</div>
		</div>
	</div>
</template>


<script>
module.exports = {
	data() {
		return {};
	},
	components: {
		myupload: httpVueLoader("upload.vue"),
	},
};
</script>

<style scoped>
#bigpanel {
	height: auto;
}
</style>
