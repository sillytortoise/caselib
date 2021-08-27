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

   

4. 创建任务 createtask?username= &name= 

5. 打开任务（总页面为0）直接返回analysis.html

   "pages"的长度为0

   

6. 打开任务(总页面不为0)

   task?name= 

   ```
   {
   	"pages":[
   		{"num":..., "title":...},
   		{"num":..., "title":...},
   		...
   	],
       ”lastpage":...
   }
   ```

   

7. 查看竞品分析页面<font color="red">某一页</font>(总页面数不为0)

   ​		task?username= &name= &page=

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





