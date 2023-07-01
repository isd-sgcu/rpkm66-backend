package seed

func (s Seed) EventSeed1658584832819() error {
	return s.db.Save(&events).Error
}
