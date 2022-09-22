package main

import (
	"bytes"
	"fmt"
	"github.com/agambondan/wedding-be/app/lib"
	"gopkg.in/gomail.v2"
	"log"
	"text/template"
)

//func main() {
//	//result := make(map[string]interface{})
//	//response, err := lib.HttpRequest(nil, result).SetBody(`{"email": "asololejosbanget@gmail.com"}`).Post("https://wiza.co/bridge/verify_email")
//	//if err != nil {
//	//	log.Println(err)
//	//}
//	//log.Println(string(response.Body()), result, response.StatusCode())
//
//	//email := "agam.pro234@gmail.com"
//	//i := strings.LastIndexByte(email, '@')
//	//account := email[:i]
//	//host := email[i+1:]
//	//err := lib.ValidateHostAndUser(host, account, "agamgantengbanget@gmail.com")
//	//if smtpErr, ok := err.(checkmail.SmtpError); ok && err != nil {
//	//	fmt.Printf("\nCode: %s, Msg: %s\n", smtpErr.Code(), smtpErr)
//	//}
//
//	verifier := emailverifier.NewVerifier()
//	email := "asololejosbangetlah@gmail.com"
//
//	ret, err := verifier.EnableGravatarCheck().EnableSMTPCheck().EnableAutoUpdateDisposable().EnableDomainSuggest().Verify(email)
//	if err != nil {
//		fmt.Println("verify email address failed, error is: ", err)
//		return
//	}
//	if !ret.Syntax.Valid {
//		fmt.Println("email address syntax is invalid")
//		return
//	}
//	// ret.Reachable is how to check mail exists or not
//
//	mapes := make(map[string]interface{})
//	bytes, err := json.Marshal(ret)
//	if err != nil {
//		log.Println(err)
//	}
//	json.Unmarshal(bytes, &mapes)
//	fmt.Printf("email validation result :\n%v\n", mapes)
//	email = "agamgans234@gmail.com"
//
//	ret, err = verifier.EnableGravatarCheck().EnableSMTPCheck().EnableAutoUpdateDisposable().EnableDomainSuggest().Verify(email)
//	if err != nil {
//		fmt.Println("verify email address failed, error is: ", err)
//		return
//	}
//	if !ret.Syntax.Valid {
//		fmt.Println("email address syntax is invalid")
//		return
//	}
//	mapes = make(map[string]interface{})
//	bytes, err = json.Marshal(ret)
//	if err != nil {
//		log.Println(err)
//	}
//	json.Unmarshal(bytes, &mapes)
//	fmt.Printf("email validation result :\n%v\n", mapes)
//}

type BodylinkEmail struct {
	URL string
}

func main() {
	//templateData := BodylinkEmail{
	//	URL: "https://detik.id/",
	//}
	//type Menu struct {
	//	Name *string
	//}
	//menus := []Menu{
	//	Menu{Name: lib.Strptr("Profile")},
	//	Menu{Name: lib.Strptr("Dashboard")},
	//	Menu{Name: lib.Strptr("Setting")},
	//}
	to := "agamgans89@gmail.com"
	var activationCode = make(map[string]interface{})
	activationCode["activation_code"] = "123456"
	activationCode["email"] = to
	activationCode["redirect_url"] = fmt.Sprintf("http://localhost:3000/client/activation?token=%s&activation_code=%s", lib.GeneratePassword(128, 30, 30, 30), activationCode["activation_code"])
	activationCode["url_home"] = fmt.Sprintf("")
	// Activation Code
	SendEmailVerification(to, "../app/pages/views/mail/activation_code.html", activationCode)

	fmt.Println("Mail Sent!")
}

func SendEmail(to string, subject string, cc map[string]string, data interface{}, templateFile string) error {
	result, err := ParseTemplate(templateFile, data)
	if err != nil {
		panic(err)
	}
	m := gomail.NewMessage()
	m.SetHeader("From", lib.ConfigSenderName)
	m.SetHeader("To", to)
	for v, k := range cc {
		m.SetAddressHeader("Cc", v, k)
	}
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", result)
	// m.Attach(templateFile) // attach whatever you want
	log.Println(lib.ConfigSmtpHost, lib.ConfigSmtpPort, lib.ConfigAuthEmail, lib.ConfigAuthPassword)
	d := gomail.NewDialer(lib.ConfigSmtpHost, lib.ConfigSmtpPort, lib.ConfigAuthEmail, lib.ConfigAuthPassword)
	err = d.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		fmt.Println(err)
		return "", err
	}
	return buf.String(), nil
}

func SendEmailVerification(to, template string, data interface{}) {
	var err error
	subject := "sample email"
	err = SendEmail(to, subject, nil, data, template)
	if err == nil {
		fmt.Println("send email '" + subject + "' success")
	} else {
		fmt.Println(err)
	}
}
