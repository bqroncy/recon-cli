package cmd

import (
	"fmt"
	"os"

	"recon-cli/core"

	"github.com/spf13/cobra"
)

var target string

var rootCmd = &cobra.Command{
	Use:   "recon",
	Short: "Otonom Bilgi Toplama ve Recon Aracı",
	Run: func(cmd *cobra.Command, args []string) {
		if target == "" {
			fmt.Println("❌ Hata: Bir hedef belirtmelisiniz!")
			cmd.Usage()
			os.Exit(1)
		}
		
		fmt.Printf("🔍 Hedef analiz ediliyor: %s\n", target)
		fmt.Println("⚙️ Otonom recon modülleri sırayla başlatılacak...")

		// Modülleri çalıştırıp çıktılarını değişkenlere topluyoruz
		nmapResult := core.RunNmapScan(target)
		httpxResult := core.RunHttpxScan(target)

		// Raporlama modülüne verileri gönderiyoruz
		core.SaveReport(target, nmapResult, httpxResult)
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&target, "target", "t", "", "Hedef domain veya IP adresi (Zorunlu)")
}