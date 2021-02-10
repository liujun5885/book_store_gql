package model

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/liujun5885/book_store_gql/constants"
)

/*
	The following is the DB and Model mapping doc:
	https://pg.uptrace.dev/models/
*/

type BookAuthor struct {
	BookID    string     `json:"book_id"`
	AuthorID  string     `json:"author_id"`
	CreatedAt *time.Time `json:"createdAt"`
}

type BookPublisher struct {
	BookID      string     `json:"book_id"`
	PublisherID string     `json:"publisher_id"`
	CreatedAt   *time.Time `json:"createdAt"`
}

type BookTopic struct {
	BookID    string     `json:"book_id"`
	TopicID   string     `json:"topic_id"`
	CreatedAt *time.Time `json:"createdAt"`
}

type Author struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Nationality string    `json:"nationality"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`

	// customized fields
	//Books       []*Book    `json:"books"`
	BookAuthor *BookAuthor `json:"-" pg:"-"`
}

type Publisher struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Headquarter string    `json:"headquarter"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`

	// customized fields
	//Books       []*Book   `json:"books"`
	BookPublisher *BookPublisher `json:"-" pg:"-"`
}

type Topic struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Score     *int      `json:"score"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	// customized fields
	//Books     []*Book   `json:"books"`
	BookTopic *BookTopic `json:"-" pg:"-"`
}

type Book struct {
	ID                 string    `json:"id"`
	Title              string    `json:"title"`
	Description        string    `json:"description"`
	DescriptionTrimmed string    `json:"descriptionTrimmed" pg:"-"`
	Pages              int       `json:"pages"`
	Language           string    `json:"language"`
	Rating             int       `json:"rating"`
	Reviews            int       `json:"reviews"`
	CoverURL           *string   `json:"coverURL" pg:"-"`
	URL                string    `json:"url"`
	IssuedAt           time.Time `json:"issuedAt"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
	Type               string    `json:"type" pg:"-"`

	// customized fields
	//Authors            []*Author    `json:"authors"`
	//Publishers         []*Publisher `json:"publishers"`
	//Topics             []*Topic     `json:"topics"`
	BookAuthor    *BookAuthor    `json:"-" pg:"-"`
	BookPublisher *BookPublisher `json:"-" pg:"-"`
	BookTopic     *BookTopic     `json:"-" pg:"-"`
}

type User struct {
	ID          string     `json:"id"`
	Email       string     `json:"email"`
	Password    string     `json:"-"`
	PhoneNumber string     `json:"phoneNumber"`
	FirstName   string     `json:"firstName"`
	LastName    string     `json:"lastName"`
	Verified    bool       `json:"verified"`
	CreatedAt   time.Time  `json:"createdAt"`
	LastLogin   *time.Time `json:"lastLogin"`
}

func (u *User) GenToken() (*AuthToken, error) {
	expiredAt := time.Now().Add(time.Hour * 24)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expiredAt.Unix(),
		Id:        u.ID,
		IssuedAt:  time.Now().Unix(),
		Issuer:    "Jun Liu",
	})
	accessToken, err := token.SignedString([]byte(os.Getenv(constants.JWTSecret)))
	if err != nil {
		return nil, err
	}

	return &AuthToken{
		AccessToken: accessToken,
		Expiration:  expiredAt,
	}, nil
}
