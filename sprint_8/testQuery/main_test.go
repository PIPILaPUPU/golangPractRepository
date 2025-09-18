package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		return
	}
	defer db.Close()

	clientID := 1

	// напиши тест здесь
	cl, err := selectClient(db, clientID)
	if err != nil {
		return
	}

	assert.Equal(t, clientID, cl.ID)
	assert.NotEmpty(t, cl.FIO)
	assert.NotEmpty(t, cl.Login)
	assert.NotEmpty(t, cl.Birthday)
	assert.NotEmpty(t, cl.Email)
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		return
	}
	defer db.Close()

	clientID := -1

	// напиши тест здесь
	cl, err := selectClient(db, clientID)

	assert.EqualError(t, err, sql.ErrNoRows.Error())
	assert.Equal(t, cl.ID, 0)
	assert.Empty(t, cl.FIO)
	assert.Empty(t, cl.Login)
	assert.Empty(t, cl.Birthday)
	assert.Empty(t, cl.Email)
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		return
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	cl.ID, err = insertClient(db, cl)
	require.NotNil(t, cl.ID)
	require.Nil(t, err)

	client, err := selectClient(db, cl.ID)
	require.Nil(t, err)
	assert.Equal(t, client.ID, cl.ID)
	assert.Equal(t, client.FIO, cl.FIO)
	assert.Equal(t, client.Login, cl.Login)
	assert.Equal(t, client.Birthday, cl.Birthday)
	assert.Equal(t, client.Email, cl.Email)
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		return
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	cl.ID, err = insertClient(db, cl)
	require.NotNil(t, cl.ID)
	require.Nil(t, err)

	client, err := selectClient(db, cl.ID)
	require.Nil(t, err)

	err = deleteClient(db, client.ID)

	_, err = selectClient(db, cl.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}
