/*
 * Copyright (c) 2015 Nhac Nguyen
 * Copyright (C) 2016 Samson Ugwuodo
 * Copyright (c) 2016 Jonathan Anderson
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package db

import (
	"database/sql"
	"fmt"
	"time"
//	"upper.io/db"
	 _"github.com/mattn/go-sqlite3"
)

type Reservation struct {
	Id      int
	db      *DB
	user    int
	machine int
	Start   time.Time
	End     time.Time
	Ended   time.Time
	PxePath string
	NfsRoot string
}

func initReservations(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE Reservations(
		id integer primary key,
		user integer,
		machine integer,
		start datetime not null,
		end datetime not null,
		ended datetime,
		pxepath text not null,
		nfsroot text not null,

		FOREIGN KEY(user) REFERENCES Users(id),
		FOREIGN KEY(machine) REFERENCES Machines(id)
	);
	`)

	return err
}

func (d DB) CreateReservation(machine string, user string,
	start time.Time, end time.Time, pxepath string, nfsroot string) error {
	       _, err := d.sql.Exec (`BEGIN;
	      		INSERT INTO Reservations (machine,user,start,end,pxepath,nfsroot)
			SELECT (SELECT id from Machines where name=?),(SELECT id from Users where username=?),?,?,?,?
			where not exists (select r.id from Reservations r 
			INNER JOIN Machines m 
			ON m.id = r.machine where m.name = ? AND (? >= start AND ? <= end) OR (? <= start AND ? >= end));COMMIT;`,machine,user,start,end,pxepath,nfsroot,machine,start,start,end,end)
	return err
}

func (d DB) GetReservations() ([]Reservation, error) {
	rows, err := d.sql.Query(`
		SELECT id,user, machine, start, end, ended, pxepath, nfsroot
		FROM Reservations
		ORDER BY end DESC
	
			`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	reservations := []Reservation{}
	for rows.Next() {
		r := Reservation{db: &d}
		//edited ended variable
		var ended *time.Time

		err = rows.Scan(&r.Id,&r.user, &r.machine,
			&r.Start, &r.End, &ended, &r.PxePath, &r.NfsRoot)
		if err != nil {
			return nil, err
		}

		if ended != nil {
			r.Ended = *ended
		}

		reservations = append(reservations, r)
	}

	return reservations, rows.Err()
}
//CREATING END-RESERVATION......
//
func (d DB) EndReservation(id int)error{

	now := time.Now()
	_, err := d.sql.Exec(`
			UPDATE Reservations
			SET ended = ?
			WHERE id = ?
	
	`,now, id)

	return err

}
//Filter reservation by dates
func(d DB) FilterByDates(start time.Time, end time.Time) ([]Reservation,error){
         rows, err := d.sql.Query(`
		SELECT id,user, machine, start, end, ended, pxepath, nfsroot
		FROM Reservations
		WHERE start > ? AND end < ?		
	`, start, end)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	filter_r := []Reservation{}
	for rows.Next() {
		r := Reservation{db: &d}
		var ended *time.Time

		err = rows.Scan(&r.Id,&r.user, &r.machine,
			&r.Start, &r.End, &ended, &r.PxePath, &r.NfsRoot)
		if err != nil {
			return nil, err
		}

		if ended != nil {
			r.Ended = *ended
		}

		filter_r = append(filter_r, r)
	}

	return filter_r, rows.Err()


}
func (d DB) Sort_End()([]Reservation, error){
	rows, err := d.sql.Query(`
		SELECT id,user, machine, start, end, ended, pxepath, nfsroot
		FROM Reservations
		ORDER BY end ASC 	
	`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	sort_e := []Reservation{}
	for rows.Next() {
		r := Reservation{db: &d}
		var ended *time.Time

		err = rows.Scan(&r.Id,&r.user, &r.machine,
			&r.Start, &r.End, &ended, &r.PxePath, &r.NfsRoot)
		if err != nil {
			return nil, err
		}

		if ended != nil {
			r.Ended = *ended
		}

		sort_e = append(sort_e, r)
	}

	return sort_e, rows.Err()



}
func (d DB) Sort_Ended()([]Reservation, error){
	rows, err := d.sql.Query(`
		SELECT id,user, machine, start, end, ended, pxepath, nfsroot
		FROM Reservations
		ORDER BY ended ASC 	
	`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	sort_e := []Reservation{}
	for rows.Next() {
		r := Reservation{db: &d}
		var ended *time.Time

		err = rows.Scan(&r.Id,&r.user, &r.machine,
			&r.Start, &r.End, &ended, &r.PxePath, &r.NfsRoot)
		if err != nil {
			return nil, err
		}

		if ended != nil {
			r.Ended = *ended
		}

		sort_e = append(sort_e, r)
	}

	return sort_e, rows.Err()

}

func (d DB) Sort_Start()([]Reservation, error){
	rows, err := d.sql.Query(`
		SELECT id,user, machine, start, end, ended, pxepath, nfsroot
		FROM Reservations
		ORDER BY start DESC 
	`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	sort_start := []Reservation{}
	for rows.Next() {
		r := Reservation{db: &d}
		var ended *time.Time

		err = rows.Scan(&r.Id,&r.user, &r.machine,
			&r.Start, &r.End, &ended, &r.PxePath, &r.NfsRoot)
		if err != nil {
			return nil, err
		}

		if ended != nil {
			r.Ended = *ended
		}

		sort_start = append(sort_start, r)
	}

	return sort_start, rows.Err()



}

func (d DB) Sort_By_Name()([]Reservation, error){
	rows, err := d.sql.Query(`
		SELECT id,user, machine, start, end, ended, pxepath, nfsroot
		FROM Reservations
		ORDER BY LOWER (machine) COLLATE NOCASE ASC	
	`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	sort_start := []Reservation{}
	for rows.Next() {
		r := Reservation{db: &d}
		var ended *time.Time

		err = rows.Scan(&r.Id,&r.user, &r.machine,
			&r.Start, &r.End, &ended, &r.PxePath, &r.NfsRoot)
		if err != nil {
			return nil, err
		}

		if ended != nil {
			r.Ended = *ended
		}

		sort_start = append(sort_start, r)
	}

	return sort_start, rows.Err()



}
func (d DB) Filter_By_Pxe_Nfs(pxe string, nfs string)([]Reservation, error){
	rows, err := d.sql.Query(`
		SELECT id,user, machine, start, end, ended, pxepath, nfsroot
		FROM Reservations
		WHERE pxepath LIKE ?||'%' OR nfsroot LIKE ?||'%'
	`, pxe, nfs)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	filter_pn := []Reservation{}
	for rows.Next() {
		r := Reservation{db: &d}
		var ended *time.Time

		err = rows.Scan(&r.Id,&r.user, &r.machine,
			&r.Start, &r.End, &ended, &r.PxePath, &r.NfsRoot)
		if err != nil {
			return nil, err
		}

		if ended != nil {
			r.Ended = *ended
		}

		filter_pn = append(filter_pn, r)
	}

	return filter_pn, rows.Err()

}

func (d DB) GetReservationsFor(col string, id int,
	start time.Time, end time.Time) ([]Reservation, error) {

	var err error
	if end.IsZero() {
		end, err = time.Parse("02 Jan 2006", "01 Jan 3000")
		if err != nil {
			return nil, err
		}
	}

	rows, err := d.sql.Query(`
		SELECT user, machine, start, end, ended, pxepath, nfsroot
		FROM Reservations
		WHERE `+col+` = ? AND end >= ? AND start <= ?
		ORDER BY end DESC
	`, id, start, end)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	reservations := []Reservation{}
	for rows.Next() {
		r := Reservation{db: &d}
		var ended *time.Time

		err = rows.Scan(&r.user, &r.machine,
			&r.Start, &r.End, &ended, &r.PxePath, &r.NfsRoot)
		if err != nil {
			return nil, err
		}

		if ended != nil {
			r.Ended = *ended
		}

		reservations = append(reservations, r)
	}

	return reservations, rows.Err()
}

func (r Reservation) User() (*User, error) {
	return r.db.GetUser(r.user)
}

func (r Reservation) Machine() (*Machine, error) {
	return r.db.GetMachine("id", r.machine)
}
func (r Reservation) String() string {
	end := r.End
	if !r.Ended.IsZero() {
		end = r.Ended
	}

	var username string
	user, err := r.User()
	if err != nil {
		username = fmt.Sprintf("<error: %s>", err)
	} else {
		username = user.Username
	}

	var machineName string
	machine, err := r.Machine()
	if err != nil {
		machineName = fmt.Sprintf("<error: %s>", err)
	} else {
		machineName = machine.Name
	}

	return fmt.Sprintf("%-12s %-8s %12s to %12s  %-s",
		machineName, username,
		r.Start.Format("1504h 02 Jan"),
		end.Format("1504h 02 Jan"),
		r.NfsRoot)
}
