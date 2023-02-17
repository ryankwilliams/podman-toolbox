package podman

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/podman/v4/pkg/bindings"
)

// Create a connection to the podman socket
func CreateConnection() (context.Context, error) {
	socket := fmt.Sprintf("unix:%s/podman/podman.sock", os.Getenv("XDG_RUNTIME_DIR"))
	connectionCxt, err := bindings.NewConnection(context.Background(), socket)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to podman socket, error: %w", err)
	}
	return connectionCxt, nil
}
