package clients

import (
    "github.com/slaghuis/YC-Product/pkg/api"
    "github.com/slaghuis/YC-Product/pkg/services"
)

type identityClient struct {
    apiServer *api.APIServer
}

func NewIdentityClient(apiServer *api.APIServer) services.IdentityClient {
    return &identityClient{apiServer: apiServer}
}

func (c *identityClient) GetContext(token string) (*services.IdentityContext, error) {
    var out services.IdentityContext
    err := c.apiServer.APICall("GET",
        "/internal/accounts/context",
        token,
        nil, // no request body
        &out,
    )
    if err != nil {
        return nil, err
    }
    return &out, nil
}
