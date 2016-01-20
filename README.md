# Sourceforge API Library

Based on https://github.com/google/go-github

## Example

```go
sfClient = sfapi.NewClient(nil, project)

tickets, _, err := sfClient.Tracker.Info(category)

for _, ticket := range tickets.Tickets {
	ticket, _, err := sfClient.Tracker.Get(category, ticket.TicketNum)

	fmt.Println(ticket)
}
```
