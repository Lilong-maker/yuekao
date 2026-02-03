package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"yuekao/bff/basic/config"
	"yuekao/bff/handler/request"
	"yuekao/bff/pkg"
	__ "yuekao/srv/basic/proto"
	"yuekao/srv/handler/model"

	"github.com/gin-gonic/gin"
	"github.com/gospacex/gospacex/core/storage/cache/redis"
)

func Login(c *gin.Context) {
	var form request.Login
	// 根据 Content-Type Header 推断使用哪个绑定器。
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	r, err := config.UserClient.Login(c, &__.LoginReq{
		Name:     form.Name,
		Password: form.Password,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	handler, err := pkg.TokenHandler(strconv.FormatInt(r.Id, 10))
	c.JSON(http.StatusBadRequest, gin.H{
		"code": r.Code,
		"msg":  r.Msg,
		"data": handler,
	})
	return
}
func GetGoods(c *gin.Context) {
	var from request.GetGoods
	err := c.ShouldBind(&from)
	if err != nil {
		c.JSON(400, gin.H{
			"msg":  "参数有误",
			"code": 500,
		})
		return
	}
	cacheKey := fmt.Sprintf("goods:%s", from.Id)
	cachedData, err := redis.RC.Get(context.Background(), cacheKey).Result()
	if err == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"data": json.RawMessage(cachedData),
		})
		return
	}
	var findGoods model.Goods
	goods := findGoods.FindGoodsId(config.DB, from.Id)
	marshal, _ := json.Marshal(goods)
	err = redis.RC.SetEx(context.Background(), cacheKey, marshal, 10*time.Minute).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "缓存存储失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": goods,
	})
}
func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		return
	}
	uploads, err := pkg.Uploads(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "上传失败",
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"code": 200,
		"msg":  "上传成功",
		"data": uploads,
	})
	return
}
func GoodsAdd(c *gin.Context) {
	var form request.GoodsAdd
	// 根据 Content-Type Header 推断使用哪个绑定器。
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	uid, _ := c.Get("userId")
	userId := uid.(string)
	sprintf := fmt.Sprintf("idempotent:Goods_add:%s:%s", userId, form.Name)
	lock, err := redis.RC.SetNX(context.Background(), sprintf, "1", 5*time.Minute).Result()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"code": 409,
		})
		return
	}
	if !lock {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "请求已处理，请勿重复提交",
			"code": 409,
		})
		return
	}
	r, err := config.UserClient.GoodsAdd(c, &__.GoodsAddReq{
		Name:  form.Name,
		Price: uint32(form.Price),
		Num:   int64(form.Num),
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"code": r.Code,
		"msg":  r.Msg,
	})
	return
}
func Sx(c *gin.Context) {
	header := c.GetHeader("token")
	if header == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "请登录",
			"code": 400,
		})
		return
	}
	xin, err := pkg.ShuXin(header)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"code": 400,
			"data": xin,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "token更新成功",
		"code": 200,
		"data": header,
	})
	return
}
