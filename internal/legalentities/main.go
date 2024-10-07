package legalentities

type LegalEService struct {
	service LegalERepository
}

func NewService(repo LegalERepository) *LegalEService {
	return &LegalEService{service: repo}
}

func (s *LegalEService) CreateBankAccount(account LegalEntities) (LegalEntities, error) {
	return s.service.CreateBankAccount(account)
}

func (s *LegalEService) GetAllBankAccounts() ([]LegalEntities, error) {
	return s.service.GetAllBankAccounts()
}

func (s *LegalEService) UpdateBankAccountByID(id int, account LegalEntities) (LegalEntities, error) {
	return s.service.UpdateBankAccount(id, account)
}

func (s *LegalEService) DeleteBankAccountByID(id int) error {
	return s.service.DeleteBankAccount(id)
}
