package core

import (
	"fmt"
	"os/exec"
)

func RunNmapScan(target string) string {
	fmt.Println("\n[+] Nmap Modülü Başlatıldı...")
	fmt.Println("[*] Temel port taraması yapılıyor, lütfen bekleyin...")

	cmd := exec.Command("nmap", "-F", target)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Nmap hatası: %v", err)
	}

	fmt.Println("✅ Nmap Taraması Tamamlandı.")
	return string(output)
}