# Go-ConfigManager
configManagerModule


## Usage:
```Go
package main

import (
	"fmt"
	"github.com/ozaiithejava/Go-ConfigManager/config"
)

func main() {
	// ConfigManager örneği oluştur
	configManager := config.NewConfigManager("config/config.yml")

	// .yml dosyasını yükle
	err := configManager.Load()
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	// .yml dosyasındaki değerlere eriş
	val, err := configManager.GetString("exampleKey")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	// Değerleri ekrana yazdır
	fmt.Println("exampleKey değeri:", val)

	// Yeni bir değer ekleyip kaydet
	configManager.Add("newKey", "Yeni Değer")
	err = configManager.Save()
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	// Eklenen değeri kontrol et
	if configManager.Check("newKey") {
		fmt.Println("newKey değeri:", configManager.config["newKey"])
	} else {
		fmt.Println("newKey bulunamadı.")
	}

	// Bir değeri silip kaydet
	configManager.Delete("exampleKey")
	err = configManager.Save()
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	// Silinen değeri kontrol et
	if configManager.Check("exampleKey") {
		fmt.Println("exampleKey değeri:", configManager.config["exampleKey"])
	} else {
		fmt.Println("exampleKey bulunamadı.")
	}
}
```
