package service

import "github.com/tockn/emukone/domain/repository"

type HashID interface {
	Encode(int64) string
	Decode(string) (int64, error)
}

type hashID struct {
	Repository repository.HashID
}

func NewHashID(repo repository.HashID) HashID {
	return &hashID{repo}
}

func (h *hashID) Encode(id int64) string {
	return h.Repository.Encode(id)
}

func (h *hashID) Decode(idStr string) (int64, error) {
	return h.Repository.Decode(idStr)
}
