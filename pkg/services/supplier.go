package services

import (
    "context"
    "slices"

    "github.com/slaghuis/YC-Product/pkg/dto"
    "github.com/slaghuis/YC-Product/pkg/models"
    "github.com/slaghuis/YC-Product/pkg/repositories"
)

type supplierService struct {
    repo              *repositories.SupplierRepository
    preferenceClient  PreferenceClient
}

func NewSupplierService(repo *repositories.SupplierRepository, preference PreferenceClient) SupplierService {
    return &supplierService{
      repo: repo,
      preferenceClient: preference,
    }
}

func (s *supplierService) CreateSupplier(ctx context.Context, name, repName, phone, email string) (*models.Supplier, error) {
  item := &models.Supplier {
    Name:     name,
    RepName:  repName,
    Phone:    phone,
    Email:    email,
  }
  err := s.repo.SaveSupplier(item)
  return item, err
}

func (s *supplierService) UpdateSupplier(ctx context.Context, id, name, repName, logoPath, phone string, phoneVerified bool, email string, sms, whatsApp, emailNotify, appPush bool) (*models.Supplier, error) {
  supplierID, err := StrToUint(id)
  if err != nil {
      return nil, err
  }

  item := &models.Supplier {
    ID           : supplierID,
    Name         : name,
    RepName      : repName,
    LogoPath     : logoPath,
    Phone        : phone,
    PhoneVerified: phoneVerified,
    Email        : email,
    SMS          : sms,
    WhatsApp     : whatsApp,
    EmailNotify  : emailNotify,
    AppPush      : appPush,
  }
  err = s.repo.SaveSupplier(item)
  return item, err
}

func (s *supplierService) GetSupplier(ctx context.Context, id string) (*models.Supplier, error) {
  supplierID, err := StrToUint(id)
  if err != nil {
      return nil, err
  }
  return s.repo.GetSupplier(supplierID)
}

func (s *supplierService) GetSupplierBasic(ctx context.Context, id string) (*dto.SupplierBasic, error) {
  supplierID, err := StrToUint(id)
  if err != nil {
      return nil, err
  }
  return s.repo.GetSupplierBasic(supplierID)
}

func (s *supplierService) GetBasicSupplierList(ctx context.Context, token string) ([]dto.SupplierBasic, error) {
    // Get all the suppliers
    suppliers, err :=  s.repo.GetSupplierBasicList()
    if err != nil {
        return nil, err
    }

    // Call the preference repo to find the supplier preferencs and set the flag
    preferences, err := s.preferenceClient.GetUserPreferences(token)
    if err != nil {
        return nil, err
    }

    // Step through the preferences, and if there is a supplier preference, mark it in the suppliers list
    for _, f := range preferences.Filters {
    if f.FilterType == "supplier" {
        for i := range suppliers {
//            fmt.Printf("Testing supplier [%s] against %q\n", suppliers[i].Name, f.Values)
            suppliers[i].FilterSelected = slices.Contains(f.Values, suppliers[i].Name)
        }
        break
      }
    }

    return suppliers, err
}

func (s *supplierService) GetMerchantProductSupliers(page, limit, sort string) (*models.Pagination, error) {
  pagination := &models.Pagination{
    Page  : StrToInt(page, 1),
    Limit : StrToInt(limit, 20),
    Sort  : sort,
  }

  supplierList, err := s.repo.GetMerchantProductSupliers(pagination)

  pagination.Rows = supplierList
  pagination.Context = "suppliers"   //Idea here is to say "category=7" or "vegetarian"
  return pagination, err
}


func (s *supplierService) DeleteSupplier(ctx context.Context, id string) error {
  supplierId, err := StrToUint(id)
  if err != nil {
    return err
  }

  return s.repo.DeleteSupplier(supplierId)
}
