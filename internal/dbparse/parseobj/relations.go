package parseobj

type Relationship struct {
	Type             *RelationshipType `"ref" ":" @@`
	ReferenceToField *ReferenceField   `@@`
}

type RelationshipType struct {
	OneToOne   bool `  @ "-"`
	ManyToMany bool `| @ "<>"`
	ManyToOne  bool `| @ ">"`
	OneToMany  bool `| @ "<"`
}
