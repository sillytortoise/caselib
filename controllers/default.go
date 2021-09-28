package controllers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/url"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var dbUri string = "neo4j://localhost:7687"

type MainController struct {
	beego.Controller
}

type LoginController struct {
	beego.Controller
}

type RegisterController struct {
	beego.Controller
}

type AssessTargetController struct {
	beego.Controller
}

type Modlevel1Controller struct {
	beego.Controller
}

type Modlevel2Controller struct {
	beego.Controller
}

type VueController struct {
	beego.Controller
}

type PicsController struct {
	beego.Controller
}

type TplController struct {
	beego.Controller
}

type CssController struct {
	beego.Controller
}

type JsController struct {
	beego.Controller
}

type InfoController struct {
	beego.Controller
}

type FuncController struct {
	beego.Controller
}
type MysqlController struct {
	beego.Controller
}

type ImageController struct {
	beego.Controller
}

type TaskController struct {
	beego.Controller
}

type PageController struct {
	beego.Controller
}

type Node struct {
	Id       string `json:"id"`
	Item     string `json:"item"`
	Target   string `json:"target"`
	Problem  string `json:"problem"`
	Advice   string `json:"advice"`
	Priority string `json:"p"`
	T        string `json:"t"`
}

type Item struct {
	Origin   string        `json:"origin"`
	Id       int           `json:"id"`
	Label    string        `json:"label"`
	Level    int           `json:"level"`
	Problem  string        `json:"problem"`
	Children []interface{} `json:"children"`
}

type Taskitem struct {
	Name     string `json:"name"`
	Created  string `json:"created"`
	Modified string `json:"modified"`
	Owner    string `json:"owner"`
	Type     string `json:"type"`
}

type PageOutline struct {
	Title string `json:"title"`
	Order int64  `json:"order"`
}

type Page struct { //竞品分析页面
	Title   string `json:"title"`
	Target  string `json:"target"`
	Problem string `json:"problem"`
	Advice  string `json:"advice"`
	Pics    []Pic  `json:"pics"`
}

type PageCase struct { //特色化案例库页面
	Title         string `json:"title"`
	Order         int64  `json:"order"`
	Case_num      string `json:"case_num"`
	Product_class string `json:"product_class"`
	Name          string `json:"name"`
	Version       string `json:"version"`
	App_class     string `json:"app_class"`
	Date          string `json:"date"`
	Feature       string `json:"feature"`
	Abstract      string `json:"abstract"`
	Page_num      string `json:"page_num"`
	Detail        string `json:"detail"`
	Idea          string `json:"idea"`
	Recorder      string `json:"recorder"`
	Username      string `json:"username"`
	Pics          []Pic  `json:"pics"`
}

type Bank struct {
	BankName string `json:"label"`
	Value    string `json:"value"`
	Vers     []Ver  `json:"children"`
}

type Ver struct {
	Vername string `json:"label"`
	Value   string `json:"value"`
}

type Pic struct {
	Path  string `json:"path"`
	Bank  string `json:"bank"`
	Ver   string `json:"ver"`
	Order int64  `json:"order"` //order本身存储在contain关系里。这里用来传给前端
	Title string `json:"title"`
}

type TotalPic struct {
	Total int64 `json:"total"`
	Imgs  []Pic `json:"imgs"`
}

type Contains struct {
	Order int `json:"order"`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/versions?charset=utf8&loc=Local")
}

func (c *MainController) Get() {
	fmt.Println(c.Ctx.GetCookie("username"))
	if c.Ctx.GetCookie("username") == "" {
		c.TplName = "signin.html"
	} else {
		c.TplName = "index.html"
	}
}

func (c *VueController) Get() {
	if c.Ctx.GetCookie("username") == "" {
		c.TplName = "signin.html"
		return
	}
	uri := c.Ctx.Request.RequestURI
	b, err := ioutil.ReadFile("./views" + uri) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	c.Ctx.WriteString(string(b))
}

func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	remember := c.GetString("remember")
	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "980115", ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, _ := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run("match(n:User) where n.username=$username and n.password=$password return n.username as u", map[string]interface{}{"username": username, "password": password})
		if err != nil {
			panic(err)
		}
		if records.Next() { //验证通过
			if remember == "true" {
				c.Ctx.SetCookie("username", username, time.Hour*24*10)
				c.Ctx.SetCookie("password", password, time.Hour*24*10)
			} else {
				c.Ctx.SetCookie("username", username, time.Second*3600*24)
				c.Ctx.SetCookie("password", password, time.Second*3600*24)

			}
			return "1", err
		} else { //验证不通过
			return "0", err
		}

	})
	flag, _ := result.(string)
	c.Ctx.WriteString(flag)
}

func (c *RegisterController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "980115", ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, _ := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run("match(n:User) where n.username=$username return n.username as u", map[string]interface{}{"username": username})
		if err != nil {
			panic(err)
		}
		if records.Next() {
			return 0, err //已被注册
		} else {
			return 1, err
		}
	})
	flag, _ := result.(int)
	fmt.Println(flag)
	if flag == 0 {
		c.Ctx.WriteString("0")
	} else {
		_, _ = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			r, e := tx.Run("create(n:User{username:$username,password:$password,privilege:'plain'})", map[string]interface{}{"username": username, "password": password})
			if e != nil {
				panic(e)
			}
			return r, e
		})
		c.Ctx.WriteString("1")
	}

}

