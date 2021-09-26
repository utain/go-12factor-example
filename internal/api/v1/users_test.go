package v1_test

import (
	"database/sql"
	"fmt"
	v1 "go-example/internal/api/v1"
	"go-example/internal/errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserAPITestSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	api    v1.UserAPI
	router *gin.Engine
}

func (s *UserAPITestSuite) SetupSuite() {
	fmt.Println("SetupSuite")
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
	s.api = v1.NewUserAPI(s.DB)
	s.mock.ExpectQuery(
		regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`)).
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "username"}).
			AddRow("1", "utain"))
	s.mock.ExpectQuery(
		regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`)).
		WithArgs("x").
		WillReturnError(gorm.ErrRecordNotFound)
	router := gin.Default()
	router.Use(errors.GinError())
	router.GET("/api/v1/users/:id", s.api.GetUser)
	router.GET("/api/v1/users", s.api.GetAllUser)
	s.router = router
}

func (s *UserAPITestSuite) TestHTTPGetAllUsers() {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users", nil)
	s.router.ServeHTTP(res, req)
	s.Equal(http.StatusOK, res.Code, "Status must be 200:OK")
}

func (s *UserAPITestSuite) TestHTTPGetUsers() {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users/1", nil)
	s.router.ServeHTTP(res, req)
	s.Equal(http.StatusOK, res.Code, "Status must be 200:OK")
}

func (s *UserAPITestSuite) TestHTTPGetUsersWithEmptyData() {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users/x", nil)
	s.router.ServeHTTP(res, req)
	s.Equal(http.StatusNotFound, res.Code, "Status must be 404:NotFound")
}

func TestUserAPITestSuite(t *testing.T) {
	suite.Run(t, new(UserAPITestSuite))
}
