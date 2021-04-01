package acquiring

import (
    "fmt"
    "log"
    "time"

    "github.com/google/uuid"
)

// Session TODO
type Session struct {
    gateway   SessionGateway
    ID        uuid.UUID  `json:"id"`
    DeviceID  uuid.UUID  `json:"device_id"`
    StartedAt *time.Time `json:"started_at,omitempty"`
    StoppedAt *time.Time `json:"stopped_at,omitempty"`
}

// NewSession TODO
func NewSession(
    gateway Gateway,
) *Session {
    session := &Session{
        gateway:  gateway,
        ID:       uuid.NewUUID().String(),
        DeviceID: device.ID,
    }

    return session
}

// Start TODO
func (session *Session) Start(annotation string) ([]Event, error) {
    if session.StartedAt != nil {
        return nil, fmt.Errorf("already started at: %v", session.StartedAt)
    }

    if err := session.gateway.Start(); err != nil {
        return nil, fmt.Errorf("failed to start with error: %v", err)
    }

    event := &SessionStartedEvent{
        Time:       time.Now(),
        Annotation: annotation,
    }

    return []Event{event}, nil
}

// Stop TODO
func (session *Session) Stop(annotation string) ([]Event, error) {
    if session.StartedAt == nil {
        return nil, fmt.Errorf("session not yet started")
    }

    if session.StoppedAt != nil {
        return nil, fmt.Errorf("already stopped at: %v", session.StoppedAt)
    }

    if err := session.gateway.Stop(); err != nil {
        return nil, fmt.Errorf("failed to stop with error: %v", err)
    }

    event := &SessionStoppedEvent{
        Time:       time.Now(),
        Annotation: annotation,
    }

    return []Event{event}, nil
}

// Apply TODO
func (session *Session) Apply(event Event) {
    switch e := event.(type) {
    case acquiring.SessionStartedEvent:
        session.StartedAt = e.Time
    case acquiring.SessionStoppedEvent:
        session.StoppedAt = e.Time
    default:
        log.Printf("unsupported event type: %v\n")
    }
}