func (c *AssessTargetController) Get() {
	if c.Ctx.GetCookie("username") == "" {
		c.TplName = "signin.html"
		return
	}
	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "980115", ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run("match(n:Assess) return n.id as id, n.item as item, n.target as target, n.problem as problem, n.advice as advice, n.priority as p, n.type  as t order by n.id", map[string]interface{}{})
		if err != nil {
			panic(err)
		}
		var results []Node
		for records.Next() {
			record := records.Record()
			id, _ := record.Get("id")
			item, _ := record.Get("item")
			if item == nil {
				item = ""
			}
			target, _ := record.Get("target")
			if target == nil {
				target = ""
			}
			problem, _ := record.Get("problem")
			if problem == nil {
				problem = ""
			}
			advice, _ := record.Get("advice")
			if advice == nil {
				advice = ""
			}
			priority, _ := record.Get("p")
			t, _ := record.Get("t")
			results = append(results, Node{id.(string), item.(string), target.(string), problem.(string), advice.(string), priority.(string), t.(string)})
		}

		return results, nil
	})
	if err != nil {
		fmt.Print(err)
	}
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *Modlevel1Controller) Post() {
	var ob Item
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "980115", ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	reg1 := regexp.MustCompile(`[0-9]+`)
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}

	for _, i := range ob.Children {
		result1 := reg1.FindAllStringSubmatch(i.(map[string]interface{})["problem"].(string), -1)
		_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			result, e := tx.Run("match(n:Assess) where n.id=$o set n.id=$id", map[string]interface{}{"id": i.(map[string]interface{})["problem"], "o": ob.Origin + result1[0][0]})
			return result, e
		})
		if err != nil {
			panic(err)
		}

	}

	c.Ctx.WriteString("")
}

func (c *Modlevel2Controller) Post() {
	name := c.GetString("name")
	item := c.GetString("item")
	target := c.GetString("target")
	problem := c.GetString("problem")
	advice := c.GetString("advice")
	p := c.GetString("p")
	t := c.GetString("t")

	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "980115", ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	_, _ = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run("match(n:Assess) where n.id=$id set n.item=$item, n.target=$target, n.problem=$problem, n.advice=$advice, n.priority=$p, n.type=$t", map[string]interface{}{"id": name, "item": item, "target": target, "problem": problem, "advice": advice, "p": p, "t": t})
		if err != nil {
			panic(err)
		}
		return result, nil
	})
	c.Ctx.WriteString("success")
}

func (c *PicsController) Post() {
	f, h, _ := c.GetFile("file") //获取上传的文件
	ext := path.Ext(h.Filename)

	//创建目录
	uploadDir := "static/upload/" //+ time.Now().Format("2006/01/02/")
	err := os.MkdirAll(uploadDir, 777)
	if err != nil {
		c.Ctx.WriteString(fmt.Sprintf("%v", err))
		return
	}
	//构造文件名称
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
	hashName := md5.Sum([]byte(time.Now().Format("2006_01_02_15_04_05_") + randNum))

	fileName := fmt.Sprintf("%x", hashName) + ext
	//c.Ctx.WriteString(  fileName )

	fpath := uploadDir + fileName
	defer f.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	err = c.SaveToFile("file", fpath)
	if err != nil {
		c.Ctx.WriteString(fmt.Sprintf("%v", err))
	}
	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "980115", ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	_, _ = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run("create(n:Pic{path:$p})", map[string]interface{}{"p": fileName})

		s := c.GetString("select")
		ss := strings.Split(s, ",")

		for _, i := range ss {
			fmt.Println(i)
			_, err = tx.Run("match (n:Assess{id:$s}),(m:Pic{path:$p}) create (n)-[p:POINTS]->(m)", map[string]interface{}{"s": i, "p": fpath})
		}
		if err != nil {
			panic(err)
		}
		return result, nil
	})
	c.Ctx.WriteString("")
}

func (c *TplController) Get() {
	if c.Ctx.GetCookie("username") == "" {
		c.TplName = "signin.html"
		return
	}
	uri := c.Ctx.Input.Param(":filename")
	b, err := ioutil.ReadFile("./views/" + uri) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	c.Ctx.WriteString(string(b))
}

func (c *CssController) Get() {
	if c.Ctx.GetCookie("username") == "" {
		c.TplName = "signin.html"
		return
	}
	uri := c.Ctx.Request.RequestURI
	b, err := ioutil.ReadFile("./static/css" + uri) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	c.Ctx.WriteString(string(b))
}

func (c *ImageController) Get() {
	if c.Ctx.GetCookie("username") == "" {
		c.TplName = "signin.html"
		return
	}
	uri := c.Ctx.Request.RequestURI
	b, err := ioutil.ReadFile("./static/upload" + uri) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	c.Ctx.WriteString(string(b))
}

func (c *JsController) Get() {
	if c.Ctx.GetCookie("username") == "" {
		c.TplName = "signin.html"
		return
	}
	uri := c.Ctx.Request.RequestURI
	b, err := ioutil.ReadFile("./static/js" + uri) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	c.Ctx.WriteString(string(b))
}

