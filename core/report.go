package core

import (
	"fmt"
	"os"
	"time"
)

// SaveReport, tarama sonuçlarını Markdown formatında bir dosyaya kaydeder
func SaveReport(target string, nmapOut string, httpxOut string) {
	fmt.Println("\n[+] Raporlama Modülü Başlatıldı...")

	// Dosya adını hedef domain ve tarih ile benzersiz yapıyoruz
	currentTime := time.Now().Format("2006-01-02_15-04")
	fileName := fmt.Sprintf("recon_%s_%s.md", target, currentTime)

	// Markdown şablonu oluşturuyoruz
	reportContent := fmt.Sprintf(`# Recon Analiz Raporu - %s
Tarama Tarihi: %s

## 🔍 Hedef Bilgisi
* **Domain/IP:** %s

## 🛡️ Nmap Port Tarama Sonuçları
%s
%s
%s

## 🌐 HTTPX Web Servis Analiz Sonuçları
%s
%s
%s

---
*Bu rapor Recon-CLI tarafından otonom olarak üretilmiştir.*
`, target, time.Now().Format("02.01.2006 15:04"), target, "```text", nmapOut, "```", "```text", httpxOut, "```")

	// Dosyayı diske yazıyoruz
	err := os.WriteFile(fileName, []byte(reportContent), 0644)
	if err != nil {
		fmt.Println("❌ Hata: Rapor dosyası kaydedilemedi:", err)
		return
	}

	fmt.Printf("🎯 Rapor başarıyla oluşturuldu: %s 📂\n", fileName)
}