package core

import (
	"fmt"
	"os/exec"
)

func RunHttpxScan(target string) string {
	fmt.Println("\n[+] HTTPX Modülü Başlatıldı...")
	fmt.Println("[*] Web servis detayları inceleniyor...")

	cmd := exec.Command("httpx", "-u", target, "-status-code", "-title", "-no-color")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("HTTPX hatası: %v", err)
	}

	fmt.Println("✅ HTTPX Analizi Tamamlandı.")
	return string(output)
}