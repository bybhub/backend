package services

import "strings"

func ArrayToString(arr []string) string {
	return strings.Join(arr, ",")
}
