package behavior

type Behavior interface {
	Think()
	ShortPause()
	LongPause()
}
