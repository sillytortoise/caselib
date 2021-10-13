<template>
	<div id="mytask">
		<div class="panel panel-default" style="height: 6%">
			<div class="panel-body container" style="padding: 4px">
				<div class="row">
					<span class="col-md-1"></span>
					<button
						class="btn btn-primary col-md-2"
						data-toggle="modal"
						data-target="#mymodal"
					>
						新建
					</button>
					<span class="col-md-1"></span>
					<div class="input-group col-md-4" id="sort">
						<input
							type="text"
							placeholder="搜索"
							style="width: 80%"
							class="uk-button uk-button-small"
						/>
						<button
							type="button"
							class="input-group-addon glyphicon glyphicon-search"
							style="height: 35px; width: 35px"
						></button>
					</div>
					<span class="col-md-1"></span>
					<div class="btn-group btn-group-sm col-md-3" style="margin: 0 auto">
						<button
							type="button"
							class="btn btn-default"
							@click="create_time_sort"
						>
							创建时间
							<span class="glyphicon glyphicon-sort-by-attributes-alt"></span>
						</button>

						<button
							type="button"
							class="btn btn-primary"
							@click="modified_time_sort"
						>
							最后修改<span
								class="glyphicon glyphicon-sort-by-attributes-alt"
							></span>
						</button>
					</div>
				</div>
				<div class="modal fade" tabindex="-1" role="dialog" id="mymodal">
					<div class="modal-dialog" role="document">
						<div class="modal-content">
							<div class="modal-header">
								<button
									type="button"
									class="close"
									data-dismiss="modal"
									aria-label="Close"
								>
									<span aria-hidden="true">&times;</span>
								</button>
								<h4 class="modal-title">创建任务</h4>
							</div>
							<div class="modal-body">
								<el-input v-model="input" placeholder="任务名称"></el-input
								><br />
								<el-radio v-model="radio" label="1" checked>竞品分析</el-radio>
								<el-radio v-model="radio" label="2" checked
									>特色化案例库</el-radio
								>
							</div>
							<div class="modal-footer">
								<button
									type="button"
									class="btn btn-default"
									data-dismiss="modal"
								>
									关闭
								</button>
								<button
									type="button"
									class="btn btn-primary"
									@click="submit_task"
								>
									提交
								</button>
							</div>
						</div>
						<!-- /.modal-content -->
					</div>
					<!-- /.modal-dialog -->
				</div>
				<!-- /.modal -->
			</div>
		</div>
		<div class="panel panel-default" style="height: 100%; text-align: center">
			<div
				class="panel-body container"
				style="text-align: center; width: 95%; height: 600px"
			>
				<table
					v-loading="loading"
					class="table table-hover"
					style="text-align: left; width: 98%"
				>
					<thead>
						<tr>
							<th>名称</th>
							<th>创建时间</th>
							<th>上次修改</th>
							<th>所有者</th>
							<th>类型</th>
							<th>操作</th>
						</tr>
					</thead>
					<tbody>
						<tr
							v-for="item in tasks"
							@click="choose_task"
							:key="item.name"
							:name="item.name"
						>
							<td>{{ item.name }}</td>
							<td>{{ item.created }}</td>
							<td>{{ item.modified }}</td>
							<td>{{ item.owner }}</td>
							<td>
								<span v-if="item.type == 1">竞品分析</span>
								<span v-else>特色化案例库</span>
							</td>
							<td>
								<!--拥有者和当前用户不一致，不显示删除-->
								<button
									v-if="username === item.owner"
									class="btn btn-sm btn-warning"
									@click.stop="delete_task"
									:name="item.name"
								>
									<span
										class="glyphicon glyphicon-trash"
										:name="item.name"
									></span>
								</button>
							</td>
						</tr>
					</tbody>
				</table>
			</div>
			<el-pagination
				background
				@current-change="handleCurrentChange"
				:current-page.sync="currentPage"
				:page-size="10"
				layout="prev, pager, next, jumper"
				:total="totalnum"
			>
			</el-pagination>
		</div>
	</div>
</template>

