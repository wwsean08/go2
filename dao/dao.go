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
	"database/sql"
	"github.com/pkg/errors"
	"time"
)

type RedirectorDAO interface {
	AddURL(url, title string) (int, error)
	AddKeyword(keyword string, isRegex bool) (int, error)
	AddURLHistoryEvent() error
	AssociateKeywordURL(keywordID, urlID int) error

	DeleteKeyword() error
	DeleteURL() error
	DisassociateKeywordURL(keywordID, urlID int) error
}

type psqlDAO struct {
	conn *sql.DB
}

type Keyword struct {
	Id         int
	Keyword    string
	ResultMode int
	IsRegex    bool
}

type URL struct {
	Id          int
	URL         string
	ClickCount  int
	LastClicked time.Time
	Title       string
}

func NewPostgresDAO(conn *sql.DB) RedirectorDAO {
	dao := psqlDAO{
		conn: conn,
	}
	return dao
}

func (dao psqlDAO) AddKeyword(keyword string, isRegex bool) (int, error) {
	stmt := `INSERT INTO keywords (keyword, is_regex) VALUES ($1, $2) RETURNING id`
	var id int
	err := dao.conn.QueryRow(stmt, keyword, isRegex).Scan(&id)
	return id, err
}

func (dao psqlDAO) AddURL(url, title string) (int, error) {
	stmt := `INSERT INTO url (url, title) VALUES ($1, $2) RETURNING id`
	var id int
	err := dao.conn.QueryRow(stmt, url, title).Scan(&id)
	return id, err
}

func (dao psqlDAO) AssociateKeywordURL(keywordID, urlID int) error {
	stmt := `INSERT INTO keyword_url (keyword_id, url_id) VALUES ($1, $2)`
	_, err := dao.conn.Exec(stmt, keywordID, urlID)
	return err
}

func (dao psqlDAO) DeleteKeyword() error {
	// TODO: Implement
	return errors.New("Not Implemented")
}

func (dao psqlDAO) DeleteURL() error {
	// TODO: Implement
	return errors.New("Not Implemented")
}

func (dao psqlDAO) DisassociateKeywordURL(keywordID, urlID int) error {
	stmt := `DELETE FROM keyword_url WHERE keyword_id = $1 AND url_id = $2`
	_, err := dao.conn.Exec(stmt, keywordID, urlID)
	return err
}

func (dao psqlDAO) AddURLHistoryEvent() error {
	// TODO: Implement
	return errors.New("Not Implemented")
}
