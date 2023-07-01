package utils

import (
	"strconv"
	"time"

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

func GetCurrentTimePtr() *time.Time {
	tmp := time.Now()
	return &tmp
}

func GetCacheKey(userid string, eventType int32) string {
	return userid + ":" + strconv.Itoa(int(eventType))
}
