package api

import (
	"github.com/gin-gonic/gin"
	mail "github.com/xhit/go-simple-mail/v2"
	"golang.org/x/net/html"
	"net/http"
	"selfletter-backend/config"
	"selfletter-backend/db"
	"strings"
)

type adminSendRequest struct {
	Key     string `json:"key" binding:"required"`
	Topic   string `json:"topic" binding:"required"`
	Message string `json:"message" binding:"required"`
	Subject string `json:"subject" binding:"required"`
}

type adminSendResponse struct {
	Errors        []string `json:"errors"`
	CriticalError string   `json:"criticalError"`
}

func AdminSend(c *gin.Context) {
	cfg := config.GetConfig()
	var request adminSendRequest
	var errors []string
	dbHandle := db.GetDatabaseHandle()

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, adminSendResponse{
			Errors:        nil,
			CriticalError: "bad json",
		})
		return
	}

	if dbHandle.First(&db.AdminKey{}, "key = ?", request.Key).RowsAffected == 0 {
		c.JSON(http.StatusForbidden, adminSendResponse{
			Errors:        nil,
			CriticalError: "invalid admin key",
		})
		return
	}

	_, err = html.Parse(strings.NewReader(request.Message))
	if err != nil {
		c.JSON(http.StatusBadRequest, adminSendResponse{
			Errors:        nil,
			CriticalError: "bad message html",
		})
		return
	}

	var users []db.User
	var usersTopicsRel []db.UsersTopicsRel
	err = dbHandle.Where("topic = ?", request.Topic).Find(&usersTopicsRel).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, adminSendResponse{
			Errors:        nil,
			CriticalError: "database error",
		})
		return
	}
	for _, userTopicRel := range usersTopicsRel {
		var user db.User
		if dbHandle.Where("email = ?", userTopicRel.Email).Find(&user).RowsAffected == 0 {
			errors = append(errors, "user not found: "+userTopicRel.Email)
			continue
		}
		users = append(users, user)
	}

	messageBodyEnclosingIndex := strings.LastIndex(request.Message, "</body>")
	messageBeforeBodyEnclosing := request.Message[:messageBodyEnclosingIndex]
	messageStartingFromBodyEnclosing := request.Message[messageBodyEnclosingIndex:]

	for _, user := range users {
		mailServer := mail.NewSMTPClient()
		mailServer.Host = cfg.Email.Auth.Host
		mailServer.Port = cfg.Email.Auth.Port
		mailServer.Username = cfg.Email.Auth.Username
		mailServer.Password = cfg.Email.Auth.Password
		mailEncryption := cfg.Email.Auth.Encryption
		switch mailEncryption {
		case "SSL/TLS":
			mailServer.Encryption = mail.EncryptionSSLTLS
		case "TLS":
			mailServer.Encryption = mail.EncryptionTLS
		case "STARTTLS":
			mailServer.Encryption = mail.EncryptionSTARTTLS
		case "SSL":
			mailServer.Encryption = mail.EncryptionSSL
		case "None":
			mailServer.Encryption = mail.EncryptionNone
		default:
			c.JSON(http.StatusInternalServerError, adminSendResponse{
				Errors:        nil,
				CriticalError: "mail server auth encryption set incorrectly",
			})
			return
		}

		smtpClient, err := mailServer.Connect()
		if err != nil {
			errors = append(errors, err.Error())
		}

		// todo: change to frontend
		unsubscribeAddress := "http://" + cfg.Domain + cfg.UrlPrefix + "/unsubscribe?token=" + user.Token
		unsubscribeString := "<a href=\"" + unsubscribeAddress + "\" target=\"_blank\">Unsubscribe</a>\n</body>"

		messageEnd := strings.Replace(messageStartingFromBodyEnclosing, "</body>", unsubscribeString, 1)
		message := messageBeforeBodyEnclosing + messageEnd

		email := mail.NewMSG()
		email.
			SetFrom(cfg.Email.From).
			AddTo(user.Email).
			SetSubject(request.Subject).
			SetListUnsubscribe(unsubscribeAddress).
			SetBody(mail.TextHTML, message)

		err = email.Send(smtpClient)
		if err != nil {
			errors = append(errors, err.Error())
		}
	}

	if len(errors) != 0 {
		c.JSON(http.StatusMultiStatus, adminSendResponse{
			Errors:        errors,
			CriticalError: "",
		})
		return
	}

	c.JSON(http.StatusOK, adminSendResponse{
		Errors:        nil,
		CriticalError: "",
	})
}
