package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"secondLife/utils"

	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
)

// GenerateRandomCode generates a random verification code

// SendVerificationCode sends a verification code to the specified email address
func SendVerificationCode(email, code string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "peaktew.technology@gmail.com")
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "Verification Code")

	// HTML email template
	htmlTemplate := `
	<!DOCTYPE html>
<html>

<head>
    <meta charset="UTF-8">
    <title>Email Verification</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f5f5f5;
            padding: 20px;
        }

        .container {
            background-color: #ffffff;
            border-radius: 5px;
            padding: 20px;
            max-width: 600px;
            margin: 0 auto;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        }

        h1 {
            color: #333333;
            margin-top: 0;
        }

        p {
            color: #555555;
            margin-bottom: 20px;
        }

        .code {
            font-size: 24px;
            font-weight: bold;
            color: #007bff;
        }

        .logo {
            display: block;
            max-width: 150px;
            margin: 0 auto 20px;
        }

        .slogan {
            color: #777777;
            font-size: 18px;
            font-style: italic;
        }
    </style>
</head>

<body>
    <div class="container">
        <img src="https://drive.google.com/uc?id=15zpKpdRIZVjTAUhLxonmZ9ueadmMIQyH" alt="Company Logo" class="logo">
        <h1>Email Verification</h1>
        <p>Thank you for signing up with <strong>PeakTew</strong>!</p>
        <p>Please use the following verification code to verify your email address:</p>
        <p class="code">` + code + `</p>
        <p class="slogan">"Together, unique in our own way"</p>
    </div>
</body>

</html>

	`
	mailer.SetBody("text/html", htmlTemplate)
	// mailer.AddAlternative("text/plain", "Your verification code is: "+code)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, "peaktew.technology@gmail.com", "iszadksaoikccoew")

	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}
	return nil
}

func (h *APIServer) handleSendVerificationCode(w http.ResponseWriter, r *http.Request) error {
	email := r.FormValue("email")
	if email == "" {
		http.Error(w, "Email address not provided", http.StatusBadRequest)
		return nil
	}

	code := utils.GenerateRandomCode()
	if err := SendVerificationCode(email, code); err != nil {
		http.Error(w, "Failed to send verification code", http.StatusInternalServerError)
		return fmt.Errorf("error sending verification code: %v", err)
	}

	response := struct {
		Message string `json:"message"`
		Code    string `json:"verification_code"`
	}{
		Message: "Verification code sent successfully",
		Code:    code,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return fmt.Errorf("error encoding JSON response: %v", err)
	}

	user, err := h.store.GetUserByEmail(email)
	if err != nil {
		return err
	}

	responseData := struct {
		UserID         uuid.UUID `json:"user_id"`
		ProfilePicture int       `json:"profile_picture"`
		UserName       string    `json:"user_name"`
	}{
		UserID:         user.UserID,
		ProfilePicture: user.ProfilePicture,
		UserName:       user.UserName,
	}

	return utils.WriteJSON(w, http.StatusOK, responseData)
}
