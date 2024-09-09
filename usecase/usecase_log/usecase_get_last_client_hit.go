package usecase_log

func (usecase *LogUsecaseImpl) GetLastClientHitUsecase(ipClient string) (int, error) {
	lastClientHit, err := usecase.DatabaseRepository.GetLastClientHitRepository(ipClient)
	return lastClientHit.Jumlah, err
}
