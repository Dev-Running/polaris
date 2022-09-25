package repositories

import (
	"context"
	"fmt"
	"github.com/gobuffalo/envy"
	"github.com/golang-jwt/jwt/v4"
	"github.com/laurentino14/user/graph/model"
	"github.com/laurentino14/user/prisma"
	"github.com/laurentino14/user/prisma/connect"
	"time"
)

type IAuthRepository interface {
	Auth(input *model.AuthenticationInput, ctx context.Context) (*model.User, error)
	GenerateToken(id string) (string, error)
	IsValid(t string) bool
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
		exec, err := r.DB.Client.User.FindFirst(prisma.User.TokenUser.Equals(*input.Token)).Exec(ctx)
		if err != nil {
			return nil, fmt.Errorf("usuário não encontrado, faca login novamente")
		}

		refreshToken, err := r.GenerateToken(exec.ID)
		if err != nil {
			return nil, fmt.Errorf("erro na validacao do token")
		}

		_, err = r.DB.Client.User.FindMany(prisma.User.ID.Equals(exec.ID)).Update(
			prisma.User.TokenUser.Set(refreshToken),
		).Exec(ctx)
		if err != nil {
			fmt.Errorf("erro na validacao do token")
			return nil, err
		}

		user := &model.User{
			ID:         exec.ID,
			Firstname:  exec.Firstname,
			Lastname:   exec.Lastname,
			Email:      exec.Email,
			Password:   exec.Password,
			Cellphone:  exec.Cellphone,
			TokenUser:  &refreshToken,
			Enrollment: nil,
		}

		return user, nil

	}

	if input.Password != nil && input.Password != nil && input.Token == nil {
		exec, err := r.DB.Client.User.FindUnique(prisma.User.Email.Equals(*input.Email)).Exec(ctx)
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
		user := &model.User{
			ID:         exec.ID,
			Firstname:  exec.Firstname,
			Lastname:   exec.Lastname,
			Email:      exec.Email,
			Password:   exec.Password,
			Cellphone:  exec.Cellphone,
			TokenUser:  &refreshToken,
			Enrollment: nil,
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
			time.Now().Add(time.Hour * 4),
		},
		NotBefore: &jwt.NumericDate{time.Time{}},
		IssuedAt:  &jwt.NumericDate{time.Time{}},
		ID:        id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.Claims(claims))
	tokenString, err := token.SignedString([]byte(r.Secret))
	if err != nil {
		fmt.Println(tokenString)
		fmt.Println(err)
		return tokenString, fmt.Errorf("Error in generating key")
	}

	return tokenString, nil
}

func (r *AuthRepository) IsValid(t string) bool {

	_, err := jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Token Invalid")
		}
		return []byte(r.Secret), nil
	})
	return err == nil
}