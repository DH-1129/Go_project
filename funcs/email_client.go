package funcs

import (
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"regexp"
	"time"

	// "sync"

	EM "github.com/jordan-wright/email"
)

// izvpcfcmjjxlhjfh 邮箱授权码

func init() {
	rand.Seed(time.Now().UnixNano())
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandStringBytes(num int) string {
	b := make([]byte, num)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// 同步
func Send_Email(email string) (c string, err error) {
	// 简单设置 log 参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	em := EM.NewEmail()
	// 设置发送方邮箱
	em.From = "1577108595@qq.com"
	// 设置接收方邮箱
	em.To = []string{email}
	// 设置主题
	em.Subject = "This is a verification!"
	// 设置文件发送neir
	code := RandStringBytes(6)
	em.Text = []byte(fmt.Sprintf("This is verification code: %v . Expires after two minutes.", code))
	// 添加附件
	// em.AttachFile("./test.html")
	// 设置服务器相关配置
	err = em.Send("smtp.qq.com:25", smtp.PlainAuth("", "1577108595@qq.com", "izvpcfcmjjxlhjfh", "smtp.qq.com"))
	if err != nil {
		return "", err
	}
	return code, nil
}

func VerificationFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// 连接池+协程  运用连接池，创建 一个 有 5 个缓冲的通道，让 3 个协程去通道里面获取数据，然后发送邮件

// func Send_Code(email string) bool {
// 	log.SetFlags(log.Lshortfile | log.LstdFlags)
// 	// 创建有5 个缓冲的通道，数据类型是  *email.Email
// 	ch := make(chan *EM.Email, 5)
// 	// 连接池
// 	p, err := EM.NewPool(
// 		"stmp.qq.com:25",
// 		3, // 数量设置为三个
// 		smtp.PlainAuth("", "1577108595@qq.com", "izvpcfcmjjxlhjfh", "smtp.qq.com"),
// 	)
// 	if err != nil {
// 		log.Fatal("email.NewPool error: ", err)
// 	}
// 	// sync 控制同步
// 	var wg sync.WaitGroup
// 	wg.Add(3)
// 	for i := 0; i < 3; i++ {
// 		go func() {
// 			defer wg.Done()
// 			for e := range ch {
// 				// 超时时间 10 秒
// 				err := p.Send(e, 10*time.Second)
// 				if err != nil {
// 					log.Printf("p.Send error : %v , e = %v , i = %d\n", err, e, i)
// 				}
// 			}
// 		}()
// 	}
// 	for i := 0; i < 5; i++ {
// 		e := EM.NewEmail()
// 		// 设置发送邮件的基本信息
// 		e.From = "1577108595@qq.com"
// 		e.To = []string{email}
// 		e.Subject = "This verification Code!"
// 		code := RandStringBytes()
// 		e.Text = []byte(fmt.Sprintf("This is verification code: %v", code))
// 		ch <- e
// 		// 关闭通道
// 		close(ch)
// 		// 等待子协程退出
// 		wg.Wait()
// 		// log.Println("send successfully ... ")
// 	}
// 	return true
// }

// 连接redis设置过期时间
