package controllers

import (
	"encoding/json"
	"final/initializers"
	"final/middlewares"
	"final/models"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type SupportStruct struct {
	Message  string `bson:"message" json:"message"`
	Subject  string `bson:"subject" json:"subject"`
	FullName string `bson:"full_name" json:"full_name"`
	Email    string `bson:"email" json:"email"`
}

func SupportPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := initializers.InitTemplates()
		headerData := models.HeaderData{CurrentSite: "Support"}
		objectId, err := middlewares.ParseObjectIdFromJWT(r)
		if err == nil {
			headerData.ProfileID = objectId.Hex()
		}
		headData := models.HeadData{
			HeadTitle: "Support",
			StyleName: "Support",
		}
		data := models.PageData{
			HeaderData: headerData,
			HeadData:   headData,
		}

		err = tmpl.ExecuteTemplate(w, "Support.html", data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal(err)
		}
	} else {
		var req SupportStruct
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			initializers.LogError("decoding support request", err, nil)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(bson.M{"is_sent": false})
		}
		emailTemplate := `<!DOCTYPE html>
<html>
<head>
    <title>Support Request Response</title>
</head>
<body>
    <h2>Response to Your Support Request - %s</h2>
    <p>Dear %s,</p>
    <p>Thank you for reaching out to us with your concern: "<strong>%s</strong>"</p>
    <p>We have received your request regarding "<strong>%s</strong>", and we wanted to assure you that we are on it. Your satisfaction and experience with our services/products are of utmost importance to us.</p>
    <p>Here's what we're doing to address your concerns:</p>
    <ol>
        <li>Immediate Action: We are reviewing the details you've provided to understand the issue better.</li>
        <li>Investigation: Our support team is investigating the issue, and we aim to have an update or resolution as soon as possible.</li>
        <li>Follow-up: We will keep you updated on our progress and any steps you might need to take.</li>
    </ol>
    <p>In the meantime, if you have any more questions or additional information that might help us resolve your issue faster, please feel free to reply to this message or contact us directly at <a href="mailto:support@example.com">support@example.com</a>.</p>
    <p>We appreciate your patience and understanding as we work to resolve your issue. Rest assured, we are committed to providing you with the quality service you deserve.</p>
    <p>Thank you for being a valued part of our community.</p>
    <p>Warm regards,</p>
    <p>Your Support Team<br>support@example.com</p>
</body>
</html>
`
		message := fmt.Sprintf(emailTemplate, req.Subject, req.FullName, req.Message, req.Message)
		err = SendMessage(req.Email, req.Subject, message)
		if err != nil {
			initializers.LogError("sending message", err, nil)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(bson.M{"is_sent": false})
		}
		fmt.Println(req)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bson.M{"is_sent": true})
	}

}
