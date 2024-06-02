package data

import "time"

type ListingPickupPoint struct {
	Id        uint32
	Name      string
	Altitude  float64
	Longitude float64
}

type Listing struct {
	Id              uint32
	UserId          uint32
	Title           string
	Description     string
	UserPickupPoint ListingPickupPoint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (d *Dal) CreateListing(userId uint32, title, description string, pickupPointId uint32) error {
	_, err := d.DB.Exec(
		`INSERT INTO listings (user_id, title, description, user_pickup_point_id)`, userId, title, description, pickupPointId,
	)
	return err
}

func (d *Dal) DeleteListing(userId, listingId uint32) error {
	_, err := d.DB.Exec(
		`DELETE FROM listings WHERE user_id = $1 AND id = $2`, userId, listingId)
	return err
}

func (d *Dal) UpdateListing(userId, listingId uint32, title, description string, userPickupPoint uint32) error {
	_, err := d.DB.Exec(
		`UPDATE listings SET title = $1, description = $2, user_pickup_point_id = $3 WHERE user_id = $4 AND id = $5`, title, description, userPickupPoint, userId, listingId,
	)
	return err
}

func (d *Dal) FindUserItems(userId uint32) ([]Listing, error) {
	var items []Listing
	rows, err := d.DB.Query(
		`SELECT l.id, l.user_id, l.title, l.description, p.name, p.altitude, p.longitude, l.created_at, l.updated_at
		 JOIN user_pickup_point p ON l.user_pickup_point_id = p.id
		 FROM listings l WHERE l.user_id = $1`,
		userId,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var item Listing
		item.UserPickupPoint = ListingPickupPoint{}
		err = rows.Scan(
			&item.Id,
			&item.UserId,
			&item.Title,
			&item.Description,
			&item.UserPickupPoint.Name,
			&item.UserPickupPoint.Altitude,
			&item.UserPickupPoint.Longitude,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
