<template>
	<div name="container" class="block">
		<span class="功能模块"></span>
		<el-cascader
			:options="options"
			filterable="true"
			:props="{ multiple: true }"
			:show-all-levels="false"
			@change="handleChange"
			clearable
		></el-cascader>
	</div>
</template>


<script>
module.exports = {
	data() {
		return {
			options: null,
		};
	},

	emits: ["changevalue"],
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
				node.label = tags[i];
				node.value = tags[i];
				node.children = [];
				for (let t in res.data) {
					label = /[A-Z]+/.exec(res.data[t].id)[0];
					if (node.label === label) {
						child = {};
						child.label = res.data[t].problem;
						child.value = res.data[t].id;
						node.children.push(child);
					}
				}
				data.push(node);
			}
			this.options = data;
		});
	},
	methods: {
		handleChange(value) {
			this.$emit("changevalue", value);
		},
	},
};
</script>

<style scoped>
div[name="container"] {
	margin-bottom: 20px;
}
</style>