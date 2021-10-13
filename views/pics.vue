<template>
	<div id="bigpanel" class="panel panel-default">
		<div class="panel-body">
			<el-switch
				v-model="switchvalue"
				active-text="手动输入"
				inactive-text="选择"
			>
			</el-switch>
			<myupload :switch="switchvalue" @afterupload="afterupload"></myupload>
		</div>
		<el-divider></el-divider>
		<selfunc ref="imgb_selfunc" @change="conschange" :strict="true"></selfunc>

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
			switchvalue: true,
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
		axios.get(`/imgs?p=1`).then((res) => {
			this.total = res.data.total;
			this.imgs = res.data.imgs;
		});
	},
	methods: {
		/*改变页数*/
		handleCurrentChange(val) {
			let func = this.$refs.imgb_selfunc.value1;
			let value = this.$refs.imgb_selfunc.value2;
			let app;
			let ver;
			if (value.length > 0) {
				app = value[0];
			} else {
				app = "";
			}
			if (value.length === 2) {
				ver = value[1];
			} else {
				ver = "";
			}
			axios
				.get(`/imgs?p=${this.currentPage}&func=${func}&app=${app}&ver=${ver}`)
				.then((res) => {
					this.imgs = res.data.imgs;
				});
		},
		/*改变筛选条件*/
		conschange(val) {
			//axios 新的条件请求图片
			let func = this.$refs.imgb_selfunc.value1;
			let value = this.$refs.imgb_selfunc.value2;
			let app;
			let ver;
			if (value.length > 0) {
				app = value[0];
			} else {
				app = "";
			}
			if (value.length === 2) {
				ver = value[1];
			} else {
				ver = "";
			}
			axios.get(`/imgs?p=1&func=${func}&app=${app}&ver=${ver}`).then((res) => {
				this.imgs = res.data.imgs;
				this.total = res.data.total;
				this.currentPage = 1;
			});
		},
		afterupload() {
			this.$refs.imgb_selfunc.init();
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

.el-switch {
	position: absolute;
	right: 20px;
}
</style>
