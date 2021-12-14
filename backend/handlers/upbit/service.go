package upbit

import (
	"toTheMoon/backend/model"
	service2 "toTheMoon/backend/model/exchange/service"
)

// GetWalletStatus 입출금 현황
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token(JWT)
func (u *Upbit) GetWalletStatus() ([]*service2.Wallet, *model.Remaining, error) {
	api, e := GetApiInfo(FuncGetWalletStatus)
	if e != nil {
		return nil, nil, e
	}

	req, e := u.createRequest(api.Method, BaseURI+api.Url, nil, api.Section)
	if e != nil {
		return nil, nil, e
	}

	resp, e := u.do(req, api.Group)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	wallets, e := service2.WalletsFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return wallets, model.RemainingFromHeader(resp.Header), nil
}

// GetApiKeys API 키 리스트 조회
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token(JWT)
func (u *Upbit) GetApiKeys() ([]*service2.ApiKey, *model.Remaining, error) {
	api, e := GetApiInfo(FuncGetApiKeys)
	if e != nil {
		return nil, nil, e
	}

	req, e := u.createRequest(api.Method, BaseURI+api.Url, nil, api.Section)
	if e != nil {
		return nil, nil, e
	}

	resp, e := u.do(req, api.Group)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	apiKeys, e := service2.ApiKeysFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return apiKeys, model.RemainingFromHeader(resp.Header), nil
}
