package repositories

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/disintegration/imaging"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
	"github.com/laurentino14/user/prisma/connect"
	"io"
	"os"
	"path"
	"strings"
)

type IUserRepository interface {
	Avatar(input model.NewUser, imageType string) string
	Create(input model.NewUser, ctx context.Context, file *graphql.Upload) (*model.User, error)
	GetAll(ctx context.Context) ([]*model.User, error)
}

type UserRepository struct {
	DB *connect.DB
}

func NewUserRepository(db *connect.DB) *UserRepository {
	return &UserRepository{DB: db}
}
func (r *UserRepository) Avatar(i model.NewUser, imageType string) string {
	if i.File.Size != 0 {
		return strings.Join([]string{"http://localhost:3131/static/", *i.Username, imageType}, "")
	}
	return "http://localhost:3131/static/default.jpg"
}
func (r *UserRepository) Create(input model.NewUser, ctx context.Context) (*model.User, error) {

	image := ""
	if strings.Contains(input.File.Filename, ".png") {
		image = ".png"
	}

	if strings.Contains(input.File.Filename, ".jpg") {
		image = ".jpg"
	}

	if strings.Contains(input.File.Filename, ".jpeg") {
		image = ".jpeg"
	}

	verify := r.Avatar(input, image)

	if verify != "http://localhost:3131/static/default.jpg" {
		fileName := path.Base(*input.Username + image)

		dest, _ := os.Create("assets/" + fileName)

		_, err := io.Copy(dest, input.File.File)
		if err != nil {
			panic(err)
		}

		img, err := imaging.Open("assets/" + fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// crop from center
		centercropimg := imaging.CropCenter(img, 650, 650)

		// save cropped image

		err = imaging.Save(centercropimg, "assets/"+fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// crop out a rectangular region

	}

	exec, err := r.DB.Client.User.CreateOne(
		prisma.User.Firstname.Set(*input.Firstname),
		prisma.User.Lastname.Set(*input.Lastname),
		prisma.User.Username.Set(*input.Username),
		prisma.User.Avatar.Set(r.Avatar(input, image)),
		prisma.User.Email.Set(*input.Email),
		prisma.User.Password.Set(*input.Password),
		prisma.User.Cellphone.Set(*input.Cellphone),
		prisma.User.TokenUser.Set(""),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	a := ""
	aa := r.Avatar(input, image)
	userData := &model.User{
		ID:         exec.ID,
		Firstname:  exec.Firstname,
		Lastname:   exec.Lastname,
		Email:      exec.Email,
		Avatar:     &aa,
		Username:   exec.Username,
		Password:   exec.Password,
		Cellphone:  exec.Cellphone,
		TokenUser:  a,
		Enrollment: nil,
	}

	return userData, nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*model.User, error) {
	exec, err := r.DB.Client.User.FindMany().Exec(ctx)

	if err != nil {
		return nil, err
	}

	var allUsers []*model.User
	for _, list := range exec {
		user := &model.User{
			ID:         list.ID,
			Firstname:  list.Firstname,
			Lastname:   list.Lastname,
			Email:      list.Email,
			Username:   list.Username,
			Avatar:     &list.Avatar,
			Password:   list.Password,
			Cellphone:  list.Cellphone,
			TokenUser:  list.TokenUser,
			Enrollment: nil,
		}
		allUsers = append(allUsers, user)
	}

	return allUsers, nil
}
