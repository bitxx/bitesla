package email

import (
	"fmt"
	"net"
	"net/smtp"
	"os"
	"strings"
)

var whoareemail = make(map[string]string)

func init() {
	var dev = []string{"192.168.1.220", "192.168.1.215"}                                                                     // 开发
	var test = []string{"192.168.1.221", "192.168.1.222", "192.168.1.203"}                                                   // 测试
	var preRelease = []string{"10.139.154.119"}                                                                              // 准正式（存管版本使用）
	var release = []string{"10.139.102.202", "10.253.0.218", "10.253.2.231", "10.253.2.94", "10.253.2.194", "10.253.13.189"} // 正式
	for i := 0; i < len(dev); i++ {
		whoareemail[dev[i]] = "dev@zcmlc.com"
	}
	for i := 0; i < len(test); i++ {
		whoareemail[test[i]] = "test@zcmlc.com"
	}
	for i := 0; i < len(preRelease); i++ {
		whoareemail[preRelease[i]] = "pre_release@zcmlc.com"
	}
	for i := 0; i < len(release); i++ {
		whoareemail[release[i]] = "release@zcmlc.com"
	}

}

// WhoAreEmail 根据IP获取 对应邮箱
func WhoAreEmail(IP string) string {
	return whoareemail[IP]
}

// SendEmail 发送邮件
func SendEmail(title, content, toUser string) {
	host := "smtp.exmail.qq.com:25"
	to := strings.Split(toUser, ";") //收件人  ;号隔开
	contentType := "Content-Type: text/html; charset=UTF-8"
	var err error
	ip := getBdIP()
	if ip == "not found" {
		msg := []byte("To: " + toUser + "\r\nFrom: zhulj@zcmlc.com\r\nSubject:" + title + "\r\n" + contentType + "\r\n\r\n" + content)
		err = smtp.SendMail(host, smtp.PlainAuth("", "zhulj@zcmlc.com", "Zlj82324572", "smtp.exmail.qq.com"), "zhulj@zcmlc.com", to, []byte(msg))
	} else {
		senduser := WhoAreEmail(ip)
		msg := []byte("To: " + toUser + "\r\nFrom: " + senduser + "\r\nSubject:" + title + "\r\n" + contentType + "\r\n\r\n" + content)
		err = smtp.SendMail(host, smtp.PlainAuth("", senduser, "Zcmlc2017", "smtp.exmail.qq.com"), senduser, to, []byte(msg))
	}
	if err != nil {
		fmt.Println(err)
	}
}

// getBdIP 获取工具包所在服务器的IP地址
func getBdIP() string {
	ip := "not found"
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return ip
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			}
		}
	}
	return ip
}
