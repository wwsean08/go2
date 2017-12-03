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
