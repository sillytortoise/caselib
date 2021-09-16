<template>
	<el-upload
		class="upload-demo"
		ref="upload"
		:action="`${task}/pic?page=${page}&p=${index}`"
		:file-list="fileList"
		list-type="picture"
		:on-change="handleChange"
		accept=".jpg,.jpeg,.png"
		:auto-upload="false"
		:data="selvalue"
	>
		<el-button size="small" type="primary">选择图片</el-button>
	</el-upload>
</template>

<script>
module.exports = {
	data() {
		return {
			fileList: [],
			task: window.location.pathname.split("/")[2],
		};
	},
	props: ["index", "selvalue", "page"],
	components: {
		selfunc: httpVueLoader("/selfunc.vue"),
	},
	methods: {
		submitUpload() {
			this.$refs.upload.submit();
		},
		handleRemove(file, fileList) {
			console.log(file, fileList);
		},
		handleChange(file, fileList) {
			this.fileList = fileList.slice(-1);
		},
	},
};
</script>

<style scoped>
.el-select,
.el-cascader {
	width: 200px;
	height: 40px;
}

input[type="file"] {
	display: none;
}
</style>

