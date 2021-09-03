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
			placeholder="选择银行与版本"
			v-model="value2"
			:options="options2"
			:props="{ expandTrigger: 'hover' }"
			@change="handleChange2"
		>
		</el-cascader>
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
	emits: ["change"],
	created: function () {
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
	methods: {
		handleChange1(value) {
			this.$emit("change", [value, this.value2]);
		},
		handleChange2(value) {
			this.$emit("change", [this.value1, value]);
		},
	},
};
</script>