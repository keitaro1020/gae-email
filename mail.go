package app

import (
	"github.com/nissy/bon"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"io/ioutil"
	"net/http"
	"net/mail"
)

func init() {

	r := bon.NewRouter()

	r.Post("/_ah/mail/:ToAddress", func(w http.ResponseWriter, r *http.Request) {
		c := appengine.NewContext(r)
		defer r.Body.Close()

		m, err := mail.ReadMessage(r.Body)
		if err != nil {
			log.Errorf(c, "Error reading r.body: %v", err)
			return
		}

		header := m.Header
		log.Infof(c, "Date:", header.Get("Date"))
		log.Infof(c, "From:", header.Get("From"))
		log.Infof(c, "To:", header.Get("To"))
		log.Infof(c, "Subject:", header.Get("Subject"))

		body, err := ioutil.ReadAll(m.Body)
		if err != nil {
			log.Errorf(c, "Error reading m.body: %v", err)
		}
		log.Infof(c, "Body:", string(body))
	})

	http.Handle("/", r)
}
