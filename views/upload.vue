<template>
	<el-upload
		class="upload-demo"
		drag
		action="/posts/"
		accept=".jpg,.jpeg,.png"
		multiple
		:on-remove="handleRemove"
		list-type="picture"
	>
		<i class="el-icon-upload"></i>
		<div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
		<div class="el-upload__tip" slot="tip">只能上传jpg/jpeg/png文件</div>
	</el-upload>
</template>
    
<script>
module.exports = {
	data() {
		return {
			fileList: [],
		};
	},
	methods: {
		submitUpload() {
			this.$refs.upload.submit();
		},
		handleRemove(file, fileList) {},
		handlePreview(file) {
			console.log(file);
		},
	},
};
</script>

<style scoped>
</style>