func (c *InfoController) Get() {
	if c.Ctx.GetCookie("username") == "" {
		c.TplName = "signin.html"
		return
	}
	id := c.GetString("id")
	fmt.Println(id)
	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "980115", ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, _ := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run("match(n:Assess) where n.id=$id return n.item as item, n.target as target, n.problem as problem, n.advice as advice, n.priority as p, n.type as t", map[string]interface{}{"id": id})
		if err != nil {
			panic(err)
		}
		result.Next()
		r := result.Record()
		item, _ := r.Get("item")
		if item == nil {
			item = ""
		}
		target, _ := r.Get("target")
		if target == nil {
			target = ""
		}
		problem, _ := r.Get("problem")
		if problem == nil {
			problem = ""
		}
		advice, _ := r.Get("advice")
		if advice == nil {
			advice = ""
		}
		priority, _ := r.Get("p")
		t, _ := r.Get("t")

		return Node{id, item.(string), target.(string), problem.(string), advice.(string), priority.(string), t.(string)}, nil
	})
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *FuncController) Get() {
	if c.Ctx.GetCookie("username") == "" {
		c.TplName = "signin.html"
		return
	}
	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "980115", ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, _ := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run("match (n:Func) return n.name as name", map[string]interface{}{})
		if err != nil {
			panic(err)
		}
		var functions []string
		for records.Next() {
			r := records.Record()
			rr, _ := r.Get("name")
			functions = append(functions, rr.(string))
		}
		return functions, err
	})
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *TaskController) Post() { //创建任务
	username := c.GetString("username")
	name := c.GetString("name")
	t := c.GetString("type")
	create_time := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "980115", ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	r, _ := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		r, err := tx.Run("match (n:User{username:$username})-[r:Control{permission:'own'}]->(t:Task{name:$name}) return t.name as name", map[string]interface{}{"username": username, "name": name})
		if err != nil {
			panic(err)
		}
		if r.Next() {
			return 0, err //有同名的
		} else {
			return 1, err
		}
	})

	flag, _ := r.(int)
	if flag == 0 {
		c.Ctx.WriteString("0")
		return
	}

	_, _ = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run("match (n:User{username:$username}) create (n)-[r:Control{permission:'own'}]->(t:Task{name:$name,t:$t,created:$c,modified:$c,owner:$o,last:1,totalnum:1})", map[string]interface{}{"username": username, "name": name, "t": t, "c": create_time, "o": username})
		if err != nil {
			panic(err)
		}
		if t == "1" {
			_, err = tx.Run("match (n:User{username:$username})-[r:Control{permission:'own'}]->(t:Task{name:$name}) create (p:Page{order:1,title:'',targetid:'',problem:'',advice:'',pic:''})<-[i:Include]-(t)", map[string]interface{}{"username": username, "name": name})
		} else {
			t := time.Now()
			s := t.Format("2006-01-02")
			_, err = tx.Run(`match (n:User{username:$username})-[r:Control{permission:'own'}]->(t:Task{name:$name}) 
			create (t)-[i:Include]->(p:Page{order:1,title:'',case_num:'',product_class:'',name:'',version:'',app_class:'', 
			date:$d, feature:'', abstract:'', page_name:'', detail:'', idea:'', username:$username})`, map[string]interface{}{"username": username, "name": name, "d": s})
		}
		if err != nil {
			panic(err)
		}
		return result, err
	})
	c.Ctx.WriteString("1")
}

func (c *TaskController) Get() { //获取任务列表
	if c.Ctx.GetCookie("username") == "" {
		c.TplName = "signin.html"
		return
	}
	page := c.GetString("page")
	//sort := c.GetString("sort")
	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "980115", ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	results, _ := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		pagenum, _ := strconv.Atoi(page)
		r, err := tx.Run("match (t:Task) return t.name as name, t.t as ty, t.created as c, t.modified as m, t.owner as owner order by t.modified desc skip ($page-1)*10 limit 10", map[string]interface{}{"page": pagenum})
		if err != nil {
			panic(err)
		}
		var objs []Taskitem
		for r.Next() {
			result := r.Record()
			name, _ := result.Get("name")
			ty, _ := result.Get("ty")
			created, _ := result.Get("c")
			modified, _ := result.Get("m")
			owner, _ := result.Get("owner")
			objs = append(objs, Taskitem{name.(string), created.(string), modified.(string), owner.(string), ty.(string)})
		}
		return objs, nil
	})
	c.Data["json"] = results
	c.ServeJSON()
}

func (c *TaskController) Total() {
	if c.Ctx.GetCookie("username") == "" {
		c.TplName = "signin.html"
		return
	}
	username := c.GetString("username")
	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "980115", ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	results, _ := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		r, err := tx.Run("match (t:Task)<-[r:Control]-(u:User{username:$u}) return count(*) as count", map[string]interface{}{"u": username})
		if err != nil {
			panic(err)
		}
		if r.Next() {
			record := r.Record()
			r, _ := record.Get("count")
			return r, err
		} else {
			return 0, err
		}
	})
	rstr := strconv.Itoa(int(results.(int64)))
	c.Ctx.WriteString(rstr)
}

func (c *TaskController) Delete() {
	if c.Ctx.GetCookie("username") == "" {
		c.TplName = "signin.html"
		return
	}
	username := c.Ctx.GetCookie("username")
	task_name := c.GetString("name")

	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "980115", ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	_, _ = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		r, err := tx.Run("match (t:Task{name:$name})<-[r:Control]-(u:User{username:$username}), (t)-[i:Include]->(p:Page) delete r,i,p,t", map[string]interface{}{"name": task_name, "username": username})
		if err != nil {
			panic(err)
		}
		return r, err
	})
	c.Ctx.WriteString("1")
}

