<template>
	<div style="display: block">
		<div>
			<div
				style="display: inline-block; width: 45%;height=200px;vertical-align:top"
			>
				<selfunc ref="upload_sel" :strict="false"></selfunc>
				<div style="text-align: left">
					<p>没有？在下面输入</p>
					<el-input v-model="input_app" placeholder="APP名称"></el-input>
					<el-input v-model="input_ver" placeholder="版本"></el-input>
					<el-input v-model="input_func" placeholder="功能模块"></el-input>
				</div>
			</div>
			<div style="display: inline-block; width: 45%;height=200px">
				<el-upload
					ref="upload"
					class="upload-demo"
					action="/posts/"
					:data="selectObj"
					:auto-upload="false"
					:file-list="fileList"
					:before-upload="beforeUpload"
					:on-success="onSuccess"
					accept=".jpg,.png,.jpeg"
				>
					<el-button slot="trigger" size="small" type="primary"
						>选取文件</el-button
					>
					<el-button
						style="margin-left: 10px"
						size="small"
						type="success"
						@click="submitUpload"
						>上传到服务器</el-button
					>
					<div class="el-upload__tip" slot="tip">只能上传jpg/jpeg/png文件</div>
				</el-upload>
			</div>
		</div>
	</div>
</template>
    
<script>
module.exports = {
	data() {
		return {
			fileList: [],
			input_app: "",
			input_ver: "",
			input_func: "",
		};
	},
	props: ["switch"],
	emits: ["afterupload"],
	components: {
		selfunc: httpVueLoader("/selfunc.vue"),
	},
	computed: {
		selectObj: function () {
			if (this.switch) {
				//手动输入
				return {
					func: this.input_func,
					app: this.input_app,
					ver: this.input_ver,
				};
			} else {
				return {
					func: this.$refs.upload_sel.value1,
					app: this.$refs.upload_sel.value2[0],
					ver: this.$refs.upload_sel.value2[1],
				};
			}
		},
	},
	methods: {
		beforeUpload(file, fileList) {
			if (this.switch) {
				//手动输入
				if (
					this.input_app == "" ||
					this.input_ver == "" ||
					this.input_func == ""
				) {
					this.$message.error("失败：请输入信息");
				}
			} else {
				if (
					this.$refs.upload_sel.value1 == "" ||
					this.$refs.upload_sel.value2.length == 0
				) {
					this.$message.error("失败：请选择完整的信息");
				}
			}
		},
		submitUpload() {
			this.$refs.upload.submit();
		},
		// removeItem(event) {
		// 	let index = $("#upload-list li").index($(event.target).parent());
		// 	this.fileList.splice(index, 1);
		// },
		onSuccess(response, file, fileList) {
			this.$refs.upload_sel.init();
			this.$emit("afterupload");
		},
	},
};
</script>

<style scoped>
label {
	margin-right: 10px;
}
.glyphicon-remove {
	float: right;
}

.el-input {
	margin-top: 10px;
	width: 221.4px;
}
</style>