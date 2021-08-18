<template>
	<div id="mytask">
		<div class="panel panel-default" style="height: 6%">
			<div class="panel-body container" style="padding: 4px">
				<div class="row">
					<span class="col-md-1"></span>
					<button class="btn btn-primary col-md-2">新建</button>
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
						<button type="button" class="btn btn-default">按钮 1</button>
						<button type="button" class="btn btn-default">按钮 2</button>
						<button type="button" class="btn btn-default">按钮 3</button>
					</div>
				</div>
			</div>
		</div>
		<div class="panel panel-default" style="height: 100%">
			<div class="panel-body container">
				<table class="table table-striped">
					<thead>
						<tr>
							<th>名称</th>
							<th>创建时间</th>
							<th>上次修改</th>
							<th>所有者</th>
							<th>权限</th>
						</tr>
					</thead>
					<tbody>
						<tr>
							<td>Tanmay</td>
							<td>Bangalore</td>
							<td>560001</td>
							<td></td>
							<td></td>
						</tr>
						<tr>
							<td>Sachin</td>
							<td>Mumbai</td>
							<td>400003</td>
							<td></td>
							<td></td>
						</tr>
						<tr>
							<td>Uma</td>
							<td>Pune</td>
							<td>411027</td>
							<td></td>
							<td></td>
						</tr>
						<tr>
							<td>Tanmay</td>
							<td>Bangalore</td>
							<td>560001</td>
							<td></td>
							<td></td>
						</tr>
						<tr>
							<td>Sachin</td>
							<td>Mumbai</td>
							<td>400003</td>
							<td></td>
							<td></td>
						</tr>
						<tr>
							<td>Uma</td>
							<td>Pune</td>
							<td>411027</td>
							<td></td>
							<td></td>
						</tr>
						<tr>
							<td>Tanmay</td>
							<td>Bangalore</td>
							<td>560001</td>
							<td></td>
							<td></td>
						</tr>
						<tr>
							<td>Sachin</td>
							<td>Mumbai</td>
							<td>400003</td>
							<td></td>
							<td></td>
						</tr>
						<tr>
							<td>Uma</td>
							<td>Pune</td>
							<td>411027</td>
							<td></td>
							<td></td>
						</tr>
						<tr>
							<td>Tanmay</td>
							<td>Bangalore</td>
							<td>560001</td>
							<td></td>
							<td></td>
						</tr>
						<tr>
							<td>Sachin</td>
							<td>Mumbai</td>
							<td>400003</td>
							<td></td>
							<td></td>
						</tr>
						<tr>
							<td>Uma</td>
							<td>Pune</td>
							<td>411027</td>
							<td></td>
							<td></td>
						</tr>
					</tbody>
				</table>
				<div style="width: 100%; text-align: center">
					<ul class="pagination">
						<li><a href="#">&laquo;</a></li>
						<li><a href="#">1</a></li>
						<li><a href="#">2</a></li>
						<li><a href="#">3</a></li>
						<li><a href="#">4</a></li>
						<li><a href="#">5</a></li>
						<li><a href="#">&raquo;</a></li>
					</ul>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
module.exports = {
	data: function () {
		return {};
	},
	created: function () {
		console.log("this is task");
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
	border-color: rgb(203, 234, 238);
}

#mytask {
	height: 100%;
}

tr,
td {
	height: 50px;
}
</style>
