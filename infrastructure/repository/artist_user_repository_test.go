package repository

// TODO mysql packageへうつす
//func TestArtistUserRepository_FindByID(t *testing.T) {
//	db := mysql.InitConn()
//	repo := NewArtistUser(db)
//	_, err := repo.FindByID(1)
//	if err != nil {
//		panic(err)
//	}
//}
//
//func TestArtistUserRepository_Update(t *testing.T) {
//	db, mock, err := getConnAndMock()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	repo := NewArtistUser(db)
//
//	du := &entity.ArtistUser{
//		ID:              external.Encode(1),
//		Name:            "name",
//		IconURL:         "icon",
//		Identifier:      ARTISTIDF,
//		HeaderImageURL:  "header",
//		MetaDescription: "meta",
//		WhyDescription:  "why",
//		HowDescription:  "how",
//		FreeDescription: "free",
//	}
//
//	mock.ExpectExec("UPDATE `users`").
//		WithArgs(du.HeaderImageURL, du.IconURL, du.MetaDescription, du.Name, AnyTime{}, 1).
//		WillReturnResult(sqlmock.NewResult(1, 1))
//	mock.ExpectExec("UPDATE `artist_users`").
//		WithArgs(du.FreeDescription, du.HowDescription, AnyTime{}, du.WhyDescription, 1).
//		WillReturnResult(sqlmock.NewResult(1, 1))
//
//	if _, err = repo.Update(du); err != nil {
//		t.Fatal(err)
//	}
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("there were unfulfilled expectations: %s", err)
//	}
//
//}
