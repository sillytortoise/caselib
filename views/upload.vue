<template>
	<div style="display: block">
		<div>
			<div style="display: inline-block; width: 45%;height=200px">
				<el-upload
					class="upload-demo"
					drag
					action="/posts/"
					multiple
					:data="selectObj"
					:show-file-list="false"
					accept=".jpg,.png,.jpeg"
					:before-upload="changeFile"
				>
					<i class="el-icon-upload"></i>
					<div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
					<div class="el-upload__tip" slot="tip">只能上传jpg/jpeg/png文件</div>
				</el-upload>
			</div>
			<div
				style="display: inline-block; width: 45%;height=200px;vertical-align:top"
			>
				<sel @changevalue="setOptions"></sel>
			</div>
		</div>

		<div class="panel panel-default">
			<div class="panel-body">
				<ul id="upload-list" class="list-group">
					<li v-for="item in fileList" class="list-group-item">
						{{ item.name }}
						<label v-for="l in item.labels" class="label label-default">{{
							l
						}}</label>
						<span class="glyphicon glyphicon-remove" @click="removeItem"></span>
					</li>
				</ul>
			</div>
		</div>
	</div>
</template>
    
<script>
module.exports = {
	data() {
		return {
			fileList: [],
			selected: null,
			selectObj: {},
		};
	},
	components: {
		sel: httpVueLoader("/sel.vue"),
	},
	methods: {
		handleRemove(file, fileList) {},
		handlePreview(file) {
			console.log(file);
		},
		changeFile(file, fileList) {
			let labels = [];
			for (let i in this.selected) {
				labels.push(this.selected[i][1]);
			}
			this.fileList.push({ name: file.name, labels: labels });
		},
		submitUpload() {
			console.log(this.fileList);
			this.$refs.upload.submit();
		},
		setOptions(value) {
			this.selected = value;
			let param = [];
			for (let i in value) {
				param.push(value[i][1]);
			}
			this.selectObj = { select: param };
		},
		removeItem(event) {
			let index = $("#upload-list li").index($(event.target).parent());
			this.fileList.splice(index, 1);
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
</style>