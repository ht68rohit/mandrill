package messaging

import (
	"encoding/json"
	"fmt"
	result "github.com/heaptracetechnology/microservice-mandrill/result"
	"github.com/mattbaird/gochimp"
	"net/http"
	"os"
)

type Email struct {
	Subject      string `json:"subject,omitempty"`
	Message      string `json:"message,omitempty"`
	From         string `json:"from,omitempty"`
	To           string `json:"to,omitempty"`
	TemplateName string `json:"templateName,omitempty"`
}

//Send Email
func Send(responseWriter http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)
	var param Email
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	apiKey := os.Getenv("API_KEY")

	mandrillApi, err := gochimp.NewMandrill(apiKey)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	templateName := param.TemplateName
	contentVar := gochimp.Var{"main", param.Message}
	content := []gochimp.Var{contentVar}

	_, err = mandrillApi.TemplateAdd(templateName, fmt.Sprintf("%s", contentVar.Content), true)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	defer mandrillApi.TemplateDelete(templateName)
	renderedTemplate, err := mandrillApi.TemplateRender(templateName, content, nil)

	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	recipients := []gochimp.Recipient{
		gochimp.Recipient{Email: param.To},
	}

	message := gochimp.Message{
		Html:      renderedTemplate,
		Subject:   param.Subject,
		FromEmail: param.From,
		To:        recipients,
	}

	sendResponse, respErr := mandrillApi.MessageSend(message, false)
	if respErr != nil {
		result.WriteErrorResponse(responseWriter, respErr)
		return
	}

	bytes, _ := json.Marshal(sendResponse)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
