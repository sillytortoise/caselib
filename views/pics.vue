<template>
	<div id="bigpanel" class="panel panel-default">
		<div class="panel-body">
			<myupload></myupload>
		</div>
		<selfunc ref="imgb_selfunc" @change="conschange"></selfunc>
		<div class="block">
			<!--分页-->
			<el-pagination
				@current-change="handleCurrentChange"
				:current-page.sync="currentPage"
				:page-size="20"
				layout="prev, pager, next"
				:total="total"
			>
			</el-pagination>
		</div>
		<div style="width: 100%">
			<div
				v-for="(t, i) in imgs"
				:key="i"
				style="display: inline-block; height: auto; width: auto"
			>
				<span
					class="glyphicon glyphicon-ok-sign"
					style="visibility: hidden"
				></span>
				<el-image
					style="max-width: 140px; height: 300px"
					:src="'/' + t.path"
					:preview-src-list="['/' + t.path]"
				>
				</el-image>
			</div>
		</div>
	</div>
</template>


<script>
module.exports = {
	data() {
		return {
			total: 0,
			imgs: [],
			currentPage: null,
			selectpic: null,
		};
	},
	components: {
		myupload: httpVueLoader("/upload.vue"),
		selfunc: httpVueLoader("/selfunc.vue"),
	},
	created: function () {
		/* {
			total:...,
			imgs:[     每一页显示的图片
				{
					path:...,
					bank:...,
					ver:...
				},
				{
					path:...,
					bank:...,
					ver:...
				},
				...
			]
		}
		*/
		axios.get(`/allimages?p=1`).then((res) => {
			this.total = res.data.total;
			this.imgs = res.data.imgs;
		});
	},
	methods: {
		handleCurrentChange(val) {
			axios.get(`/allimages?p=${this.currentPage}`).then((res) => {
				this.imgs = res.data.imgs;
			});
		},
		conschange(val) {
			//axios 新的条件请求图片
		},
	},
};
</script>

<style scoped>
#bigpanel {
	height: auto;
	min-height: 80%;
}

.el-image {
	margin: 20px;
}

.el-pagination {
	text-align: center;
}
</style>
