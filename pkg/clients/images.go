package clients

import (
  "fmt"
  "github.com/slaghuis/YC-Product/pkg/api"
  "github.com/slaghuis/YC-Product/pkg/services"
)

type imagesClient struct {
    apiServer *api.APIServer
}

func NewImagesClient(apiServer *api.APIServer) services.ImagesClient {
    return &imagesClient{apiServer: apiServer}
}

// Calls the immaging microsevice
// /images/primary?reference_id="2"&reference_type="product"

func (c *imagesClient) GetPrimaryImage(reference_id, reference_type, token string) (services.ImagesResponse, error) {
    var out services.ImagesResponse
    url := fmt.Sprintf("/internal/images/primary?reference_id=%s&reference_type=%s",reference_id,reference_type)
    err := c.apiServer.APICall("GET",
        url,
        token,
        nil,   // request body
        &out,  // response body
    )
    if err != nil {
        return services.ImagesResponse{}, err
    }
    return out, nil
}
