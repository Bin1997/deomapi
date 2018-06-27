package controller

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/Bin1997/deomapi/entity"
	"strconv"
	"net/http"

)
var (
	db  *gorm.DB
	err error
)

//创建User
func AddUser(c *gin.Context) {

	fmt.Println(c)
	//将网页获取的字符串转为uint类型
	id,_:=strconv.Atoi(c.PostForm("id"))
	age, _ := strconv.Atoi(c.PostForm("age"))
	fmt.Println(age)
	//添加数据
	us := entity.User{ID:uint(id),Name: c.PostForm("name"), Age: uint(age), Gender: c.PostForm("gender"), Email: c.PostForm("email")}
	fmt.Println(us)

	//数据库连接
	db, err = gorm.Open("mysql", "bin:Bin123.@tcp(localhost:3306)/test")
	if err != nil {
		panic("failed to connect database")
		fmt.Println(err)
	}
	//保存到数据库中
	db.Create(us)
	defer db.Close()

	c.JSON(200, us)
	//c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "massage": "to add"})

}

//查看user内容
func FetchAll(c *gin.Context) {
	fmt.Println(c)
	var (

		db  *gorm.DB
		us []entity.User
	)

	//数据库连接
	db, err = gorm.Open("mysql", "bin:Bin123.@tcp(localhost:3306)/test")
	if err != nil {
		panic("failed to connect database")
		fmt.Println(err)
	}
	db.Find(&us)
	fmt.Println(db)
	defer db.Close()
	//如果us的元素小于等于0，就返回找不到user
	if len(us) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}
	//遍历us内容，返回数据
	for _, item := range us {
		us = append(us, entity.User{ID: item.ID, Name: item.Name, Age: item.Age, Gender: item.Gender, Email: item.Email})
	}
	c.JSON(200,us)
	//c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": us})

}

//根据id修改内容
func UpdateUser(c *gin.Context) {

	var (
		err error
		db  *gorm.DB
		us entity.User
	)
//获取前端返回字符串
	id := c.Params.ByName("id")

	//数据库连接
	db, err = gorm.Open("mysql", "bin:Bin123.@tcp(localhost:3306)/test")
	if err != nil {
		panic("failed to connect database")
		fmt.Println(err)
	}
	//根据前台返回id，查看内容，并判断是否为空
	if err = db.Where("id = ?", id).First(&us).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Print(err)
	}
		usID := c.Param("id")

		db.First(&us, usID)

		if us.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
			return
		}

		_id, _ := strconv.Atoi(c.PostForm("id"))
		db.Model(us).Update("id",_id)
		db.Model(&us).Update("Name", c.PostForm("name"))
		age, _ := strconv.Atoi(c.PostForm("age"))
		db.Model(&us).Update("Age", age)
		db.Model(&us).Update("Gender", c.PostForm("gender"))
		db.Model(&us).Update("Email", c.PostForm("email"))
		c.JSON(200,us)
		//c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "user updated successfully!"})

}

//根据id删除
func RemoveUser(c *gin.Context) {

	var us entity.User

	//数据库连接
	db, err = gorm.Open("mysql", "bin:Bin123.@tcp(localhost:3306)/test")
	if err != nil {
		panic("failed to connect database")
		fmt.Println(err)
	}
	//获取id
	id := c.Params.ByName("id")
	//根据获取的id进行删除
	d := db.Where("id=?", id).Delete(&us)
	defer db.Close()
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})

}
