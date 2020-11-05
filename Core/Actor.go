package Core

type Actor interface {
	OnGlance() string
	OnInspect() string
	OnProvoke() string
	OnDeath() string
	OnIdle() string
	OnMove() string
	OnInjury() string
	OnSurrender() string
	OnEscape() string

	OnApprentice()

	TriggerEvent(eventId int)
}