func (c *PageController) Get() { //changePage ?p=
	if c.Ctx.GetCookie("username") == "" {
		c.TplName = "signin.html"
		return
	}
	page := c.GetString("p")
	task_type := c.GetString("type")
	encode_user := c.Ctx.Input.Param(":user")
	encode_name := c.Ctx.Input.Param(":name")
	user, _ := url.QueryUnescape(encode_user)
	name, _ := url.QueryUnescape(encode_name)

	if page == "" {
		if task_type == "1" {
			c.TplName = "analysis.html"
		} else {
			c.TplName = "case.html"
		}
		return
	}
	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "980115", ""))
	if err != nil {
		panic(err)
	}

	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	result, _ := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		p, _ := strconv.Atoi(page)
		var r neo4j.Result
		if task_type == "1" { //竞品分析
			r, err = tx.Run(`match (u:User{username:$user})-[r:Control]->(t:Task{name:$name}),
			(t)-[i:Include]->(p:Page{order:$order}) 
			return p.title as title, 
			p.targetid as target, 
			p.problem as problem, 
			p.advice as advice`, map[string]interface{}{"user": user, "name": name, "order": p})
			if err != nil {
				panic(err)
			}
			var obj Page
			if r.Next() {
				record := r.Record()
				title, _ := record.Get("title")
				target, _ := record.Get("target")
				problem, _ := record.Get("problem")
				advice, _ := record.Get("advice")
				obj.Title, _ = title.(string)
				obj.Target, _ = target.(string)
				obj.Problem, _ = problem.(string)
				obj.Advice, _ = advice.(string)
			}

			r, err = tx.Run(`match (u:User{username:$user})-[r:Control]->(t:Task{name:$name}),
			(t)-[i:Include]->(p:Page{order:$order}),
			(p)-[c:Contains]->(pic:Pic)
			return c.order as order,c.title as title, pic.path as path order by order`, map[string]interface{}{"user": user, "name": name, "order": p})
			if err != nil {
				panic(err)
			}
			for r.Next() {
				record := r.Record()
				path, _ := record.Get("path")
				pathstr, _ := path.(string)
				order, _ := record.Get("order")
				orderint, _ := order.(int64)
				title, _ := record.Get("title")
				titlestr, _ := title.(string)
				obj.Pics = append(obj.Pics, Pic{Path: pathstr, Order: orderint, Title: titlestr})
			}

			return obj, nil
		} else { //特色化案例库
			r, err = tx.Run(`match (u:User{username:$user})-[r:Control]->(t:Task{name:$name}),
			(t)-[i:Include]->(p:Page{order:$order}) 
			return p.title as title,
			p.case_num as case_num,
			p.product_class as product_class,
			p.name as name,
			p.version as version,
			p.app_class as app_class,
			p.date as mdate,
			p.feature as feature,
			p.abstract as abstract,
			p.page_num as page_num,
			p.detail as detail,
			p.idea as idea,
			p.recorder as recorder,
			p.username as username`, map[string]interface{}{"user": user, "name": name, "order": p})
			if err != nil {
				panic(err)
			}
			var obj PageCase
			if r.Next() {
				record := r.Record()
				title, _ := record.Get("title")
				case_num, _ := record.Get("case_num")
				product_class, _ := record.Get("product_class")
				name, _ := record.Get("name")
				version, _ := record.Get("version")
				app_class, _ := record.Get("app_class")
				date, _ := record.Get("mdate")
				feature, _ := record.Get("feature")
				abstract, _ := record.Get("abstract")
				page_num, _ := record.Get("page_num")
				detail, _ := record.Get("detail")
				idea, _ := record.Get("idea")
				recorder, _ := record.Get("recorder")
				username, _ := record.Get("username")
				obj.Title, _ = title.(string)
				obj.Case_num, _ = case_num.(string)
				obj.Product_class, _ = product_class.(string)
				obj.Name, _ = name.(string)
				obj.Version, _ = version.(string)
				obj.App_class, _ = app_class.(string)
				obj.Date, _ = date.(string)
				obj.Feature, _ = feature.(string)
				obj.Abstract, _ = abstract.(string)
				obj.Page_num, _ = page_num.(string)
				obj.Detail, _ = detail.(string)
				obj.Idea, _ = idea.(string)
				obj.Recorder, _ = recorder.(string)
				obj.Username = username.(string)
			}

			r, err = tx.Run(`match (u:User{username:$user})-[r:Control]->(t:Task{name:$name}),
		(t)-[i:Include]->(p:Page{order:$order}),
		(p)-[c:Contains]->(pic:Pic) 
		return c.order as order, c.title as title, pic.path as path order by order`, map[string]interface{}{"user": user, "name": name, "order": p})
			if err != nil {
				panic(err)
			}
			for r.Next() {
				record := r.Record()
				path, _ := record.Get("path")
				order, _ := record.Get("order")
				orderint, _ := order.(int64)
				pathstr, _ := path.(string)
				title, _ := record.Get("title")
				titlestr, _ := title.(string)
				obj.Pics = append(obj.Pics, Pic{Path: pathstr, Order: orderint, Title: titlestr})
			}
			return obj, nil

		}

	})
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *PageController) Get_pages() {
	if c.Ctx.GetCookie("username") == "" {
		c.TplName = "signin.html"
		return
	}
	encode_user := c.Ctx.Input.Param(":user")
	encode_name := c.Ctx.Input.Param(":name")
	user, _ := url.QueryUnescape(encode_user)
	name, _ := url.QueryUnescape(encode_name)

	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "980115", ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, _ := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		r, err := tx.Run("match (u:User{username:$user})-[r:Control]->(t:Task{name:$name}), (t)-[i:Include]->(p:Page) return p.title as t, p.order as o order by o", map[string]interface{}{"name": name, "user": user})
		if err != nil {
			panic(err)
		}

		var pages []PageOutline
		for r.Next() {
			results := r.Record()
			title, _ := results.Get("t")
			order, _ := results.Get("o")
			pages = append(pages, PageOutline{title.(string), order.(int64)})
		}
		return pages, err
	})
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *PageController) Addtoend() {
	if c.Ctx.GetCookie("username") == "" {
		c.TplName = "signin.html"
		return
	}
	encode_user := c.Ctx.Input.Param(":user")
	encode_name := c.Ctx.Input.Param(":name")
	user, _ := url.QueryUnescape(encode_user)
	name, _ := url.QueryUnescape(encode_name)
	t := c.GetString("type")
	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "980115", ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	if t == "1" {
		_, _ = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			_, err := tx.Run("match (u:User{username:$user})-[r:Control]->(t:Task{name:$name}) create (t)-[i:Include]->(p:Page{title:'',order:t.totalnum+1}) set t.totalnum=t.totalnum+1", map[string]interface{}{"name": name, "user": user})
			if err != nil {
				panic(err)
			}

			return nil, nil
		})
	} else {
		d := time.Now()
		s := d.Format("2006-01-02")
		_, _ = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			_, err = tx.Run(`match (n:User{username:$username})-[r:Control]->(t:Task{name:$name}) 
			create (t)-[i:Include]->(p:Page{order:t.totalnum+1, title:'',case_num:'',product_class:'',name:'',version:'',app_class:'', 
			date:$d, feature:'', abstract:'', page_name:'', detail:'', idea:'', username:$username}) set t.totalnum=t.totalnum+1`, map[string]interface{}{"username": user, "name": name, "d": s})
			if err != nil {
				panic(err)
			}

			return nil, nil
		})
	}
	c.Ctx.WriteString("1")
}

