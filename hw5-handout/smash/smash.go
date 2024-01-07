WeChat: cstutorcs
QQ: 749389476
Email: tutorcs@163.com
package smash

import (
	"io"
)

type word string

func Smash(r io.Reader, smasher func(word) uint32) map[uint32]uint {
	m := make(map[uint32]uint)
	// TODO: Incomplete!
	return m
}
