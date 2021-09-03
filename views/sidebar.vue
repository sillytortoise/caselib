<template>
	<div style="width: 100px; display: inline-block; float: left">
		<el-button @click="drawer = true" type="default" style="margin-left: 16px">
			<span class="glyphicon glyphicon-th"></span>
		</el-button>

		<el-drawer
			title="我是标题"
			:visible.sync="drawer"
			direction="ltr"
			@open="handleOpen"
			:before-close="handleClose"
			:append-to-body="false"
			:modal="false"
			:wrapper-Closable="false"
			size="100%"
		>
			<el-button
				type="primary"
				plain
				icon="el-icon-plus"
				@click="addtoend"
			></el-button>
			<ul class="list-group" style="text-align: left">
				<li
					class="list-group-item"
					v-for="(item, i) in data"
					:class="{ selected: i == currentPage - 1 }"
					@click="choosePage"
				>
					{{ item.title !== "" ? item.title : `第${i + 1}页 未命名` }}
				</li>
			</ul>
		</el-drawer>
	</div>
</template>

<script>
module.exports = {
	data() {
		return {
			drawer: false,
			direction: "rtl",
			currentPage: 1,
			data: null,
			task: window.location.pathname.split("/")[2],
		};
	},
	emits: ["changepage"],
	methods: {
		handleClose(done) {
			let height = $(".panel-default").css("height").slice(0, -2);
			let width = $(".panel-default").css("width").slice(0, -2);
			width *= 1.25;
			height *= 1.25;
			$(".panel-default").css("height", height.toFixed(2) + "px");
			$(".panel-default").css("width", width.toFixed(2) + "px");
			done();
		},
		handleOpen() {
			let height = $(".panel-default").css("height").slice(0, -2);
			let width = $(".panel-default").css("width").slice(0, -2);
			width *= 0.8;
			height *= 0.8;
			$(".panel-default").css("height", height.toFixed(2) + "px");
			$(".panel-default").css("width", width.toFixed(2) + "px");
		},
		addtoend() {
			axios.get(`${this.task}/addtoend`).then((res) => {
				this.get_pages();
				this.$message({
					type: "success",
					message: "新增成功!",
					duration: 2000,
				});
			});
		},
		addtonext() {
			axios.get(`${this.task}/addtonext?p=${this.currentPage}`).then((res) => {
				this.get_pages();
				this.$message({
					type: "success",
					message: "新增成功!",
					duration: 2000,
				});
			});
		},
		get_pages() {
			axios.get(`${this.task}/get_pages`).then((res) => {
				this.data = res.data;
			});
		},
		choosePage(event) {
			$(".selected").removeClass("selected");
			$(event.target).addClass("selected");
			let page = $(event.target).index();
			this.currentPage = page + 1;
			this.$emit("changepage", this.currentPage);
		},
		deletePage(num) {
			axios.get(`${this.task}/deletepage?p=num`).then((res) => {});
		},
		changeItemColor(value) {
			$(".selected").removeClass("selected");
			$($("li")[value - 1]).addClass("selected");
			this.currentPage = value;
		},
	},
	created: function () {
		this.get_pages();
	},
};
</script>

<style scoped>
.list-group {
	text-align: left;
}

.list-group-item:hover {
	background: #409eff;
	color: white;
}

.selected {
	background: rgb(204, 102, 0);
	color: white;
}
</style>