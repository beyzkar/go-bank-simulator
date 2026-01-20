package repositorys

/*
	Sadece veri erişim katmanıdır
	DB ile konuşur
	CRUD yapar (Create, Read, Update, Delete)
	Karar vermez
*/

import (
	"strings"

	"github.com/beyza/go-bank-simulator/database" // DB bağlantısına ihtiyacımız var
	"github.com/beyza/go-bank-simulator/models"   // go da models/customer olarak çağıramayız, sadece package'ları çağırabiliriz
)

// =======================
// CREATE
// =======================

// Yeni müşteri oluşturur
func CreateCustomer(customer *models.Customer) error {
	/*
		“customer’ı her kullandığımda models içindeki Customer ŞABLONUNU kullanarak oluşturulmuş bir NESNE ile çalışıyorum
		ve repository’de bu NESNEYİ DB’ye ekliyorum.”
	*/
	return database.DB.Create(customer).Error
	/*
		database.DB: Global DB bağlantısı
		Create(customer): Bu customer’ı DB’ye yaz
	*/
}

// =======================
// READ
// =======================

// ID ile müşteri getirir
func GetCustomerByID(id uint) (*models.Customer, error) { // (id uint) DB'deki customers.id ile eşleşir
	var customer models.Customer
	err := database.DB.First(&customer, id).Error
	/*
		First(&customer, 10)  -> “ID = 10 olan kaydı getir (varsayılan sıralama ile).”
		Last(&customer, 10)   -> “ID = 10 olan kaydı getir ama sondan bakarak.”
		Take(&customer, 10)   -> “ID = 10 olan kaydı getir, sıralamayı dikkate alma.”
	*/
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

// Birebir isimle tek müşteri getirir (İLK kaydı döndürür)
func FindCustomerByName(name string) (*models.Customer, error) {
	var customer models.Customer
	result := database.DB.Where("name = ?", strings.TrimSpace(name)).First(&customer)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}

// Tüm müşterileri getirir
func GetAllCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	err := database.DB.Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

// =======================
// DELETE (CASCADE için yardımcılar)
// =======================

// ✅ Customer silinmeden önce o müşteriye ait hesapları silmek için
func DeleteAccountsByCustomerID(customerID uint) error {
	return database.DB.
		Where("customer_id = ?", customerID).
		Delete(&models.Account{}).Error
}

// ✅ Asıl müşteri silme fonksiyonu (SENDE EKSİKTİ)
// services.DeleteCustomer() bunu çağıracak
func DeleteCustomer(id uint) error {
	return database.DB.Delete(&models.Customer{}, id).Error
}

// =======================
// SEARCH
// =======================

// 🔍 İsimle arama (Zeynep → Zeynep Demir, Zeynep Kaya vs)
func SearchCustomersByName(q string) ([]models.Customer, error) {
	var customers []models.Customer
	q = strings.ToLower(strings.TrimSpace(q))

	err := database.DB.
		Where("LOWER(name) LIKE ?", "%"+q+"%").
		Order("id ASC").
		Find(&customers).Error

	return customers, err
}

// 🎯 Birebir isimle arama (aynı isimli birden fazla kişi dönebilir)
func FindCustomerByExactName(name string) ([]models.Customer, error) {
	var customers []models.Customer
	name = strings.TrimSpace(name)

	err := database.DB.
		Where("name = ?", name).
		Order("id ASC").
		Find(&customers).Error

	return customers, err
}
