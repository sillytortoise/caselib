package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"

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

func (c *LoginController) Get() {
	c.TplName = "signin.html"
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
	_, h, _ := c.GetFile("file")
	fmt.Println(h)
	c.Ctx.WriteString("")
}
