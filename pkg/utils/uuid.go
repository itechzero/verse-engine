package utils

import "github.com/google/uuid"

func NewUUID() string {
	return uuid.NewString()
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func NewUUIDs(n int) []string {
	uuids := make([]string, n)
	for i := 0; i < n; i++ {
		uuids[i] = NewUUID()
	}

	return uuids
}
