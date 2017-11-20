tester_webapp
===

This is a simple Go webapp with a Docker container for any use.

# Run
The following command will make a Docker container for the webapp, make the webbapp, and start it on the container. The webapp will become available at [localhost:5000](http://localhost:5000).

```
$ ./run.sh
```

# Source
The main Go files are located in `tester/src/webapp/`. This repo uses [dep](https://github.com/golang/dep) for dependency management. You can modify any source to make your own webapp.

```
$ tree tester/src
tester/src/webapp/
├── app.go
├── Gopkg.lock
├── Gopkg.toml
└── views
    ├── base.html
    └── index.html
```

# License
MIT.
