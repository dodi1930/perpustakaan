package repositories

import (
	"perpustakaan-x-cgpt/config"
	"perpustakaan-x-cgpt/models"
)

// Fungsi untuk mendapatkan nilai denda per hari
func GetFinePerDay() (float64, error) {
	var fineSetting models.FineSetting

	// Ambil nilai denda dari database
	if err := config.DB.First(&fineSetting).Error; err != nil {
		return 0, err
	}
	return fineSetting.FinePerDay, nil
}

// Fungsi untuk memperbarui denda per hari
func UpdateFinePerDay(newFine float64) error {
	var fineSetting models.FineSetting

	// Ambil data pertama dari tabel fine_settings
	if err := config.DB.First(&fineSetting).Error; err != nil {
		return err
	}

	// Perbarui nilai denda
	fineSetting.FinePerDay = newFine

	// Simpan perubahan ke database
	if err := config.DB.Save(&fineSetting).Error; err != nil {
		return err
	}

	return nil
}
