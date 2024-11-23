package database

import (
	"encoding/json"

	badger "github.com/dgraph-io/badger/v4"
)

type dBEntry string

const (
	accessKeyEntry dBEntry = "accesskey"
	endtimeEntry   dBEntry = "endtime"
	blocklistEntry dBEntry = "blocklist"
)

type DB struct {
	path string
	db   *badger.DB
}

func NewDB(path string) (DB, error) {
	// path = env.DBPath

	options := badger.DefaultOptions(path).WithLogger(nil)
	db, err := badger.Open(options)

	if err != nil {
		return DB{}, err
	}

	return DB{
		path,
		db,
	}, nil
}

func (d *DB) CloseDB() {
	println("close")
	d.db.Close()
}

// endtime

func (d *DB) GetEndtime() (string, error) {
	var valCopy []byte

	err := d.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(endtimeEntry))
		if err != nil {
			return err
		}

		valCopy, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return string(valCopy), nil
}

func (d *DB) SetEndtime(endtime string) error {
	err := d.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(endtimeEntry), []byte(endtime))
		return err
	})

	return err
}

// blocklist

func (d *DB) GetBlocklist() ([]string, error) {
	var valCopy []byte

	err := d.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(blocklistEntry))
		if err != nil {
			return err
		}

		valCopy, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	slice, err := decodeStringSlice(valCopy)
	if err != nil {
		return nil, err
	}

	return slice, nil
}

func (d *DB) SetBlocklist(blocklist []string) error {
	err := d.db.Update(func(txn *badger.Txn) error {
		bytelist, err := encodeStringSlice(blocklist)
		if err != nil {
			return err
		}

		err = txn.Set([]byte(blocklistEntry), bytelist)
		return err
	})

	return err
}

// helpers

func encodeStringSlice(data []string) ([]byte, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

func decodeStringSlice(data []byte) ([]string, error) {
	var result []string
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
