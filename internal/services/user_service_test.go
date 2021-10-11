package services_test

import (
	"database/sql"

	"go-example/internal/entities"
	"go-example/internal/services"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	s.DB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}))
	require.NoError(s.T(), err)
	entities.AutoMigrate(s.DB)
	require.NoError(s.T(), err)

	s.service = services.NewUserService(s.DB)
}

func (s *UserServiceTestSuite) TestServiceGetUser() {
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`)).
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "username"}).
			AddRow("1", "utain"))

	user, err := s.service.GetUser("1")
	s.Assert().NoError(err)
	s.Assert().Contains(user.ID, "1")
	s.Assert().Contains(user.Username, "utain")

	userX, err := s.service.GetUser("x")
	s.Assert().Error(err, "User id=x should not found")
	s.Assert().Nil(userX)
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
