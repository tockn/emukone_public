package repository

import (
	"github.com/speps/go-hashids"
	"github.com/tockn/emukone/infrastructure/repository/mysql"
	"testing"
)

func TestUser_FindIDByExternalService(t *testing.T) {
	db := InitConn()
	es := mysql.ExternalServiceUID{
		Service: "twitter",
		UID:     "twitter_uid",
		UserID:  1,
	}
	if err := db.Create(&es).Error; err != nil {
		t.Fatal("create error:", err)
	}

	hh, _ := hashids.New()
	h := NewHashID(hh)
	repo := NewUser(db, h)

	id, ok, err := repo.FindIDByExternalService(es.ToDomain())
	if err != nil || !ok {
		db.Delete(&es)
		t.Fatal("find error:", err, ok)
	}

	if id != h.Encode(es.UserID) {
		db.Delete(&es)
		t.Fatal("expect:", es.UserID, "got:", id)
	}
	db.Delete(&es)
}
