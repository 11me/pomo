package lib

import "time"

// Reply type is just an alisas to a string
// and defined just for semantical reason.
type Reply string

// PomoAPI type for interaction with pomo.
// This type exposed and registered as specified in RPC.
type PomoAPI struct {
	Pomo *Pomo
}

// Start API call.
// The first argument provided by caller is always nil.
func (api *PomoAPI) Start(_ bool, r *Reply) error {

	api.Pomo.Start()

	return nil
}

// Stop API call.
// The first argument provided by caller is always nil.
func (api *PomoAPI) Stop(_ bool, r *Reply) error {

	api.Pomo.Stop()

	return nil
}

// SetDuration API call.
// Takes a duration time format and calls SetDuration pomo's method.
func (api *PomoAPI) SetDuration(d time.Duration, r *Reply) error {

	api.Pomo.SetDuration(d)

	return nil
}

// GetDuration API call.
// The first argument provided by caller is always nil.
func (api *PomoAPI) GetDuration(_ bool, r *Reply) error {
	*r = Reply(api.Pomo.GetDuration())
	return nil
}

// Restart API call.
func (api *PomoAPI) Restart(_ bool, r *Reply) error {

	api.Pomo.Restart()

	return nil
}
