package reports

import (
	"sync"

	"sdgestreaming/internal/database"
	"sdgestreaming/internal/models"
)

type ReportGenerator struct {
	db *database.DB
}

func NewReportGenerator(db *database.DB) *ReportGenerator {
	return &ReportGenerator{db: db}
}

// Reporte de usuarios registrados
func (rg *ReportGenerator) UsersReport() ([]models.User, error) {
	rows, err := rg.db.Conn.Query(
		"SELECT id, email, role FROM users",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Email, &u.Role); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

// Reporte de contenido disponible
func (rg *ReportGenerator) ContentReport() ([]models.Content, error) {
	rows, err := rg.db.Conn.Query(
		"SELECT id, title, category, genre, duration, year, artist, rating, min_age FROM contents",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contents []models.Content
	for rows.Next() {
		var c models.Content
		if err := rows.Scan(
			&c.ID,
			&c.Title,
			&c.Category,
			&c.Genre,
			&c.Duration,
			&c.Year,
			&c.Artist,
			&c.Rating,
			&c.MinAge,
		); err != nil {
			return nil, err
		}
		contents = append(contents, c)
	}

	return contents, nil
}

// Reporte simulado de ingresos por suscripciones
func (rg *ReportGenerator) IncomeReport() (float64, error) {
	row := rg.db.Conn.QueryRow(
		"SELECT IFNULL(SUM(amount), 0) FROM payments",
	)

	var total float64
	if err := row.Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}

// Generación concurrente de todos los reportes
func (rg *ReportGenerator) GenerateAllReports() (map[string]interface{}, error) {
	var (
		users    []models.User
		contents []models.Content
		income   float64
	)

	var wg sync.WaitGroup
	var err error

	wg.Add(3)

	go func() {
		defer wg.Done()
		users, err = rg.UsersReport()
	}()

	go func() {
		defer wg.Done()
		contents, err = rg.ContentReport()
	}()

	go func() {
		defer wg.Done()
		income, err = rg.IncomeReport()
	}()

	wg.Wait()

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"users":    users,
		"contents": contents,
		"income":   income,
	}, nil
}
