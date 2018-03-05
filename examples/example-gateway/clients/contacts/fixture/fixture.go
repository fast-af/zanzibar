package fixture

import (
	mc "github.com/uber/zanzibar/examples/example-gateway/build/clients/contacts/mock-client"
	gen "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/contacts/contacts"
)

var saveContactsFixtures = &mc.SaveContactsScenarios{
	Success: &mc.SaveContactsFixture{
		Arg0Any: true,
		Arg1Any: true,
		Arg2: &gen.SaveContactsRequest{
			UserUUID: "foo",
		},

		Ret0: &gen.SaveContactsResponse{},
	},
}

// Fixture ...
var Fixture = &mc.ClientFixture{
	SaveContacts: saveContactsFixtures,
}
