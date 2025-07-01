package models

import "time"

type ResourceMetadata struct {
	ID        string    `json:"id"`        
	Name      string    `json:"name"`      
	Type      string    `json:"type"`      
	Provider  string    `json:"provider"`  
	Region    string    `json:"region"`    
	Status    string    `json:"status"`    
	CreatedAt time.Time `json:"createdAt"` 
	UpdatedAt time.Time `json:"updatedAt"` 
}
