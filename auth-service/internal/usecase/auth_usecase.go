package usecase

import (
	"errors"
	"strconv"

	"auth-service/internal/domain"
	"auth-service/internal/models"
	rabbitmq "auth-service/pkg/rabbitMq"
	"auth-service/pkg/redis"

	userrpb "github.com/khbdev/what-food-proto/proto/userr"
)

type AuthUsecase struct {
	userClient domain.UserService
	redis      *redis.Service
	rabbit     *rabbitmq.Publisher
}

func NewAuthUsecase(u domain.UserService, r *redis.Service, mq *rabbitmq.Publisher) *AuthUsecase {
	return &AuthUsecase{
		userClient: u,
		redis:      r,
		rabbit:     mq,
	}
}

//////////////////////////////////////////////////////
// REGISTER
//////////////////////////////////////////////////////

func (uc *AuthUsecase) Register(req models.RegisterRequest) error {

	// user mavjudmi
	res, _ := uc.userClient.GetUserByPhone(&userrpb.GetUserByPhoneRequest{
		Phone: req.Phone,
	})
	if res != nil && res.User != nil {
		return errors.New("user already exists")
	}

	// OTP generate
	code := otp.GenerateCode()

	// Redisga saqlash
	if err := uc.redis.SetOTP(int64(code), req); err != nil {
		return err
	}

	// RabbitMQ push
	if err := uc.rabbit.PublishAuthMessage("auth.otp", models.AuthMessage{
		Task:  "register",
		Phone: req.Phone,
		OTP:   strconv.Itoa(code),
	}); err != nil {
		return err
	}

	return nil
}

//////////////////////////////////////////////////////
// LOGIN
//////////////////////////////////////////////////////

func (uc *AuthUsecase) Login(req models.LoginRequest) error {

	// user borligini tekshiramiz
	res, _ := uc.userClient.GetUserByPhone(&userrpb.GetUserByPhoneRequest{
		Phone: req.Phone,
	})
	if res == nil || res.User == nil {
		return errors.New("user not found")
	}

	// OTP
	code := otp.GenerateCode()

	// login uchun ham RegisterRequest ishlatyapmiz (redis struct shu)
	data := models.RegisterRequest{
		Phone: req.Phone,
	}

	if err := uc.redis.SetOTP(int64(code), data); err != nil {
		return err
	}

	if err := uc.rabbit.PublishAuthMessage("auth.otp", models.AuthMessage{
		Task:  "login",
		Phone: req.Phone,
		OTP:   strconv.Itoa(code),
	}); err != nil {
		return err
	}

	return nil
}

//////////////////////////////////////////////////////
// VERIFY
//////////////////////////////////////////////////////

func (uc *AuthUsecase) Verify(code int64) (string, string, error) {

	// Redisdan olish
	data, err := uc.redis.GetOTP(code)
	if err != nil {
		return "", "", err
	}
	if data == nil {
		return "", "", errors.New("otp expired or invalid")
	}

	// user create (har doim create qilamiz — sen aytgandek flow)
	userRes, err := uc.userClient.CreateUser(&userrpb.CreateUserRequest{
		Name:    data.FullName,
		Phone:   data.Phone,
		Age:     int32(data.Age),
		Address: data.Address,
	})
	if err != nil {
		return "", "", err
	}

	// token model
	tokenModel := models.TokenModel{
		UserID:   uint(userRes.User.Id),
		UserName: userRes.User.Name,
		Role:     "user",
	}

	// access
	access, err := jwt.GenerateAccessToken(tokenModel)
	if err != nil {
		return "", "", err
	}

	// refresh
	refresh, err := jwt.GenerateRefreshToken(tokenModel)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}