package Core

type Combo struct {
	Text       string
	ActionType int
}

type Actor interface {
	OnGlance() string
	OnInspect() string
	OnProvoke(actor *Actor) string
	OnProvoked(actor *Actor) string
	OnDeath() string
	OnIdle() string
	OnMove() string
	OnInjury() string
	OnSurrender() string
	OnEscape() string

	OnApprentice()

	TriggerEvent(eventId int)

	GetName() string
}
