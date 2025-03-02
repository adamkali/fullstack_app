package controllers

import (
	"errors"

	"github.com/adamkali/fullstack_app/responses"
	"github.com/adamkali/fullstack_app/services"
	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Name             string
	AuthService      *services.AuthService
	UserService      *services.UserService
	ValidatorService *services.ValidatorService
}

func BuildUserController(p *ControllerParams) UserController {
	return UserController{
		Name:             "/api/users",
		AuthService:      services.CreateAuthService(*p.CTX, p.DB),
		UserService:      services.CreateUserService(*p.CTX, p.DB),
		ValidatorService: &services.ValidatorService{},
	}
}

// @Summary Delete User by their UUID
// @Description get string by ID
//
// @ID          DeleteUserByUUID
// @Tags        Users
// @Produce     json
// @Param       id                  path         string                         true "User Id"          default("e38e78a4-2ca3-4c59-a3ea-a2019866e593")
// @Param       Authorization       header       string                         true "admin header"     default("Bearer token")
// @Success     200                 {object}     responses.DeleteUserResponse
// @Router      /users/{user_id}    [delete]
func (UserController UserController) DeleteUser(ctx echo.Context) error {
	jwt_token := ctx.Get("user").(services.CustomJwt)
	user_id := jwt_token.UserId
	admin_user, err := UserController.UserService.Get(user_id)
	if err != nil {
		return responses.NewDeleteUserResponse().Fail(ctx, 500, err)
	}
	if !admin_user.Admin {
		return responses.NewDeleteUserResponse().Fail(ctx, 403, errors.New("Not Admin"))
	}
	delete_user_id := ctx.Param("user_id")
	delete_user_id_parsed, err := uuid.Parse(delete_user_id)
	if err := UserController.UserService.Remove(delete_user_id_parsed); err != nil {
		return responses.NewDeleteUserResponse().Fail(ctx, 500, err)
	}
	return responses.NewDeleteUserResponse().Successful(ctx, delete_user_id_parsed)
}

// @Summary Signup to the app
// @Description Signup using the requests.NewUserRequest
//
// @ID          Signup
// @Tags        Users
// @Accept      json
// @Produce     json
// @Param       SignupRequest   body        NewUserRequest          true "Signup Request"
// @Success     200             {object}    responses.LoginResponse
// @Failure     400             {object}    responses.LoginResponse
// @Failure     500             {object}    responses.LoginResponse
// @Router      /users/signup   [post]
func (UserController *UserController) Signup(ctx echo.Context) error {
	signupRequest, err := UserController.ValidatorService.ValidateNewUserRequest(ctx)
	if err != nil {
		return responses.NewLoginResponse().Fail(ctx, 400, err)
	}
	user, err := UserController.UserService.Create(signupRequest)
	if err != nil {
		return responses.NewLoginResponse().Fail(ctx, 500, err)
	}
	token, err := UserController.AuthService.Create(user)
	if err != nil {
		return responses.NewLoginResponse().Fail(ctx, 500, err)
	}
	return responses.NewLoginResponse().Successful(ctx, user, *token)
}

// @Summary Login
// @Description to a user account with either email or username
//
// @ID          Login
// @Tags        Users
// @Accept      json
// @Produce     json
// @Param       SignupRequest   body        LoginRequest            true "Signup Request"
// @Success     200             {object}    responses.LoginResponse
// @Failure     400             {object}    responses.LoginResponse
// @Failure     401             {object}    responses.LoginResponse
// @Failure     500             {object}    responses.LoginResponse
// @Router      /users/login    [post]
func (UserController *UserController) Login(ctx echo.Context) error {
	loginRequest, err := UserController.ValidatorService.ValidateLoginRequest(ctx)
	if err != nil {
		return responses.NewLoginResponse().Fail(ctx, 400, err)
	}
	user, err := UserController.UserService.Login(loginRequest)
	if err != nil {
		return responses.NewLoginResponse().Fail(ctx, 401, err)
	}
	token, err := UserController.AuthService.Update(*user)
	if err != nil {
		return responses.NewLoginResponse().Fail(ctx, 500, err)
	}
	return responses.NewLoginResponse().Successful(ctx, user, *token)
}

// @Summary Get Current User
// @Description Get the Current User by the uuid storred in the Claims header
//
// @ID          GetCurrent
// @Tags        Users
// @Produce     json
// @Param       Authorization   header       string                         true "admin header"     default("Bearer token")
// @Success     200             {object}     responses.UserResponse
// @Failure     400             {object}     responses.UserResponse
// @Failure     401             {object}     responses.UserResponse
// @Router      /users/current  [get]
func (UserController *UserController) GetCurrent(ctx echo.Context) error {
	user := ctx.Get("user").(services.CustomJwt)
	claims := user.RegisteredClaims
	err := UserController.AuthService.CheckToken(claims.ID)
	if err != nil {
		return responses.NewUserResponse().Fail(ctx, 401, err)
	}
	user_data, err := UserController.UserService.Get(user.UserId)
	if err != nil {
		return responses.NewUserResponse().Fail(ctx, 404, err)
	}
	return responses.NewUserResponse().Successful(ctx, user_data)
}

func (uc UserController) Attatch(e *echo.Echo) {
	g := e.Group(uc.Name)
	g.GET("/current", uc.GetCurrent)
	g.POST("/login", uc.Login)
	g.POST("/signup", uc.Signup)
	g.DELETE("/:user_id", uc.DeleteUser)
}
