//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=cfg.yaml ../../../specs/petstore.1.0.0.yaml

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

// ServerInterface represents xall server handlers.
type ServerInterface interface {
	FindPets(w http.ResponseWriter, r *http.Request, params FindPetsParams)
	AddPet(w http.ResponseWriter, r *http.Request)
	DeletePet(w http.ResponseWriter, r *http.Request, id uuid.UUID)
	FindPetByID(w http.ResponseWriter, r *http.Request, id uuid.UUID)
}

type PetStore struct {
	Pets   map[string]Pet
	NextId uuid.UUID
	Lock   sync.Mutex
}

var _ ServerInterface = (*PetStore)(nil)

func NewPetStore() *PetStore {
	return &PetStore{
		Pets:   make(map[string]Pet),
		NextId: uuid.New(),
	}
}

// This function wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func sendPetstoreError(w http.ResponseWriter, code int, message string) {
	petErr := Error{
		Code:    int32(code),
		Message: message,
	}
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(petErr)
}

// Here, we implement all of the handlers in the ServerInterface
func (p *PetStore) FindPets(w http.ResponseWriter, r *http.Request, params FindPetsParams) {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	var result []Pet

	for _, pet := range p.Pets {
		if params.Tags != nil {
			// If we have tags,  filter pets by tag
			for _, t := range *params.Tags {
				if pet.Tag != nil && (*pet.Tag == t) {
					result = append(result, pet)
				}
			}
		} else {
			// Add all pets if we're not filtering
			result = append(result, pet)
		}

		if params.Limit != nil {
			l := int(*params.Limit)
			if len(result) >= l {
				// We're at the limit
				break
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (p *PetStore) AddPet(w http.ResponseWriter, r *http.Request) {
	// We expect a NewPet object in the request body.
	var newPet NewPet
	if err := json.NewDecoder(r.Body).Decode(&newPet); err != nil {
		sendPetstoreError(w, http.StatusBadRequest, "Invalid format for NewPet")
		return
	}

	// We now have a pet, let's add it to our "database".

	// We're always asynchronous, so lock unsafe operations below
	p.Lock.Lock()
	defer p.Lock.Unlock()

	// We handle pets, not NewPets, which have an additional ID field
	var pet Pet
	pet.Name = newPet.Name
	pet.Tag = newPet.Tag
	pet.Id = p.NextId
	p.NextId = uuid.New()

	// Insert into map
	p.Pets[pet.Id.String()] = pet

	// Now, we have to return the NewPet
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pet)
}

func (p *PetStore) FindPetByID(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	pet, found := p.Pets[id.String()]
	if !found {
		sendPetstoreError(w, http.StatusNotFound, fmt.Sprintf("Could not find pet with ID %d", id))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pet)
}

func (p *PetStore) DeletePet(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	_, found := p.Pets[id.String()]
	if !found {
		sendPetstoreError(w, http.StatusNotFound, fmt.Sprintf("Could not find pet with ID %d", id))
		return
	}
	delete(p.Pets, id.String())

	w.WriteHeader(http.StatusNoContent)
}
