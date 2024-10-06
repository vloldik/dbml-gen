package converts

import (
	"fmt"

	"guthub.com/vloldik/dbml-gen/internal/dbparse/models"
	"guthub.com/vloldik/dbml-gen/internal/dbparse/parseobj"
)

func (c *ParseObjectToModelConverter) applySettings(to any, settings *parseobj.Settings) error {
	for _, unknown := range settings.SettingList {
		switch destination := to.(type) {
		case *models.Field:
			return c.applyFieldSetting(destination, unknown)
		case *models.Table:
			return c.applyTableSetting(destination, unknown)
		case *models.Index:
			return c.applyIndexSettings(destination, unknown)
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
					field.TableName.Namespace,
					field.TableName.BaseName,
					field.Name,
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
	case *parseobj.HeaderColorSetting:
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
		index.Name = setting.Value
	default:
		return ErrorUnknownSetting(setting, "models.Index")
	}

	return nil
}
