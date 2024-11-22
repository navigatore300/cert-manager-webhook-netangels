// webhook.go

package main

import (
	"context"

	"github.com/cert-manager/webhook-example/api/v1alpha1"
)

type Webhook struct {
	NetangelsAPI *NetangelsAPI
}

func (w *Webhook) Present(ctx context.Context, chal v1alpha1.ChallengeRequest) (*v1alpha1.ChallengeResponse, error) {
	// Обработка запроса на создание или обновление сертификата
	domain := chal.GetIdentifier().GetDNSName()
	recordType := "TXT"
	recordValue := chal.GetValidationRecord()[0].GetBody()

	err := w.NetangelsAPI.CreateDNSRecord(domain, recordType, recordValue)
	if err != nil {
		return nil, err
	}

	token, err := w.NetangelsAPI.GetToken(domain)
	if err != nil {
		return nil, err
	}

	return &v1alpha1.ChallengeResponse{
		Type:  "dns-01",
		Token: token,
	}, nil
}

func (w *Webhook) CleanUp(ctx context.Context, chal v1alpha1.ChallengeRequest) (*v1alpha1.ChallengeResponse, error) {
	// Обработка запроса на удаление сертификата
	domain := chal.GetIdentifier().GetDNSName()
	recordType := "TXT"
	recordValue := chal.GetValidationRecord()[0].GetBody()

	err := w.NetangelsAPI.DeleteDNSRecord(domain, recordType, recordValue)
	if err != nil {
		return nil, err
	}

	return &v1alpha1.ChallengeResponse{}, nil
}
