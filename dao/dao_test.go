package dao

import (
	"github.com/stretchr/testify/require"
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
