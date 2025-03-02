package services

import (
	"context"
	"os"
	"time"

	"github.com/adamkali/fullstack_app/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type CustomJwt struct {
	UserId               uuid.UUID `json:"user_id"`
	ProfilePic           string    `json:"profile_pic"`
	jwt.RegisteredClaims `json:"claims"`
}

type AuthService struct {
	ctx  context.Context
	conn *pgx.Conn
}

func newExpiration() time.Time { return time.Now().Add(time.Hour * 72) }

func jwtFromUser(user *repository.User) *CustomJwt {
	return &CustomJwt{
		user.ID,
		*user.ProfilePicUrl,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(newExpiration()),
		},
	}
}

func CreateAuthService(ctx context.Context, conn *pgx.Conn) *AuthService {
	return &AuthService{ctx, conn}
}

func (AuthService *AuthService) Create(user *repository.User) (*string, error) {
	jwttoken := jwtFromUser(user)
	tx, err := AuthService.conn.Begin(AuthService.ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(AuthService.ctx)
	expiration := jwttoken.ExpiresAt

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodPS512.SigningMethodRSA, jwttoken)
	// Sign it with the server JWT_TOKEN
	t, err := token.SignedString([]byte(os.Getenv("JWT_TOKEN")))
	if err != nil {
		return nil, err
	}

	params := repository.CreateTokenParams{
        UserID:             user.ID,
		ExpirationDatetime: &expiration.Time,
		Token:              t,
	}
	repo := repository.New(tx)
	row, err := repo.CreateToken(AuthService.ctx, params)
	if err != nil {
		return nil, err
	}

	return &row.Token, nil
}

func (AuthService *AuthService) CheckToken(token string) error {
	tx, err := AuthService.conn.Begin(AuthService.ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(AuthService.ctx)
	repo := repository.New(tx)
	_, err = repo.FindTokenByToken(AuthService.ctx, token)
	if err != nil {
		return err
	}
	return nil
}

func (AuthService *AuthService) Update(user repository.User) (*string, error) {
	jwttoken := jwtFromUser(&user)
	tx, err := AuthService.conn.Begin(AuthService.ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(AuthService.ctx)
	expiration := jwttoken.ExpiresAt

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodPS512.SigningMethodRSA, jwttoken)
	// Sign it with the server JWT_TOKEN
	t, err := token.SignedString([]byte(os.Getenv("JWT_TOKEN")))
	if err != nil {
		return nil, err
	}

	params := repository.UpdateTokenByUserIdParams{
		UserID:             user.ID,
		ExpirationDatetime: &expiration.Time,
		Token:              t,
	}
	repo := repository.New(tx)
	err = repo.UpdateTokenByUserId(AuthService.ctx, params)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
