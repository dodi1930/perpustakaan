package repositories

import (
	"fmt"
	"log"
	"perpustakaan-x-cgpt/config"
	"perpustakaan-x-cgpt/models"
	"time"
)

func GetLoanHistory(userID uint) ([]models.Loan, error) {
	var loans []models.Loan
	result := config.DB.Where("user_id = ? AND return_date IS NOT NULL", userID).Find(&loans)
	if result.Error != nil {
		return nil, result.Error
	}
	return loans, nil
}

func ReturnBook(loanID uint) error {
	// Ambil waktu sekarang
	now := time.Now()

	// Update return_date dan returned = true
	result := config.DB.Model(&models.Loan{}).
		Where("id = ?", loanID).
		Updates(map[string]interface{}{
			"return_date": now,
			"returned":    true,
		})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetLoanStatistics() (map[string]interface{}, error) {
	var totalLoansToday int64
	var totalLoansThisWeek int64
	var totalLoansThisMonth int64

	// Hitung peminjaman hari ini
	today := time.Now().Format("2006-01-02")
	config.DB.Model(&models.Loan{}).Where("DATE(borrow_date) = ?", today).Count(&totalLoansToday)

	// Hitung peminjaman minggu ini
	startOfWeek := time.Now().AddDate(0, 0, -int(time.Now().Weekday())).Format("2006-01-02")
	config.DB.Model(&models.Loan{}).Where("DATE(borrow_date) >= ?", startOfWeek).Count(&totalLoansThisWeek)

	// Hitung peminjaman bulan ini
	startOfMonth := time.Now().Format("2006-01") + "-01"
	config.DB.Model(&models.Loan{}).Where("DATE(borrow_date) >= ?", startOfMonth).Count(&totalLoansThisMonth)

	// Buku yang paling sering dipinjam
	var mostBorrowedBooks []struct {
		BookID uint
		Count  int
	}
	config.DB.Table("loans").Select("book_id, COUNT(*) as count").Group("book_id").Order("count DESC").Limit(5).Scan(&mostBorrowedBooks)

	// Pengguna paling aktif meminjam
	var topActiveUsers []struct {
		UserID uint
		Count  int
	}
	config.DB.Table("loans").Select("user_id, COUNT(*) as count").Group("user_id").Order("count DESC").Limit(5).Scan(&topActiveUsers)

	// Kembalikan hasil statistik
	return map[string]interface{}{
		"total_loans_today":      totalLoansToday,
		"total_loans_this_week":  totalLoansThisWeek,
		"total_loans_this_month": totalLoansThisMonth,
		"most_borrowed_books":    mostBorrowedBooks,
		"top_active_users":       topActiveUsers,
	}, nil
}

func GetPendingLoans(userID uint) ([]models.Loan, error) {
	var loans []models.Loan

	log.Println("🔍 Mencari peminjaman untuk user:", userID)

	// Pastikan koneksi database tersedia
	if config.DB == nil {
		log.Println("❌ Database belum terkoneksi!")
		return nil, fmt.Errorf("database tidak tersedia")
	}

	// Jalankan query untuk mengambil data peminjaman
	result := config.DB.Unscoped().Where("user_id = ?", userID).Find(&loans)

	// Debugging hasil query
	if result.Error != nil {
		log.Println("❌ Query gagal:", result.Error)
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		log.Println("⚠️ Tidak ada data peminjaman ditemukan!")
	}

	log.Println("✅ Peminjaman ditemukan:", loans)
	return loans, nil
}
