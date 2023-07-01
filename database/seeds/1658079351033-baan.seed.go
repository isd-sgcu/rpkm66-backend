package seed

func (s Seed) BaanSeed1658079351033() error {
	for _, b := range baans {
		err := s.db.Save(&b).Error

		if err != nil {
			return err
		}
	}
	return nil
}
