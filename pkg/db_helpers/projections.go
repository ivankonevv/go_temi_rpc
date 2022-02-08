package db_helpers

import "go.mongodb.org/mongo-driver/bson"

// GetDoorsProjection - metal doors with excluded prices and data.
var GetDoorsProjection = bson.D{
	{"price.k_opt", 0},        // exclude k opt from result.
	{"price.vip", 0},          // exclude vip from result.
	{"price.opt", 0},          // exclude opt from result.
	{"variants.size.data", 0}, // exclude size data from result.
}
