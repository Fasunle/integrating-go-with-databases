package data

import (
	"context"
	"errors"
	"math/rand"
	"time"
)

type Password struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Code      string `json:"code"`
	Password  string `json:"-"`
	Used      bool   `json:"used"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (p *Password) ValidateCode(c string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
	select id, email, code, password, used, created_at, updated_at 
	from passwords 
	where email = $1, used = false
	`

	row := db.QueryRowContext(ctx, query, p.Email)

	err := row.Scan(
		&p.ID,
		&p.Email,
		&p.Code,
		&p.Password,
		&p.Used,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		return false, err
	}

	if p.Code != c {
		return false, errors.New("invalid code")
	}

	return true, nil
}

func (p *Password) FindByPassword(password string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
	select id, email, code, password, used, created_at, updated_at
	from passwords
	where password = $1 returning (id, email, code, password, used)
	`

	row := db.QueryRowContext(ctx, query, password)

	err := row.Scan(
		&p.ID,
		&p.Email,
		&p.Code,
		&p.Password,
		&p.Used,
	)

	if err != nil || p.Password != password {
		return false, err
	}

	return true, nil
}

func (p *Password) Insert(email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
	insert into passwords (email, code, used, created_at, updated_at)
	values ($1, $2, $3, $4, $5)
	`
	code := generateRandomString(6)

	_, err := db.ExecContext(ctx, query, email, code, false, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
