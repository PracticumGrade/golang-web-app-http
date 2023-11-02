package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const storageFilename = "rest-api-storage.json"

var (
	ErrNotFound      = fmt.Errorf("data not found")
	ErrAlreadyExists = fmt.Errorf("data already exists")
)

type UserData struct {
	UserID    int64  `json:"user_id,omitempty"`
	UserName  string `json:"user_name,omitempty"`
	UserEmail string `json:"user_email,omitempty"`
}

type (
	StorageI interface {
		// RecoverAll recovers all users from file.
		RecoverAll() ([]UserData, error)
		// Recover recovers specific user from file. Returns error if user not found
		Recover(id int64) (UserData, error)
		// Store stores new user to file. Returns error if user already exists.
		Store(userData UserData) (id int64, err error)
		// Update updates user in file. Returns error if user not found.
		Update(userData UserData) error
		// Delete deletes user from file. Returns error if user not found.
		Delete(userData UserData) error
	}

	fileStorageManagement struct {
		fullPath string
	}
)

func New(path string) StorageI {
	return &fileStorageManagement{fullPath: filepath.Join(path, storageFilename)}
}

func (f *fileStorageManagement) RecoverAll() ([]UserData, error) {
	storage, err := os.Open(f.fullPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []UserData{}, nil
		}

		return nil, fmt.Errorf("os.Open: %w", err)
	}

	defer storage.Close()

	var users []UserData

	err = json.NewDecoder(storage).Decode(&users)
	if err != nil {
		return nil, fmt.Errorf("json.NewDecoder().Decode(): %w", err)
	}

	return users, nil
}

func (f *fileStorageManagement) Recover(id int64) (UserData, error) {
	users, err := f.RecoverAll()
	if err != nil {
		return UserData{}, fmt.Errorf("f.RecoverAll(): %w", err)
	}

	for _, user := range users {
		if user.UserID == id {
			return user, nil
		}
	}

	return UserData{}, fmt.Errorf("%w: id=%d", ErrNotFound, id)
}

func (f *fileStorageManagement) Store(userData UserData) (int64, error) {
	users, err := f.RecoverAll()
	if err != nil {
		return 0, fmt.Errorf("f.RecoverAll(): %w", err)
	}

	var lastID int64

	for _, user := range users {
		lastID = user.UserID

		if user.UserID == userData.UserID ||
			user.UserName == userData.UserName ||
			user.UserEmail == userData.UserEmail {
			return user.UserID, fmt.Errorf("%w: id=%d", ErrAlreadyExists, lastID)
		}
	}

	userData.UserID = lastID + 1

	users = append(users, userData)

	err = f.saveUsers(users)
	if err != nil {
		return userData.UserID, fmt.Errorf("f.saveUsers: %w", err)
	}

	return userData.UserID, nil
}

func (f *fileStorageManagement) Update(userData UserData) error {
	users, err := f.RecoverAll()
	if err != nil {
		return fmt.Errorf("f.RecoverAll(): %w", err)
	}

	found := false
	for i, user := range users {
		if user.UserID == userData.UserID {
			users[i] = userData
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("%w: id=%d", ErrNotFound, userData.UserID)
	}

	err = f.saveUsers(users)
	if err != nil {
		return fmt.Errorf("f.saveUsers: %w", err)
	}

	return nil
}

func (f *fileStorageManagement) Delete(userData UserData) error {
	users, err := f.RecoverAll()
	if err != nil {
		return fmt.Errorf("f.RecoverAll(): %w", err)
	}

	found := false
	for i, user := range users {
		if user.UserID == userData.UserID {
			users = append(users[:i], users[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("%w: id=%d", ErrNotFound, userData.UserID)
	}

	err = f.saveUsers(users)
	if err != nil {
		return fmt.Errorf("f.saveUsers: %w", err)
	}

	return nil
}

func (f *fileStorageManagement) saveUsers(users []UserData) error {
	storage, err := os.Create(f.fullPath)
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}

	defer storage.Close()

	err = json.NewEncoder(storage).Encode(users)
	if err != nil {
		return fmt.Errorf("json.NewEncoder().Encode(): %w", err)
	}

	return nil
}
