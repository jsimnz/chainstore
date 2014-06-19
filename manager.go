package chainstore

// StoreManager interface
// Used to create managers of stores like lrumanager and batchmgr
// By calling Attach(...) the manager becomes respsonsible for those
// stores
type StoreManager interface {
	Store
	Attach(store ...Store)
}

// Default implementation of Manager
// Should be embedded as a pointer in you own manager implementation
// and overwrite any methods you need your self.
// It exposes a []Store called Stores to your implementation which is
// set via Attach(Store...). These are stores that the manager is responsible
// for
type DefaultManager struct {
	Stores []Store
}

func (m *DefaultManager) Attach(stores ...Store) {
	m.Stores = stores
}

/*func (m *DefaultManager) Skip() bool {
	return false
}

func (m *DefaultManager) Continue() bool {
	return true
}

func (m *DefaultManager) BeforeNext(store Store) (err error) {
	return //noop
}

func (m *DefaultManager) AfterNext(store Store)*/

// s3store := s3store.New(...)
// batchManager.Attach(s3store)
// batchManager.Use(s3store)
// batchManager.Manage
