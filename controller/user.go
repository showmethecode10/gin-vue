package controller

import (
	"fmt"
	"gin-vue/dao"
	"gin-vue/dto"
	"gin-vue/model"
	"gin-vue/response"
	"gin-vue/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	//获取参数
	username := c.PostForm("username")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	fmt.Println("用户信息为：", username, password, phone)

	//数据校验
	//手机号校验
	if len(phone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号要不少于11位！")
		fmt.Println("电话号码少于11位，", len(phone))
		fmt.Println(phone)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "the phone number must be 11 digits!",
		})
		return
	}
	//密码校验
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码长度少于6位")
		fmt.Println("密码长度小于6位，", len(password))
		fmt.Println(password)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "password cannot be less than 6 digits!",
		})
		return
	}
	//用户名校验
	if len(username) == 0 {
		fmt.Println("用户名为空，", len(username))
		fmt.Println(username)
		username = tools.RandomString(10)
	}
	log.Println(username, password, phone)
	//手机号是否已注册的校验
	if tools.IsPhoneExist(phone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号已被注册！")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "该手机号已注册!",
		})
		return
	}

	makePassword := tools.MakePassword(password)
	fmt.Println("加密之后的密码：", makePassword)

	//创建用户
	user := model.User{
		Model:    gorm.Model{},
		Username: username,
		Password: makePassword,
		Phone:    phone,
	}
	dao.Mgr.Register(&user)

	//返回结果

	c.JSON(http.StatusOK, gin.H{
		"msg":  "register success!",
		"user": user,
	})
	response.Success(c, nil, "注册成功！")
}

func Login(c *gin.Context) {
	//获取参数
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	//数据验证
	if len(phone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号必须为11位！",
		})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能少于6位",
		})
		return
	}
	//判断手机号是否存在
	if dao.Mgr.GetUserByPhone(phone).Phone == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户不存在!",
		})
		return
	}
	//判断密码是否正确
	if tools.MakePassword(password) != dao.Mgr.GetUserByPhone(phone).Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "密码错误！",
		})
		return
	}

	user := dao.Mgr.GetUserByPhone(phone)

	//发放token
	token, err := tools.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "System err!",
		})
		log.Println("token generate err:", err)
		return
	}
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"token": token,
		"msg":   "登录成功！",
	})
	response.Success(c, gin.H{"token": token}, "登录成功！")
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"user": dto.ToUserDto(user.(model.User)),
		},
	})
}
