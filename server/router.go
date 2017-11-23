package server

func (s *JTServer) loadRoutes() {
	s.Handle(cmdAuth, auth, mArgsEqual(2))

	if password == "" {
		// server routes
		s.Handle(cmdPing, ping, mArgsEqual(1))
		s.Handle(cmdSave, save, mArgsEqual(1))
		// keys routes
		s.Handle(cmdDel, del, mArgsEqual(2))
		s.Handle(cmdRename, rename, mArgsEqual(3))
		s.Handle(cmdTTL, ttl, mArgsEqual(2))
		s.Handle(cmdPersist, persist, mArgsEqual(2))
		s.Handle(cmdExpire, expire, mArgsEqual(3))
		s.Handle(cmdType, hType, mArgsEqual(2))
		s.Handle(cmdKeys, keys, mArgsEqual(2))
		s.Handle(cmdExists, exists, mArgsMoreThan(2))
		// string routes
		s.Handle(cmdSet, set, mArgsEqual(3))
		s.Handle(cmdGet, get, mArgsEqual(2))
		s.Handle(cmdAppend, hAppend, mArgsEqual(3))
		s.Handle(cmdGetSet, getset, mArgsEqual(3))
		s.Handle(cmdStrlen, strlen, mArgsEqual(2))
		s.Handle(cmdIncr, incr, mArgsEqual(2))
		s.Handle(cmdIncrBy, incrBy, mArgsEqual(3))
		// list routes
		s.Handle(cmdLPush, lpush, mArgsEqual(3))
		s.Handle(cmdRPush, rpush, mArgsEqual(3))
		s.Handle(cmdLPop, lpop, mArgsEqual(2))
		s.Handle(cmdRPop, rpop, mArgsEqual(2))
		s.Handle(cmdLRem, lrem, mArgsEqual(4))
		s.Handle(cmdLIndex, lindex, mArgsEqual(3))
		s.Handle(cmdLRange, lrange, mArgsEqual(4))
		s.Handle(cmdLLen, llen, mArgsEqual(2))
		// dict routes
		s.Handle(cmdDSet, dset, mArgsMoreThan(4), mArgsCountEven())
		s.Handle(cmdDGet, dget, mArgsMoreThan(3))
		s.Handle(cmdDDel, ddel, mArgsEqual(3))
		s.Handle(cmdDExists, dexists, mArgsEqual(3))
		s.Handle(cmdDLen, dlen, mArgsEqual(2))
		s.Handle(cmdDIncrBy, dincrBy, mArgsEqual(4))
		s.Handle(cmdDIncrByFloat, dincrByFloat, mArgsEqual(4))
	}
	if password != "" {
		// server routes
		s.Handle(cmdPing, ping, mAuth(), mArgsEqual(1))
		s.Handle(cmdSave, save, mAuth(), mArgsEqual(1))
		// keys routes
		s.Handle(cmdDel, del, mAuth(), mArgsEqual(2))
		s.Handle(cmdRename, rename, mAuth(), mArgsEqual(3))
		s.Handle(cmdTTL, ttl, mAuth(), mArgsEqual(2))
		s.Handle(cmdPersist, persist, mAuth(), mArgsEqual(2))
		s.Handle(cmdExpire, expire, mAuth(), mArgsEqual(3))
		s.Handle(cmdType, hType, mAuth(), mArgsEqual(2))
		s.Handle(cmdKeys, keys, mAuth(), mArgsEqual(2))
		s.Handle(cmdExists, exists, mAuth(), mArgsMoreThan(2))
		// string routes
		s.Handle(cmdSet, set, mAuth(), mArgsEqual(3))
		s.Handle(cmdGet, get, mAuth(), mArgsEqual(2))
		s.Handle(cmdAppend, hAppend, mAuth(), mArgsEqual(3))
		s.Handle(cmdGetSet, getset, mAuth(), mArgsEqual(3))
		s.Handle(cmdStrlen, strlen, mAuth(), mArgsEqual(2))
		s.Handle(cmdIncr, incr, mAuth(), mArgsEqual(2))
		s.Handle(cmdIncrBy, incrBy, mAuth(), mArgsEqual(3))
		// list routes
		s.Handle(cmdLPush, lpush, mAuth(), mArgsEqual(3))
		s.Handle(cmdRPush, rpush, mAuth(), mArgsEqual(3))
		s.Handle(cmdLPop, lpop, mAuth(), mArgsEqual(2))
		s.Handle(cmdRPop, rpop, mAuth(), mArgsEqual(2))
		s.Handle(cmdLRem, lrem, mAuth(), mArgsEqual(4))
		s.Handle(cmdLIndex, lindex, mAuth(), mArgsEqual(3))
		s.Handle(cmdLRange, lrange, mAuth(), mArgsEqual(4))
		s.Handle(cmdLLen, llen, mAuth(), mArgsEqual(2))
		// dict routes
		s.Handle(cmdDSet, dset, mAuth(), mArgsMoreThan(4), mArgsCountEven())
		s.Handle(cmdDGet, dget, mAuth(), mArgsMoreThan(3))
		s.Handle(cmdDDel, ddel, mAuth(), mArgsEqual(3))
		s.Handle(cmdDExists, dexists, mAuth(), mArgsEqual(3))
		s.Handle(cmdDLen, dlen, mAuth(), mArgsEqual(2))
		s.Handle(cmdDIncrBy, dincrBy, mAuth(), mArgsEqual(4))
		s.Handle(cmdDIncrByFloat, dincrByFloat, mAuth(), mArgsEqual(4))
	}
}
