func (s *TripService) CreateTrip(req struct, userID string) (any, error) {
	tripMapper := &models.TripMapper{}
	trip := albumMapper.ToAlbum(req,userID)

	
	result, err := s.TripRepo.CreateTrip(trip)
	if err != nil {
		return nil, err
	}

	return result, nil
}