func (c *PageController) Addtonext() {
	if c.Ctx.GetCookie("username") == "" {
		c.TplName = "signin.html"
		return
	}
	encode_user := c.Ctx.Input.Param(":user")
	encode_name := c.Ctx.Input.Param(":name")
	user, _ := url.QueryUnescape(encode_user)
	name, _ := url.QueryUnescape(encode_name)
	p := c.GetString("p")
	t := c.GetString("type")
	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "980115", ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	pint, _ := strconv.Atoi(p)
	_, _ = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := tx.Run("match (u:User{username:$user})-[r:Control]->(t:Task{name:$name}), (t)-[i:Include]->(p:Page) where p.order>$order set p.order=p.order+1", map[string]interface{}{"name": name, "user": user, "order": pint})
		if err != nil {
			panic(err)
		}
		_, _ = tx.Run("match (u:User{username:$user})-[r:Control]->(t:Task{name:$name}) set t.totalnum=t.totalnum+1", map[string]interface{}{"name": name, "user": user})

		if t == "1" {
			_, err = tx.Run("match (u:User{username:$user})-[r:Control]->(t:Task{name:$name}) create (t)-[i:Include]->(p:Page{title:'', order:$order})", map[string]interface{}{"user": user, "name": name, "order": pint + 1})

		} else {
			s := time.Now().Format("2006-01-02")
			_, err = tx.Run(`match (n:User{username:$username})-[r:Control]->(t:Task{name:$name}) 
			create (t)-[i:Include]->(p:Page{order:$order, title:'',case_num:'',product_class:'',name:'',version:'',app_class:'', 
			date:$d, feature:'', abstract:'', page_name:'', detail:'', idea:'', username:$username})`, map[string]interface{}{"username": user, "name": name, "d": s, "order": pint + 1})

		}
		if err != nil {
			panic(err)
		}
		return nil, err
	})
	c.Ctx.WriteString("1")
}

func (c *PageController) DeletePage() {
	if c.Ctx.GetCookie("username") == "" {
		c.TplName = "signin.html"
		return
	}
	p := c.GetString("p")
	encode_user := c.Ctx.Input.Param(":user")
	encode_name := c.Ctx.Input.Param(":name")
	user, _ := url.QueryUnescape(encode_user)
	name, _ := url.QueryUnescape(encode_name)
	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "980115", ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	pint, _ := strconv.Atoi(p)
	_, _ = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := tx.Run(`match (u:User{username:$user})-[r:Control]->(t:Task{name:$name}),
		(t)-[i:Include]->(p:Page{order:$order}),
		 (p)-[c:Contains]->(pic:Pic) delete c`, map[string]interface{}{"name": name, "user": user, "order": pint})
		if err != nil {
			panic(err)
		}
		_, err = tx.Run("match (u:User{username:$user})-[r:Control]->(t:Task{name:$name}),(t)-[i:Include]->(p:Page{order:$order}) delete i,p", map[string]interface{}{"name": name, "user": user, "order": pint})
		if err != nil {
			panic(err)
		}
		_, err = tx.Run("match (u:User{username:$user})-[r:Control]->(t:Task{name:$name}),(t)-[i:Include]->(p:Page) where p.order>$p set p.order=p.order-1", map[string]interface{}{"name": name, "user": user, "p": pint})
		if err != nil {
			panic(err)
		}
		_, err = tx.Run("match (u:User{username:$user})-[r:Control]->(t:Task{name:$name}) set t.totalnum=t.totalnum-1", map[string]interface{}{"name": name, "user": user})
		return nil, err
	})
	c.Ctx.WriteString("1")
}

