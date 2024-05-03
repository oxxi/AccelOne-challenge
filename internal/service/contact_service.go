package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"sync"

	"github.com/oxxi/accel-one/internal/model"
	"github.com/oxxi/accel-one/pkg/dto"
)

// define a contact
type IContactService interface {
	// Interface for handling contact related service
	GetById(ctx context.Context, id int) (dto.ContactDto, error)
	Save(ctx context.Context, contactDto dto.ContactDto) (dto.ContactDto, error)
	Update(ctx context.Context, id int, contactDto dto.ContactDto) (dto.ContactDto, error)
	Delete(ctx context.Context, id int) error
}

type contactService struct {
	users  map[int]model.Contact // map to keep in memory
	mu     sync.RWMutex
	nextID int
}

// Delete implements IContactService.
func (c *contactService) Delete(ctx context.Context, id int) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.users[id]; !ok {
		return fmt.Errorf("contact with ID %d not found", id)
	}

	delete(c.users, id)
	return nil
}

// GetById implements IContactService.
func (c *contactService) GetById(ctx context.Context, id int) (dto.ContactDto, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if contact, ok := c.users[id]; ok {
		return contact.ToDto(), nil
	}

	return dto.ContactDto{}, fmt.Errorf("contact with ID %d not found", id)
}

// Save implements IContactService.
func (c *contactService) Save(ctx context.Context, contactDto dto.ContactDto) (dto.ContactDto, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	id := generateID()
	model := model.Contact{
		ID:    id,
		Name:  contactDto.Name,
		Email: contactDto.Email,
		Phone: contactDto.Phone,
	}
	c.users[id] = model
	return model.ToDto(), nil
}

// Update implements IContactService.
func (c *contactService) Update(ctx context.Context, id int, contactDto dto.ContactDto) (dto.ContactDto, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.users[id]; !ok {
		return dto.ContactDto{}, fmt.Errorf("contact with ID %d not found", id)
	}
	contact := model.Contact{
		ID:    id,
		Name:  contactDto.Name,
		Email: contactDto.Email,
		Phone: contactDto.Phone,
	}

	c.users[id] = contact

	return contact.ToDto(), nil
}

var once sync.Once
var instance *contactService

// create new instance
func NewContractService() IContactService {
	once.Do(func() {
		instance = &contactService{
			users:  make(map[int]model.Contact),
			nextID: 1,
		}
	})
	return instance
}

// generate ID
func generateID() int {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	id := int(b[0])<<24 | int(b[1])<<16 | int(b[2])<<8 | int(b[3])
	return id
}
