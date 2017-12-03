// Copyright © 2017 Sean Smith <sean@wwsean08.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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
	keyword := "foo"
	isRegex := true

	mock.ExpectQuery("INSERT INTO keywords").WithArgs(keyword, isRegex).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
	id, err := dao.AddKeyword(keyword, isRegex)
	require.NoError(t, err, "Unexpected Error [%v]", err)
	require.Equal(t, 1, id)

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

	mock.ExpectQuery("INSERT INTO url").WithArgs(url, title).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1000"))

	id, err := dao.AddURL(url, title)
	require.NoError(t, err, "Unexpected Error [%v]", err)
	require.Equal(t, 1000, id)

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
