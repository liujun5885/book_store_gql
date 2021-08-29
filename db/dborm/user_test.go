package dborm

import (
	"github.com/liujun5885/book_store_gql/graph/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func TestCreateUserAndFetchUserByEmail(t *testing.T) {
	dsn := "host=localhost port=5432 user=book_store password=aaaaa dbname=book_store_ugc sslmode=disable TimeZone=Asia/Shanghai"
	dbConn, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	user := User{
		DB: dbConn,
	}
	email := "test@test.com"
	password := "password"
	phoneNumber := ""
	firstName := ""
	lastName := ""

	userModel1, err := user.CreateUser(&email, &password, &phoneNumber, &firstName, &lastName)
	if err != nil {
		t.Error(err)
		return
	}
	userModel2, err := user.FetchUserByEmail(&email)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, userModel1.ID, userModel2.ID)

	userModel2.FirstName = "updated first name"
	userModel3, err := user.UpdateUser(userModel2)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, userModel3.FirstName, userModel2.FirstName)
}

func TestUpdateAndFetchProfilesByUserID(t *testing.T) {
	dsn := "host=localhost port=5432 user=book_store password=aaaaa dbname=book_store_ugc sslmode=disable TimeZone=Asia/Shanghai"
	dbConn, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	user := User{
		DB: dbConn,
	}
	userID := "5dc6927f-56a1-4132-9e48-e9c6b1ee2f49"
	addr := "test Addr"
	city := "Beijing"
	country := "China"
	profiles := model.UserProfile{
		UserID:   userID,
		Address:  &addr,
		City:     &city,
		Province: nil,
		Country:  &country,
		Job:      nil,
		School:   nil,
	}
	ret, err := user.UpdateUserProfiles(&profiles)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, country, *ret.Country)

	profile, err := user.FetchProfilesByUserID(&userID)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, &country, profile.Country)
}

func TestUpdateAndFetchUserSettingsByUserID(t *testing.T) {
	dsn := "host=localhost port=5432 user=book_store password=aaaaa dbname=book_store_ugc sslmode=disable TimeZone=Asia/Shanghai"
	dbConn, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	user := User{
		DB: dbConn,
	}
	userID := "5dc6927f-56a1-4132-9e48-e9c6b1ee2f49"
	email := "a@a.com"
	settings := &model.UserSettings{
		UserID:        userID,
		KindleAccount: &email,
	}
	ret, err := user.UpdateUserSettings(settings)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, email, *ret.KindleAccount)

	settings, err = user.FetchSettingsByUserID(&userID)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, email, *settings.KindleAccount)
}
