# 🔍 Recon-CLI: Otonom Zafiyet Keşif Modülü

Recon-CLI, siber güvenlik araştırmacıları ve bug bounty avcıları için geliştirilmiş, ağ tarama ve web servis analiz araçlarını tek merkezden yöneten **otonom bir CLI aracıdır**.

Proje Discovery ve Nmap altyapılarını kullanarak hedefleri tarar, çıktıları normalize eder ve otomatik olarak Markdown (.md) raporları üretir.

## 🛠️ Özellikler
* ⚡ **Hızlı & Güçlü:** Go (Golang) dili ve Cobra CLI framework'ü ile geliştirilmiştir.
* 🛡️ **Otonom Orkestrasyon:** Tek komutla hem port taraması (`nmap`) hem de web servis analizi (`httpx`) yapar.
* 📊 **Otomatik Raporlama:** Tarama sonuçlarını tarih ve hedef bazlı olarak temiz bir Markdown dosyası olarak diske kaydeder.

## 🚀 Kurulum & Gereksinimler

Projenin arka planda kullandığı bağımlılıkların sisteminizde yüklü olması gerekir:
* [Go](https://go.dev/) (v1.16+)
* [Nmap](https://nmap.org/)
* [Httpx](https://github.com/projectdiscovery/httpx)

```bash
# Repoyu klonlayın
git clone [https://github.com/bqroncy/recon-cli.git](https://github.com/bqroncy/recon-cli.git)
cd recon-cli

# Bağımlılıkları yükleyin
go mod tidy