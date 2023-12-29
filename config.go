package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// ConfigManager, .yml dosyalarını yönetmek için bir yapı
type ConfigManager struct {
	filePath string
	config   map[string]interface{}
}

// NewConfigManager, yeni bir ConfigManager örneği oluşturur
func NewConfigManager(filePath string) *ConfigManager {
	return &ConfigManager{
		filePath: filePath,
		config:   make(map[string]interface{}),
	}
}

// Load, .yml dosyasını okur ve içeriği ConfigManager'a yükler
func (cm *ConfigManager) Load() error {
	content, err := ioutil.ReadFile(cm.filePath)
	if err != nil {
		return fmt.Errorf("Dosya okuma hatası: %v", err)
	}

	err = yaml.Unmarshal(content, &cm.config)
	if err != nil {
		return fmt.Errorf("YAML çözme hatası: %v", err)
	}

	return nil
}

// Save, ConfigManager'daki değişiklikleri .yml dosyasına kaydeder
func (cm *ConfigManager) Save() error {
	content, err := yaml.Marshal(&cm.config)
	if err != nil {
		return fmt.Errorf("YAML kodlama hatası: %v", err)
	}

	err = ioutil.WriteFile(cm.filePath, content, 0644)
	if err != nil {
		return fmt.Errorf("Dosya yazma hatası: %v", err)
	}

	return nil
}

// GetString, belirtilen anahtarın değerini string olarak döndürür
func (cm *ConfigManager) GetString(key string) (string, error) {
	val, ok := cm.config[key].(string)
	if !ok {
		return "", fmt.Errorf("'%s' anahtarı bir string değil", key)
	}
	return val, nil
}

// Add, yeni bir anahtar-değer çifti ekler
func (cm *ConfigManager) Add(key string, value interface{}) {
	cm.config[key] = value
}

// Delete, belirtilen anahtarı siler
func (cm *ConfigManager) Delete(key string) {
	delete(cm.config, key)
}

// Check, belirtilen anahtarın varlığını kontrol eder
func (cm *ConfigManager) Check(key string) bool {
	_, ok := cm.config[key]
	return ok
}
