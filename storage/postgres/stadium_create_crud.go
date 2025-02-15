package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	pb "stadion/genproto/stadium"
	"stadion/pkg/logger"
	"stadion/storage"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type StadiumRepository struct {
	Db  *sql.DB
	Log *slog.Logger
}

func NewStadiumRepository(db *sql.DB) storage.IStadiumStorage {
	return &StadiumRepository{Db: db, Log: logger.NewLogger()}
}

func (s *StadiumRepository) CreateStadium(ctx context.Context, req *pb.CreateStadiumRequest) (*pb.CreateStadiumResponse, error) {
	tx, err := s.Db.BeginTx(ctx, nil)
	if err != nil {
		s.Log.ErrorContext(ctx, fmt.Sprintf("error starting transaction: %v", err.Error()))
		return nil, err
	}

	query_stadium := `insert into stadium (
		id,user_id,name,location,address,phonenummer,
		price,length,width,situation,created_at,updated_at
	  ) values (
			$1, $2, $3, ST_SetSRID(ST_MakePoint($4, $5), 4326), $6, $7, $8, $9, $10, $11, $12, $13
		)`

	stadium_id := uuid.NewString()
	newtime := time.Now()

	_, err = tx.ExecContext(ctx, query_stadium, stadium_id, req.UserId, req.Name, req.Latitude, req.Longitude, req.Address,
		req.Phonenummer, req.Price, req.Length, req.Width, req.Situation, newtime, newtime)

	if err != nil {
		tx.Rollback()
		s.Log.ErrorContext(ctx, fmt.Sprintf("error adding property: %v", err.Error()))
		return nil, err
	}

	query_stadium_images := `insert into stadium_images (
		id, stadium_id, image_url, created_at, updated_at
	) values (
		$1, $2, $3, $4, $5)`

	for _, image := range req.ImageUrls {
		image_id := uuid.NewString()
		_, err = tx.ExecContext(ctx, query_stadium_images, image_id, stadium_id, image, newtime, newtime)
		if err != nil {
			tx.Rollback()
			s.Log.ErrorContext(ctx, fmt.Sprintf("error adding property images: %v", err.Error()))
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		s.Log.ErrorContext(ctx, fmt.Sprintf("error in commenting the transaction: %v", err.Error()))
		return nil, err
	}

	response := &pb.Stadium{
		Id:          stadium_id,
		UserId:      req.UserId,
		Name:        req.Name,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		Address:     req.Address,
		Phonenummer: req.Phonenummer,
		Price:       req.Price,
		Length:      req.Length,
		Width:       req.Width,
		Situation:   req.Situation,
		ImageUrls:   req.ImageUrls,
	}

	return &pb.CreateStadiumResponse{Stadium: response}, nil
}

func (s *StadiumRepository) UpdateStadium(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	tx, err := s.Db.BeginTx(ctx, nil)
	if err != nil {
		s.Log.ErrorContext(ctx, fmt.Sprintf("error starting transaction: %v", err.Error()))
		return nil, err
	}

	query_get_stadium := `SELECT id,user_id,name,ST_AsText(location),address,phonenummer,
									price,length,width,situation
						   FROM stadium WHERE id = $1 AND deleted_at = 0`

	query_get_stadium_images := `SELECT id, image_url FROM stadium_images WHERE stadium_id = $1 AND deleted_at = 0`

	var oldStadium struct {
		id          string
		user_id     string
		location    string
		name        string
		address     string
		phonenummer string
		price       float64
		length      float64
		width       float64
		situation   string
	}

	err = tx.QueryRowContext(ctx, query_get_stadium, req.Id).Scan(
		&oldStadium.id, &oldStadium.user_id, &oldStadium.name, &oldStadium.location, &oldStadium.address,
		&oldStadium.phonenummer, &oldStadium.price, &oldStadium.length, &oldStadium.width, &oldStadium.situation)
	if err != nil {
		tx.Rollback()
		s.Log.ErrorContext(ctx, fmt.Sprintf("error reading property: %v", err.Error()))
		return nil, err
	}

	if req.Name == "" || req.Name == oldStadium.name {
		req.Name = oldStadium.name
	}

	if req.Address == "" || req.Address == oldStadium.address {
		req.Address = oldStadium.address
	}

	if req.Phonenummer == "" || req.Phonenummer == oldStadium.phonenummer {
		req.Phonenummer = oldStadium.phonenummer
	}

	if req.Price == 0 || req.Price == float32(oldStadium.price) {
		req.Price = float32(oldStadium.price)
	}

	if req.Length == 0 || req.Length == float32(oldStadium.length) {
		req.Length = float32(oldStadium.length)
	}

	if req.Width == 0 || req.Width == float32(oldStadium.width) {
		req.Width = float32(oldStadium.width)
	}

	if req.Situation == "" || req.Situation == oldStadium.situation {
		req.Situation = oldStadium.situation
	}

	oldImages := make(map[string]string) // Map to store old image URLs with their IDs
	rows, err := tx.QueryContext(ctx, query_get_stadium_images, req.Id)
	if err != nil {
		tx.Rollback()
		s.Log.ErrorContext(ctx, fmt.Sprintf("error reading stadium images: %v", err.Error()))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id, imageURL string
		if err := rows.Scan(&id, &imageURL); err != nil {
			tx.Rollback()
			s.Log.ErrorContext(ctx, fmt.Sprintf("error scanning stadium image: %v", err.Error()))
			return nil, err
		}
		oldImages[imageURL] = id
	}
	var oldLat, oldLon float64
	fmt.Sscanf(oldStadium.location, "POINT(%f %f)", &oldLon, &oldLat)

	if req.Latitude == 0 || req.Longitude == 0 || (req.Latitude == float32(oldLat) && req.Longitude == float32(oldLon)) {
		// Keep old location if new values are empty or the same
		req.Latitude = float32(oldLat)
		req.Longitude = float32(oldLon)
	}

	query_update_stadium := `UPDATE stadium 
SET name = $1, 
    location = ST_SetSRID(ST_MakePoint($2, $3), 4326), 
    address = $4, 
    phonenummer = $5, 
    price = $6, 
    length = $7, 
    width = $8, 
    situation = $9, 
    updated_at = NOW() 
WHERE id = $10`

	_, err = tx.ExecContext(ctx, query_update_stadium, req.Name, req.Longitude, req.Latitude, req.Address,
		req.Phonenummer, req.Price, req.Length, req.Width, req.Situation, req.Id)
	if err != nil {
		tx.Rollback()
		s.Log.ErrorContext(ctx, fmt.Sprintf("error updating stadium: %v", err.Error()))
		return nil, err
	}

	// Process Images: Update existing ones if they differ
	for _, imageURL := range req.ImageUrls {
		if _, exists := oldImages[imageURL]; exists {
			// If the image exists and is the same, skip update
			delete(oldImages, imageURL)
		} else {
			for oldURL, imageID := range oldImages {
				if oldURL != imageURL {
					query_update_image := `UPDATE stadium_images SET image_url = $1, updated_at = NOW() WHERE id = $2`
					_, err = tx.ExecContext(ctx, query_update_image, imageURL, imageID)
					if err != nil {
						tx.Rollback()
						s.Log.ErrorContext(ctx, fmt.Sprintf("error updating image: %v", err.Error()))
						return nil, err
					}
					delete(oldImages, oldURL)
					break
				}
			}
		}
	}

	if err = tx.Commit(); err != nil {
		s.Log.ErrorContext(ctx, fmt.Sprintf("error committing transaction: %v", err.Error()))
		return nil, err
	}

	response := &pb.Stadium{
		Id:          req.Id,
		UserId:      req.UserId,
		Name:        req.Name,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		Address:     req.Address,
		Phonenummer: req.Phonenummer,
		Price:       req.Price,
		Length:      req.Length,
		Width:       req.Width,
		Situation:   req.Situation,
		ImageUrls:   req.ImageUrls,
	}

	return &pb.UpdateResponse{Stadium: response}, nil
}

func (s *StadiumRepository) GetStadium(ctx context.Context, req *pb.GetStadiumRequest) (*pb.GetStadiumResponse, error) {
	query_stadium := `select s.id, s.user_id, s.name, ST_X(s.location::geometry) as longitude, ST_Y(s.location::geometry) as latitude,
						s.address, s.phonenummer, s.price, s.length, s.width, s.situation,
						Array(select image_url from stadium_images where stadium_id = s.id and deleted_at = 0) as image_urls 
					from stadium s 
					where s.id = $1 and deleted_at = 0`

	var stadium pb.GetStadiumResponse
	var imageUrls []string

	err := s.Db.QueryRowContext(ctx, query_stadium, req.Id).Scan(&stadium.Stadium.Id, &stadium.Stadium.UserId, &stadium.Stadium.Name,
		&stadium.Stadium.Longitude, &stadium.Stadium.Latitude, &stadium.Stadium.Address, &stadium.Stadium.Phonenummer,
		&stadium.Stadium.Price, &stadium.Stadium.Length, &stadium.Stadium.Width, &stadium.Stadium.Situation, &imageUrls)

	if err != nil {
		s.Log.ErrorContext(ctx, fmt.Sprintf("error reading stadium: %v", err.Error()))
		return nil, err
	}

	stadium.Stadium.ImageUrls = imageUrls

	return &stadium, nil
}

func (s *StadiumRepository) GetStadiums(ctx context.Context, req *pb.GetStadiumsRequest) (*pb.GetStadiumsResponse, error) {
	query_stadium := `
		SELECT 
			s.id, 
			s.user_id, 
			s.name, 
			ST_X(s.location::geometry) AS longitude, 
			ST_Y(s.location::geometry) AS latitude,
			s.address, 
			s.phonenummer, 
			s.price, 
			s.length, 
			s.width, 
			s.situation,
			ARRAY(
				SELECT image_url 
				FROM stadium_images 
				WHERE stadium_id = s.id AND deleted_at = 0
			) AS image_urls 
		FROM stadium s 
		WHERE s.user_id = $1 AND s.deleted_at = 0`

	rows, err := s.Db.QueryContext(ctx, query_stadium, req.UserId)
	if err != nil {
		s.Log.ErrorContext(ctx, fmt.Sprintf("error querying stadiums: %v", err.Error()))
		return nil, err
	}
	defer rows.Close()

	stadiums := []*pb.Stadium{}
	for rows.Next() {
		var stadium pb.Stadium
		var imageUrls []string

		err = rows.Scan(
			&stadium.Id, &stadium.UserId, &stadium.Name, &stadium.Longitude, &stadium.Latitude,
			&stadium.Address, &stadium.Phonenummer, &stadium.Price, &stadium.Length, &stadium.Width, &stadium.Situation,
			pq.Array(&imageUrls),
		)
		if err != nil {
			s.Log.ErrorContext(ctx, fmt.Sprintf("error scanning stadium row: %v", err.Error()))
			return nil, err
		}

		stadium.ImageUrls = imageUrls
		stadiums = append(stadiums, &stadium)
	}

	if err = rows.Err(); err != nil {
		s.Log.ErrorContext(ctx, fmt.Sprintf("error iterating stadium rows: %v", err.Error()))
		return nil, err
	}

	return &pb.GetStadiumsResponse{Stadiums: stadiums}, nil
}

func (s *StadiumRepository) DeleteStadium(ctx context.Context, req *pb.DeleteStadiumRequest) (*pb.DeleteStadiumResponse, error) {
	query_delete_stadium := `UPDATE stadium_images SET deleted_at = NOW() WHERE stadium_id = $1`
	_, err := s.Db.ExecContext(ctx, query_delete_stadium, req.Id)
	if err != nil {
		s.Log.ErrorContext(ctx, fmt.Sprintf("error deleting stadium images: %v", err.Error()))
		return nil, err
	}

	query_delete_stadium = `UPDATE stadium SET deleted_at = NOW() WHERE id = $1`
	_, err = s.Db.ExecContext(ctx, query_delete_stadium, req.Id)
	if err != nil {
		s.Log.ErrorContext(ctx, fmt.Sprintf("error deleting stadium: %v", err.Error()))
		return nil, err
	}

	return &pb.DeleteStadiumResponse{}, nil
}

func (s *StadiumRepository) CreateOrderStadium(ctx context.Context, req *pb.CreateOrderStadiumRequest) (*pb.CreateOrderStadiumResponse, error) {
	const layoutISO = "2006-01-02"
	date, err := time.Parse(layoutISO, req.Date)
	if err != nil {
		s.Log.ErrorContext(ctx, fmt.Sprintf("invalid date format: %v", err.Error()))
		return nil, fmt.Errorf("invalid date format, expected YYYY-MM-DD")
	}

	query_order := `insert into order_stadium (
		id, user_id, stadium_id, start_time, end_time, price, status, date, created_at, updated_at
	  ) values (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
		)`

	order_id := uuid.NewString()
	newtime := time.Now()

	_, err = s.Db.ExecContext(ctx, query_order, order_id, req.UserId, req.StadiumId, req.StartTime, req.EndTime, req.Price, req.Status, date.Format(layoutISO), newtime, newtime)
	if err != nil {
		s.Log.ErrorContext(ctx, fmt.Sprintf("error adding order: %v", err.Error()))
		return nil, err
	}

	response := &pb.OrderStadium{
		Id:        order_id,
		UserId:    req.UserId,
		StadiumId: req.StadiumId,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Price:     req.Price,
		Status:    req.Status,
	}

	return &pb.CreateOrderStadiumResponse{OrderStadium: response}, nil
}

func (s *StadiumRepository) GetOrderStadiums(ctx context.Context, req *pb.GetOrderStadiumsRequest) (*pb.GetOrderStadiumsResponse, error) {
	query_order := `SELECT id, user_id, stadium_id, date, start_time, end_time, price, status 
					FROM order_stadium 
					WHERE user_id = $1 and deleted_at = 0`

	rows, err := s.Db.QueryContext(ctx, query_order, req.UserId)
	if err != nil {
		s.Log.ErrorContext(ctx, fmt.Sprintf("error querying orders: %v", err.Error()))
		return nil, err
	}
	defer rows.Close()

	var orders []*pb.OrderStadium
	for rows.Next() {
		var order pb.OrderStadium
		err := rows.Scan(&order.Id, &order.UserId, &order.StadiumId, &order.Date,
			&order.StartTime, &order.EndTime, &order.Price, &order.Status)
		if err != nil {
			s.Log.ErrorContext(ctx, fmt.Sprintf("error scanning order: %v", err.Error()))
			return nil, err
		}
		orders = append(orders, &order)
	}

	if err = rows.Err(); err != nil {
		s.Log.ErrorContext(ctx, fmt.Sprintf("rows iteration error: %v", err.Error()))
		return nil, err
	}

	orderRes := &pb.GetOrderStadiumsResponse{
		OrderStadiums: orders,
	}

	return orderRes, nil
}

func (s *StadiumRepository) GetOrderStadium(ctx context.Context, req *pb.GetOrderStadiumRequest) (*pb.GetOrderStadiumResponse, error) {
	query_order := `SELECT id, user_id, stadium_id, date, start_time, end_time, price, status 
					FROM order_stadium 
					WHERE id = $1 and deleted_at = 0`

	var order pb.OrderStadium
	err := s.Db.QueryRowContext(ctx, query_order, req.Id).Scan(&order.Id, &order.UserId, &order.StadiumId, &order.Date, &order.StartTime, &order.EndTime, &order.Price, &order.Status)
	if err != nil {
		s.Log.ErrorContext(ctx, fmt.Sprintf("error reading order: %v", err.Error()))
		return nil, err
	}

	return &pb.GetOrderStadiumResponse{OrderStadium: &order}, nil
}

func (s *StadiumRepository) UpdateOrderStadium(ctx context.Context, req *pb.UpdateOrderStadiumRequest) (*pb.UpdateOrderStadiumResponse, error) {
	query_order := `SELECT id, user_id, stadium_id, date, start_time, end_time, price, status 
					FROM order_stadium 
					WHERE id = $1 and deleted_at = 0`

	var oldStadiumOrder struct {
		ID         string
		User_id    string
		Stadium_id string
		Date       string
		Start_time string
		End_time   string
		Status     string
		Price      float64
	}

	err := s.Db.QueryRowContext(ctx, query_order, req.Id).Scan(&oldStadiumOrder.ID, &oldStadiumOrder.User_id, &oldStadiumOrder.Stadium_id, &oldStadiumOrder.Date, &oldStadiumOrder.Start_time, &oldStadiumOrder.End_time, &oldStadiumOrder.Price, &oldStadiumOrder.Status)
	if err != nil {
		s.Log.ErrorContext(ctx, fmt.Sprintf("error reading order: %v", err.Error()))
		return nil, err
	}

	if req.StartTime == "" || req.StartTime == oldStadiumOrder.Start_time {
		req.StartTime = oldStadiumOrder.Start_time
	}

	if req.EndTime == "" || req.EndTime == oldStadiumOrder.End_time {
		req.EndTime = oldStadiumOrder.End_time
	}

	if req.Price == 0 || req.Price == float32(oldStadiumOrder.Price) {
		req.Price = float32(oldStadiumOrder.Price)
	}

	if req.Status == "" || req.Status == oldStadiumOrder.Status {
		req.Status = oldStadiumOrder.Status
	}

	if req.Date == "" || req.Date == oldStadiumOrder.Date {
		req.Date = oldStadiumOrder.Date
	}

	query_order_update := `UPDATE order_stadium SET start_time = $1, end_time = $2, price = $3, status = $4, updated_at = NOW() WHERE id = $5`

	_, err = s.Db.ExecContext(ctx, query_order_update, req.StartTime, req.EndTime, req.Price, req.Status, req.Id)
	if err != nil {
		s.Log.ErrorContext(ctx, fmt.Sprintf("error updating order: %v", err.Error()))
		return nil, err
	}

	response := &pb.OrderStadium{
		Id:        req.Id,
		UserId:    req.UserId,
		StadiumId: req.StadiumId,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Price:     req.Price,
		Status:    req.Status,
	}

	return &pb.UpdateOrderStadiumResponse{OrderStadium: response}, nil
}

func (s StadiumRepository) DeleteOrderStadium(ctx context.Context, req *pb.DeleteOrderStadiumRequest) (*pb.DeleteOrderStadiumResponse, error) {
	query_delete_order := `UPDATE order_stadium SET deleted_at = NOW() WHERE id = $1`
	_, err := s.Db.ExecContext(ctx, query_delete_order, req.Id)
	if err != nil {
		s.Log.ErrorContext(ctx, fmt.Sprintf("error deleting order: %v", err.Error()))
		return nil, err
	}

	return &pb.DeleteOrderStadiumResponse{
		Message: "Order deleted successfully",
	}, nil
}

func (s StadiumRepository) GetDeletedOrderStadiums(ctx context.Context, req *pb.GetDeletedOrderStadiumsRequest) (*pb.GetDeletedOrderStadiumsResponse, error) {
	query_order := `SELECT id, user_id, stadium_id, date, start_time, end_time, price, status 
					FROM order_stadium 
					WHERE user_id = $1 and deleted_at != 0`

	rows, err := s.Db.QueryContext(ctx, query_order, req.UserId)
	if err != nil {
		s.Log.ErrorContext(ctx, fmt.Sprintf("error querying orders: %v", err.Error()))
		return nil, err
	}
	defer rows.Close()

	var orders []*pb.OrderStadium
	for rows.Next() {
		var order pb.OrderStadium
		err := rows.Scan(&order.Id, &order.UserId, &order.StadiumId, &order.Date,
			&order.StartTime, &order.EndTime, &order.Price, &order.Status)
		if err != nil {
			s.Log.ErrorContext(ctx, fmt.Sprintf("error scanning order: %v", err.Error()))
			return nil, err
		}
		orders = append(orders, &order)
	}

	if err = rows.Err(); err != nil {
		s.Log.ErrorContext(ctx, fmt.Sprintf("rows iteration error: %v", err.Error()))
		return nil, err
	}

	orderRes := &pb.GetDeletedOrderStadiumsResponse{
		OrderStadiums: orders,
	}

	return orderRes, nil
}

func (s StadiumRepository) GetAllStadium(ctx context.Context, req *pb.GetAllStadiumRequest) (*pb.GetAllStadiumResponse, error) {
	query_stadium := `
		SELECT 
			s.id, 
			s.user_id, 
			s.name, 
			ST_X(s.location::geometry) AS longitude, 
			ST_Y(s.location::geometry) AS latitude,
			s.address, 
			s.phonenummer, 
			s.price, 
			s.length, 
			s.width, 
			s.situation,
			ARRAY(
				SELECT image_url 
				FROM stadium_images 
				WHERE stadium_id = s.id AND deleted_at = 0
			) AS image_urls 
		FROM stadium s 
		WHERE s.deleted_at = 0
		ORDER BY s.created_at DESC
		LIMIT $1 OFFSET $2`

	limit := req.Limit
	if limit <= 0 {
		limit = 10 // Default limit
	}

	page := req.Page
	if page <= 0 {
		page = 1 // Default page
	}

	offset := (page - 1) * limit
	fmt.Println("Offset:", offset)
	fmt.Println("Limit:", limit)

	rows, err := s.Db.QueryContext(ctx, query_stadium, limit, offset)
	if err != nil {
		s.Log.ErrorContext(ctx, fmt.Sprintf("error querying stadiums: %v", err.Error()))
		return nil, err
	}
	defer rows.Close()

	var stadiums []*pb.Stadium

	for rows.Next() {
		var stadium pb.Stadium
		var imageUrls []string

		err = rows.Scan(
			&stadium.Id, &stadium.UserId, &stadium.Name, &stadium.Longitude, &stadium.Latitude,
			&stadium.Address, &stadium.Phonenummer, &stadium.Price, &stadium.Length, &stadium.Width, &stadium.Situation,
			pq.Array(&imageUrls),
		)
		if err != nil {
			s.Log.ErrorContext(ctx, fmt.Sprintf("error scanning stadium row: %v", err.Error()))
			return nil, err
		}

		stadium.ImageUrls = imageUrls
		stadiums = append(stadiums, &stadium)
	}

	if err = rows.Err(); err != nil {
		s.Log.ErrorContext(ctx, fmt.Sprintf("error iterating stadium rows: %v", err.Error()))
		return nil, err
	}

	// Debug ma'lumot chiqarish
	fmt.Printf("Returning %d stadiums\n", len(stadiums))

	return &pb.GetAllStadiumResponse{Stadiums: stadiums}, nil
}
