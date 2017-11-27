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
		s.Handle(cmdPing, ping, mIsAuthorized(), mArgsEqual(1))
		s.Handle(cmdSave, save, mIsAuthorized(), mArgsEqual(1))
		// keys routes
		s.Handle(cmdDel, del, mIsAuthorized(), mArgsEqual(2))
		s.Handle(cmdRename, rename, mIsAuthorized(), mArgsEqual(3))
		s.Handle(cmdTTL, ttl, mIsAuthorized(), mArgsEqual(2))
		s.Handle(cmdPersist, persist, mIsAuthorized(), mArgsEqual(2))
		s.Handle(cmdExpire, expire, mIsAuthorized(), mArgsEqual(3))
		s.Handle(cmdType, hType, mIsAuthorized(), mArgsEqual(2))
		s.Handle(cmdKeys, keys, mIsAuthorized(), mArgsEqual(2))
		s.Handle(cmdExists, exists, mIsAuthorized(), mArgsMoreThan(2))
		// string routes
		s.Handle(cmdSet, set, mIsAuthorized(), mArgsEqual(3))
		s.Handle(cmdGet, get, mIsAuthorized(), mArgsEqual(2))
		s.Handle(cmdAppend, hAppend, mIsAuthorized(), mArgsEqual(3))
		s.Handle(cmdGetSet, getset, mIsAuthorized(), mArgsEqual(3))
		s.Handle(cmdStrlen, strlen, mIsAuthorized(), mArgsEqual(2))
		s.Handle(cmdIncr, incr, mIsAuthorized(), mArgsEqual(2))
		s.Handle(cmdIncrBy, incrBy, mIsAuthorized(), mArgsEqual(3))
		// list routes
		s.Handle(cmdLPush, lpush, mIsAuthorized(), mArgsEqual(3))
		s.Handle(cmdRPush, rpush, mIsAuthorized(), mArgsEqual(3))
		s.Handle(cmdLPop, lpop, mIsAuthorized(), mArgsEqual(2))
		s.Handle(cmdRPop, rpop, mIsAuthorized(), mArgsEqual(2))
		s.Handle(cmdLRem, lrem, mIsAuthorized(), mArgsEqual(4))
		s.Handle(cmdLIndex, lindex, mIsAuthorized(), mArgsEqual(3))
		s.Handle(cmdLRange, lrange, mIsAuthorized(), mArgsEqual(4))
		s.Handle(cmdLLen, llen, mIsAuthorized(), mArgsEqual(2))
		// dict routes
		s.Handle(cmdDSet, dset, mIsAuthorized(), mArgsMoreThan(4), mArgsCountEven())
		s.Handle(cmdDGet, dget, mIsAuthorized(), mArgsMoreThan(3))
		s.Handle(cmdDDel, ddel, mIsAuthorized(), mArgsEqual(3))
		s.Handle(cmdDExists, dexists, mIsAuthorized(), mArgsEqual(3))
		s.Handle(cmdDLen, dlen, mIsAuthorized(), mArgsEqual(2))
		s.Handle(cmdDIncrBy, dincrBy, mIsAuthorized(), mArgsEqual(4))
		s.Handle(cmdDIncrByFloat, dincrByFloat, mIsAuthorized(), mArgsEqual(4))
	}
}
