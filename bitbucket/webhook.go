package bitbucket

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	EventIDHeader        = "X-Request-Id"
	EventKeyHeader       = "X-Event-Key"
	EventSignatureHeader = "X-Hub-Signature"
)

const maxPayloadSize = 10 * 1024 * 1024 // 10 MiB

func ParsePayload(r *http.Request, key []byte) (interface{}, error) {
	p, err := validateSignature(r, key)
	if err != nil {
		return nil, err
	}

	evk := r.Header.Get(EventKeyHeader)
	if evk == "" {
		return nil, fmt.Errorf("unable find event key in request")
	}
	k := EventKey(evk)
	var event interface{}
	switch k {
	case EventKeyRepoRefsChanged:
		event = &RepositoryPushEvent{}
	case EventKeyPullRequestOpened, EventKeyPullRequestFrom, EventkeyPullRequestModified, EventKeyPullRequestDeclined, EventKeyPullRequestDeleted, EventKeyPullRequestMerged:
		event = &PullRequestEvent{}
	default:
		return nil, fmt.Errorf("event type not supported: %s", k)
	}

	err = json.Unmarshal(p, event)
	if err != nil {
		return nil, fmt.Errorf("unable to parse event payload: %w", err)
	}

	return event, nil
}

func validateSignature(r *http.Request, key []byte) ([]byte, error) {
	sig := r.Header.Get(EventSignatureHeader)
	if sig == "" {
		return nil, fmt.Errorf("no signature found")
	}

	payload, err := io.ReadAll(io.LimitReader(r.Body, maxPayloadSize))
	if err != nil {
		return nil, fmt.Errorf("unable to parse payload: %w", err)
	}

	sp := strings.Split(sig, "=")
	if len(sp) != 2 {
		return nil, fmt.Errorf("signatur format invalid")
	}

	if sp[0] != "sha256" {
		return nil, fmt.Errorf("unsupported hash algorithm: %s", sp[0])
	}

	sd, err := hex.DecodeString(sp[1])
	if err != nil {
		return nil, fmt.Errorf("unable to parse signature data: %w", err)
	}

	h := hmac.New(sha256.New, key)
	h.Write([]byte(payload))

	if !hmac.Equal(h.Sum(nil), sd) {
		return nil, fmt.Errorf("signature does not match")
	}

	return payload, nil
}
