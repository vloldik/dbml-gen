package parseobj

// Union type for all Settingsettings
type Setting interface {
	setting()
}

type Settings struct {
	SettingList []Setting `("[" @@* ( "," @@ )* ","? "]")?`
}

type SettingPrimaryKey struct {
	*settingImpl
	Value bool `@("pk" | ("primary" "key"))`
}

type SettingIncrement struct {
	*settingImpl
	Value bool `@"increment"`
}

type SettingUnique struct {
	*settingImpl
	Value bool `@"unique"`
}

type SettingNotNull struct {
	*settingImpl
	Value bool `@("not_null" | ("not" "null"))`
}

type SettingReference struct {
	*settingImpl
	Value *Relationship `@@`
}

type SettingNote struct {
	*settingImpl
	Value string `"note" ":" @String`
}

type SettingDefaultValue struct {
	*settingImpl
	Value string `"default" ":" (@String | @Number | @Ident | @DBStatement)`
}

type SettingName struct {
	*settingImpl
	Value string `"name" ":" @String`
}

type SettingsIndexType struct {
	*settingImpl
	Value string `"type" ":" @Ident`
}

type HeaderColorSetting struct {
	*settingImpl
	Value string `"headercolor" ":" @Color`
}

type settingImpl struct{}

func (*settingImpl) setting() {}
