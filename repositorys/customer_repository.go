package repositorys

/*
	Sadece veri erişim katmanıdır
	DB ile konuşur
	CRUD yapar (Create, Read, Update, Delete)
	Karar vermez
*/
import (
	"github.com/beyza/go-bank-simulator/database" //DB bağlantısına ihtiyacımız var
	"github.com/beyza/go-bank-simulator/models"   //go da models/customer olarak çağıramayız, sadece packageları çağırabiliriz
)

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

// ID ile müşteri getirir
func GetCustomerByID(id uint) (*models.Customer, error) { //(id uint) bu kısım modelsın içerisindeki costumerın içerisindeki id ile eşleşiyor
	var customer models.Customer
	err := database.DB.First(&customer, id).Error
	/*
		First(&customer, 10)
		“ID = 10 olan kaydı getir, ama sıralamayı da dikkate al.”

		err := database.DB.Last(&customers, 10).Error --->  “ID = 10 olan kaydı getir, ama sondan bakarak.”
		err := database.DB.Take(&customers, 10).Error --->  “ID = 10 olan kaydı getir, sıralamayı dikkate alma.”
	*/
	if err != nil {
		return nil, err // eğer hata varsa nill döndür ve hatayı yaz
	}

	return &customer, nil //eğer bir hata yoksa customerı döntür hatayı da nill olarak döndür
}
func FindCustomerByName(name string) (*models.Customer, error) {
	var customer models.Customer
	result := database.DB.Where("name = ?", name).First(&customer)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}

// Tüm müşterileri getirir
func GetAllCustomers() ([]models.Customer, error) {
	var customers []models.Customer           //[] models.Customer: customer listesi oluştur, boş bir slice oluştur
	err := database.DB.Find(&customers).Error //“DB’ye git, bütün customers kayıtlarını bul, bu slice’ın içine doldur.”
	if err != nil {
		return nil, err
	}
	return customers, nil

}

// Müşteri siler
func DeleteCustomer(id uint) error {
	return database.DB.Delete(&models.Customer{}, id).Error
}
