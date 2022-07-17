package utils

import "github.com/google/uuid"

func BoolAdr(b bool) *bool {
	return &b
}

func UUIDAdr(in uuid.UUID) *uuid.UUID {
	return &in
}
