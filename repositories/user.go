package repositories

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
	"github.com/laurentino14/user/prisma/connect"
	"github.com/laurentino14/user/repositories/utils"
)

type IUserRepository interface {
	Avatar(input model.NewUser, imageType string) string
	Create(input model.NewUser, ctx context.Context) (*model.User, error)
	CreateUserGITHUB(input model.NewUserGithub, ctx context.Context) (*model.User, error)
	CreateUserGOOGLE(input model.NewUserGoogle, ctx context.Context) (*model.User, error)
	GetAll(ctx context.Context) ([]*model.User, error)
	GetUserByID(id string, ctx context.Context) (*model.User, error)
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
	if i.File == nil {
		return "http://localhost:3131/static/default.jpg"
	}
	return "http://localhost:3131/static/default.jpg"
}
func (r *UserRepository) Create(input model.NewUser, ctx context.Context) (*model.User, error) {

	image := ""
	if strings.Contains(input.File.Filename, ".png") {
		image = ".png"
	}

	if strings.Contains(input.File.Filename, ".PNG") {
		image = ".png"
	}

	if strings.Contains(input.File.Filename, ".jpg") {
		image = ".jpg"
	}

	if strings.Contains(input.File.Filename, ".JPG") {
		image = ".jpg"
	}

	if strings.Contains(input.File.Filename, ".jpeg") {
		image = ".jpeg"
	}

	if strings.Contains(input.File.Filename, ".JPEG") {
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
		prisma.User.EmailVerified.Set(false),
		prisma.User.Avatar.Set(r.Avatar(input, image)),
		prisma.User.Email.Set(*input.Email),
		prisma.User.TokenUser.Set(""),
		prisma.User.Platform.Set(prisma.PlatformDR),
		prisma.User.Password.Set(*input.Password),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	a := ""
	slEnroll := []*model.Enrollment{}
	sl := exec.RelationsUser.Enrollment

	for _, enr := range sl {

		slEnroll = append(slEnroll, &model.Enrollment{
			ID:        enr.ID,
			UserID:    enr.UserID,
			CreatedAt: enr.CreatedAt.String(),
			UpdatedAt: utils.ExtractData(enr.UpdatedAt),
			CourseID:  enr.CourseID,
			DeletedAt: utils.ExtractData(enr.DeletedAt),
		})

	}

	userData := &model.User{
		ID:         exec.ID,
		Firstname:  exec.Firstname,
		Lastname:   exec.Lastname,
		Email:      exec.Email,
		Avatar:     exec.Avatar,
		Github:     utils.ExtractString(exec.Github),
		Bio:        utils.ExtractString(exec.Bio),
		Location:   utils.ExtractString(exec.Location),
		Twitter:    utils.ExtractString(exec.Twitter),
		Site:       utils.ExtractString(exec.Site),
		Username:   exec.Username,
		TokenUser:  a,
		Role:       model.Role(exec.Role),
		Enrollment: slEnroll,
	}

	return userData, nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*model.User, error) {
	exec, err := r.DB.Client.User.FindMany().With(prisma.UserRelationWith(prisma.User.Enrollment.Fetch())).Exec(ctx)

	if err != nil {
		return nil, err
	}

	var allUsers []*model.User
	for _, list := range exec {

		slEnroll := []*model.Enrollment{}
		sl := list.RelationsUser.Enrollment

		for _, enr := range sl {

			slEnroll = append(slEnroll, &model.Enrollment{
				ID:        enr.ID,
				UserID:    enr.UserID,
				CreatedAt: enr.CreatedAt.String(),
				UpdatedAt: utils.ExtractData(enr.UpdatedAt),
				CourseID:  enr.CourseID,
				DeletedAt: utils.ExtractData(enr.DeletedAt),
			})

		}
		user := &model.User{
			ID:         list.ID,
			Firstname:  list.Firstname,
			Lastname:   list.Lastname,
			Email:      list.Email,
			Username:   list.Username,
			Avatar:     list.Avatar,
			Platform:   model.Platform(list.Platform),
			Github:     utils.ExtractString(list.Github),
			Bio:        utils.ExtractString(list.Bio),
			Twitter:    utils.ExtractString(list.Twitter),
			Site:       utils.ExtractString(list.Site),
			Location:   utils.ExtractString(list.Location),
			Role:       model.Role(list.Role),
			Password:   utils.ExtractString(list.Password),
			TokenUser:  list.TokenUser,
			Enrollment: slEnroll,
		}

		allUsers = append(allUsers, user)
	}

	return allUsers, nil
}

func (r *UserRepository) CreateUserGITHUB(input model.NewUserGithub, ctx context.Context) (*model.User, error) {
	exec, err := r.DB.Client.User.CreateOne(
		prisma.User.Firstname.Set(*input.Firstname),
		prisma.User.Lastname.Set(*input.Lastname),
		prisma.User.Username.Set(*input.Username),
		prisma.User.EmailVerified.Set(false),
		prisma.User.Avatar.Set(*input.Avatar),
		prisma.User.Email.Set(*input.Email),
		prisma.User.TokenUser.Set(""),
		prisma.User.Platform.Set(prisma.PlatformGITHUB),
		prisma.User.Site.Set(*input.Site),
		prisma.User.Bio.Set(*input.Bio),
		prisma.User.Location.Set(*input.Location),
		prisma.User.Github.Set(*input.Github),
		prisma.User.Twitter.Set(*input.Twitter),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	enrollments := []*model.Enrollment{}
	for _, enr := range exec.RelationsUser.Enrollment {
		enrollments = append(enrollments, &model.Enrollment{
			ID:        enr.ID,
			UserID:    enr.UserID,
			CreatedAt: enr.CreatedAt.String(),
			UpdatedAt: utils.ExtractData(enr.UpdatedAt),
			CourseID:  enr.CourseID,
			DeletedAt: utils.ExtractData(enr.DeletedAt),
		})
	}

	userData := &model.User{
		ID:         exec.ID,
		Firstname:  exec.Firstname,
		Lastname:   exec.Lastname,
		Role:       model.Role(exec.Role),
		Email:      exec.Email,
		Avatar:     exec.Avatar,
		Platform:   model.Platform(exec.Platform),
		Github:     utils.ExtractString(exec.Github),
		Bio:        utils.ExtractString(exec.Bio),
		Twitter:    utils.ExtractString(exec.Twitter),
		Site:       utils.ExtractString(exec.Site),
		Password:   utils.ExtractString(exec.Password),
		Username:   exec.Username,
		Location:   utils.ExtractString(exec.Location),
		TokenUser:  exec.TokenUser,
		Enrollment: enrollments,
	}

	return userData, nil
}

func (r *UserRepository) CreateUserGOOGLE(input model.NewUserGoogle, ctx context.Context) (*model.User, error) {
	exec, err := r.DB.Client.User.CreateOne(
		prisma.User.Firstname.Set(*input.Firstname),
		prisma.User.Lastname.Set(*input.Lastname),
		prisma.User.Username.Set(*input.Username),
		prisma.User.EmailVerified.Set(false),
		prisma.User.Avatar.Set(*input.Avatar),
		prisma.User.Email.Set(*input.Email),
		prisma.User.TokenUser.Set(""),
		prisma.User.Platform.Set(prisma.PlatformGOOGLE),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	enrollments := []*model.Enrollment{}
	for _, enr := range exec.RelationsUser.Enrollment {
		enrollments = append(enrollments, &model.Enrollment{
			ID:        enr.ID,
			UserID:    enr.UserID,
			CreatedAt: enr.CreatedAt.String(),
			UpdatedAt: utils.ExtractData(enr.UpdatedAt),
			CourseID:  enr.CourseID,
			DeletedAt: utils.ExtractData(enr.DeletedAt),
		})
	}

	userData := &model.User{
		ID:         exec.ID,
		Firstname:  exec.Firstname,
		Lastname:   exec.Lastname,
		Role:       model.Role(exec.Role),
		Email:      exec.Email,
		Avatar:     exec.Avatar,
		Platform:   model.Platform(exec.Platform),
		Github:     utils.ExtractString(exec.Github),
		Bio:        utils.ExtractString(exec.Bio),
		Twitter:    utils.ExtractString(exec.Twitter),
		Site:       utils.ExtractString(exec.Site),
		Password:   utils.ExtractString(exec.Password),
		Username:   exec.Username,
		Location:   utils.ExtractString(exec.Location),
		TokenUser:  exec.TokenUser,
		Enrollment: enrollments,
	}
	return userData, nil
}

func (r *UserRepository) GetUserByID(id string, ctx context.Context) (*model.User, error) {
	exec, err := r.DB.Client.User.FindUnique(prisma.User.ID.Equals(id)).Exec(ctx)
	if err != nil {
		return nil, err
	}

	userData := &model.User{
		ID:         exec.ID,
		Firstname:  exec.Firstname,
		Lastname:   exec.Lastname,
		Role:       model.Role(exec.Role),
		Email:      exec.Email,
		Avatar:     exec.Avatar,
		Platform:   model.Platform(exec.Platform),
		Github:     utils.ExtractString(exec.Github),
		Bio:        utils.ExtractString(exec.Bio),
		Twitter:    utils.ExtractString(exec.Twitter),
		Site:       utils.ExtractString(exec.Site),
		Username:   exec.Username,
		Location:   utils.ExtractString(exec.Location),
		TokenUser:  exec.TokenUser,
		Enrollment: nil,
	}

	return userData, nil
}
