<template>
	<div class="custom-tree-container">
		<div class="block">
			<div style="width=100%;">
				<el-button @click="appendNode">新建分类</el-button>
			</div>
			<br />
			<el-input placeholder="输入关键字进行过滤" v-model="filterText">
			</el-input
			><br />
			<el-tree
				:data="data"
				class="filter-tree"
				node-key="id"
				:filter-node-method="filterNode"
				:expand-on-click-node="false"
				ref="tree"
				props="defaultProps"
				@node-drag-start="handleDragStart"
				@node-drag-enter="handleDragEnter"
				@node-drag-leave="handleDragLeave"
				@node-drag-over="handleDragOver"
				@node-drag-end="handleDragEnd"
				@node-drop="handleDrop"
				draggable
				:allow-drop="allowDrop"
				:allow-drag="allowDrag"
			>
				<span class="custom-tree-node" slot-scope="{ node, data }">
					<div
						style="
							width: 40%;
							white-space: normal;
							word-break: break-all;
							word-wrap: break-word;
						"
					>
						{{ node.label }}
					</div>
					<span v-if="data.level === 1">{{ data.children.length }}个子项</span>
					<span>{{ data.problem }}</span>
					<span>
						<el-button
							v-if="data.level === 1"
							type="text"
							size="mini"
							@click="() => append(data)"
						>
							新增子项
						</el-button>
						<el-button
							type="text"
							size="mini"
							@click="() => remove(node, data)"
						>
							删除
						</el-button>
						<el-button type="text" size="mini" @click="edit(data)">
							编辑
						</el-button>
					</span>
				</span>
			</el-tree>
		</div>

		<!--一级模态框-->
		<div
			class="modal fade"
			id="modal-level1"
			tabindex="-1"
			role="dialog"
			aria-labelledby="myModalLabel"
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
						<h4 class="modal-title" id="myModalLabel1"></h4>
					</div>
					<div class="modal-body container">
						<div class="row">
							<span class="col-md-1">缩写</span>
							<input id="modalinput1" class="col-md-2" />
						</div>
					</div>
					<div class="modal-footer">
						<button type="button" class="btn btn-default" data-dismiss="modal">
							关闭
						</button>
						<button
							id="modalbtn1"
							@click="submit1"
							type="button"
							class="btn btn-primary"
						>
							提交更改
						</button>
					</div>
				</div>
				<!-- /.modal-content -->
			</div>
			<!-- /.modal -->
		</div>

		<!--二级模态框-->
		<div
			class="modal fade"
			id="modal-level2"
			tabindex="-1"
			role="dialog"
			aria-labelledby="myModalLabel"
			aria-hidden="true"
		>
			<div class="modal-dialog" style="width: 700px">
				<form target="myiframe" id="myform">
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
							<h4 class="modal-title" id="myModalLabel2"></h4>
						</div>
						<div class="modal-body container" style="height: 300px">
							<div class="row" style="width: 90%; vertical-align: middle">
								<span class="col-sm-2">评测项</span>
								<textarea class="col-md-2" name="item"></textarea>
								<span class="col-sm-2">指标</span>
								<textarea class="col-md-2" name="target"></textarea>
							</div>
							<br />
							<div class="row" style="width: 90%">
								<span class="col-sm-2">问题</span>
								<textarea class="col-md-2" name="problem"></textarea>
								<span class="col-sm-2">建议</span>
								<textarea class="col-md-2" name="advice"></textarea>
							</div>
							<div class="row" style="width: 90%; margin-top: 30px">
								<span class="col-sm-2">优先级</span>
								<select class="col-md-2" name="p">
									<option value="高">高</option>
									<option value="中">中</option>
									<option value="低">低</option>
								</select>
								<span class="col-sm-2">问题类型</span>
								<input class="col-md-2" name="t" />
							</div>
						</div>
						<div class="modal-footer">
							<button
								type="button"
								class="btn btn-default"
								data-dismiss="modal"
							>
								关闭
							</button>
							<button @click="submit2" type="button" class="btn btn-primary">
								提交更改
							</button>
						</div>
					</div>
				</form>
				<iframe id="myiframe" name="myiframe" style="display: none"></iframe>
				<!-- /.modal-content -->
			</div>
			<!-- /.modal -->
		</div>
	</div>
</template>

<script>
let id = 1;

