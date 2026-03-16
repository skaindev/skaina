BUILD_FLAGS = -tags "$(BUILD_TAGS)" -ldflags "

build:
	@ echo "Building Skaina Client..."
	@ go build -o $(GOPATH)/bin/skaina ./chain/skaina/
	@ echo "Done building."
#.PHONY: skaina
skaina:
	@ echo "Building Skaina Client..."
	@ go build -o $(GOPATH)/bin/skaina ./chain/skaina/
	@ echo "Done building."
	@ echo "Run ./skaina to launch the client. "

install:
	@ echo "Installing..."
	@ go install -mod=readonly $(BUILD_FLAGS) ./chain/skaina
	@ echo "Install success."