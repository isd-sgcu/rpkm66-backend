package utils

import (
	"github.com/google/uuid"
)

func BoolAdr(b bool) *bool {
	return &b
}

func UUIDAdr(in uuid.UUID) *uuid.UUID {
	return &in
}

func IsDuplicatedString(in []string) bool {
	checker := map[string]struct{}{}

	for _, e := range in {
		if _, ok := checker[e]; !ok {
			checker[e] = struct{}{}
			continue
		}
		return true
	}

	return false
}
