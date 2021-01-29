package furni

import (
	"database/sql"
	"fmt"
	"time"
)

// Arcturusadder stores furni to the Arcturus database.
type ArcturusAdder struct {
	DB *sql.DB
}

func (a *ArcturusAdder) Add(furni Furni) error {
	// get catalog page
	var catalogPage int
	err := a.DB.QueryRow("select furniripper.catalog_id FROM furniripper inner join catalog_pages on catalog_pages.id = furniripper.catalog_id LIMIT 1").Scan(&catalogPage)
	if err != nil {
		return err
	}

	spriteId := time.Now().Unix()

	// add to items_base
	query := "INSERT INTO items_base (sprite_id, public_name, item_name, stack_height, width, length) VALUES (?, ?, ?, ?, ?, ?)"
	res, err := a.DB.Exec(query, spriteId, furni.Name, furni.Name, furni.Height, furni.Width, furni.Length)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	itemId, err := res.LastInsertId()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("%d rows affected instead of 1 while inserting into items_base", rowsAffected)
	}

	res, err = a.DB.Exec("INSERT INTO catalog_items (item_ids, page_id, catalog_name) VALUES (?, ?, ?)", itemId, catalogPage, furni.Name)
	if err != nil {
		return err
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("%d rows affected instead of 1 while inserting into catalog_items", rowsAffected)
	}

	xmlData := fmt.Sprintf(`<furnitype id="%d" classname="%s"><revision>1337</revision><defaultdir>0</defaultdir><xdim>1</xdim><ydim>1</ydim><partcolors /><name>%s</name><description>Added by Habbo.ovh furni ripper</description><adurl /><offerid>-1</offerid><buyout>0</buyout><rentofferid>-1</rentofferid><rentbuyout>0</rentbuyout><bc>0</bc><excludeddynamic>0</excludeddynamic><customparams /><specialtype>1</specialtype><canstandon>0</canstandon><cansiton>0</cansiton><canlayon>0</canlayon></furnitype>`, spriteId, furni.Name, furni.Name)

	res, err = a.DB.Exec("INSERT INTO furnidata (data, type) VALUES (?, ?)", xmlData, "room")
	if err != nil {
		return err
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("%d rows affected instead of 1 while working on furnidata table", rowsAffected)
	}

	return nil
}
