package dborm

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/liujun5885/book_store_gql/graph/model"
	"github.com/liujun5885/book_store_gql/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type User struct {
	DB *gorm.DB
}

func (u *User) FetchUserByID(userID *string) (*model.User, error) {
	if userID == nil || *userID == "" {
		return nil, nil
	}
	user := model.User{}
	err := u.DB.Where("id = ?", *userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) FetchUserByEmail(email *string) (*model.User, error) {
	if email == nil || *email == "" {
		return nil, nil
	}
	user := model.User{}
	err := u.DB.Where("email = ?", strings.ToLower(*email)).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) FetchUserByPhoneNumber(phoneNumber *string) (*model.User, error) {
	if phoneNumber == nil || *phoneNumber == "" {
		return nil, nil
	}
	user := model.User{}
	err := u.DB.Where("phone_number = ?", strings.ToLower(*phoneNumber)).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) CreateUser(email, password, phoneNumber, firstName, lastName *string) (*model.User, error) {
	if email == nil || *email == "" {
		return nil, errors.New("email cannot be empty")
	}
	if password == nil || *password == "" {
		return nil, errors.New("password cannot be empty")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	user := &model.User{
		ID:       uuid.New().String(),
		Email:    *email,
		Password: string(hashedPassword),
	}
	if phoneNumber != nil {
		user.PhoneNumber = *phoneNumber
	}
	if firstName != nil {
		user.FirstName = *firstName
	}
	if lastName != nil {
		user.LastName = *lastName
	}
	err := u.DB.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) Login(email, password string) (*model.User, error) {
	user := &model.User{}
	err := u.DB.Where("email = ?", email).First(user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New(fmt.Sprintf("user %s does not exist", email))
	} else if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New(fmt.Sprintf("invalid password"))
	}
	return user, nil
}

func (u *User) UpdateUser(user *model.User) (*model.User, error) {
	if user.ID == "" {
		return nil, nil
	}
	columns := utils.GetFieldsWithValue(user)
	err := u.DB.Model(&user).Select(columns).Updates(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) UpdateUserProfiles(profile *model.UserProfile) (*model.UserProfile, error) {
	if profile.UserID == "" {
		return nil, nil
	}
	columns := utils.GetFieldsWithValue(profile)
	err := u.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns(columns),
	}).Select(columns).Create(profile).Error

	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (u *User) FetchProfilesByUserID(userID *string) (*model.UserProfile, error) {
	if userID == nil || *userID == "" {
		return nil, nil
	}
	profile := model.UserProfile{}
	err := u.DB.Where("user_id = ?", *userID).First(&profile).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &profile, nil
}

func (u *User) UpdateUserSettings(settings *model.UserSettings) (*model.UserSettings, error) {
	if settings.UserID == "" {
		return nil, nil
	}
	columns := utils.GetFieldsWithValue(settings)
	err := u.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns(columns),
	}).Select(columns).Create(settings).Error

	if err != nil {
		return nil, err
	}
	return settings, nil
}

func (u *User) FetchSettingsByUserID(userID *string) (*model.UserSettings, error) {
	if userID == nil || *userID == "" {
		return nil, nil
	}
	settings := model.UserSettings{}
	err := u.DB.Where("user_id = ?", *userID).First(&settings).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &settings, nil
}
