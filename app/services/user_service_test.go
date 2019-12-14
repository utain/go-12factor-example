package services_test

import (
	"database/sql"
	"go-example/app/models"
	"go-example/app/services"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	service services.UserService
	person  *models.User
}

func (s *UserServiceTestSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)
	s.service = services.NewUserService(s.DB)
}

func (s *UserServiceTestSuite) TestServiceGetUser() {
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users"  WHERE "users"."deleted_at" IS NULL AND (("users"."username" = $1)) ORDER BY "users"."id" ASC LIMIT 1`)).
		WithArgs("utain").
		WillReturnRows(sqlmock.NewRows([]string{"id", "username"}).
			AddRow("1", "utain"))

	res := s.service.GetUser("utain")
	s.Assert().Contains(res.ID, "1")
	s.Assert().Contains(res.Username, "utain")
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
