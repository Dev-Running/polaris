package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/gobuffalo/envy"
	"github.com/golang-jwt/jwt/v4"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
	"github.com/laurentino14/user/prisma/connect"
	"github.com/laurentino14/user/repositories/utils"
)

type IAuthRepository interface {
	Auth(input *model.AuthenticationInput, ctx context.Context) (*model.User, error)
	GenerateToken(id string) (string, error)
	IsValid(t string) bool
	GetUserAuthenticated(input *model.GetUserAuthInput, ctx context.Context) (*model.UserAuthenticated, error)
}

type AuthRepository struct {
	DB     *connect.DB
	Secret string
	Issuer string
}

func NewAuthRepository(db *connect.DB) *AuthRepository {
	return &AuthRepository{
		DB:     db,
		Secret: envy.Get("SECRET", ""),
		Issuer: "api-user",
	}
}

type Input struct {
	Token    *string
	Email    *string
	Password *string
}

func (r *AuthRepository) Auth(input *model.AuthenticationInput, ctx context.Context) (*model.User, error) {
	if input.Token != nil {
		t := r.IsValid(*input.Token)
		if !t {
			return nil, fmt.Errorf("token invalido ou expirado, faca login novamente")
		}
		exec, err := r.DB.Client.User.FindFirst(prisma.User.TokenUser.Equals(*input.Token)).With(prisma.User.Enrollment.Fetch()).Exec(ctx)
		if err != nil {
			return nil, fmt.Errorf("usuário não encontrado, faca login novamente")
		}

		refreshToken, err := r.GenerateToken(exec.ID)
		if err != nil {
			return nil, fmt.Errorf("erro na geracao do token")
		}

		_, err = r.DB.Client.User.FindMany(prisma.User.ID.Equals(exec.ID)).Update(
			prisma.User.TokenUser.Set(refreshToken),
		).Exec(ctx)
		if err != nil {
			return nil, fmt.Errorf("erro na validacao do token")
		}

		enrollments := []*model.Enrollment{}
		for _, enr := range exec.RelationsUser.Enrollment {
			enrollments = append(enrollments, &model.Enrollment{
				ID:        enr.ID,
				CourseID:  enr.CourseID,
				CreatedAt: enr.CreatedAt.String(),
				UpdatedAt: utils.ExtractData(enr.UpdatedAt),
				DeletedAt: utils.ExtractData(enr.DeletedAt),
				UserID:    enr.UserID,
			})
		}

		user := &model.User{
			ID:         exec.ID,
			Firstname:  exec.Firstname,
			Lastname:   exec.Lastname,
			Role:       model.Role(exec.Role),
			Avatar:     exec.Avatar,
			Platform:   model.Platform(exec.Platform),
			Github:     utils.ExtractString(exec.Github),
			Bio:        utils.ExtractString(exec.Bio),
			Location:   utils.ExtractString(exec.Location),
			Twitter:    utils.ExtractString(exec.Twitter),
			Site:       utils.ExtractString(exec.Site),
			Username:   exec.Username,
			Email:      exec.Email,
			Password:   utils.ExtractString(exec.Password),
			TokenUser:  refreshToken,
			Enrollment: enrollments,
		}

		return user, nil

	}

	if input.Password != nil && input.Email != nil && input.Token == nil {
		exec, err := r.DB.Client.User.FindUnique(prisma.User.Email.Equals(*input.Email)).With(prisma.User.Enrollment.Fetch()).Exec(ctx)
		if err != nil {
			return nil, fmt.Errorf("erro de conexão com o banco de dados")
		}

		refreshToken, err := r.GenerateToken(exec.ID)
		if err != nil {
			return nil, fmt.Errorf("erro na validacao do token")
		}

		_, err = r.DB.Client.User.FindMany(prisma.User.ID.Equals(exec.ID)).Update(prisma.User.TokenUser.Set(refreshToken)).Exec(ctx)
		if err != nil {
			return nil, fmt.Errorf("erro na validacao do token")
		}

		enrollments := []*model.Enrollment{}
		for _, enr := range exec.RelationsUser.Enrollment {
			enrollments = append(enrollments, &model.Enrollment{
				ID:        enr.ID,
				CourseID:  enr.CourseID,
				CreatedAt: enr.CreatedAt.String(),
				UpdatedAt: utils.ExtractData(enr.UpdatedAt),
				DeletedAt: utils.ExtractData(enr.DeletedAt),
				UserID:    enr.UserID,
			})
		}

		user := &model.User{
			ID:         exec.ID,
			Firstname:  exec.Firstname,
			Lastname:   exec.Lastname,
			Email:      exec.Email,
			Password:   utils.ExtractString(exec.Password),
			TokenUser:  refreshToken,
			Enrollment: enrollments,
		}

		return user, nil

	}

	return nil, fmt.Errorf("dados inválidos")
}

func (r *AuthRepository) GenerateToken(id string) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:   r.Issuer,
		Subject:  "",
		Audience: nil,
		ExpiresAt: &jwt.NumericDate{
			Time: time.Now().Add(time.Hour * 48),
		},
		NotBefore: &jwt.NumericDate{},
		IssuedAt:  &jwt.NumericDate{},
		ID:        id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.Claims(claims))
	tokenString, err := token.SignedString([]byte(r.Secret))
	if err != nil {
		fmt.Println(tokenString)
		fmt.Println(err)
		return tokenString, fmt.Errorf("error in generating key")
	}

	return tokenString, nil
}

func (r *AuthRepository) IsValid(t string) bool {

	_, err := jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("token Invalid")
		}
		return []byte(r.Secret), nil
	})
	return err == nil
}

type UserAuthenticated struct {
	Id        string
	Firstname string
	Lastname  string
	Email     string
	Cellphone string
	TokenUser string
}

func (r *AuthRepository) GetUserAuthenticated(input *model.GetUserAuthInput, ctx context.Context) (*model.UserAuthenticated, error) {
	user, err := r.DB.Client.User.FindFirst(prisma.User.TokenUser.Equals(*input.Token)).With(prisma.User.Enrollment.Fetch()).Exec(ctx)

	if err != nil {
		return nil, err
	}

	enrolls := []*model.Enrollment{}

	for _, enr := range user.RelationsUser.Enrollment {
		enrolls = append(enrolls, &model.Enrollment{
			ID:        enr.ID,
			UserID:    enr.UserID,
			CreatedAt: enr.CreatedAt.String(),
			UpdatedAt: utils.ExtractData(enr.UpdatedAt),
			CourseID:  enr.CourseID,
			DeletedAt: utils.ExtractData(enr.DeletedAt),
		})

	}
	userData := &model.UserAuthenticated{
		ID:         user.ID,
		Firstname:  user.Firstname,
		Lastname:   user.Lastname,
		Email:      user.Email,
		Avatar:     user.Avatar,
		Github:     utils.ExtractString(user.Github),
		Bio:        utils.ExtractString(user.Bio),
		Location:   utils.ExtractString(user.Location),
		Twitter:    utils.ExtractString(user.Twitter),
		Site:       utils.ExtractString(user.Site),
		Username:   user.Username,
		TokenUser:  user.TokenUser,
		Role:       model.Role(user.Role),
		Enrollment: enrolls,
	}
	return userData, nil
}
