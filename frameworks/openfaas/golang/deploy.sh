faas-cli template pull https://github.com/alexellis/golang-http-template
faas-cli new --lang golang-http hello-world
faas-cli build -f helloworld-yml
faas-cli push -f helloworld-yml
faas-cli deploy -f helloworld-yml
