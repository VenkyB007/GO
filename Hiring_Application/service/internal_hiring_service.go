package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VenkyB007/go/Hiring_Application/domain"
)

type Candidate struct {
	CandidateId int      `json:"candidateid"`
	Fullname    string   `json:"fullname"`
	Skillsets   []string `json:"skillsets"`
}

var internalcandidates []Candidate

// domain.Candidate details for internal hiring

func InternalPostRequirement(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var internalcandidate Candidate
	json.NewDecoder(r.Body).Decode(&internalcandidate)
	internalcandidates = append(internalcandidates, internalcandidate)

	fmt.Fprintf(w, "Success")
}

// Getting Internal domain.Candidate details

func InternalAllCandidates(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(internalcandidates)
}

//Matching the skills with domain.Candidate Skills

var requirementint domain.RequirementInt

func InternalMatchRequirement(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&requirementint)

	var matched []Candidate
	var matchedext []ExtCandidate
	for i := 0; i < len(internalcandidates); i++ {
		if internalcandidates[i].MatchBothSkills() == true {
			matched = append(matched, internalcandidates[i])
		}
	}
	for i := 0; i < len(externalcandidates); i++ {
		if externalcandidates[i].ExternalHiring() == true {
			matchedext = append(matchedext, externalcandidates[i])
		}
	}
	if len(matched) == 0 {
		json.NewEncoder(w).Encode(matchedext)
	} else {
		json.NewEncoder(w).Encode(matched)
	}

}

func (cand *Candidate) MatchBothSkills() bool {

	c := 0
	for j := 0; j < len(cand.Skillsets); j++ {
		for k := 0; k < len(requirementint.RequiredSkillsets); k++ {
			if cand.Skillsets[j] == requirementint.RequiredSkillsets[k] {
				c++
			}
		}
		for k := 0; k < len(requirementint.OptionalSkillsets); k++ {
			if cand.Skillsets[j] == requirementint.OptionalSkillsets[k] {
				c++
			}
		}
	}
	if c == (len(requirementint.RequiredSkillsets) + len(requirementint.OptionalSkillsets)) {
		return true
	} else {
		c := 0
		for j := 0; j < len(cand.Skillsets); j++ {
			for k := 0; k < len(requirementint.RequiredSkillsets); k++ {
				if cand.Skillsets[j] == requirementint.RequiredSkillsets[k] {
					c++
				}
			}
			if c == len(requirementint.RequiredSkillsets) {
				return true
			}
		}
	}
	return false
}

func (extcand *ExtCandidate) ExternalHiring() bool {

	c := 0
	for j := 0; j < len(extcand.Skillsets); j++ {
		for k := 0; k < len(requirementint.RequiredSkillsets); k++ {
			if extcand.Skillsets[j] == requirementint.RequiredSkillsets[k] {
				c++
			}
		}
		for k := 0; k < len(requirementint.OptionalSkillsets); k++ {
			if extcand.Skillsets[j] == requirementint.OptionalSkillsets[k] {
				c++
			}
		}
	}
	if c == (len(requirementint.RequiredSkillsets) + len(requirementint.OptionalSkillsets)) {
		return true
	}
	return false
}
