package roles

const (
	metalService      = "/metal.MetalDoorsApi/"        // Metal doors service routes.
	woodService       = "/wood.doors.WoodDoorsApi/"    // Wood doors service routes.
	woodColorsService = "/wood.color.ColorsApi/"       // Wood colors service routes.
	userService       = "/user.UserService/"           // User service routes.
	inventoryService  = "/inventory.InventoryService/" // Inventory service routes.

	dealerService = "/dealer.ReserveService/" // TODO: unimplemented
)

const (
	boss    = "boss"
	admin   = "admin"
	manager = "manager"
	dealer  = "dealer"
)

// AccessibleRoles returns map of routes with available roles.
func AccessibleRoles() map[string][]string {
	return map[string][]string{
		inventoryService + "NewInventory":     {admin},          // TODO:refactor and test
		inventoryService + "ApproveInventory": {admin},          // TODO:implement
		inventoryService + "GetInventory":     {admin, manager}, // TODO:implement
		inventoryService + "GetInventories":   {admin, manager}, // TODO:implement
		inventoryService + "UpdateInventory":  {admin},          // TODO:implement
		inventoryService + "DeleteInventory":  {admin},          // TODO:implement

		metalService + "CreatePost":  {admin, manager}, // TODO:refactor and test
		metalService + "UpdateMetal": {admin, manager}, // TODO:implement
		metalService + "DeleteMetal": {admin, manager}, // TODO:implement

		woodService + "CreateWoodDoor": {admin, manager}, // TODO:refactor and test

		woodColorsService + "CreateColor": {admin, manager}, // TODO:refactor and test

		userService + "CreateUser":  {admin},
		userService + "GetAllUsers": {admin},
		userService + "DeleteUser":  {admin}, // TODO:implement
		userService + "GetUser":     {admin, manager, dealer},

		dealerService + "AddCatalogItem": {admin},                  // TODO:implement
		dealerService + "DropAll":        {admin},                  // TODO:implement
		dealerService + "NewReserve":     {admin, dealer},          // TODO:implement
		dealerService + "DeleteReserve":  {admin, manager, dealer}, // TODO:implement
		dealerService + "GetCatalog":     {admin, manager, dealer}, // TODO:implement
	}
}
