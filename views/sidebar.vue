<template>
	<div style="width: 100px; display: inline-block; float: left">
		<el-button @click="drawer = true" type="default" style="margin-left: 16px">
			<span class="glyphicon glyphicon-th"></span>
		</el-button>

		<el-drawer
			title="我是标题"
			:visible.sync="drawer"
			direction="ltr"
			:before-close="handleClose"
			:append-to-body="false"
			:modal="false"
			:wrapper-Closable="false"
			size="100%"
		>
			<el-button type="primary" plain icon="el-icon-plus"></el-button>
			<ul class="list-group">
				<li class="list-group-item">Cras justo odio</li>
				<li class="list-group-item">Dapibus ac facilisis in</li>
				<li class="list-group-item">Morbi leo risus</li>
				<li class="list-group-item">Porta ac consectetur ac</li>
				<li class="list-group-item">Vestibulum at eros</li>
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
		};
	},
	methods: {
		handleClose(done) {
			done();
		},
	},
	created: function () {},
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
</style>