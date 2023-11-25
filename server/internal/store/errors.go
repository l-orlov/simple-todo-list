package store

import (
	"errors"
	"strings"
)

var (
	// ErrNotFound - запись не найдена
	ErrNotFound = errors.New("record not found")
	// ErrViolatesUniqConst - запись нарушает ограничение на уникальность
	ErrViolatesUniqConst = errors.New("violates unique constraint")
)

func dbError(err error) error {
	if strings.Contains(err.Error(), "no rows in result set") {
		return ErrNotFound
	}
	if strings.Contains(err.Error(), "violates unique constraint") {
		return ErrViolatesUniqConst
	}

	return err
}
