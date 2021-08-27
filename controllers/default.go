package controllers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

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
type BVController struct {
	beego.Controller
}

type ImageController struct {
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

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (c *VueController) Get() {
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
	uploadDir := "static/upload/" + time.Now().Format("2006/01/02/")
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
		result, err := tx.Run("create(n:Picture{path:$p})", map[string]interface{}{"p": fileName})

		s := c.GetString("select")
		ss := strings.Split(s, ",")

		for _, i := range ss {
			fmt.Println(i)
			_, err = tx.Run("match (n:Assess{id:$s}),(m:Picture{path:$p}) create (n)-[p:POINTS]->(m)", map[string]interface{}{"s": i, "p": fpath})
		}
		if err != nil {
			panic(err)
		}
		return result, nil
	})
	c.Ctx.WriteString("")
}

func (c *TplController) Get() {
	uri := c.Ctx.Input.Param(":filename")
	b, err := ioutil.ReadFile("./views/" + uri) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	c.Ctx.WriteString(string(b))
}

func (c *CssController) Get() {
	uri := c.Ctx.Request.RequestURI
	b, err := ioutil.ReadFile("./static/css" + uri) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	c.Ctx.WriteString(string(b))
}

func (c *ImageController) Get() {
	uri := c.Ctx.Request.RequestURI
	b, err := ioutil.ReadFile("./static/img" + uri) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	c.Ctx.WriteString(string(b))
}

func (c *JsController) Get() {
	uri := c.Ctx.Request.RequestURI
	b, err := ioutil.ReadFile("./static/js" + uri) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	c.Ctx.WriteString(string(b))
}

func (c *InfoController) Get() {
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
