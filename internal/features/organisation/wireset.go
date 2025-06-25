package organisation

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewOrganisationHandler, NewOrganisationService)
