package main

import (
	"net/http"
	"testing"

	"snippetbox.xmxxmx.us/internal/assert"
)

func TestPing(t *testing.T) {
	// Create a new instance of our application struct. For now, this just
	// contains a structured logger (which uses the slog.DiscardHandler handler
	// and will discard anything written to it with no action).
	// app := &application{
	// 	logger: slog.New(slog.DiscardHandler),
	// }
	app := newTestApplication(t)

	// We then use the httptest.NewTLSServer() function to create a new test
	// server, passing in the value returned by our app.routes() method as the
	// handler for the server. This starts up an HTTPS server which listens on a
	// randomly-chosen port of your local machine for the duration of the test.
	// Notice that we defer a call to ts.Close() so that the server is shut down
	// when the test finishes.
	// ts := httptest.NewTLSServer(app.routes())
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	// The network address that the test server is listening on is contained in
	// the ts.URL field. We can  use this along with the ts.Client().Get() method
	// to make a GET /ping request against the test server. This returns a
	// http.Response struct containing the response.
	// rs, err := ts.Client().Get(ts.URL + "/ping")
	// if err != nil {
	// 	t.Fatal(err)
	// }
	code, _, body := ts.get(t, "/ping")

	// Check that the status code written by the ping handler was 200.
	assert.Equal(t, code, http.StatusOK)

	// And we can check that the response body written by the ping handler
	// equals "OK".
	// defer rs.Body.Close()
	// body, err := io.ReadAll(rs.Body)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// body = bytes.TrimSpace(body)

	assert.Equal(t, body, "OK")
}
