package dao

import (
	"github.com/stretchr/testify/require"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestPsqlDAO_DeleteKeywordNotImplemented(t *testing.T) {
	dao := NewPostgresDAO(nil)
	err := dao.DeleteKeyword()
	require.EqualError(t, err, "Not Implemented", "Expected not implemented error but got %s", err.Error())
}

func TestPsqlDAO_DeleteURLNotImplemented(t *testing.T) {
	dao := NewPostgresDAO(nil)
	err := dao.DeleteURL()
	require.EqualError(t, err, "Not Implemented", "Expected not implemented error but got %s", err.Error())
}

func TestPsqlDAO_AddURLHistoryEventNotImplemented(t *testing.T) {
	dao := NewPostgresDAO(nil)
	err := dao.AddURLHistoryEvent()
	require.EqualError(t, err, "Not Implemented", "Expected not implemented error but got %s", err.Error())
}

func TestPsqlDAO_AddKeyword(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err, "Unexpected Error [%v]", err)
	defer db.Close()
	dao := NewPostgresDAO(db)

	mock.ExpectExec("INSERT INTO keyword").WithArgs("foo", true).WillReturnResult(
		sqlmock.NewResult(1, 1))
	err = dao.AddKeyword("foo", true)
	require.NoError(t, err, "Unexpected Error [%v]", err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err, "Expectations were not met [%v]", err)
}

func TestPsqlDAO_AddURL(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err, "Unexpected Error [%v]", err)
	defer db.Close()
	dao := NewPostgresDAO(db)
	url := "https://example.com"
	title := "example"

	mock.ExpectExec("INSERT INTO url").WithArgs(url, title).WillReturnResult(
		sqlmock.NewResult(1, 1))
	err = dao.AddURL(url, title)
	require.NoError(t, err, "Unexpected Error [%v]", err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err, "Expectations were not met [%v]", err)
}

func TestPsqlDAO_AssociateKeywordURL(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err, "Unexpected Error [%v]", err)
	defer db.Close()
	dao := NewPostgresDAO(db)
	kwID := 42
	urlID := 1001

	mock.ExpectExec("INSERT INTO keyword_url").WithArgs(kwID, urlID).WillReturnResult(
		sqlmock.NewResult(1, 1))
	err = dao.AssociateKeywordURL(kwID, urlID)
	require.NoError(t, err, "Unexpected Error [%v]", err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err, "Expectations were not met [%v]", err)
}

func TestPsqlDAO_DisassociateKeywordURL(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err, "Unexpected Error [%v]", err)
	defer db.Close()
	dao := NewPostgresDAO(db)
	kwID := 42
	urlID := 1001

	mock.ExpectExec("DELETE FROM keyword_url").WithArgs(kwID, urlID).WillReturnResult(
		sqlmock.NewResult(1, 1))
	err = dao.DisassociateKeywordURL(kwID, urlID)
	require.NoError(t, err, "Unexpected Error [%v]", err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err, "Expectations were not met [%v]", err)
}
