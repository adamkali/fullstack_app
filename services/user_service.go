package services

import (
	"context"

	"github.com/adamkali/fullstack_app/internal/repository"
	"github.com/adamkali/fullstack_app/requests"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func verifyPassword(storedHash, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(providedPassword))
	if err != nil {
		return false
	}
	return true
}

type UserService struct {
	ctx  context.Context
	conn *pgx.Conn
}

// Returns a refrence to a new UserService to be used in the controller
func CreateUserService(ctx context.Context, conn *pgx.Conn) *UserService {
	return &UserService{ctx, conn}
}

func (UserService *UserService) SignUp(params *requests.NewUserRequest) (*repository.User, error) {
	BCryptHash, err := hashPassword(params.Password)
	if err != nil {
		return nil, err
	}
	if params.IsAdmin {
		return UserService.addNewUserAdmin(repository.CreateUserAdminParams{
			BCryptHash: BCryptHash,
			Username:   params.Username,
			Email:      params.Email,
		})
	} else {
		return UserService.addNewUser(repository.CreateUserParams{
			BCryptHash: BCryptHash,
			Username:   params.Username,
			Email:      params.Email,
		})
	}
}

func (UserService *UserService) addNewUser(
	params repository.CreateUserParams,
) (*repository.User, error) {
	var user repository.User
	tx, err := UserService.conn.Begin(UserService.ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(UserService.ctx)
	repo := repository.New(tx)
	user, err = repo.CreateUser(UserService.ctx, params)
	if err != nil {
		return nil, err
	}
    tx.Commit(UserService.ctx)
	return &user, nil
}

func (UserService *UserService) addNewUserAdmin(
	params repository.CreateUserAdminParams,
) (*repository.User, error) {
	var user repository.User
	tx, err := UserService.conn.Begin(UserService.ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(UserService.ctx)
	repo := repository.New(tx)
	user, err = repo.CreateUserAdmin(UserService.ctx, params)
	if err != nil {
		return nil, err
	}
    tx.Commit(UserService.ctx)
	return &user, nil
}

func (UserService *UserService) Login(params *requests.LoginRequest) (*repository.User, error) {
	var user repository.User
	tx, err := UserService.conn.Begin(UserService.ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(UserService.ctx)
	var BCryptHash string
	repo := repository.New(tx)
	if params.Email != "" {
		BCryptHash, err = repo.FindBCryptHashByEmail(UserService.ctx, params.Email)
		if err != nil {
			return nil, err
		}
		if verifyPassword(BCryptHash, params.Password) {
			user, err = repo.FindUserByEmail(UserService.ctx, params.Email)
			if err != nil {
				return nil, err
			}
		}
	} else if params.Username != "" {
		BCryptHash, err = repo.FindBCryptHashByUsername(UserService.ctx, params.Username)
		if err != nil {
			return nil, err
		}
		if verifyPassword(BCryptHash, params.Password) {
			user, err = repo.FindUserByUsername(UserService.ctx, params.Username)
			if err != nil {
				return nil, err
			}
		}
	}
    tx.Commit(UserService.ctx)
	return &user, nil
}

func (UserService *UserService) RemoveUser(user_id uuid.UUID) error {
	tx, err := UserService.conn.Begin(UserService.ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(UserService.ctx)
	repo := repository.New(tx)
	if err := repo.DeleteUserByID(UserService.ctx, user_id); err != nil {
		return err
	}
    tx.Commit(UserService.ctx)
	return nil
}

func (UserService *UserService) GetUserById(user_id uuid.UUID) (*repository.User, error) {
	tx, err := UserService.conn.Begin(UserService.ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(UserService.ctx)
    var user repository.User
	repo := repository.New(tx)
    if user, err = repo.FindUserByID(UserService.ctx,user_id); err != nil {
        return nil, err
    }
    tx.Commit(UserService.ctx)
    return &user, nil
}
