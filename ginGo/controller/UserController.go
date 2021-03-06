package controller

import (
	"fmt"
	"ginGo/common"
	"ginGo/dto"
	"ginGo/model"
	"ginGo/response"
	"ginGo/util"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// 查询手机号
func isTelephoneExis(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

// 判断手机号和密码的长度是否正确
func isRight(telephone string,password string,ctx *gin.Context) bool{
	if len(telephone) != 11 {
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"手机号必须为11位")
		return false
	}

	if len(password) < 6 {
		response.Response(ctx,http.StatusUnprocessableEntity,423,nil,"密码不能少于6位")
		return false
	}
	return true
}

// 用户注册
func Register(ctx *gin.Context) {
	DB := common.GetDB()
	var requestUser = model.User{}
	err := ctx.Bind(&requestUser)
	if err !=nil{
		response.Fail(ctx,err.Error(),nil)
		return
	}
	// 获取参数
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	// 数据验证
	isReturn := isRight(telephone,password,ctx)
	if !isReturn{
		return
	}
	// 如果名称没有传给一个随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, password, telephone)
	// 判断手机号是否存在
	if isTelephoneExis(DB, telephone) {
		response.Response(ctx,http.StatusUnprocessableEntity,424,nil,"用户已存在")
		return
	}
	// 用户不存在就创建用户
	hasedPassword,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err !=nil{
		response.Response(ctx,http.StatusInternalServerError,500,nil,"加密发送错误")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)

	// 返回结果
	response.Success(ctx,nil,"注册成功")

}

// 用户登录
func Login(ctx *gin.Context){
	DB := common.GetDB()
	// 获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	isReturn := isRight(telephone,password,ctx)
	if !isReturn{
		return
	}
	// 判断手机号是否存在
	var user model.User
	DB.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx,http.StatusUnprocessableEntity,424,nil,"用户不存在")
		return
	}

	// 判断密码是否正确
	fmt.Println(user.Password,"--")
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err!=nil{
		fmt.Println(err)
		response.Response(ctx,http.StatusBadRequest,400,nil,"密码错误")
		return
	}
	// 发送token
	token,err := common.ReleaseToken(user)
	if err !=nil{
		response.Response(ctx,http.StatusInternalServerError,500,nil,"系统异常")
		fmt.Println(err)
		return
	}
	// 返回结果
	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":gin.H{"token":token},
		"msg":"登录成功",
	})
}

// 用户信息
func Info(ctx *gin.Context){
	user,_ := ctx.Get("user")

	ctx.JSON(http.StatusOK,gin.H{"code":200,"data":gin.H{"user":dto.ToUserDto(user.(model.User))}})
}