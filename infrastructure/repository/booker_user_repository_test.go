package repository

// TODO mysql packageにうつす
//func TestBookerUserRepository_FindByID(t *testing.T) {
//	db := initConn()
//	repo := NewBookerUser(db)
//	log.Println(repo.FindByID(3))
//}
//
//func TestBookerUserRepository_Update(t *testing.T) {
//	db, mock, err := getConnAndMock()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	repo := NewBookerUser(db)
//
//	du := &entity.BookerUser{
//		ID:              external.Encode(1),
//		Name:            "name",
//		IconURL:         "icon",
//		Identifier:      BOOKERIDF,
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
//	mock.ExpectExec("UPDATE `booker_users`").
//		WithArgs(du.FreeDescription, du.HowDescription, AnyTime{}, du.WhyDescription, 1).
//		WillReturnResult(sqlmock.NewResult(1, 1))
//
//	if _, err = repo.Update(du); err != nil {
//		t.Fatal(err)
//	}
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("there were unfulfilled expectations: %s", err)
//	}
//}
