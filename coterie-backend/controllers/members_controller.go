package controllers

import (
	"coterie/models"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/qkgo/yin"
)

//NESTED

//GetMembers will retrieve all members for an org
func GetMembers(memberTable *models.MemberTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		organizationID := chi.URLParam(r, "organizationID")

		members, err := memberTable.MembersLister(organizationID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(members)
	}
}

//AddMember adds new member for a given org
func AddMember(memberTable *models.MemberTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		body := map[string]string{}
		req.BindBody(&body)
		organizationID := chi.URLParam(r, "organizationID")

		orgID, _ := strconv.Atoi(organizationID)
		funds, _ := strconv.Atoi(body["funds_raised"])
		member := models.Member{
			Name:           body["name"],
			Birthdate:      body["birthdate"],
			Class:          body["class"],
			Email:          body["email"],
			FundsRaised:    funds,
			OrganizationID: orgID,
			CreatedAt:      time.Now().String(),
			UpdatedAt:      time.Now().String(),
		}

		result, err := memberTable.MemberAdder(member)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(result)
	}
}

//UNNESTED

//GetMember is show action
func GetMember(memberTable *models.MemberTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		memberID := chi.URLParam(r, "memberID")

		member, err := memberTable.MemberGetter(memberID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(member)
	}
}

//UpdateMember is update action
func UpdateMember(memberTable *models.MemberTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, req := yin.Event(w, r)
		memberID := chi.URLParam(r, "memberID")
		body := map[string]string{}
		req.BindBody(&body)

		memID, _ := strconv.Atoi(memberID)
		funds, _ := strconv.Atoi(body["funds_raised"])
		member := models.Member{
			ID:          memID,
			Class:       body["class"],
			Email:       body["email"],
			FundsRaised: funds,
			UpdatedAt:   time.Now().String(),
		}

		result, err := memberTable.MemberUpdater(member)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		res.SendJSON(result)
	}
}

//DeleteMember is destroy action
func DeleteMember(memberTable *models.MemberTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		memberID := chi.URLParam(r, "memberID")

		err := memberTable.MemberDeleter(memberID)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		res.SendStatus(200)
	}
}
