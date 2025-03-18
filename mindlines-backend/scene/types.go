package scene

import (
	"fmt"
	"github.com/fredericobormann/mindlines-web/mindlines-backend/helper"
	"github.com/open-spaced-repetition/go-fsrs/v3"
	"time"
)

type MetaScene struct {
	Name       string
	Identifier string
	Index      uint8
}

type Scene struct {
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
	Index      uint8  `json:"index"`
	Content    []Line `json:"content"`
}

type Line struct {
	Character string    `json:"character"`
	Line      string    `json:"line"`
	Card      fsrs.Card `json:"card"`
}

type ReviewTimesDto struct {
	Again string `json:"again"`
	Hard  string `json:"hard"`
	Good  string `json:"good"`
	Easy  string `json:"easy"`
}

type LineDto struct {
	Character string `json:"character"`
	Line      string `json:"line"`
	//CorrectAnswers uint16 `json:"correctAnswers"`
	ReviewTimes ReviewTimesDto `json:"reviewTimes"`
	DueTime     time.Time      `json:"dueTime"`
}

type SceneDto struct {
	Name       string    `json:"name"`
	Identifier string    `json:"identifier"`
	Index      uint8     `json:"index"`
	Content    []LineDto `json:"content"`
}

type MetaSceneDto struct {
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
	Index      uint8  `json:"index"`
}

func (l Line) ToDto(fsrs2 fsrs.FSRS) LineDto {
	now := time.Now()
	repeatingCards := fsrs2.Repeat(l.Card, now)
	return LineDto{
		Character:   l.Character,
		Line:        l.Line,
		ReviewTimes: ReviewTimesDtoFromCard(repeatingCards, now),
		DueTime:     l.Card.Due,
	}
}

func SchedulingInfoToDueString(card fsrs.Card, now time.Time) string {
	if card.ScheduledDays == 0 {
		return card.Due.Sub(now).String()
	}
	return fmt.Sprintf("%d days", card.ScheduledDays)
}

func ReviewTimesDtoFromCard(schedulingInfo fsrs.RecordLog, now time.Time) ReviewTimesDto {
	return ReviewTimesDto{
		Again: SchedulingInfoToDueString(schedulingInfo[fsrs.Again].Card, now),
		Hard:  SchedulingInfoToDueString(schedulingInfo[fsrs.Hard].Card, now),
		Good:  SchedulingInfoToDueString(schedulingInfo[fsrs.Good].Card, now),
		Easy:  SchedulingInfoToDueString(schedulingInfo[fsrs.Easy].Card, now),
	}
}

func (s Scene) ToDto(fsrs2 fsrs.FSRS) SceneDto {
	return SceneDto{
		Name:       s.Name,
		Identifier: s.Identifier,
		Index:      s.Index,
		Content:    helper.Map(s.Content, func(el Line) LineDto { return el.ToDto(fsrs2) }),
	}
}

func (m MetaScene) ToDto() MetaSceneDto {
	return MetaSceneDto{
		Name:       m.Name,
		Identifier: m.Identifier,
		Index:      m.Index,
	}
}
