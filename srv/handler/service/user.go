package service

import (
	"context"
	"fmt"
	"yuekao/srv/basic/config"
	"yuekao/srv/handler/model"

	__ "yuekao/srv/basic/proto"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedUserServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) Login(_ context.Context, in *__.LoginReq) (*__.LoginResp, error) {

	var user model.User
	err := user.FindUser(config.DB, in.Name)
	if err != nil {
		return &__.LoginResp{
			Msg:  "用户不存在",
			Code: 400,
		}, nil
	}
	if user.Password != in.Password {
		return &__.LoginResp{
			Msg:  "密码错误",
			Code: 400,
		}, nil
	}
	return &__.LoginResp{
		Msg:  "登录成功",
		Code: 200,
		Id:   int64(user.ID),
	}, nil
}

func (s *Server) GoodsAdd(_ context.Context, in *__.GoodsAddReq) (*__.GoodsAddResp, error) {

	var goods model.Goods
	err := goods.FindGoods(config.DB, in.Name)
	if err != nil {
		return &__.GoodsAddResp{
			Msg:  "商品不存在",
			Code: 400,
		}, nil
	}
	goods.Name = in.Name
	goods.Price = float32(in.Price)
	goods.Num = int(in.Num)
	err = goods.GoodsAdd(config.DB)
	if err != nil {
		return &__.GoodsAddResp{
			Msg:  "商品添加成功",
			Code: 200,
		}, nil
	}
	go func() {
		EsAdd := map[string]interface{}{
			"Name":  goods.Name,
			"Price": goods.Price,
			"Num":   goods.Num,
		}
		_, err = config.Elastic.Index().Index("goods").BodyJson(EsAdd).Do(context.Background())
		if err != nil {
			return
		}
		fmt.Println("es同步失败")
	}()

	return &__.GoodsAddResp{
		Msg:  "商品添加成功",
		Code: 200,
	}, nil
}
