Really basic go file server.

Serves files in `www`.

Requests to `/` will return the contenxt of `www/index.html`, if it exists.

Includes basic `FROM scratch` Dockerfile to run the program.
