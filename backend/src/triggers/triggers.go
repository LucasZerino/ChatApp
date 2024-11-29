package triggers

import (
	"gorm.io/gorm"
)

func CreateTriggers(db *gorm.DB) error {
	// Trigger para criar a sequência quando um novo registro for inserido em accounts
	triggerSQL1 := `
		CREATE TRIGGER accounts_after_insert_row_tr
		AFTER INSERT ON accounts
		FOR EACH ROW
		EXECUTE FUNCTION create_sequence_if_not_exists(NEW.id);
	`
	err := db.Exec(triggerSQL1).Error
	if err != nil {
		return err
	}

	// Trigger para gerar display_id antes da inserção em conversations
	triggerSQL2 := `
		CREATE TRIGGER conversations_before_insert_row_tr
		BEFORE INSERT ON conversations
		FOR EACH ROW
		EXECUTE FUNCTION set_display_id_for_conversations(NEW.account_id);
	`
	err = db.Exec(triggerSQL2).Error
	if err != nil {
		return err
	}

	// Trigger para criar a sequência quando um novo registro for inserido em accounts (camp)
	triggerSQL3 := `
		CREATE TRIGGER camp_dpid_before_insert
		AFTER INSERT ON accounts
		FOR EACH ROW
		EXECUTE FUNCTION create_sequence_if_not_exists(NEW.id);
	`
	err = db.Exec(triggerSQL3).Error
	if err != nil {
		return err
	}

	// Trigger para gerar display_id antes da inserção em campaigns
	triggerSQL4 := `
		CREATE TRIGGER campaigns_before_insert_row_tr
		BEFORE INSERT ON campaigns
		FOR EACH ROW
		EXECUTE FUNCTION set_display_id_for_campaigns(NEW.account_id);
	`
	err = db.Exec(triggerSQL4).Error
	if err != nil {
		return err
	}

	return nil
}
