# 设计说明

## 1. 数据库实体

![数据库对象设计](C:\Users\steve\Desktop\数据库对象设计.png)



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





