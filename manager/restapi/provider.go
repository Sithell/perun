package restapi

import (
	"context"
	"fmt"
	"github.com/go-openapi/runtime"
	"github.com/sithell/perun/internal/database"
	"github.com/sithell/perun/manager/restapi/operations"
	"github.com/sithell/perun/provider/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

type registerProviderResponder struct {
	params operations.RegisterProviderParams
	app    *App
}

func (rs registerProviderResponder) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	provider := database.Provider{}
	providerHost := fmt.Sprintf("%s:%d", *rs.params.Host.Host, *rs.params.Host.Port)

	if err := checkProvider(rs.params.HTTPRequest.Context(), providerHost); err != nil {
		log.Printf("failed to verify provider: %v", err)
		operations.NewRegisterProviderForbidden().WriteResponse(rw, producer)
		return
	}

	result := rs.app.DB.
		Where(database.Provider{Host: providerHost}).
		Assign(database.Provider{Status: "active"}).
		FirstOrCreate(&provider)
	if result.Error != nil {
		log.Fatalf("failed to update provider in db: %v", result.Error)
	}
	operations.NewRegisterProviderOK().WriteResponse(rw, producer)
}

func checkProvider(ctx context.Context, target string) error {
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to establish grpc connection: %w", err)
	}
	client := pb.NewProviderClient(conn)
	_, err = client.Ping(ctx, &pb.PingRequest{})
	if err != nil {
		return fmt.Errorf("failed to get ping response from provider: %w", err)
	}
	defer func(conn *grpc.ClientConn) {
		connCloseErr := conn.Close()
		if connCloseErr != nil {
			log.Printf("failed to close grpc connection: %v", connCloseErr)
		}
	}(conn)
	return nil
}
