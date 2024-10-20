package parseobj

type DBML struct {
	Structures []DBMLStructure `@@*`
}

type StructureTable struct {
	*dbmlStructureImpl
	Name     *NamespacedName `"Table" @@`
	As       *string         `("as" @Ident)?`
	Settings *Settings       `@@`

	Content *TableContent `"{"  @@? "}"`
}

type StructureEnum struct {
	*dbmlStructureImpl
	Name   *NamespacedName `"Enum" @@ "{"`
	Values []*EnumValue    `@@* "}"`
}

type StructureFullReference struct {
	*dbmlStructureImpl
	Field            *ReferenceField   `"ref" ":" @@`
	Type             *RelationshipType `@@`
	ReferenceToField *ReferenceField   `@@`
	Settings         *Settings         `@@`
}

type DBMLStructure interface {
	structure()
}

type dbmlStructureImpl struct{}

func (dbmlStructureImpl) structure() {}
