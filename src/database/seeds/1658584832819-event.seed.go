package seed

func (s Seed) EventSeed1658584832819() error {
	for _, event := range events {
		err := s.db.Create(&event).Error

		if err != nil {
			return err
		}
	}
	return nil
}
