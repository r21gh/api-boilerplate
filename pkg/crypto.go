package root

// Crypto interface
type Crypto interface {
	Salt(s string) (error, string)
	Compare(hash string, s string) (error, bool)
}