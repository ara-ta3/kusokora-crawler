package kusokora

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type KusokoraRepositoryOnSQLite struct {
	DB *sql.DB
}

func NewKusokoraRepositoryOnSQLite(db *sql.DB) *KusokoraRepositoryOnSQLite {
	return &KusokoraRepositoryOnSQLite{DB: db}
}

func (kr *KusokoraRepositoryOnSQLite) GetAll() ([]Kusokora, error) {
	rows, err := kr.DB.Query("SELECT id, pictureUrl, FROM kusokoras;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	krs := []Kusokora{}
	for rows.Next() {
		var id int
		var pictureURL string
		err = rows.Scan(&id, &pictureURL)
		if err != nil {
			return nil, err
		}
		krs = append(krs, Kusokora{
			ID:         id,
			PictureURL: pictureURL,
		})
	}
	return krs, nil
}

func (kr *KusokoraRepositoryOnSQLite) Put(k Kusokora) error {
	tx, err := kr.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("INSERT INTO kusokoras(pictureUrl) values(?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(k.PictureURL)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}
