package repository_test

import (
	"clean_arch/domain"
	userdto "clean_arch/dto/userDto"

	"clean_arch/src/user/repository"
	"clean_arch/utils/db"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	userRepo domain.UserRepository
	database *gorm.DB
	mock     sqlmock.Sqlmock
	userData []domain.User
)

func init() {
	database, mock = db.DbMock()
	userRepo = repository.NewUserRepo(database)
	userData = []domain.User{
		{
			ID:   "1",
			Name: "Adrian",
		},
		{
			ID:   "2",
			Name: "Oemar",
		},
	}
}

func TestGetAll(t *testing.T) {
	// t.Parallel()

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(userData[0].ID, userData[0].Name).
		AddRow(userData[1].ID, userData[1].Name)

	// mock.ExpectQuery("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL").WillReturnRows(rows)
	mock.ExpectQuery("SELECT (.*)").WillReturnRows(rows)

	data, err := userRepo.GetAll()

	userDataParsed := []domain.UserGetAllAPI{}
	copier.Copy(&userDataParsed, data)

	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, userDataParsed, data)
}

func TestGetById(t *testing.T) {
	id := "1"

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(userData[0].ID, userData[0].Name)

	mock.ExpectQuery("SELECT (.*)").WithArgs(id).WillReturnRows(rows)

	data, err := userRepo.GetById(id)

	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, userData[0], data)
}

func TestCreate(t *testing.T) {
	id := userData[0].ID

	mock.ExpectBegin()
	mock.ExpectExec("INSERT (.*)").WithArgs(id, userData[0].Name, sqlmock.AnyArg(), sqlmock.AnyArg(), nil).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := userRepo.Create(id, userdto.CreateUserDto{
		Name: userData[0].Name,
	})

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())

}
