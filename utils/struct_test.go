package utils

import (
	"github.com/liujun5885/book_store_gql/graph/model"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestGetFieldsWithValue(t *testing.T) {
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
	ret := GetFieldsWithValue(&profiles)
	sort.Strings(ret)
	assert.Equal(t, []string{"address", "city", "country", "user_id"}, ret)
}
