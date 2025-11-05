package timeline

type DeleteTimelineRepository interface {
	DeleteTimeline(id string) error
}
