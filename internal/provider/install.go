package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/c3os-io/c3os/pkg/bus"
	"github.com/mudler/go-nodepair"
	"github.com/mudler/go-pluggable"
)

func Install(e *pluggable.Event) pluggable.EventResponse {
	cfg := &bus.InstallPayload{}
	err := json.Unmarshal([]byte(e.Data), cfg)
	if err != nil {
		return pluggable.EventResponse{Error: fmt.Sprintf("Failed reading JSON input: %s", err.Error())}
	}

	r := map[string]string{}
	ctx := context.Background()
	if err := nodepair.Receive(ctx, &r, nodepair.WithToken(cfg.Token)); err != nil {
		return pluggable.EventResponse{Error: fmt.Sprintf("Failed reading JSON input: %s", err.Error())}
	}

	payload, err := json.Marshal(r)
	if err != nil {
		return pluggable.EventResponse{Error: fmt.Sprintf("Failed marshalling JSON input: %s", err.Error())}
	}

	return pluggable.EventResponse{
		State: "",
		Data:  string(payload),
		Error: "",
	}
}
