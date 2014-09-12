package gravatar

import "crypto/md5"
import "fmt"
import "io"
import "net/url"

const (
	Scheme = "http"
	Host   = "www.gravatar.com"
	Path   = "avatar"
)

// Hash converts an email to md5
func Hash(email string) string {
	h := md5.New()
	io.WriteString(h, email)

	return fmt.Sprintf("%x", h.Sum(nil))
}

// Gravatar
type Gravatar struct {
	Email string
	Hash  string
	URL   *url.URL
}

// Lookup returns a Gravatar object for the given email
func Lookup(email string) (grav *Gravatar) {
	grav = &Gravatar{
		Email: email,
		Hash:  Hash(email),
	}

	grav.URL = &url.URL{
		Scheme: Scheme,
		Host:   Host,
		Path:   fmt.Sprintf("%s/%s", Path, grav.Hash),
	}

	return
}
