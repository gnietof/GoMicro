package members

import (
	model "gomicro/internal/model"

	"gomicro/internal/shared/db2"
)

type MembersRepository struct {
	DB *db2.DB2Client
}

func NewMembersRepository() *MembersRepository {
	db, _ := db2.NewDB2Client()
	return &MembersRepository{DB: db}
}

func (m *MembersRepository) GetMembers() ([]model.Member, error) {

	rows, _ := m.DB.DB.Query("SELECT ID,FIRSTNAME,LASTNAME,EMAIL FROM ER_POPULATIONS.MEMBERS")

	// var members []map[string]interface{}
	var members []model.Member

	for rows.Next() {
		var member model.Member
		rows.Scan(&member.Id, &member.FirstName, &member.LastName, &member.EMail)

		// result = append(result, map[string]interface{}{
		// 	"id":        id,
		// 	"firstname": firstname,
		// 	"lastname":  lastname,
		// 	"email":     email,
		// })
		members = append(members, member)

	}

	return members, nil
}

func (m *MembersRepository) GetMemberById(id string) (model.Member, error) {

	row := m.DB.DB.QueryRow("SELECT FIRSTNAME,LASTNAME,EMAIL FROM ER_POPULATIONS.MEMBERS WHERE ID=?", id)

	var member model.Member
	err := row.Scan(&member.FirstName, &member.LastName, &member.EMail)
	if err != nil {
		return member, err
	}

	// result := map[string]interface{}{
	// 	"id":        id,
	// 	"firstname": firstname,
	// 	"lastname":  lastname,
	// 	"email":     email,
	// }

	// json.NewEncoder(w).Encode(result)
	return member, nil
}
