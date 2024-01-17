package basic

func (s *service) Run() {
	s.run()
}

func (s *service) run() {
	s.logger.Debug("Start service")
	defer s.logger.Debug("End service")

	for i := 0; i < len(s.functions); i++ {
		s.functions[i]()
	}
}
