package emqx_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"graduate_design/define"
	"io"
	"net/http"
)

func CreateAuthUser(in *CreateAuthUserRequest) error {
	data, _ := json.Marshal(in)
	body := bytes.NewReader(data)
	req, _ := http.NewRequest(http.MethodPost, define.EmqxAddr+"/authentication/password_based%3Abuilt_in_database/users", body)
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(define.EmqxKey, define.EmqxSec)
	client := http.Client{}
	rep, err := client.Do(req)
	if err != nil {
		return err
	}
	resp := new(CreateAuthUserResponse)
	data, err = io.ReadAll(rep.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAuthUser(clientId string) error {
	req, _ := http.NewRequest(http.MethodDelete, define.EmqxAddr+"/authentication/password_based%3Abuilt_in_database/users/"+clientId, nil)
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(define.EmqxKey, define.EmqxSec)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	data, _ := io.ReadAll(resp.Body)
	if len(data) > 0 {
		return errors.New("invalid client")
	}
	return nil
}
