package v1_test

import (
	"database/sql"
	"fmt"
	v1 "go-example/internal/api/v1"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
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

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.api = v1.NewUserAPI(s.DB)
	s.mock.ExpectQuery(
		regexp.QuoteMeta(`SELECT * FROM "users"  WHERE "users"."deleted_at" IS NULL AND (("users"."id" = $1)) ORDER BY "users"."id" ASC LIMIT 1`)).
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "username"}).
			AddRow("1", "utain"))

	router := gin.Default()
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
