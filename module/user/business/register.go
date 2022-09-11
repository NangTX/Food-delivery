package userbiz

import (
	"context"
	"errors"
	"fmt"
	"food-delivery/common"
	usermodel "food-delivery/module/user/model"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(context context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBiz(registerStorage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{registerStorage: registerStorage, hasher: hasher}
}

func (biz *registerBusiness) RegisterUser(ctx context.Context, data *usermodel.UserCreate) error {

	userFound, _ := biz.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if userFound != nil {
		return usermodel.ErrEmailExisted
	}

	if data.Email == "" {
		return errors.New("email is not empty")
	}

	salt := common.GenSalt(50)

	fmt.Print("salt: ", salt)
	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user" // hard code

	if err := biz.registerStorage.CreateUser(ctx, data); err != nil {
		return err
	}
	return nil
}
