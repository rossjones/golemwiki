package main

import (
	"unicode"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Group struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func slugify(name string) string {
	buffer := make([]rune, 0, len(name))
	for _, r := range name {
		switch {
		case unicode.IsLetter(r):
			r = unicode.ToLower(r)
		case unicode.IsSpace(r):
			r = '-'
		case unicode.IsNumber(r):
			r = unicode.ToLower(r)
		}
		buffer = append(buffer, r)
	}

	if i := len(buffer) - 1; i >= 0 && buffer[i] == '-' {
		buffer = buffer[:i]
	}

	return string(buffer)
}
