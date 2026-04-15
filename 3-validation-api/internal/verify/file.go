package verify

import (
	"encoding/json"
	"os"
)

const EMAIL_DATA = "emailData.json"

type EmailData struct {
	Email string
	Hash  string
}

func SaveEmailData(emailData EmailData) error {

	emailDataList, err := getEmailData()

	if err != nil {
		return err
	}

	for index, emailDataItem := range emailDataList {
		if emailDataItem.Email == emailData.Email {
			emailDataList[index].Hash = emailData.Hash
			return nil
		}
	}

	emailDataList = append(emailDataList, emailData)

	return writeEmailData(emailDataList)
}

func ReadEmailData(hash string) (bool, error) {
	emailDataList, err := getEmailData()
	isFound := false

	if err != nil {
		return false, err
	}

	var updatedEmailDataList []EmailData

	for _, emailDataItem := range emailDataList {
		if emailDataItem.Hash == hash {
			isFound = true
			continue
		}
		updatedEmailDataList = append(updatedEmailDataList, emailDataItem)
	}

	return isFound, writeEmailData(updatedEmailDataList)
}

func getEmailData() ([]EmailData, error) {
	content, err := os.ReadFile(EMAIL_DATA)
	if err != nil {
		if os.IsNotExist(err) {
			return []EmailData{}, nil
		}
		return nil, err
	}

	var emailDataList []EmailData

	unmarshalErr := json.Unmarshal(content, &emailDataList)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return emailDataList, nil
}

func writeEmailData(emailDataList []EmailData) error {
	updatedContent, err := json.MarshalIndent(emailDataList, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(EMAIL_DATA, updatedContent, 0644)
}
