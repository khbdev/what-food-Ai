package usecase

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"auth-service/internal/domain"
	"auth-service/internal/models"
	"auth-service/pkg/jwt"
	"auth-service/pkg/otp"
	rabbitmq "auth-service/pkg/rabbitMq"
	"auth-service/pkg/redis"

	userrpb "github.com/khbdev/what-food-proto/proto/userr"
)

type AuthUsecase struct {
	userClient domain.UserService
	redis      *redis.Service
	rabbit     *rabbitmq.Publisher
	phoneCache *redis.PhoneCache

	ctx context.Context
}

func NewAuthUsecase(
	u domain.UserService,
	r *redis.Service,
	mq *rabbitmq.Publisher,
	p *redis.PhoneCache,
) *AuthUsecase {

	ctx, _ := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)

	return &AuthUsecase{
		userClient: u,
		redis:      r,
		rabbit:     mq,
		phoneCache: p,
		ctx:        ctx,
	}
}
//////////////////////////////////////////////////////
// REGISTER
//////////////////////////////////////////////////////

func (uc *AuthUsecase) Register(req models.RegisterRequest) error {

	// redis check
	exists, err := uc.phoneCache.Get(uc.ctx, req.Phone)
	if err != nil {
		return err
	}

	// agar otp allaqachon yuborilgan bo‘lsa
	if exists {
		return errors.New("otp already sent")
	}

	res, _ := uc.userClient.GetUserByPhone(&userrpb.GetUserByPhoneRequest{
		Phone: req.Phone,
	})

	if res != nil && res.User != nil {
		return errors.New("user already exists")
	}

	// phone redisga save
	if err := uc.phoneCache.Set(uc.ctx,req.Phone); err != nil {
		return err
	}

	code := otp.GenerateCode()
	fmt.Println("OTP:", code)

	if err := uc.redis.SetOTP(int64(code), req); err != nil {
		return err
	}

	if err := uc.rabbit.PublishAuthMessage(
		os.Getenv("AUTH_ROUTING_KEY"),
		models.AuthMessage{
			Task:  "register",
			Phone: req.Phone,
			OTP:   strconv.Itoa(code),
		},
	); err != nil {
		return err
	}

	return nil
}

//////////////////////////////////////////////////////
// LOGIN
//////////////////////////////////////////////////////

func (uc *AuthUsecase) Login(req models.LoginRequest) error {

	// redis check
	exists, err := uc.phoneCache.Get(uc.ctx,req.Phone)
	if err != nil {
		return err
	}

	// agar otp allaqachon yuborilgan bo‘lsa
	if exists {
		return errors.New("otp already sent")
	}

	res, _ := uc.userClient.GetUserByPhone(&userrpb.GetUserByPhoneRequest{
		Phone: req.Phone,
	})

	if res == nil || res.User == nil {
		return errors.New("user not found")
	}

	// phone redisga save
	if err := uc.phoneCache.Set(uc.ctx, req.Phone); err != nil {
		return err
	}

	code := otp.GenerateCode()

	data := models.RegisterRequest{
		Phone: req.Phone,
	}

	if err := uc.redis.SetOTP(int64(code), data); err != nil {
		return err
	}

	if err := uc.rabbit.PublishAuthMessage(
		os.Getenv("AUTH_ROUTING_KEY"),
		models.AuthMessage{
			Task:  "login",
			Phone: req.Phone,
			OTP:   strconv.Itoa(code),
		},
	); err != nil {
		return err
	}

	return nil
}
//////////////////////////////////////////////////////
// VERIFY
//////////////////////////////////////////////////////
func (uc *AuthUsecase) Verify(code int64) (string, string, error) {

	data, err := uc.redis.GetOTP(code)
	if err != nil {
		return "", "", err
	}
	if data == nil {
		return "", "", errors.New("otp expired or invalid")
	}

	var (
		userID   uint
		userName string
		roleStr  string
	)

	// check user
	res, _ := uc.userClient.GetUserByPhone(&userrpb.GetUserByPhoneRequest{
		Phone: data.Phone,
	})

	if res != nil && res.User != nil {

		// LOGIN CASE
		userID = uint(res.User.Id)
		userName = res.User.Name

		// enum -> string convert
		switch res.User.Role {
		case userrpb.Role_ROLE_ADMIN:
			roleStr = "admin"
		case userrpb.Role_ROLE_USER:
			roleStr = "user"
		default:
			roleStr = "user"
		}

	} else {

		// REGISTER CASE
		userRes, err := uc.userClient.CreateUser(&userrpb.CreateUserRequest{
			Name:    data.FullName,
			Phone:   data.Phone,
			Age:     int32(data.Age),
			Address: data.Address,
		})
		if err != nil {
			return "", "", err
		}
		if userRes == nil || userRes.User == nil {
			return "", "", errors.New("user yaratishda xato")
		}

		userID = uint(userRes.User.Id)
		userName = userRes.User.Name

		switch userRes.User.Role {
		case userrpb.Role_ROLE_ADMIN:
			roleStr = "admin"
		case userrpb.Role_ROLE_USER:
			roleStr = "user"
		default:
			roleStr = "user"
		}
	}

	// token model
	tokenModel := models.TokenModel{
		UserID:   userID,
		UserName: userName,
		Role:     roleStr,
	}

	access, err := jwt.GenerateAccessToken(tokenModel)
	if err != nil {
		return "", "", err
	}

	refresh, err := jwt.GenerateRefreshToken(tokenModel)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}