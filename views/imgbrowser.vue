<template>
	<div style="margin-top: 15px">
		<div
			class="modal fade"
			id="myModalimgb"
			tabindex="-1"
			role="dialog"
			aria-labelledby="myModalimgbLabel"
			aria-hidden="true"
		>
			<div class="modal-dialog">
				<div class="modal-content">
					<div class="modal-header">
						<button
							type="button"
							class="close"
							data-dismiss="modal"
							aria-hidden="true"
						>
							&times;
						</button>
						<h4 class="modal-title" id="myModalimgbLabel">选择</h4>
					</div>
					<div class="modal-body">
						<selfunc
							ref="imgb_selfunc"
							@change="conschange"
							:strict="true"
						></selfunc>
						<button type="button" class="btn btn-primary" @click="sure">
							确定
						</button>
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
									style="max-width: 150px; height: 300px"
									:src="'/' + t.path"
									:preview-src-list="['/' + t.path]"
									@contextmenu.prevent="select"
								>
								</el-image>
							</div>
						</div>
					</div>
				</div>
				<!-- /.modal-content -->
			</div>
			<!-- /.modal -->
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
	emits: ["choosepic"],
	components: {
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
		select(event) {
			$(".glyphicon-ok-sign").css("visibility", "hidden");
			$($(event.target).parent().prev()).css("visibility", "");
			this.selectpic = $(event.target).attr("src");
		},
		sure() {
			this.$emit("choosepic", this.selectpic.substring(1));
		},
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
	},
};
</script>

<style scoped>
.modal {
	height: 600px;
}

.modal-dialog {
	height: 100%;
	width: 1000px;
}

span {
	position: relative;
	bottom: 150px;
	left: 50px;
	z-index: 9999;
}
</style>
