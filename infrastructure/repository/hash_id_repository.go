package repository

import (
	"errors"
	"github.com/speps/go-hashids"
	"github.com/tockn/emukone/domain/repository"
)

var (
	ErrDecode = errors.New("decode error")
)

type hashID struct {
	hash *hashids.HashID
}

func NewHashID(h *hashids.HashID) repository.HashID {
	return &hashID{h}
}

func (i *hashID) Decode(idStr string) (int64, error) {
	ns, err := i.hash.DecodeInt64WithError(idStr)
	if err != nil {
		return 0, ErrDecode
	}
	if len(ns) != 1 {
		return 0, ErrDecode
	}
	return ns[0], nil
}

func (i *hashID) Encode(id int64) string {
	idStr, _ := i.hash.EncodeInt64([]int64{id})
	return idStr
}
