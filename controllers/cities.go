package controllers

import (
	"context"
	"database/sql"
	"unsia/models"
	"unsia/pb/cities"
)

// City struct
type City struct{
	DB *sql.DB;
	cities.UnimplementedCitiesServiceServer
}
// GetCity function
func (s *City) GetCity(ctx context.Context, in *cities.Id) (*cities.City, error) {
	var cityModel models.City
	err := cityModel.Get(ctx, s.DB, in)	
	return &cityModel.Pb, err
}