package converts

import (
	"errors"
	"fmt"

	"github.com/vloldik/dbml-gen/internal/dbparse/models"
	"github.com/vloldik/dbml-gen/internal/dbparse/parseobj"
	"github.com/vloldik/dbml-gen/internal/utils/strutil"
)

func (c *ParseObjectToModelConverter) applySettings(to any, settings *parseobj.Settings) error {
	if settings == nil || settings.SettingList == nil {
		return nil
	}
	for _, unknown := range settings.SettingList {
		switch destination := to.(type) {
		case *models.Field:
			return c.applyFieldSetting(destination, unknown)
		case *models.Table:
			return c.applyTableSetting(destination, unknown)
		case *models.Index:
			return c.applyIndexSettings(destination, unknown)
		case *models.Relationship:
			return c.applyRelationsSettings(destination, unknown)
		default:
			return fmt.Errorf("unknown type to apply settings for %T", destination)
		}
	}
	return nil
}

func (c *ParseObjectToModelConverter) applyFieldSetting(field *models.Field, unknown parseobj.Setting) error {
	switch setting := unknown.(type) {
	case *parseobj.SettingDefaultValue:
		field.DefaultValue = setting.Value
	case *parseobj.SettingNote:
		field.Note = setting.Value
	case *parseobj.SettingIncrement:
		field.IsIncrement = setting.Value
	case *parseobj.SettingNotNull:
		field.IsNotNull = setting.Value
	case *parseobj.SettingPrimaryKey:
		field.IsPrimaryKey = setting.Value
	case *parseobj.SettingUnique:
		field.IsUnique = setting.Value
	case *parseobj.SettingReference:
		c.referenceList = append(c.referenceList, &parseobj.StructureFullReference{
			Type:             setting.Value.Type,
			ReferenceToField: setting.Value.ReferenceToField,
			Field: &parseobj.ReferenceField{
				NameParts: []string{
					field.Table.TableName.Namespace,
					field.Table.TableName.BaseName,
					field.DBName,
				},
			},
		})
	default:
		return ErrorUnknownSetting(setting, "models.Field")
	}

	return nil
}

func (c *ParseObjectToModelConverter) applyTableSetting(table *models.Table, unknown parseobj.Setting) error {
	switch setting := unknown.(type) {
	case *parseobj.SettingHeaderColor:
		break
	case *parseobj.SettingNote:
		table.Note = setting.Value
	default:
		return ErrorUnknownSetting(setting, "models.Table")
	}
	return nil
}

func (c *ParseObjectToModelConverter) applyIndexSettings(index *models.Index, unknown parseobj.Setting) error {
	switch setting := unknown.(type) {
	case *parseobj.SettingNote:
		index.Note = setting.Value
	case *parseobj.SettingPrimaryKey:
		index.IsPrimaryKey = setting.Value
	case *parseobj.SettingUnique:
		index.IsUnique = setting.Value
	case *parseobj.SettingsIndexType:
		index.Type = setting.Value
	case *parseobj.SettingName:
		name, err := strutil.UnquoteString(setting.Value)
		if err != nil {
			return err
		}
		index.Name = name
	default:
		return ErrorUnknownSetting(setting, "models.Index")
	}

	return nil
}

func (c *ParseObjectToModelConverter) applyRelationsSettings(relation *models.Relationship, unknown parseobj.Setting) error {
	switch setting := unknown.(type) {
	case *parseobj.SettingRefOnAction:
		var settingType models.OnRefChangeAction
		if setting.Type.IsCascade {
			settingType = models.Cascade
		} else if setting.Type.IsNoAction {
			settingType = models.NoAction
		} else if setting.Type.IsRestrict {
			settingType = models.Restrict
		} else if setting.Type.IsSetDefault {
			settingType = models.SetDefault
		} else if setting.Type.IsSetNull {
			settingType = models.SetNull
		} else {
			return errors.New("unknown setting for relation")
		}
		if setting.TriggerOn.IsDelete {
			relation.OnDelete = settingType
		} else if setting.TriggerOn.IsUpdate {
			relation.OnUpdate = settingType
		} else {
			return errors.New("unknown setting trigger on")
		}
	default:
		return fmt.Errorf("unknown setting type for relation: %T", setting)
	}
	return nil
}
