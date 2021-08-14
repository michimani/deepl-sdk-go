package types

import (
	"errors"
	"strings"
)

type APIPlan int

const (
	APIPlanNone APIPlan = iota
	APIPlanFree
	APIPlanPro
)

func (p APIPlan) Valid() bool {
	return p == APIPlanFree || p == APIPlanPro
}

func (p APIPlan) ToString() string {
	switch p {
	case APIPlanFree:
		return "free"
	case APIPlanPro:
		return "pro"
	default:
		return ""
	}
}

func ToAPIPlan(planName string) (APIPlan, error) {
	planName = strings.ToLower(planName)

	switch planName {
	case APIPlanFree.ToString():
		return APIPlanFree, nil
	case APIPlanPro.ToString():
		return APIPlanPro, nil
	default:
		return APIPlanNone, errors.New("invalid plan name")
	}
}
