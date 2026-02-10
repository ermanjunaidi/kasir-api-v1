package repositories

import (
	"database/sql"
	"kasir-api/models"
	"time"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetSalesReport(startDate, endDate time.Time) (*models.SalesReport, error) {
	report := &models.SalesReport{}

	// 1. Total Revenue and Total Transaksi
	err := repo.db.QueryRow(`
		SELECT COALESCE(SUM(total_amount), 0), COUNT(id) 
		FROM transactions 
		WHERE created_at >= $1 AND created_at <= $2`,
		startDate, endDate).Scan(&report.TotalRevenue, &report.TotalTransaksi)

	if err != nil {
		return nil, err
	}

	// 2. Produk Terlaris
	topProduct := &models.TopProduct{}
	err = repo.db.QueryRow(`
		SELECT p.name, SUM(td.quantity) as total_qty
		FROM transaction_details td
		JOIN products p ON td.product_id = p.id
		JOIN transactions t ON td.transaction_id = t.id
		WHERE t.created_at >= $1 AND t.created_at <= $2
		GROUP BY p.name
		ORDER BY total_qty DESC
		LIMIT 1`,
		startDate, endDate).Scan(&topProduct.Nama, &topProduct.QtyTerjual)

	if err == sql.ErrNoRows {
		report.ProdukTerlaris = nil
	} else if err != nil {
		return nil, err
	} else {
		report.ProdukTerlaris = topProduct
	}

	return report, nil
}
