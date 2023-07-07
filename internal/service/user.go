package service

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"go-bank/domain"
	"go-bank/dto"
	"go-bank/internal/util"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo     domain.UserRepository
	cacheRepo    domain.CacheRepository
	emailService domain.EmailService
}

func NewUser(userRepo domain.UserRepository, cacheRepo domain.CacheRepository, emailService domain.EmailService) domain.UserService {
	return &userService{
		userRepo:     userRepo,
		cacheRepo:    cacheRepo,
		emailService: emailService,
	}
}

func (u userService) Register(ctx context.Context, req dto.RegisterReq) (dto.RegisterRes, error) {
	// Check if username or email already exist
	existingUsername, err := u.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		return dto.RegisterRes{}, domain.ErrInternalServerError
	}
	if existingUsername != (domain.User{}) {
		return dto.RegisterRes{}, domain.ErrUsernameAlreadyExist
	}
	existingEmail, err := u.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return dto.RegisterRes{}, domain.ErrInternalServerError
	}
	if existingEmail != (domain.User{}) {
		return dto.RegisterRes{}, domain.ErrEmailAlreadyExist
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.RegisterRes{}, domain.ErrInternalServerError
	}

	userDomain := domain.User{
		FullName: req.FullName,
		Phone:    req.Phone,
		Email:    req.Email,
		Username: req.Username,
		Password: string(hashedPassword),
	}

	user, err := u.userRepo.Create(ctx, &userDomain)
	if err != nil {
		return dto.RegisterRes{}, domain.ErrInternalServerError
	}

	id, err := u.userRepo.GetLastID(ctx)
	if err != nil {
		return dto.RegisterRes{}, domain.ErrInternalServerError
	}

	RefrenceID := util.GenerateRandomString(16)
	otp := util.GenerateRandomOTPNumber(6)
	log.Println("OTP: ", otp)

	//Send Email
	_ = u.emailService.SendEmailVerification(user.Email, otp)

	//Assign RefrenceID to Cache
	err = u.cacheRepo.Set("user-ref:"+RefrenceID, []byte(user.Username))
	if err != nil {
		return dto.RegisterRes{}, domain.ErrInternalServerError
	}

	//Assign OTP to Cache
	err = u.cacheRepo.Set("otp:"+RefrenceID, []byte(otp))
	if err != nil {
		return dto.RegisterRes{}, domain.ErrInternalServerError
	}

	// //Assign OTP Expired Time to Cache
	// expiredAt := time.Now().Add(5 * time.Minute).Format(time.RFC3339)
	// err = u.cacheRepo.Set("otp:"+RefrenceID+":expire", []byte(expiredAt))
	// if err != nil {
	// 	return dto.RegisterRes{}, domain.ErrInternalServerError
	// }

	return dto.RegisterRes{
		ID:         id,
		FullName:   user.FullName,
		Phone:      user.Phone,
		Email:      user.Email,
		Username:   user.Username,
		RefrenceID: RefrenceID,
	}, nil
}

func (u userService) Authenticate(ctx context.Context, req dto.AuthReq) (dto.AuthRes, error) {
	user, err := u.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		return dto.AuthRes{}, err
	}
	if user == (domain.User{}) {
		return dto.AuthRes{}, domain.ErrAuthFailed
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.AuthRes{}, domain.ErrAuthFailed
	}

	token := util.GenerateRandomString(72)

	userJson, _ := json.Marshal(user)
	_ = u.cacheRepo.Set("users:"+token, userJson)
	return dto.AuthRes{
		Token: token,
	}, nil
}

func (u userService) ValidateToken(ctx context.Context, token string) (dto.UserData, error) {
	data, err := u.cacheRepo.Get("users:" + token)
	if err != nil {
		return dto.UserData{}, domain.ErrAuthFailed
	}

	var user domain.User
	_ = json.Unmarshal(data, &user)

	return dto.UserData{
		ID:       user.ID,
		FullName: user.FullName,
		Phone:    user.Phone,
		Email:    user.Email,
		Username: user.Username,
	}, nil
}

func (u userService) ValidateOTP(ctx context.Context, req dto.ValidateOTPReq) error {
	otp, err := u.cacheRepo.Get("otp:" + req.RefrenceID)
	if err != nil {
		return domain.ErrInternalServerError
	}

	if string(otp) != req.OTP {
		return domain.ErrOTPInvalid
	}

	// expiredAt, err := u.cacheRepo.Get("otp:" + req.RefrenceID + ":expire")
	// if err != nil {
	// 	return domain.ErrInternalServerError
	// }

	// expiredAtTime, err := time.Parse(time.RFC3339, string(expiredAt))
	// if err != nil {
	// 	return domain.ErrInternalServerError
	// }

	// if time.Now().After(expiredAtTime) {
	// 	return domain.ErrOTPExpired
	// }

	val, err := u.cacheRepo.Get("user-ref:" + req.RefrenceID)
	if err != nil {
		return domain.ErrInternalServerError
	}

	username := string(val)
	user, err := u.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return domain.ErrInternalServerError
	}

	user.EmailVerifiedAt = time.Now()
	err = u.userRepo.Update(ctx, &user)
	if err != nil {
		log.Println("error when update user: ", err.Error())
		return domain.ErrInternalServerError
	}

	return nil
}
