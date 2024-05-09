package storage

import (
	"secondLife/types"

	"github.com/google/uuid"
)

func (s *PostgresStore) CreateOrg(org *types.Org) error {
	query := `
        INSERT INTO orgs (user_id, orgName, Location, description)
        VALUES ($1, $2, $3, $4)
    `
	_, err := s.DB.Exec(
		query,
		org.UserID,
		org.OrgName,
		org.Location,
		org.Description,
	)
	if err != nil {
		return err
	}

	return nil
}
func (s *PostgresStore) GetUserName(userID uuid.UUID, orgName string) (string, error) {
	query := `
        SELECT u.userName
        FROM orgs o
        JOIN users u ON o.user_id = u.id
        WHERE o.user_id = $1 AND o.orgName = $2
    `

	var userName string
	err := s.DB.QueryRow(query, userID, orgName).Scan(&userName)
	if err != nil {
		return "", err
	}

	return userName, nil
}

func (s *PostgresStore) GetOrgByName(orgName string) (*types.Org, error) {
	query := `
        SELECT user_id, orgName, Location, description
        FROM orgs
        WHERE orgName = $1
    `

	var org types.Org
	err := s.DB.QueryRow(query, orgName).Scan(&org.UserID, &org.OrgName, &org.Location, &org.Description)
	if err != nil {
		return nil, err
	}

	return &org, nil
}
