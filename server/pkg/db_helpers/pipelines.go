package db_helpers

import "go.mongodb.org/mongo-driver/bson"

// MWCCollectionPipeline - Manufacturers with collections of colors pipeline.
var MWCCollectionPipeline = []bson.M{
	{"$unwind": bson.M{"path": "$collections", "preserveNullAndEmptyArrays": true}},
	{"$lookup": bson.M{
		"from":         "wood-colors",
		"localField":   "collections.id",
		"foreignField": "collection.id",
		"as":           "collections.colors",
	},
	},
	{"$unwind": bson.M{"path": "$collections"}},
	{"$group": bson.M{
		"_id":          "$_id",
		"manufacturer": bson.M{"$first": "$manufacturer"},
		"collections":  bson.M{"$push": "$collections"},
	}},
}

// ProductWColorsPipeline - Product with color variants pipeline.
var ProductWColorsPipeline = []bson.M{
	{"$unwind": bson.M{"path": "$variants", "preserveNullAndEmptyArrays": true}},
	{"$lookup": bson.M{
		"from":         "woodColors",
		"localField":   "variants.color",
		"foreignField": "uuid",
		"as":           "variants.color",
	}},
	{"$unwind": bson.M{"path": "$variants.color"}},
	{"$group": bson.M{
		"_id":      "$_id",
		"uuid":     bson.M{"$first": "$uuid"},
		"variants": bson.M{"$push": "$variants"},
		"price":    bson.M{"$first": "$price.roz"},
		"title":    bson.M{"$first": "$title"},
		"tags":     bson.M{"$first": "$tags"},
	}},
}