module.exports = {
	data() {
		const filterText = "";
		const data = [];
		return {
			filterText: "",
			data: JSON.parse(JSON.stringify(data)),
		};
	},

	watch: {
		filterText(val) {
			this.$refs.tree.filter(val);
		},
	},
	created: function () {
		let tags = [];
		axios.get("/assess").then((res) => {
			for (let i in res.data) {
				let s = /[A-Z]+/.exec(res.data[i].id);
				if (tags.indexOf(s[0]) === -1) {
					tags.push(s[0]);
				}
			}

			let data = [];
			for (let i in tags) {
				let node = {};
				node.id = id++;
				node.label = tags[i];
				node.level = 1;
				node.problem = "";
				node.children = [];
				for (let t in res.data) {
					label = /[A-Z]+/.exec(res.data[t].id)[0];
					if (node.label === label) {
						child = {};
						child.id = id++;
						child.label = res.data[t].problem;
						child.level = 2;
						child.item = res.data[t].item;
						child.target = res.data[t].target;
						child.problem = res.data[t].id;
						child.advice = res.data[t].advice;
						child.p = res.data[t].p;
						child.t = res.data[t].t;
						node.children.push(child);
					}
				}
				data.push(node);
			}
			this.data = data;
		});
	},

	methods: {
		expandAll() {
			$(".el-tree-node").each(function () {
				$(this).addClass("is-expanded");
				$(this).attr("aria-expanded", true);
			});
		},
		hideAll() {
			$(".el-tree-node").attr["aria-expanded"] = false;
		},
		filterNode(value, data) {
			if (!value) return true;
			return data.label.indexOf(value) !== -1;
		},
		append(data) {
			const newChild = { id: id++, label: "testtest", children: [] };
			if (!data.children) {
				this.$set(data, "children", []);
			}
			data.children.push(newChild);
		},

		appendNode(data) {
			const newNode = {
				id: id++,
				label: "tsttest",
				level: 1,
				children: [],
			};
			this.data.push(newNode);
		},

		remove(node, data) {
			const parent = node.parent;
			const children = parent.data.children || parent.data;
			const index = children.findIndex((d) => d.id === data.id);
			children.splice(index, 1);
		},

		getCheckedNodes() {
			console.log(this.$refs.tree.getCheckedNodes());
		},

		resetChecked() {
			this.$refs.tree.setCheckedKeys([]);
		},
		handleDragStart(node, ev) {
			console.log("drag start", node);
		},
		handleDragEnter(draggingNode, dropNode, ev) {
			console.log("tree drag enter: ", dropNode.label);
		},
		handleDragLeave(draggingNode, dropNode, ev) {
			console.log("tree drag leave: ", dropNode.label);
		},
		handleDragOver(draggingNode, dropNode, ev) {
			console.log("tree drag over: ", dropNode.label);
		},
		handleDragEnd(draggingNode, dropNode, dropType, ev) {
			console.log("tree drag end: ", dropNode && dropNode.label, dropType);
		},
		handleDrop(draggingNode, dropNode, dropType, ev) {
			console.log("tree drop: ", dropNode.label, dropType);
		},
		allowDrop(draggingNode, dropNode, type) {
			if (dropNode.data.label === "二级 3-1") {
				return type !== "inner";
			} else {
				return true;
			}
		},
		allowDrag(draggingNode) {
			return draggingNode.data.label.indexOf("三级 3-2-2") === -1;
		},
		edit(data) {
			if (data.level === 1) {
				$("#modal-level1").modal();
				$("#myModalLabel1").text(data.label);
				$("#modalinput1").val(data.label);
			} else {
				$("#modal-level2").modal();
				$("#myModalLabel2").text(data.problem);
				$("textarea[name='item']").val(data.item);
				$("textarea[name='target']").val(data.target);
				$("textarea[name='problem']").val(data.label);
				$("textarea[name='advice']").val(data.advice);
				$("select[name='p']").val(data.p);
				$("input[name='t']").val(data.t);
			}
		},
		submit1() {
			for (let i in this.data) {
				let origin = $("#myModalLabel1").text();
				if (this.data[i].label === origin) {
					this.data[i].label = $("#modalinput1").val();
					for (let c in this.data[i].children) {
						this.data[i].children[c].problem =
							this.data[i].label +
							this.data[i].children[c].problem.replace(/[A-Z]+/i, "");
					}
					let data = this.data[i];
					data["origin"] = origin;
					$.ajax({
						url: "/modlevel1",
						data: JSON.stringify(data),
						type: "post",
						async: true,
						dataType: "json",
						headers: {
							contentType: "application/json; charset=utf-8",
						},
						success: function (result) {
							alert("成功");
						},
						error: function (result) {},
					});

					break;
				}
			}
			$("#modal-level1").modal("hide");
		},
		submit2() {
			let data = this.data;
			$("#myform").ajaxSubmit({
				type: "POST",
				url: `/modlevel2?name=${$("#myModalLabel2").text()}`,
				success: function () {
					let n = $("#myModalLabel2").text().search(/[0-9]/i);
					let head = $("#myModalLabel2").text().substring(0, n);
					for (let d in data) {
						if (data[d].label === head) {
							for (let c in data[d].children) {
								if (
									data[d].children[c].problem === $("#myModalLabel2").text()
								) {
									data[d].children[c].item = $("textarea[name='item']").val();
									data[d].children[c].target = $(
										"textarea[name='target']"
									).val();
									data[d].children[c].label = $(
										"textarea[name='problem']"
									).val();
									data[d].children[c].advice = $(
										"textarea[name='advice']"
									).val();
									data[d].children[c].p = $("select[name='p']").val();
									data[d].children[c].t = $("input[name='t']").val();
									break;
								}
							}
							break;
						}
					}
					$("#modal-level2").modal("hide");
					alert("成功");
				},
			});
		},
	},
};
</script>

<style>
.custom-tree-node {
	flex: 1;
	display: flex;
	align-items: center;
	justify-content: space-between;
	font-size: 14px;
	padding-right: 8px;
}

input {
	border-style: solid;
	border-width: 1px;
	border-color: gray;
}

input:focus {
	border-color: gray;
	border-width: 2px;
}

textarea {
	height: 60px;
	resize: none;
}

.modal span {
	margin-right: 0;
}
</style>