<script>
module.exports = {
	data: function () {
		return {
			radio: "1",
			input: "",
			currentPage: 1,
			totalnum: 0,
			tasks: null,
			loading: true,
			sort: "modified",
			order: "desc",
		};
	},
	emits: ["changePage"],
	created: function () {
		this.get_total();
	},
	computed: {
		username: function () {
			var arr,
				reg = new RegExp("(^| )username=([^;]*)(;|$)");
			if ((arr = document.cookie.match(reg))) return decodeURI(arr[2]);
			else return null;
		},
	},
	methods: {
		submit_task: function () {
			if (this.input == "") {
				this.$message({
					showClose: true,
					message: "任务名不可为空",
					type: "error",
					duration: 2000,
				});
				return;
			}
			new Promise((resolve, reject) => {
				$.ajax({
					type: "POST",
					url: `/createtask?username=${this.username}&name=${this.input}&type=${this.radio}`,
					success: function (res) {
						$("#mymodal").modal("hide");
						resolve(res);
					},
				});
			}).then((data) => {
				if (data == 1) {
					this.$message({
						showClose: true,
						message: "任务创建成功",
						type: "success",
						duration: 2000,
					});
					this.totalnum += 1;
					this.get_tasks(this.currentPage, "modified", "desc");
				} else {
					this.$message({
						showClose: true,
						message: "失败，已有同名任务",
						type: "error",
						duration: 2000,
					});
				}
			});
		},
		get_total: function () {
			axios.get(`/gettotal?username=${this.username}`).then((res) => {
				this.totalnum = res.data;
				this.get_tasks(1, "modified", "desc");
			});
		},
		get_tasks: function (p, sort, order) {
			this.loading = true;
			axios
				.get(
					`/gettasks?username=${this.username}&page=${p}&sort=${sort}&order=${order}`
				)
				.then((res) => {
					this.tasks = res.data;
					this.loading = false;
				});
		},
		handleCurrentChange(val) {
			this.get_tasks(val);
		},
		create_time_sort(event) {
			if ($(event.target).prop("class") === "btn btn-primary") {
				if (
					$(event.target).children().first().prop("class") ==
					"glyphicon glyphicon-sort-by-attributes"
				) {
					$(event.target)
						.children()
						.first()
						.removeClass("glyphicon-sort-by-attributes");
					$(event.target)
						.children()
						.first()
						.addClass("glyphicon-sort-by-attributes-alt");
					this.get_tasks(this.currentPage, "created", "desc");
					this.sort = "created";
					this.order = "desc";
				} else {
					$(event.target)
						.children()
						.first()
						.removeClass("glyphicon-sort-by-attributes-alt");
					$(event.target)
						.children()
						.first()
						.addClass("glyphicon-sort-by-attributes");
					this.get_tasks(this.currentPage, "created", "asc");
					this.sort = "created";
					this.order = "asc";
				}
			}
			$(event.target).removeClass("btn-default");
			$(event.target).addClass("btn-primary");
			$(event.target).next().removeClass("btn-primary");
		},
		modified_time_sort: function (event) {
			if ($(event.target).prop("class") === "btn btn-primary") {
				if (
					$(event.target).children().first().prop("class") ==
					"glyphicon glyphicon-sort-by-attributes"
				) {
					$(event.target)
						.children()
						.first()
						.removeClass("glyphicon-sort-by-attributes");
					$(event.target)
						.children()
						.first()
						.addClass("glyphicon-sort-by-attributes-alt");
					this.get_tasks(this.currentPage, "modified", "desc");
					this.sort = "modified";
					this.order = "desc";
				} else {
					$(event.target)
						.children()
						.first()
						.removeClass("glyphicon-sort-by-attributes-alt");
					$(event.target)
						.children()
						.first()
						.addClass("glyphicon-sort-by-attributes");
					this.get_tasks(this.currentPage, "modified", "asc");
					this.sort = "modified";
					this.order = "asc";
				}
			}

			$(event.target).removeClass("btn-default");
			$(event.target).addClass("btn-primary");
			$(event.target).prev().removeClass("btn-primary");
		},
		choose_task: function (event) {
			let name = $(event.target).parents("tr").attr("name");
			let task_type = $($(event.target).parents("tr").children("td")[4]).text();
			let task_type_flag;
			let owner = $($(event.target).parents("tr").children("td")[3]).text();
			if (task_type == "竞品分析") task_type_flag = 1;
			else task_type_flag = 2;
			window.open(`/${owner}/${name}?type=${task_type_flag}`);
		},
		delete_task: function (event) {
			axios
				.get(`/delete_task?name=${$(event.target).attr("name")}`)
				.then((res) => {
					if (res.data == 1) {
						this.$message({
							showClose: true,
							message: "任务已删除",
							type: "success",
							duration: 2000,
						});
						this.totalnum--;
						this.get_tasks(this.currentPage, this.sort, this.order);
					} else {
						this.$message({
							showClose: true,
							message: "删除失败",
							type: "error",
							duration: 2000,
						});
					}
				});
		},
	},
};
</script>

<style scoped>
#mytask {
	width: 100%;
	height: 100%;
}

#sort {
	float: left;
}

input {
	border-color: rgb(174, 180, 180);
	border-style: solid;
	border-width: 1px;
	height: 34px;
}

input:focus {
	border-width: 2px;
	border-color: rgb(64, 158, 255);
}

#mytask {
	height: 100%;
}

tr,
td {
	height: 50px;
}

.el-radio {
	margin-top: 30px;
}
</style>