func (c *MysqlController) Getbv() {
	if c.Ctx.GetCookie("username") == "" {
		c.TplName = "signin.html"
		return
	}
	o := orm.NewOrm()
	var lists []orm.ParamsList
	var banks []Bank
	_, err := o.Raw("select distinct bank_name from versions").ValuesList(&lists)
	if err != nil {
		panic(err)
	}
	for _, item := range lists {
		var bank Bank
		bank.BankName = item[0].(string)
		bank.Value = item[0].(string)
		var ls []orm.ParamsList
		_, e := o.Raw("select distinct ver from versions where bank_name=?", item[0].(string)).ValuesList(&ls)
		if e != nil {
			panic(e)
		}
		for _, value := range ls {
			var ver Ver
			ver.Vername = value[0].(string)
			ver.Value = value[0].(string)
			bank.Vers = append(bank.Vers, ver)
		}
		banks = append(banks, bank)
	}

	c.Data["json"] = banks
	c.ServeJSON()
}

func (c *PageController) Upload_pic() { //竞品分析&特色化案例库页面上传图片
	if c.Ctx.GetCookie("username") == "" {
		c.TplName = "signin.html"
		return
	}
	f, h, _ := c.GetFile("file") //获取上传的文件
	ext := path.Ext(h.Filename)

	//创建目录
	uploadDir := "static/upload/" // + time.Now().Format("2006/01/02/")
	err := os.MkdirAll(uploadDir, 777)
	if err != nil {
		c.Ctx.WriteString(fmt.Sprintf("%v", err))
		return
	}
	//构造文件名称
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
	hashName := md5.Sum([]byte(time.Now().Format("2006_01_02_15_04_05_") + randNum))

	fileName := fmt.Sprintf("%x", hashName) + ext
	//c.Ctx.WriteString(  fileName )

	fpath := uploadDir + fileName
	defer f.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	err = c.SaveToFile("file", fpath)
	if err != nil {
		c.Ctx.WriteString(fmt.Sprintf("%v", err))
	}

	page := c.GetString("page")
	p := c.GetString("p")
	encode_user := c.Ctx.Input.Param(":user")
	encode_name := c.Ctx.Input.Param(":task")
	user, _ := url.QueryUnescape(encode_user)
	task, _ := url.QueryUnescape(encode_name)
	fun := c.GetString("func")
	bank := c.GetString("bank")
	ver := c.GetString("ver")

	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "980115", ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	pint, _ := strconv.Atoi(page)
	rank, _ := strconv.Atoi(p)

	_, _ = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		fmt.Println(user, task, pint, rank)
		r, err := tx.Run("match (u:User{username:$user})-[r:Control]->(t:Task{name:$name}),(t)-[i:Include]->(p:Page{order:$order}),(p)-[c:Contains{order:$rank}]->(pic:Pic) return pic as pi", map[string]interface{}{"user": user, "name": task, "order": pint, "rank": rank + 1})
		if err != nil {
			panic(err)
		}
		if r.Next() { //如果原来有图  删除关系
			_, err = tx.Run("match (u:User{username:$user})-[r:Control]->(t:Task{name:$name}),(t)-[i:Include]->(p:Page{order:$order}),(p)-[c:Contains{order:$rank}]->(pic:Pic) delete c", map[string]interface{}{"user": user, "name": task, "order": pint, "rank": rank + 1})
			if err != nil {
				panic(err)
			}
		}

		/*创建关系和图片*/
		_, err = tx.Run("match (u:User{username:$user})-[r:Control]->(t:Task{name:$name}),(t)-[i:Include]->(p:Page{order:$order}) create (p)-[c:Contains{order:$rank}]->(pic:Pic{path:$path,bank:$bank,ver:$ver})", map[string]interface{}{"user": user, "name": task, "order": pint, "path": fileName, "rank": rank + 1, "bank": bank, "ver": ver})
		if err != nil {
			panic(err)
		}

		/*建立与功能模块的关系 implement*/
		_, err = tx.Run("match (f:Func{name:$name}),(p:Pic{path:$path,bank:$bank,ver:$ver}) create (p)-[i:Implements]->(f)", map[string]interface{}{"name": fun, "path": fileName, "bank": bank, "ver": ver})
		if err != nil {
			panic(err)
		}
		return nil, nil
	})
	c.Ctx.WriteString("success")
}

