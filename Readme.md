# 设计说明

## 1. 数据库实体

![数据库对象设计](https://github.com/sillytortoise/caselib/blob/master/%E6%95%B0%E6%8D%AE%E5%BA%93%E5%AF%B9%E8%B1%A1%E8%AE%BE%E8%AE%A1.png?raw=true)



### 特色化案例库页面  节点字段

order 第几页

title

case_num

product_class

name

version

app_class

date

username

feature	特色类型

abstract 特色摘要

page_num 案例页数

detail

idea 应用思路



## 2. 接口

1. 查图

   API: ？func= &bank= &ver= &page=

   json: {imgs:[]}

2. 注册 /register

   ```json
   返回string  msg: ""
   ```

   

3. 登录 post

   /login

   ```javasc
   data: {
   
   ​	username:...,
   
   ​	password:...,
   
   ​	remember:...
   
   }
   ```

   

4. 创建任务 createtask?username= &name= &type=

5. 获取任务列表 每页10条 gettasks?username= &page= &sort=

6. 删除任务 delete_task?name=

7. 打开竞品分析任务 返回analysis.html

8. 获取所有页面标题 /:user/:task_name/get_pages

   ```json
   [{name:'', order: ...},{},...]
   ```

   

9. 获取一个任务的所有页面 taskpages?username= &name= &type=

   ```
   {
   	"pages":[
   		{"num":..., "title":...},
   		{"num":..., "title":...},
   		...
   	],
       
   }
   ```

   

10. 查看竞品分析页面<font color="red">某一页</font>(总页面数不为0)

   ​		taskcontent?username= &name= &type= &page=

   ```json
   {
       "title":...,
       "target":...,
       "p":...,
       "t":...,
       "problem":...,
       "advice":...,
       "pics":[
       	{"info":..., "path":...},
   		{"info":..., "path":...},
   		...
       ]
       
       
   }
   ```

11. 自动保存





<font color="red">图片添加删除按钮</font>

图片浏览器





bug:

1. analysis 图片用el-image
2. 改图片路径
3. 改setpic
4. 改autosave
5. 改upload方法
6. 改上传时src blob情况

新增功能

1. 按照app 筛选图片（按照版本号排序）
2. 上传时指定app/版本/功能模块
