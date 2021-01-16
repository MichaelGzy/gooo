package store

type Store interface {
	// Init initialises the store. It must perform any required setup on the backing storage implementation and check that it is ready for use, returning any errors.
	Init(...Option) error
	// Options allows you to view the current options.
	Options() Options
}

type Option func(*Options)
