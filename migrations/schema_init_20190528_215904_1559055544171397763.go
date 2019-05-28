package main

import "github.com/zainul/gan/pkg/migration"

// Schema_init_20190528_215904_1559055544171397763 ...
type Schema_init_20190528_215904_1559055544171397763 struct {
	migration.Migration
}

func init() {
	// will be some migration with up and down feature
	m := &Schema_init_20190528_215904_1559055544171397763{}
	migration.Register("Schema_init_20190528_215904_1559055544171397763", m)
}

// Up is migration up
func (m *Schema_init_20190528_215904_1559055544171397763) Up() {
	m.SQL(`
	create table if not exists user_account
	(
		account_number bigint not null
			constraint user_account_pk
				unique,
		balance numeric(11,2) default 0.00 not null,
		modified timestamp with time zone default CURRENT_TIMESTAMP not null
	);


	create table if not exists transaction_log
	(
		tx_id varchar(50) not null,
		account_number bigint not null,
		date timestamp with time zone not null,
		transaction_code varchar(10) not null,
		previous_balance double precision not null,
		end_balance double precision not null,
		opponent_account_number bigint not null,
		transaction_status integer not null,
		dr_cr varchar(2) not null,
		information text,
		last_update timestamp with time zone not null,
		amount double precision not null,
		constraint transaction_log_pk
			unique (tx_id, account_number)
	);

	create table if not exists transaction_history
	(
		tx_id varchar(50) not null,
		account_number bigint not null,
		transaction_status integer not null,
		dr_cr varchar(2) not null,
		previous_balance double precision not null,
		end_balance double precision not null,
		transaction_code varchar(10) not null,
		information text not null,
		user_comment text,
		transaction_date timestamp with time zone not null,
		amount double precision not null,
		opponent_account bigint not null,
		constraint transaction_history_pk
			unique (tx_id, account_number)
	);
	`)
}

// Down migration down
func (m *Schema_init_20190528_215904_1559055544171397763) Down() {
	m.SQL(`
	drop table transaction_history;
	drop table transaction_log;
	drop table user_account;
	`)
}
