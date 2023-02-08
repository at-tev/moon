package postgresql

import (
	"database/sql"

	"github.com/abadojack/whatlanggo"
	"github.com/at-tev/moon/pkg/models"
)

type ReviewModel struct {
	DB *sql.DB
}

func (m *ReviewModel) Insert(user, review string) error {
	info := whatlanggo.Detect(review)
	lang := info.Lang.Iso6391()
	stmt := `INSERT INTO reviews (user_moon, text_moon, lang_moon) VALUES($1, $2, $3)`

	_, err := m.DB.Exec(stmt, user, review, lang)
	if err != nil {
		return err
	}

	return nil
}

func (m *ReviewModel) AllReviews(lang string) ([]*models.Review, error) {
	stmt := `SELECT user_moon, text_moon, date_moon 
	FROM reviews WHERE ok IS TRUE ORDER BY CASE lang_moon WHEN $1 THEN 0 ELSE 1 END ASC, date_moon DESC`
	rows, err := m.DB.Query(stmt, lang)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reviews := make([]*models.Review, 0)
	for rows.Next() {
		review := models.Review{}
		err = rows.Scan(&review.User, &review.Text, &review.Created)
		if err != nil {
			return nil, err
		}

		reviews = append(reviews, &review)
	}

	if err = rows.Err(); err != nil {
		return reviews, err
	}

	return reviews, nil
}

func (m *ReviewModel) Random4(lang string) ([]*models.Review, error) {
	stmt := `SELECT user_moon, text_moon, date_moon 
	FROM reviews WHERE ok IS TRUE ORDER BY CASE lang_moon WHEN $1 THEN 0 ELSE 1 END ASC, date_moon DESC LIMIT 4`
	rows, err := m.DB.Query(stmt, lang)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reviews := make([]*models.Review, 0)
	for rows.Next() {
		review := models.Review{}
		err = rows.Scan(&review.User, &review.Text, &review.Created)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, &review)
	}

	if err = rows.Err(); err != nil {
		return reviews, err
	}

	return reviews, nil
}
