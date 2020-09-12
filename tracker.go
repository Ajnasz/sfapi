package sfapi

import (
	"fmt"
	"path"
	"strconv"
	"time"
)

// Milestone represents project milestones, comes in TrackerInfo
type Milestone struct {
	Closed   int  `json:"closed"`
	Complete bool `json:"complete"`
	// Default     bool   `json:"default"` // bool or string in sourceforge
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Name        string `json:"name"`
	Total       int    `json:"total"`
}

func (m *Milestone) DueTime() time.Time {
	dueTime, _ := time.Parse("2006-01-02 15:04:05", m.DueDate)
	return dueTime
}

// TrackerInfoTicket represents tickets in the TrackerInfo response
type TrackerInfoTicket struct {
	Summary   string `json:"summary"`
	TicketNum int    `json:"ticket_num"`
}

// TrackerInfo represents information of individual ticket categories
type TrackerInfo struct {
	Count      int                 `json:"count"`
	Limit      int                 `json:"limit"`
	Page       int                 `json:"page"`
	Milestones []Milestone         `json:"milestones"`
	Tickets    []TrackerInfoTicket `json:"tickets"`
}

// TicketAttachment represents attachments of a ticket
type TicketAttachment struct {
	Bytes int    `json:"bytes"`
	URL   string `json:"url"`
}

// DiscussionPost represents comments on a Sourceforge issue
type DiscussionPost struct {
	Attachments []TicketAttachment `json:"attachments"`
	Author      string             `json:"author"`
	LastEdited  interface{}        `json:"last_edited"`
	Slug        string             `json:"slug"`
	Subject     string             `json:"subject"`
	Text        string             `json:"text"`
	Timestamp   string             `json:"timestamp"`
}

func (t DiscussionPost) TimestampTime() time.Time {
	timestampTime, _ := time.Parse("2006-01-02 15:04:05", t.Timestamp)
	return timestampTime

}

// Ticket represents Sourceforge issue
type Ticket struct {
	ID           string             `json:"_id"`
	AssignedTo   string             `json:"assigned_to"`
	AssignedToID string             `json:"assigned_to_id"`
	Attachments  []TicketAttachment `json:"attachments"`
	CreatedDate  string             `json:"created_date"`
	CustomFields struct {
		Milestone string `json:"_milestone"`
		Priority  string `json:"_priority"`
	} `json:"custom_fields"`
	Description        string `json:"description"`
	DiscussionDisabled bool   `json:"discussion_disabled"`
	DiscussionThread   struct {
		ID           string           `json:"_id"`
		DiscussionID string           `json:"discussion_id"`
		Limit        int              `json:"limit"`
		Page         interface{}      `json:"page"`
		Posts        []DiscussionPost `json:"posts"`
		Subject      string           `json:"subject"`
	} `json:"discussion_thread"`
	DiscussionThreadURL string        `json:"discussion_thread_url"`
	Labels              []string      `json:"labels"`
	ModDate             string        `json:"mod_date"`
	Private             bool          `json:"private"`
	RelatedArtifacts    []interface{} `json:"related_artifacts"`
	ReportedBy          string        `json:"reported_by"`
	ReportedByID        string        `json:"reported_by_id"`
	Status              string        `json:"status"`
	Summary             string        `json:"summary"`
	TicketNum           int           `json:"ticket_num"`
	VotesDown           int           `json:"votes_down"`
	VotesUp             int           `json:"votes_up"`
}

func (t *Ticket) CreatedTime() time.Time {
	createdTime, _ := time.Parse("2006-01-02 15:04:05", t.CreatedDate)
	return createdTime
}

func (t *Ticket) ModTime() time.Time {
	modTime, _ := time.Parse("2006-01-02 15:04:05", t.ModDate)
	return modTime
}

// TicketResponse represents a ticket response
type TicketResponse struct {
	Ticket `json:"ticket"`
}

// RequestQuery Object represents what fields can be added to query string
type RequestQuery struct {
	Page  int
	Limit int
}

// NewRequestQuery Creates a RequestQuery object with default values
func NewRequestQuery() *RequestQuery {
	return &RequestQuery{
		Page:  0,
		Limit: 100,
	}
}

// TrackerService handles communication with the ticket tracker tool releated methods of the Sourceforge API
type TrackerService struct {
	client *Client
}

// Info Downloads information of a tracker
func (s *TrackerService) Info(trackerName string, query RequestQuery) (*TrackerInfo, *Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("%s?limit=%d&page=%d", path.Join(s.client.Project, trackerName), query.Limit, query.Page), nil)

	if err != nil {
		return nil, nil, err
	}

	tickets := new(TrackerInfo)
	resp, err := s.client.Do(req, tickets)

	if err != nil {
		return nil, nil, err
	}

	return tickets, resp, err
}

// Get a single ticket
func (s *TrackerService) Get(trackerName string, id int) (*Ticket, *Response, error) {
	rel := path.Join(s.client.Project, trackerName, strconv.Itoa(id))

	req, err := s.client.NewRequest("GET", rel, nil)

	if err != nil {
		return nil, nil, err
	}

	ticketResponse := new(TicketResponse)
	resp, err := s.client.Do(req, ticketResponse)

	if err != nil {
		return nil, nil, err
	}

	return &ticketResponse.Ticket, resp, err
}
