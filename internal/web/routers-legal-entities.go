package web

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/krisch/crm-backend/internal/legalentities"
	"github.com/krisch/crm-backend/internal/web/ofederation"
)

func (a *Web) DeleteLegalEntitiesId(ctx context.Context, request ofederation.DeleteLegalEntitiesIdRequestObject) (ofederation.DeleteLegalEntitiesIdResponseObject, error) {
	id, err := strconv.Atoi(request.Id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}

	err = a.app.LegalEntitiesService.DeleteBankAccountByID(id)
	if err != nil {
		if err.Error() == "not found" {
			return ofederation.DeleteLegalEntitiesId404Response{}, nil
		}
		return nil, err
	}
	return ofederation.DeleteLegalEntitiesId204Response{}, nil
}

func (a *Web) GetLegalEntities(ctx context.Context, request ofederation.GetLegalEntitiesRequestObject) (ofederation.GetLegalEntitiesResponseObject, error) {
	entities, err := a.app.LegalEntitiesService.GetAllBankAccounts()
	if err != nil {
		return nil, err
	}

	response := make([]ofederation.BankAccountDTO, len(entities))
	for i, entity := range entities {
		id := fmt.Sprint(entity.ID)
		accountNumber := entity.RBill
		bankName := entity.Bank
		accountHolderName := entity.Comment
		address := entity.Address
		bik := entity.BIK
		kBill := entity.KBill
		rBill := entity.RBill
		currency := entity.Currency
		comment := entity.Comment

		response[i] = ofederation.BankAccountDTO{
			Id:                &id,
			AccountNumber:     &accountNumber,
			BankName:          &bankName,
			AccountHolderName: &accountHolderName,
			Address:           &address,
			Bik:               &bik,
			KBill:             &kBill,
			RBill:             &rBill,
			Currency:          &currency,
			Comment:           &comment,
		}
	}

	return ofederation.GetLegalEntities200JSONResponse(response), nil
}

func (a *Web) PostLegalEntities(ctx context.Context, request ofederation.PostLegalEntitiesRequestObject) (ofederation.PostLegalEntitiesResponseObject, error) {
	if request.Body == nil {
		log.Println("Request body is nil")
		return nil, errors.New("request body cannot be nil")
	}

	messageRequest := request.Body

	log.Printf("Received request: %+v", messageRequest)

	// Initialize a new legal entity
	newEntity := legalentities.LegalEntities{}
	var validationErrors []string 

	if messageRequest.Bik != nil {
		newEntity.BIK = *messageRequest.Bik
	} else {
		log.Println("Bik is nil")
		validationErrors = append(validationErrors, "Bik is required")
	}

	if messageRequest.BankName != nil {
		newEntity.Bank = *messageRequest.BankName
	} else {
		log.Println("BankName is nil")
		validationErrors = append(validationErrors, "BankName is required")
	}

	if messageRequest.Address != nil {
		newEntity.Address = *messageRequest.Address
	} else {
		log.Println("Address is nil")
		validationErrors = append(validationErrors, "Address is required")
	}

	if messageRequest.KBill != nil {
		newEntity.KBill = *messageRequest.KBill
	} else {
		log.Println("KBill is nil")
		validationErrors = append(validationErrors, "KBill is optional but recommended")
	}

	if messageRequest.RBill != nil {
		newEntity.RBill = *messageRequest.RBill
	} else {
		log.Println("RBill is nil")
		validationErrors = append(validationErrors, "RBill is optional but recommended")
	}

	if messageRequest.Currency != nil {
		newEntity.Currency = *messageRequest.Currency
	} else {
		log.Println("Currency is nil")
		validationErrors = append(validationErrors, "Currency is optional but recommended")
	}

	if messageRequest.Comment != nil {
		newEntity.Comment = *messageRequest.Comment
	} else {
		log.Println("Comment is nil")
	}

	if len(validationErrors) > 0 {
		return nil, fmt.Errorf("validation errors: %s", strings.Join(validationErrors, ", "))
	}

	log.Printf("Creating new entity: %+v", newEntity)

	_, err := a.app.LegalEntitiesService.CreateBankAccount(newEntity)
	if err != nil {
		log.Printf("Error creating bank account: %v", err)
		return nil, err
	}

	response := ofederation.PostLegalEntities201JSONResponse{
		Uuid: uuid.New(),
	}

	log.Printf("Response: %+v", response)

	return response, nil
}

func (a *Web) PatchLegalEntitiesId(ctx context.Context, request ofederation.PatchLegalEntitiesIdRequestObject) (ofederation.PatchLegalEntitiesIdResponseObject, error) {
	id, err := strconv.Atoi(request.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	entityToUpdate := legalentities.LegalEntities{}

	if request.Body.BankName != nil {
		entityToUpdate.Bank = *request.Body.BankName
	}
	if request.Body.Bik != nil {
		entityToUpdate.BIK = *request.Body.Bik
	}
	if request.Body.Address != nil {
		entityToUpdate.Address = *request.Body.Address
	}
	if request.Body.KBill != nil {
		entityToUpdate.KBill = *request.Body.KBill
	}
	if request.Body.RBill != nil {
		entityToUpdate.RBill = *request.Body.RBill
	}
	if request.Body.Currency != nil {
		entityToUpdate.Currency = *request.Body.Currency
	}
	if request.Body.Comment != nil {
		entityToUpdate.Comment = *request.Body.Comment
	}

	updatedEntity, err := a.app.LegalEntitiesService.UpdateBankAccountByID(id, entityToUpdate)
	if err != nil {
		if err.Error() == "not found" {
			return ofederation.PatchLegalEntitiesId404Response{}, nil
		}
		return nil, err
	}

	updatedEntityID := fmt.Sprint(updatedEntity.ID)

	response := ofederation.PatchLegalEntitiesId200JSONResponse{
		Id:       &updatedEntityID,
		BankName: &updatedEntity.Bank,
		Bik:      &updatedEntity.BIK,
		Address:  &updatedEntity.Address,
		KBill:    &updatedEntity.KBill,
		RBill:    &updatedEntity.RBill,
		Currency: &updatedEntity.Currency,
		Comment:  &updatedEntity.Comment,
	}

	return response, nil
}
