<template>
	<div>
		<el-select
			v-model="value1"
			filterable
			placeholder="选择功能模块"
			clearable
			@change="handleChange1"
		>
			<el-option
				v-for="item in options1"
				:key="item.value"
				:label="item.label"
				:value="item.value"
			>
			</el-option>
		</el-select>
		<el-cascader
			clearable
			placeholder="选择App与版本"
			v-model="value2"
			:options="options2"
			:props="{ checkStrictly: strict }"
			@change="handleChange2"
			filterable
		>
		</el-cascader>
		<slot></slot>
	</div>
</template>

<script>
module.exports = {
	data() {
		return {
			options1: [],
			value1: "",
			value2: [],
			options2: [],
		};
	},
	props: {
		strict: Boolean, //default:false
	},
	emits: ["change"],
	created: function () {
		this.init();
	},
	methods: {
		init() {
			//获取功能模块
			axios.get("/getfunc").then((res) => {
				for (let item of res.data) {
					this.options1.push({ value: item, label: item });
				}
			});
			//获取APP和版本
			axios.get("/getbv").then((res) => {
				this.options2 = res.data;
			});
		},
		handleChange1(value) {
			this.$emit("change", [value, this.value2]);
		},
		handleChange2(value) {
			this.$emit("change", [this.value1, value]);
		},
	},
};
</script>

<style scoped>
div {
	text-align: center;
}
</style>