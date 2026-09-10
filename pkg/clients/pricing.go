package clients

import (
  "github.com/slaghuis/YC-Product/pkg/api"
  "github.com/slaghuis/YC-Product/pkg/services"
)


type pricingClient struct {
    apiServer *api.APIServer
}

func NewPricingClient(apiServer *api.APIServer) services.PricingClient {
    return &pricingClient{apiServer: apiServer}
}

func (c *pricingClient) Calculate(req services.PricingRequest, token string) (services.PricingResponse, error) {
    var out services.PricingResponse
    err := c.apiServer.APICall("POST",
        "/internal/pricing/calculate",
        token,
        req,   // request body
        &out,  // response body
    )
    if err != nil {
        return services.PricingResponse{}, err
    }
    return out, nil
}

func (c *pricingClient) CreatePrice(req services.SetRequest, token string) (services.PricingSKU, error) {
    var out services.PricingSKU
    err := c.apiServer.APICall("POST",
        "/internal/pricing/skus",
        token,
        req,   // request body
        &out,  // response body
    )
    if err != nil {
        return services.PricingSKU{}, err
    }
    return out, nil
}

func (c *pricingClient) UpdatePrice(req services.PricingUpdate, token string) (services.PricingSKU, error) {
  var out services.PricingSKU
  err := c.apiServer.APICall("PUT",
      "/internal/pricing/skus",
      token,
      req,   // request body
      &out,  // response body
  )
  if err != nil {
      return services.PricingSKU{}, err
  }
  return out, nil
}
