package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const customersBasePath = "admin/customers"
const customersResourceName = "customers"

// CustomerService is an interface for interfacing with the customers endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/customer
type CustomerService interface {
	List(interface{}) ([]Customer, error)
	Count(interface{}) (int, error)
	Get(uint64, interface{}) (*Customer, error)
	Search(interface{}) ([]Customer, error)
	Create(Customer) (*Customer, error)
	Update(Customer) (*Customer, error)
	Delete(uint64) error

	// MetafieldsService used for Customer resource to communicate with Metafields resource
	MetafieldsService
}

// CustomerServiceOp handles communication with the product related methods of
// the Shopify API.
type CustomerServiceOp struct {
	client *Client
}


func NewCustomerServiceOp(cl *Client)CustomerService{
	return &CustomerServiceOp{client:cl}
}

// Customer represents a Shopify customer
type Customer struct {
	ID                  uint64                `json:"id,omitempty"`
	Email               string             `json:"email,omitempty"`
	FirstName           string             `json:"first_name,omitempty"`
	LastName            string             `json:"last_name,omitempty"`
	State               string             `json:"state,omitempty"`
	Note                string             `json:"note,omitempty"`
	VerifiedEmail       bool               `json:"verified_email,omitempty"`
	MultipassIdentifier string             `json:"multipass_identifier,omitempty"`
	OrdersCount         int                `json:"orders_count,omitempty"`
	TaxExempt           bool               `json:"tax_exempt,omitempty"`
	TotalSpent          *decimal.Decimal   `json:"total_spent,omitempty"`
	Phone               string             `json:"phone,omitempty"`
	Tags                string             `json:"tags,omitempty"`
	LastOrderId         int                `json:"last_order_id,omitempty"`
	LastOrderName       string             `json:"last_order_name,omitempty"`
	AcceptsMarketing    bool               `json:"accepts_marketing,omitempty"`
	DefaultAddress      *CustomerAddress   `json:"default_address,omitempty"`
	Addresses           []*CustomerAddress `json:"addresses,omitempty"`
	CreatedAt           *time.Time         `json:"created_at,omitempty"`
	UpdatedAt           *time.Time         `json:"updated_at,omitempty"`
	Metafields          []Metafield        `json:"metafields,omitempty"`
}

// Represents the result from the customers/X.json endpoint
type CustomerResource struct {
	Customer *Customer `json:"customer"`
}

// Represents the result from the customers.json endpoint
type CustomersResource struct {
	Customers []Customer `json:"customers"`
}

// Represents the options available when searching for a customer
type CustomerSearchOptions struct {
	Page   int    `url:"page,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Fields string `url:"fields,omitempty"`
	Order  string `url:"order,omitempty"`
	Query  string `url:"query,omitempty"`
}

// List customers
func (s *CustomerServiceOp) List(options interface{}) ([]Customer, error) {
	path := fmt.Sprintf("%s.json", customersBasePath)
	resource := new(CustomersResource)
	err := s.client.Get(path, resource, options)
	return resource.Customers, err
}

// Count customers
func (s *CustomerServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", customersBasePath)
	return s.client.Count(path, options)
}

// Get customer
func (s *CustomerServiceOp) Get(customerID uint64, options interface{}) (*Customer, error) {
	path := fmt.Sprintf("%s/%v.json", customersBasePath, customerID)
	resource := new(CustomerResource)
	err := s.client.Get(path, resource, options)
	return resource.Customer, err
}

// Create a new customer
func (s *CustomerServiceOp) Create(customer Customer) (*Customer, error) {
	path := fmt.Sprintf("%s.json", customersBasePath)
	wrappedData := CustomerResource{Customer: &customer}
	resource := new(CustomerResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Customer, err
}

// Update an existing customer
func (s *CustomerServiceOp) Update(customer Customer) (*Customer, error) {
	path := fmt.Sprintf("%s/%d.json", customersBasePath, customer.ID)
	wrappedData := CustomerResource{Customer: &customer}
	resource := new(CustomerResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Customer, err
}

// Delete an existing customer
func (s *CustomerServiceOp) Delete(customerID uint64) error {
	path := fmt.Sprintf("%s/%d.json", customersBasePath, customerID)
	return s.client.Delete(path)
}

// Search customers
func (s *CustomerServiceOp) Search(options interface{}) ([]Customer, error) {
	path := fmt.Sprintf("%s/search.json", customersBasePath)
	resource := new(CustomersResource)
	err := s.client.Get(path, resource, options)
	return resource.Customers, err
}

// List metafields for a customer
func (s *CustomerServiceOp) ListMetafields(customerID uint64, options interface{}) ([]Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: customersResourceName, resourceID: customerID}
	return metafieldService.List(options)
}

// Count metafields for a customer
func (s *CustomerServiceOp) CountMetafields(customerID uint64, options interface{}) (int, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: customersResourceName, resourceID: customerID}
	return metafieldService.Count(options)
}

// Get individual metafield for a customer
func (s *CustomerServiceOp) GetMetafield(customerID uint64, metafieldID uint64, options interface{}) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: customersResourceName, resourceID: customerID}
	return metafieldService.Get(metafieldID, options)
}

// Create a new metafield for a customer
func (s *CustomerServiceOp) CreateMetafield(customerID uint64, metafield Metafield) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: customersResourceName, resourceID: customerID}
	return metafieldService.Create(metafield)
}

// Update an existing metafield for a customer
func (s *CustomerServiceOp) UpdateMetafield(customerID uint64, metafield Metafield) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: customersResourceName, resourceID: customerID}
	return metafieldService.Update(metafield)
}

// // Delete an existing metafield for a customer
func (s *CustomerServiceOp) DeleteMetafield(customerID uint64, metafieldID uint64) error {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: customersResourceName, resourceID: customerID}
	return metafieldService.Delete(metafieldID)
}
