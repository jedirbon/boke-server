package main

import (
	"fmt"
	"github.com/dchest/captcha"
	"net/http"
)

func main() {
	// 生成验证码（默认 6 位数字）
	captchaID := captcha.New()
	fmt.Println("Captcha ID:", captchaID)

	// 在 HTTP 服务中返回验证码图片
	http.HandleFunc("/captcha", func(w http.ResponseWriter, r *http.Request) {
		captcha.WriteImage(w, captchaID, 120, 80) // 宽 120px，高 80px
	})

	// 验证用户输入
	http.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		userInput := r.FormValue("captcha")
		captchaID := r.FormValue("captcha_id")
		if captcha.VerifyString(captchaID, userInput) {
			w.Write([]byte("验证码正确！"))
		} else {
			w.Write([]byte("验证码错误！"))
		}
	})

	http.ListenAndServe(":8080", nil)
}
