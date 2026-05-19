package main

import (
	"fmt"
	"os"

	"recon-cli/cmd" // cmd klasöründeki kodlarımızı çağırıyoruz
)

func main() {
	// Root komutunu çalıştır (cmd/root.go dosyasındaki Execute fonksiyonu)
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Uygulama başlatılırken hata oluştu:", err)
		os.Exit(1)
	}
}