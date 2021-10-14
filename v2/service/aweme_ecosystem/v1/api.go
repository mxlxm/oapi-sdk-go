// Package aweme_ecosystem code generated by lark suite oapi sdk gen
package aweme_ecosystem

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v2"
)

type AwemeEcosystemService struct {
	AwemeUsers *awemeUsers
}

func New(app *lark.App) *AwemeEcosystemService {
	a := &AwemeEcosystemService{}
	a.AwemeUsers = &awemeUsers{app: app}
	return a
}

type awemeUsers struct {
	app *lark.App
}

func (a *awemeUsers) GetBindInfo(ctx context.Context, req *AwemeUserGetBindInfoReq, options ...lark.RequestOptionFunc) (*AwemeUserGetBindInfoResp, error) {
	rawResp, err := a.app.SendRequestWithAccessTokenTypes(ctx, http.MethodPost,
		"/open-apis/aweme_ecosystem/v1/aweme_users/get_bind_info", []lark.AccessTokenType{lark.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	resp := &AwemeUserGetBindInfoResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
