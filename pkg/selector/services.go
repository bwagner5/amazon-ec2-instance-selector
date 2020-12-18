package selector

import (
	"fmt"
	"strings"

	"github.com/imdario/mergo"
)

// Service is used to write custom service filter transforms
type Service interface {
	Filters(version string) (Filters, error)
}

// ServiceFiltersFn is the func type definition for the Service interface
type ServiceFiltersFn func(version string) (Filters, error)

// Filters implements the Service interface on ServiceFiltersFn
// This allows any ServiceFiltersFn to be passed into funcs accepting the Service interface
func (fn ServiceFiltersFn) Filters(version string) (Filters, error) {
	return fn(version)
}

// ServiceRegistry is used to register service filter transforms
type ServiceRegistry struct {
	services map[string]*Service
}

// NewRegistry creates a new instance of a ServiceRegistry
func NewRegistry() ServiceRegistry {
	return ServiceRegistry{
		services: make(map[string]*Service),
	}
}

// Register takes a service name and Service implementation that will be executed on an ExecuteTransforms call
func (sr ServiceRegistry) Register(name string, service Service) {
	if sr.services == nil {
		sr.services = make(map[string]*Service)
	}
	sr.services[name] = &service
}

// RegisterAWSServices registers the built-in AWS service filter transforms
func (sr ServiceRegistry) RegisterAWSServices() {
	sr.Register("eks", &EKS{
		TTLInSeconds: 600,
	})
	sr.Register("emr", &EMR{})
}

// ExecuteTransforms will execute the ServiceRegistry's registered service filter transforms
// Filters.Service will be parsed as <service-name>-<version> and passed to Service.Filters
func (sr ServiceRegistry) ExecuteTransforms(filters Filters) (Filters, error) {
	if filters.Service == nil {
		return filters, nil
	}
	serviceAndVersion := strings.ToLower(*filters.Service)
	versionParts := strings.Split(serviceAndVersion, "-")
	serviceName := versionParts[0]
	version := ""
	if len(versionParts) == 2 {
		version = versionParts[1]
	}
	service, ok := sr.services[serviceName]
	if !ok {
		return filters, fmt.Errorf("Service %s is not registered", serviceName)
	}

	serviceFilters, err := (*service).Filters(version)
	if err != nil {
		return filters, err
	}
	if err := mergo.Merge(&filters, serviceFilters); err != nil {
		return filters, err
	}
	return filters, nil
}
