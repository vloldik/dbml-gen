package parseobj

// Union type for all Settingsettings
type Setting interface {
	setting()
}

type Settings struct {
	SettingList []Setting `("[" @@? ( "," @@ )* ","? "]")?`
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

type SettingHeaderColor struct {
	*settingImpl
	Value string `"headercolor" ":" @Color`
}

type SettingRefOnAction struct {
	*settingImpl
	TriggerOn *RefActionTrigger       `@@`
	Type      *SettingOnRefActionType `@@`
}

type SettingOnRefActionType struct {
	IsSetNull    bool `  @ "set" "null"`
	IsCascade    bool `| @ "cascade"`
	IsRestrict   bool `| @ "restrict"`
	IsSetDefault bool `| @ "set" "default"`
	IsNoAction   bool `| @ "no" "action"`
}

type RefActionTrigger struct {
	IsUpdate bool `@ "update" ":"`
	IsDelete bool `| @ "delete" ":"`
}

type settingImpl struct{}

func (*settingImpl) setting() {}
