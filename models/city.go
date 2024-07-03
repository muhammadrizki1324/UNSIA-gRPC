package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"unsia/pb/cities"
)

type City struct {
    Log *log.Logger
    Pb cities.City
}
//Get City
func (u *City) Get(ctx context.Context, db *sql.DB, in *cities.Id) error {
    query := "SELECT id, name FROM cities WHERE id = $1"
    err := db.QueryRowContext(ctx, query, in.Id).Scan(&u.Pb.Id, &u.Pb.Name)
    if err != nil {
        u.Log.Println("ERROR SELECT CITIES", err)
        return err
    }
    return nil
}
//Create City
func (u *City) Create(ctx context.Context, db *sql.DB, in *cities.CityInput) error {
    query := `INSERT INTO cities (name) VALUES ($1) RETURNING id;`
    stmt, err := db.PrepareContext(ctx, query)
    if err != nil {
        return err
    }
    
    err = stmt.QueryRowContext(ctx, in.Name).Scan(&u.Pb.Id)
    if err != nil {
        return err
    }
    return nil
}
//Update City
func (u *City) Update(ctx context.Context, db *sql.DB, in *cities.City ) error {
    query := `UPDATE cities SET name = $2 WHERE id = $1`
    stmt, err := db.PrepareContext(ctx, query)
    if err != nil {
        return err
    }

    _, err = stmt.ExecContext(ctx,  in.GetId(),in.GetName())
    if err != nil {
        return  err
    }

    return nil
}
//Delete City
func (u *City) Delete(ctx context.Context, db *sql.DB, in *cities.Id) error {
    query := `DELETE FROM cities WHERE id = $1;`
    stmt, err := db.PrepareContext(ctx, query)
    if err != nil {
        return err
    }
    
    rs, err := stmt.ExecContext(ctx, in.Id) 
    if err != nil {
        return err
    }

    affected, err := rs.RowsAffected()
    if err != nil {
        return err
    }

    if affected == 0 {
        return fmt.Errorf("DATA NOT FOUND")
    }
    return nil
}
