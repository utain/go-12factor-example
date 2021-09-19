package services_test

import (
	"database/sql"

	"go-example/internal/entities"
	"go-example/internal/services"
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
		`SELECT * FROM "users"  WHERE "users"."deleted_at" IS NULL AND (("users"."id" = $1)) ORDER BY "users"."id" ASC LIMIT 1`)).
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "username"}).
			AddRow("1", "utain"))

	user := new(entities.User)
	s.service.GetUser(user, "1")
	s.Assert().Contains(user.ID, "1")
	s.Assert().Contains(user.Username, "utain")

	userX := new(entities.User)
	err := s.service.GetUser(userX, "x")
	s.Assert().NotNil(err, "User id=x should not found")
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