func (c *PageController) Autosave() {
	if c.Ctx.GetCookie("username") == "" {
		c.TplName = "signin.html"
		return
	}
	t := c.GetString("type")
	encode_user := c.Ctx.Input.Param(":user")
	encode_name := c.Ctx.Input.Param(":task")
	user, _ := url.QueryUnescape(encode_user)
	task, _ := url.QueryUnescape(encode_name)
	page := c.GetString("p")
	pint, _ := strconv.Atoi(page)

	imgtitle0 := c.GetString("imgtitle0")
	imgpath0 := c.GetString("imgpath0")
	imgtitle1 := c.GetString("imgtitle1")
	imgpath1 := c.GetString("imgpath1")
	imgtitle2 := c.GetString("imgtitle2")
	imgpath2 := c.GetString("imgpath2")
	imgtitle3 := c.GetString("imgtitle3")
	imgpath3 := c.GetString("imgpath3")
	imgtitle4 := c.GetString("imgtitle4")
	imgpath4 := c.GetString("imgpath4")
	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "980115", ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	if t == "1" { //竞品分析
		title := c.GetString("title")
		problem := c.GetString("problem")
		advice := c.GetString("advice")
		targetid := c.GetString("targetid")

		fmt.Println(title, problem, advice, targetid, pint)

		_, _ = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			_, err := tx.Run(`match (u:User{username:$user})-[r:Control]->(t:Task{name:$name}),
			(t)-[i:Include]->(p:Page{order:$order}) 
			set p.title=$title,p.problem=$problem,p.advice=$advice,p.targetid=$targetid`,
				map[string]interface{}{"user": user, "name": task, "order": pint, "title": title, "problem": problem, "advice": advice, "targetid": targetid})
			if err != nil {
				panic(err)
			}

			return nil, nil
		})
	} else { //特色化案例库
		title := c.GetString("title")
		case_num := c.GetString("case_num")
		product_class := c.GetString("product_class")
		name := c.GetString("name")
		version := c.GetString("version")
		app_class := c.GetString("app_class")
		feature := c.GetString("feature")
		abstract := c.GetString("abstract")
		page_num := c.GetString("page_num")
		detail := c.GetString("detail")
		idea := c.GetString("idea")
		recorder := c.GetString("recorder")

		_, _ = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			_, err := tx.Run(`match (u:User{username:$user})-[r:Control]->(t:Task{name:$task}),
		(t)-[i:Include]->(p:Page{order:$order})
		set p.title=$title,p.case_num=$case_num,p.product_class=$product_class,
		p.name=$name,p.version=$version,
		p.app_class=$app_class, p.feature=$feature, p.abstract=$abstract,
		p.page_num=$page_num, p.detail=$detail, p.idea=$idea, p.recorder=$recorder`,
				map[string]interface{}{"user": user, "task": task, "order": pint, "title": title, "case_num": case_num, "product_class": product_class,
					"name": name, "version": version, "app_class": app_class, "feature": feature, "abstract": abstract,
					"page_num": page_num, "detail": detail, "idea": idea, "recorder": recorder})
			if err != nil {
				panic(err)
			}

			return nil, nil
		})

	}
	_, _ = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		if imgpath0 != "" {
			imgpath0 = imgpath0[1:]
			r, err := tx.Run(`match (u:User{username:$user})-[r:Control]->(t:Task{name:$task}),
		(t)-[i:Include]->(p:Page{order:$order}),
		(p)-[c:Contains{order:1}]->(pic:Pic) return pic as pic`, map[string]interface{}{"user": user, "task": task, "order": pint})
			if err != nil {
				panic(err)
			}
			if r.Next() { //对应位置上已经有了图片，解开关系
				_, err = tx.Run(`match (u:User{username:$user})-[r:Control]->(t:Task{name:$task}), 
			(t)-[i:Include]->(p:Page{order:$order}), 
			(p)-[c:Contains{order:1}]->(pic:Pic) delete c`, map[string]interface{}{"user": user, "task": task, "order": pint})
				if err != nil {
					panic(err)
				}
			}
			r, err = tx.Run(`match (p:Pic{path:$path}) return p as pic`, map[string]interface{}{"path": imgpath0})
			if err != nil {
				panic(err)
			}
			if r.Next() { //如果这张图片已经存在
				_, err = tx.Run(`match (u:User{username:$user})-[r:Control]->(t:Task{name:$task}), 
			(t)-[i:Include]->(p:Page{order:$order}), (pic:Pic{path:$path}) create (p)-[c:Contains{title:$title,order:1}]->(pic)`, map[string]interface{}{"user": user, "task": task, "order": pint, "path": imgpath0, "title": imgtitle0})
				if err != nil {
					panic(err)
				}
			}
		}

		if imgpath1 != "" {
			imgpath1 = imgpath1[1:]
			r, err := tx.Run(`match (u:User{username:$user})-[r:Control]->(t:Task{name:$task}),
		(t)-[i:Include]->(p:Page{order:$order}),
		(p)-[c:Contains{order:2}]->(pic:Pic) return pic as pic`, map[string]interface{}{"user": user, "task": task, "order": pint})
			if err != nil {
				panic(err)
			}
			if r.Next() { //对应位置上已经有了图片，解开关系
				_, err = tx.Run(`match (u:User{username:$user})-[r:Control]->(t:Task{name:$task}), 
			(t)-[i:Include]->(p:Page{order:$order}), 
			(p)-[c:Contains{order:2}]->(pic:Pic) delete c`, map[string]interface{}{"user": user, "task": task, "order": pint})
				if err != nil {
					panic(err)
				}
			}
			r, err = tx.Run(`match (p:Pic{path:$path}) return p as pic`, map[string]interface{}{"path": imgpath1})
			if err != nil {
				panic(err)
			}
			if r.Next() { //如果这张图片已经存在
				_, err = tx.Run(`match (u:User{username:$user})-[r:Control]->(t:Task{name:$task}), 
			(t)-[i:Include]->(p:Page{order:$order}), (pic:Pic{path:$path}) create (p)-[c:Contains{title:$title,order:2}]->(pic)`, map[string]interface{}{"user": user, "task": task, "order": pint, "path": imgpath1, "title": imgtitle1})
				if err != nil {
					panic(err)
				}
			}
		}

		if imgpath2 != "" {
			imgpath2 = imgpath2[1:]
			r, err := tx.Run(`match (u:User{username:$user})-[r:Control]->(t:Task{name:$task}),
		(t)-[i:Include]->(p:Page{order:$order}),
		(p)-[c:Contains{order:3}]->(pic:Pic) return pic as pic`, map[string]interface{}{"user": user, "task": task, "order": pint})
			if err != nil {
				panic(err)
			}
			if r.Next() { //对应位置上已经有了图片，解开关系
				_, err = tx.Run(`match (u:User{username:$user})-[r:Control]->(t:Task{name:$task}), 
			(t)-[i:Include]->(p:Page{order:$order}), 
			(p)-[c:Contains{order:3}]->(pic:Pic) delete c`, map[string]interface{}{"user": user, "task": task, "order": pint})
				if err != nil {
					panic(err)
				}
			}
			r, err = tx.Run(`match (p:Pic{path:$path}) return p as pic`, map[string]interface{}{"path": imgpath2})
			if err != nil {
				panic(err)
			}
			if r.Next() { //如果这张图片已经存在
				_, err = tx.Run(`match (u:User{username:$user})-[r:Control]->(t:Task{name:$task}), 
			(t)-[i:Include]->(p:Page{order:$order}), (pic:Pic{path:$path}) create (p)-[c:Contains{title:$title,order:3}]->(pic)`, map[string]interface{}{"user": user, "task": task, "order": pint, "path": imgpath2, "title": imgtitle2})
				if err != nil {
					panic(err)
				}
			}
		}

		if imgpath3 != "" {
			imgpath3 = imgpath3[1:]
			r, err := tx.Run(`match (u:User{username:$user})-[r:Control]->(t:Task{name:$task}),
		(t)-[i:Include]->(p:Page{order:$order}),
		(p)-[c:Contains{order:4}]->(pic:Pic) return pic as pic`, map[string]interface{}{"user": user, "task": task, "order": pint})
			if err != nil {
				panic(err)
			}
			if r.Next() { //对应位置上已经有了图片，解开关系
				_, err = tx.Run(`match (u:User{username:$user})-[r:Control]->(t:Task{name:$task}), 
			(t)-[i:Include]->(p:Page{order:$order}), 
			(p)-[c:Contains{order:4}]->(pic:Pic) delete c`, map[string]interface{}{"user": user, "task": task, "order": pint})
				if err != nil {
					panic(err)
				}
			}
			r, err = tx.Run(`match (p:Pic{path:$path}) return p as pic`, map[string]interface{}{"path": imgpath3})
			if err != nil {
				panic(err)
			}
			if r.Next() { //如果这张图片已经存在
				_, err = tx.Run(`match (u:User{username:$user})-[r:Control]->(t:Task{name:$task}), 
			(t)-[i:Include]->(p:Page{order:$order}), (pic:Pic{path:$path}) create (p)-[c:Contains{title:$title,order:4}]->(pic)`, map[string]interface{}{"user": user, "task": task, "order": pint, "path": imgpath3, "title": imgtitle3})
				if err != nil {
					panic(err)
				}
			}
		}

		if imgpath4 != "" {
			imgpath4 = imgpath4[1:]
			r, err := tx.Run(`match (u:User{username:$user})-[r:Control]->(t:Task{name:$task}),
		(t)-[i:Include]->(p:Page{order:$order}),
		(p)-[c:Contains{order:5}]->(pic:Pic) return pic as pic`, map[string]interface{}{"user": user, "task": task, "order": pint})
			if err != nil {
				panic(err)
			}
			if r.Next() { //对应位置上已经有了图片，解开关系
				_, err = tx.Run(`match (u:User{username:$user})-[r:Control]->(t:Task{name:$task}), 
			(t)-[i:Include]->(p:Page{order:$order}), 
			(p)-[c:Contains{order:5}]->(pic:Pic) delete c`, map[string]interface{}{"user": user, "task": task, "order": pint})
				if err != nil {
					panic(err)
				}
			}
			r, err = tx.Run(`match (p:Pic{path:$path}) return p as pic`, map[string]interface{}{"path": imgpath4})
			if err != nil {
				panic(err)
			}
			if r.Next() { //如果这张图片已经存在
				_, err = tx.Run(`match (u:User{username:$user})-[r:Control]->(t:Task{name:$task}), 
			(t)-[i:Include]->(p:Page{order:$order}), (pic:Pic{path:$path}) create (p)-[c:Contains{title:$title,order:5}]->(pic)`, map[string]interface{}{"user": user, "task": task, "order": pint, "path": imgpath4, "title": imgtitle4})
				if err != nil {
					panic(err)
				}
			}
		}
		return nil, nil
	})

	c.Ctx.WriteString("success")
}

func (c *ImageController) Allimages() {
	if c.Ctx.GetCookie("username") == "" {
		c.TplName = "signin.html"
		return
	}
	page := c.GetString("p")
	page_int, _ := strconv.Atoi(page)
	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "980115", ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	results, _ := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(`match (p:Pic) return p.path as path, p.bank as bank, p.ver as ver skip $head limit 20`, map[string]interface{}{"head": (page_int - 1) * 20})
		if err != nil {
			panic(err)
		}
		var pics []Pic
		for records.Next() {
			record := records.Record()
			path, _ := record.Get("path")
			bank, _ := record.Get("bank")
			ver, _ := record.Get("ver")
			pathstr, _ := path.(string)
			bankstr, _ := bank.(string)
			verstr, _ := ver.(string)
			pics = append(pics, Pic{Path: pathstr, Bank: bankstr, Ver: verstr})
		}
		total, err := tx.Run("match (p:Pic) return count(p) as num", map[string]interface{}{})
		if err != nil {
			panic(err)
		}
		var num int64
		if total.Next() {
			record := total.Record()
			numi, _ := record.Get("num")
			num, _ = numi.(int64)
		}
		return TotalPic{num, pics}, nil
	})
	c.Data["json"] = results
	c.ServeJSON()
}
