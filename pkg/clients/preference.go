package clients

import (
  "github.com/slaghuis/YC-Product/pkg/api"
  "github.com/slaghuis/YC-Product/pkg/dto"
  "github.com/slaghuis/YC-Product/pkg/services"
)

type preferenceClient struct {
    apiServer *api.APIServer
}

func NewPreferenceClient(apiServer *api.APIServer) services.PreferenceClient {
    return &preferenceClient{apiServer: apiServer}
}

// Calls the preference microsevice
func (c *preferenceClient) GetUserPreferences(token string) (*dto.Preferences, error) {
  var out dto.Preferences
  err := c.apiServer.APICall("GET",
      "/internal/preference/filters",
      token,
      nil,   // request body
      &out,  // response body
  )
  if err != nil {
      return nil, err
  }
  return &out, nil
}


func (c *preferenceClient) GetUserSortOrder(token string) (*dto.SortOrder, error) {
  var out dto.SortOrder
  err := c.apiServer.APICall("GET",
      "/internal/preference/sort_order",
      token,
      nil,   // request body
      &out,  // response body
  )
  if err != nil {
      return nil, err
  }
  return &out, nil
}
