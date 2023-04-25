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
	rep, err := http.Post(define.EmqxAddr+"/authentication/password_based%3Abuilt_in_database/users", "application/json", body)
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
	req.Header.Set("Authorization", "basic/bearer")
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
