package v1_test

import (
	"database/sql"
	v1 "go-example/app/api/v1"
	"go-example/app/models"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
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
	person *models.User
}

func (s *UserAPITestSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.api = v1.NewUserAPI(s.DB, true)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users"  WHERE "users"."deleted_at" IS NULL AND (("users"."username" = $1)) ORDER BY "users"."id" ASC LIMIT 1`)).
		WithArgs("utain").
		WillReturnRows(sqlmock.NewRows([]string{"id", "username"}).
			AddRow("1", "utain"))
}
func (s *UserAPITestSuite) TestHTTPGetUsers() {
	router := gin.Default()
	router.GET("/:name", s.api.GetUser)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/utain", nil)
	router.ServeHTTP(res, req)
	log.SetOutput(os.Stdout)
	s.Equal(http.StatusOK, res.Code, "Status must be 200:OK")
}

func TestUserAPITestSuite(t *testing.T) {
	suite.Run(t, new(UserAPITestSuite))
}
