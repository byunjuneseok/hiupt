package users

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type User struct {
	Id        string `json:"id"`
	StudentId string `json:"student_id"`
	Email     string `json:"email"`
}

func (user User) setId() error {
	if user.Id != "" {
		return errors.New("already has id")
	}
	user.Id = uuid.New().String()
	return nil
}

func (user User) getSchoolEmail() string {
	return fmt.Sprintf("%s@mail.hongik.ac.kr", user.StudentId)
}
