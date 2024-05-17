// Code generated by ent, DO NOT EDIT.

package thesis

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ZEQUANR/zhulong/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Thesis {
	return predicate.Thesis(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Thesis {
	return predicate.Thesis(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Thesis {
	return predicate.Thesis(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Thesis {
	return predicate.Thesis(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Thesis {
	return predicate.Thesis(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Thesis {
	return predicate.Thesis(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Thesis {
	return predicate.Thesis(sql.FieldLTE(FieldID, id))
}

// FileName applies equality check predicate on the "file_name" field. It's identical to FileNameEQ.
func FileName(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldFileName, v))
}

// FileURL applies equality check predicate on the "file_url" field. It's identical to FileURLEQ.
func FileURL(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldFileURL, v))
}

// FileState applies equality check predicate on the "file_state" field. It's identical to FileStateEQ.
func FileState(v int) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldFileState, v))
}

// UploadTime applies equality check predicate on the "upload_time" field. It's identical to UploadTimeEQ.
func UploadTime(v time.Time) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldUploadTime, v))
}

// ChineseTitle applies equality check predicate on the "chinese_title" field. It's identical to ChineseTitleEQ.
func ChineseTitle(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldChineseTitle, v))
}

// EnglishTitle applies equality check predicate on the "english_title" field. It's identical to EnglishTitleEQ.
func EnglishTitle(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldEnglishTitle, v))
}

// Authors applies equality check predicate on the "authors" field. It's identical to AuthorsEQ.
func Authors(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldAuthors, v))
}

// Teachers applies equality check predicate on the "teachers" field. It's identical to TeachersEQ.
func Teachers(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldTeachers, v))
}

// FirstAdvance applies equality check predicate on the "first_advance" field. It's identical to FirstAdvanceEQ.
func FirstAdvance(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldFirstAdvance, v))
}

// SecondAdvance applies equality check predicate on the "second_advance" field. It's identical to SecondAdvanceEQ.
func SecondAdvance(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldSecondAdvance, v))
}

// ThirdAdvance applies equality check predicate on the "third_advance" field. It's identical to ThirdAdvanceEQ.
func ThirdAdvance(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldThirdAdvance, v))
}

// Drawback applies equality check predicate on the "drawback" field. It's identical to DrawbackEQ.
func Drawback(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldDrawback, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldCreateTime, v))
}

// FileNameEQ applies the EQ predicate on the "file_name" field.
func FileNameEQ(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldFileName, v))
}

// FileNameNEQ applies the NEQ predicate on the "file_name" field.
func FileNameNEQ(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldNEQ(FieldFileName, v))
}

// FileNameIn applies the In predicate on the "file_name" field.
func FileNameIn(vs ...string) predicate.Thesis {
	return predicate.Thesis(sql.FieldIn(FieldFileName, vs...))
}

// FileNameNotIn applies the NotIn predicate on the "file_name" field.
func FileNameNotIn(vs ...string) predicate.Thesis {
	return predicate.Thesis(sql.FieldNotIn(FieldFileName, vs...))
}

// FileNameGT applies the GT predicate on the "file_name" field.
func FileNameGT(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldGT(FieldFileName, v))
}

// FileNameGTE applies the GTE predicate on the "file_name" field.
func FileNameGTE(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldGTE(FieldFileName, v))
}

// FileNameLT applies the LT predicate on the "file_name" field.
func FileNameLT(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldLT(FieldFileName, v))
}

// FileNameLTE applies the LTE predicate on the "file_name" field.
func FileNameLTE(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldLTE(FieldFileName, v))
}

// FileNameContains applies the Contains predicate on the "file_name" field.
func FileNameContains(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldContains(FieldFileName, v))
}

// FileNameHasPrefix applies the HasPrefix predicate on the "file_name" field.
func FileNameHasPrefix(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldHasPrefix(FieldFileName, v))
}

// FileNameHasSuffix applies the HasSuffix predicate on the "file_name" field.
func FileNameHasSuffix(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldHasSuffix(FieldFileName, v))
}

// FileNameIsNil applies the IsNil predicate on the "file_name" field.
func FileNameIsNil() predicate.Thesis {
	return predicate.Thesis(sql.FieldIsNull(FieldFileName))
}

// FileNameNotNil applies the NotNil predicate on the "file_name" field.
func FileNameNotNil() predicate.Thesis {
	return predicate.Thesis(sql.FieldNotNull(FieldFileName))
}

// FileNameEqualFold applies the EqualFold predicate on the "file_name" field.
func FileNameEqualFold(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEqualFold(FieldFileName, v))
}

// FileNameContainsFold applies the ContainsFold predicate on the "file_name" field.
func FileNameContainsFold(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldContainsFold(FieldFileName, v))
}

// FileURLEQ applies the EQ predicate on the "file_url" field.
func FileURLEQ(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldFileURL, v))
}

// FileURLNEQ applies the NEQ predicate on the "file_url" field.
func FileURLNEQ(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldNEQ(FieldFileURL, v))
}

// FileURLIn applies the In predicate on the "file_url" field.
func FileURLIn(vs ...string) predicate.Thesis {
	return predicate.Thesis(sql.FieldIn(FieldFileURL, vs...))
}

// FileURLNotIn applies the NotIn predicate on the "file_url" field.
func FileURLNotIn(vs ...string) predicate.Thesis {
	return predicate.Thesis(sql.FieldNotIn(FieldFileURL, vs...))
}

// FileURLGT applies the GT predicate on the "file_url" field.
func FileURLGT(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldGT(FieldFileURL, v))
}

// FileURLGTE applies the GTE predicate on the "file_url" field.
func FileURLGTE(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldGTE(FieldFileURL, v))
}

// FileURLLT applies the LT predicate on the "file_url" field.
func FileURLLT(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldLT(FieldFileURL, v))
}

// FileURLLTE applies the LTE predicate on the "file_url" field.
func FileURLLTE(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldLTE(FieldFileURL, v))
}

// FileURLContains applies the Contains predicate on the "file_url" field.
func FileURLContains(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldContains(FieldFileURL, v))
}

// FileURLHasPrefix applies the HasPrefix predicate on the "file_url" field.
func FileURLHasPrefix(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldHasPrefix(FieldFileURL, v))
}

// FileURLHasSuffix applies the HasSuffix predicate on the "file_url" field.
func FileURLHasSuffix(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldHasSuffix(FieldFileURL, v))
}

// FileURLIsNil applies the IsNil predicate on the "file_url" field.
func FileURLIsNil() predicate.Thesis {
	return predicate.Thesis(sql.FieldIsNull(FieldFileURL))
}

// FileURLNotNil applies the NotNil predicate on the "file_url" field.
func FileURLNotNil() predicate.Thesis {
	return predicate.Thesis(sql.FieldNotNull(FieldFileURL))
}

// FileURLEqualFold applies the EqualFold predicate on the "file_url" field.
func FileURLEqualFold(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEqualFold(FieldFileURL, v))
}

// FileURLContainsFold applies the ContainsFold predicate on the "file_url" field.
func FileURLContainsFold(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldContainsFold(FieldFileURL, v))
}

// FileStateEQ applies the EQ predicate on the "file_state" field.
func FileStateEQ(v int) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldFileState, v))
}

// FileStateNEQ applies the NEQ predicate on the "file_state" field.
func FileStateNEQ(v int) predicate.Thesis {
	return predicate.Thesis(sql.FieldNEQ(FieldFileState, v))
}

// FileStateIn applies the In predicate on the "file_state" field.
func FileStateIn(vs ...int) predicate.Thesis {
	return predicate.Thesis(sql.FieldIn(FieldFileState, vs...))
}

// FileStateNotIn applies the NotIn predicate on the "file_state" field.
func FileStateNotIn(vs ...int) predicate.Thesis {
	return predicate.Thesis(sql.FieldNotIn(FieldFileState, vs...))
}

// FileStateGT applies the GT predicate on the "file_state" field.
func FileStateGT(v int) predicate.Thesis {
	return predicate.Thesis(sql.FieldGT(FieldFileState, v))
}

// FileStateGTE applies the GTE predicate on the "file_state" field.
func FileStateGTE(v int) predicate.Thesis {
	return predicate.Thesis(sql.FieldGTE(FieldFileState, v))
}

// FileStateLT applies the LT predicate on the "file_state" field.
func FileStateLT(v int) predicate.Thesis {
	return predicate.Thesis(sql.FieldLT(FieldFileState, v))
}

// FileStateLTE applies the LTE predicate on the "file_state" field.
func FileStateLTE(v int) predicate.Thesis {
	return predicate.Thesis(sql.FieldLTE(FieldFileState, v))
}

// FileStateIsNil applies the IsNil predicate on the "file_state" field.
func FileStateIsNil() predicate.Thesis {
	return predicate.Thesis(sql.FieldIsNull(FieldFileState))
}

// FileStateNotNil applies the NotNil predicate on the "file_state" field.
func FileStateNotNil() predicate.Thesis {
	return predicate.Thesis(sql.FieldNotNull(FieldFileState))
}

// UploadTimeEQ applies the EQ predicate on the "upload_time" field.
func UploadTimeEQ(v time.Time) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldUploadTime, v))
}

// UploadTimeNEQ applies the NEQ predicate on the "upload_time" field.
func UploadTimeNEQ(v time.Time) predicate.Thesis {
	return predicate.Thesis(sql.FieldNEQ(FieldUploadTime, v))
}

// UploadTimeIn applies the In predicate on the "upload_time" field.
func UploadTimeIn(vs ...time.Time) predicate.Thesis {
	return predicate.Thesis(sql.FieldIn(FieldUploadTime, vs...))
}

// UploadTimeNotIn applies the NotIn predicate on the "upload_time" field.
func UploadTimeNotIn(vs ...time.Time) predicate.Thesis {
	return predicate.Thesis(sql.FieldNotIn(FieldUploadTime, vs...))
}

// UploadTimeGT applies the GT predicate on the "upload_time" field.
func UploadTimeGT(v time.Time) predicate.Thesis {
	return predicate.Thesis(sql.FieldGT(FieldUploadTime, v))
}

// UploadTimeGTE applies the GTE predicate on the "upload_time" field.
func UploadTimeGTE(v time.Time) predicate.Thesis {
	return predicate.Thesis(sql.FieldGTE(FieldUploadTime, v))
}

// UploadTimeLT applies the LT predicate on the "upload_time" field.
func UploadTimeLT(v time.Time) predicate.Thesis {
	return predicate.Thesis(sql.FieldLT(FieldUploadTime, v))
}

// UploadTimeLTE applies the LTE predicate on the "upload_time" field.
func UploadTimeLTE(v time.Time) predicate.Thesis {
	return predicate.Thesis(sql.FieldLTE(FieldUploadTime, v))
}

// UploadTimeIsNil applies the IsNil predicate on the "upload_time" field.
func UploadTimeIsNil() predicate.Thesis {
	return predicate.Thesis(sql.FieldIsNull(FieldUploadTime))
}

// UploadTimeNotNil applies the NotNil predicate on the "upload_time" field.
func UploadTimeNotNil() predicate.Thesis {
	return predicate.Thesis(sql.FieldNotNull(FieldUploadTime))
}

// ChineseTitleEQ applies the EQ predicate on the "chinese_title" field.
func ChineseTitleEQ(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldChineseTitle, v))
}

// ChineseTitleNEQ applies the NEQ predicate on the "chinese_title" field.
func ChineseTitleNEQ(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldNEQ(FieldChineseTitle, v))
}

// ChineseTitleIn applies the In predicate on the "chinese_title" field.
func ChineseTitleIn(vs ...string) predicate.Thesis {
	return predicate.Thesis(sql.FieldIn(FieldChineseTitle, vs...))
}

// ChineseTitleNotIn applies the NotIn predicate on the "chinese_title" field.
func ChineseTitleNotIn(vs ...string) predicate.Thesis {
	return predicate.Thesis(sql.FieldNotIn(FieldChineseTitle, vs...))
}

// ChineseTitleGT applies the GT predicate on the "chinese_title" field.
func ChineseTitleGT(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldGT(FieldChineseTitle, v))
}

// ChineseTitleGTE applies the GTE predicate on the "chinese_title" field.
func ChineseTitleGTE(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldGTE(FieldChineseTitle, v))
}

// ChineseTitleLT applies the LT predicate on the "chinese_title" field.
func ChineseTitleLT(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldLT(FieldChineseTitle, v))
}

// ChineseTitleLTE applies the LTE predicate on the "chinese_title" field.
func ChineseTitleLTE(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldLTE(FieldChineseTitle, v))
}

// ChineseTitleContains applies the Contains predicate on the "chinese_title" field.
func ChineseTitleContains(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldContains(FieldChineseTitle, v))
}

// ChineseTitleHasPrefix applies the HasPrefix predicate on the "chinese_title" field.
func ChineseTitleHasPrefix(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldHasPrefix(FieldChineseTitle, v))
}

// ChineseTitleHasSuffix applies the HasSuffix predicate on the "chinese_title" field.
func ChineseTitleHasSuffix(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldHasSuffix(FieldChineseTitle, v))
}

// ChineseTitleEqualFold applies the EqualFold predicate on the "chinese_title" field.
func ChineseTitleEqualFold(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEqualFold(FieldChineseTitle, v))
}

// ChineseTitleContainsFold applies the ContainsFold predicate on the "chinese_title" field.
func ChineseTitleContainsFold(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldContainsFold(FieldChineseTitle, v))
}

// EnglishTitleEQ applies the EQ predicate on the "english_title" field.
func EnglishTitleEQ(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldEnglishTitle, v))
}

// EnglishTitleNEQ applies the NEQ predicate on the "english_title" field.
func EnglishTitleNEQ(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldNEQ(FieldEnglishTitle, v))
}

// EnglishTitleIn applies the In predicate on the "english_title" field.
func EnglishTitleIn(vs ...string) predicate.Thesis {
	return predicate.Thesis(sql.FieldIn(FieldEnglishTitle, vs...))
}

// EnglishTitleNotIn applies the NotIn predicate on the "english_title" field.
func EnglishTitleNotIn(vs ...string) predicate.Thesis {
	return predicate.Thesis(sql.FieldNotIn(FieldEnglishTitle, vs...))
}

// EnglishTitleGT applies the GT predicate on the "english_title" field.
func EnglishTitleGT(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldGT(FieldEnglishTitle, v))
}

// EnglishTitleGTE applies the GTE predicate on the "english_title" field.
func EnglishTitleGTE(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldGTE(FieldEnglishTitle, v))
}

// EnglishTitleLT applies the LT predicate on the "english_title" field.
func EnglishTitleLT(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldLT(FieldEnglishTitle, v))
}

// EnglishTitleLTE applies the LTE predicate on the "english_title" field.
func EnglishTitleLTE(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldLTE(FieldEnglishTitle, v))
}

// EnglishTitleContains applies the Contains predicate on the "english_title" field.
func EnglishTitleContains(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldContains(FieldEnglishTitle, v))
}

// EnglishTitleHasPrefix applies the HasPrefix predicate on the "english_title" field.
func EnglishTitleHasPrefix(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldHasPrefix(FieldEnglishTitle, v))
}

// EnglishTitleHasSuffix applies the HasSuffix predicate on the "english_title" field.
func EnglishTitleHasSuffix(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldHasSuffix(FieldEnglishTitle, v))
}

// EnglishTitleEqualFold applies the EqualFold predicate on the "english_title" field.
func EnglishTitleEqualFold(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEqualFold(FieldEnglishTitle, v))
}

// EnglishTitleContainsFold applies the ContainsFold predicate on the "english_title" field.
func EnglishTitleContainsFold(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldContainsFold(FieldEnglishTitle, v))
}

// AuthorsEQ applies the EQ predicate on the "authors" field.
func AuthorsEQ(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldAuthors, v))
}

// AuthorsNEQ applies the NEQ predicate on the "authors" field.
func AuthorsNEQ(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldNEQ(FieldAuthors, v))
}

// AuthorsIn applies the In predicate on the "authors" field.
func AuthorsIn(vs ...string) predicate.Thesis {
	return predicate.Thesis(sql.FieldIn(FieldAuthors, vs...))
}

// AuthorsNotIn applies the NotIn predicate on the "authors" field.
func AuthorsNotIn(vs ...string) predicate.Thesis {
	return predicate.Thesis(sql.FieldNotIn(FieldAuthors, vs...))
}

// AuthorsGT applies the GT predicate on the "authors" field.
func AuthorsGT(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldGT(FieldAuthors, v))
}

// AuthorsGTE applies the GTE predicate on the "authors" field.
func AuthorsGTE(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldGTE(FieldAuthors, v))
}

// AuthorsLT applies the LT predicate on the "authors" field.
func AuthorsLT(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldLT(FieldAuthors, v))
}

// AuthorsLTE applies the LTE predicate on the "authors" field.
func AuthorsLTE(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldLTE(FieldAuthors, v))
}

// AuthorsContains applies the Contains predicate on the "authors" field.
func AuthorsContains(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldContains(FieldAuthors, v))
}

// AuthorsHasPrefix applies the HasPrefix predicate on the "authors" field.
func AuthorsHasPrefix(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldHasPrefix(FieldAuthors, v))
}

// AuthorsHasSuffix applies the HasSuffix predicate on the "authors" field.
func AuthorsHasSuffix(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldHasSuffix(FieldAuthors, v))
}

// AuthorsEqualFold applies the EqualFold predicate on the "authors" field.
func AuthorsEqualFold(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEqualFold(FieldAuthors, v))
}

// AuthorsContainsFold applies the ContainsFold predicate on the "authors" field.
func AuthorsContainsFold(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldContainsFold(FieldAuthors, v))
}

// TeachersEQ applies the EQ predicate on the "teachers" field.
func TeachersEQ(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldTeachers, v))
}

// TeachersNEQ applies the NEQ predicate on the "teachers" field.
func TeachersNEQ(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldNEQ(FieldTeachers, v))
}

// TeachersIn applies the In predicate on the "teachers" field.
func TeachersIn(vs ...string) predicate.Thesis {
	return predicate.Thesis(sql.FieldIn(FieldTeachers, vs...))
}

// TeachersNotIn applies the NotIn predicate on the "teachers" field.
func TeachersNotIn(vs ...string) predicate.Thesis {
	return predicate.Thesis(sql.FieldNotIn(FieldTeachers, vs...))
}

// TeachersGT applies the GT predicate on the "teachers" field.
func TeachersGT(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldGT(FieldTeachers, v))
}

// TeachersGTE applies the GTE predicate on the "teachers" field.
func TeachersGTE(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldGTE(FieldTeachers, v))
}

// TeachersLT applies the LT predicate on the "teachers" field.
func TeachersLT(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldLT(FieldTeachers, v))
}

// TeachersLTE applies the LTE predicate on the "teachers" field.
func TeachersLTE(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldLTE(FieldTeachers, v))
}

// TeachersContains applies the Contains predicate on the "teachers" field.
func TeachersContains(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldContains(FieldTeachers, v))
}

// TeachersHasPrefix applies the HasPrefix predicate on the "teachers" field.
func TeachersHasPrefix(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldHasPrefix(FieldTeachers, v))
}

// TeachersHasSuffix applies the HasSuffix predicate on the "teachers" field.
func TeachersHasSuffix(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldHasSuffix(FieldTeachers, v))
}

// TeachersEqualFold applies the EqualFold predicate on the "teachers" field.
func TeachersEqualFold(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEqualFold(FieldTeachers, v))
}

// TeachersContainsFold applies the ContainsFold predicate on the "teachers" field.
func TeachersContainsFold(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldContainsFold(FieldTeachers, v))
}

// FirstAdvanceEQ applies the EQ predicate on the "first_advance" field.
func FirstAdvanceEQ(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldFirstAdvance, v))
}

// FirstAdvanceNEQ applies the NEQ predicate on the "first_advance" field.
func FirstAdvanceNEQ(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldNEQ(FieldFirstAdvance, v))
}

// FirstAdvanceIn applies the In predicate on the "first_advance" field.
func FirstAdvanceIn(vs ...string) predicate.Thesis {
	return predicate.Thesis(sql.FieldIn(FieldFirstAdvance, vs...))
}

// FirstAdvanceNotIn applies the NotIn predicate on the "first_advance" field.
func FirstAdvanceNotIn(vs ...string) predicate.Thesis {
	return predicate.Thesis(sql.FieldNotIn(FieldFirstAdvance, vs...))
}

// FirstAdvanceGT applies the GT predicate on the "first_advance" field.
func FirstAdvanceGT(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldGT(FieldFirstAdvance, v))
}

// FirstAdvanceGTE applies the GTE predicate on the "first_advance" field.
func FirstAdvanceGTE(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldGTE(FieldFirstAdvance, v))
}

// FirstAdvanceLT applies the LT predicate on the "first_advance" field.
func FirstAdvanceLT(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldLT(FieldFirstAdvance, v))
}

// FirstAdvanceLTE applies the LTE predicate on the "first_advance" field.
func FirstAdvanceLTE(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldLTE(FieldFirstAdvance, v))
}

// FirstAdvanceContains applies the Contains predicate on the "first_advance" field.
func FirstAdvanceContains(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldContains(FieldFirstAdvance, v))
}

// FirstAdvanceHasPrefix applies the HasPrefix predicate on the "first_advance" field.
func FirstAdvanceHasPrefix(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldHasPrefix(FieldFirstAdvance, v))
}

// FirstAdvanceHasSuffix applies the HasSuffix predicate on the "first_advance" field.
func FirstAdvanceHasSuffix(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldHasSuffix(FieldFirstAdvance, v))
}

// FirstAdvanceEqualFold applies the EqualFold predicate on the "first_advance" field.
func FirstAdvanceEqualFold(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEqualFold(FieldFirstAdvance, v))
}

// FirstAdvanceContainsFold applies the ContainsFold predicate on the "first_advance" field.
func FirstAdvanceContainsFold(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldContainsFold(FieldFirstAdvance, v))
}

// SecondAdvanceEQ applies the EQ predicate on the "second_advance" field.
func SecondAdvanceEQ(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldSecondAdvance, v))
}

// SecondAdvanceNEQ applies the NEQ predicate on the "second_advance" field.
func SecondAdvanceNEQ(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldNEQ(FieldSecondAdvance, v))
}

// SecondAdvanceIn applies the In predicate on the "second_advance" field.
func SecondAdvanceIn(vs ...string) predicate.Thesis {
	return predicate.Thesis(sql.FieldIn(FieldSecondAdvance, vs...))
}

// SecondAdvanceNotIn applies the NotIn predicate on the "second_advance" field.
func SecondAdvanceNotIn(vs ...string) predicate.Thesis {
	return predicate.Thesis(sql.FieldNotIn(FieldSecondAdvance, vs...))
}

// SecondAdvanceGT applies the GT predicate on the "second_advance" field.
func SecondAdvanceGT(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldGT(FieldSecondAdvance, v))
}

// SecondAdvanceGTE applies the GTE predicate on the "second_advance" field.
func SecondAdvanceGTE(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldGTE(FieldSecondAdvance, v))
}

// SecondAdvanceLT applies the LT predicate on the "second_advance" field.
func SecondAdvanceLT(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldLT(FieldSecondAdvance, v))
}

// SecondAdvanceLTE applies the LTE predicate on the "second_advance" field.
func SecondAdvanceLTE(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldLTE(FieldSecondAdvance, v))
}

// SecondAdvanceContains applies the Contains predicate on the "second_advance" field.
func SecondAdvanceContains(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldContains(FieldSecondAdvance, v))
}

// SecondAdvanceHasPrefix applies the HasPrefix predicate on the "second_advance" field.
func SecondAdvanceHasPrefix(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldHasPrefix(FieldSecondAdvance, v))
}

// SecondAdvanceHasSuffix applies the HasSuffix predicate on the "second_advance" field.
func SecondAdvanceHasSuffix(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldHasSuffix(FieldSecondAdvance, v))
}

// SecondAdvanceEqualFold applies the EqualFold predicate on the "second_advance" field.
func SecondAdvanceEqualFold(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEqualFold(FieldSecondAdvance, v))
}

// SecondAdvanceContainsFold applies the ContainsFold predicate on the "second_advance" field.
func SecondAdvanceContainsFold(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldContainsFold(FieldSecondAdvance, v))
}

// ThirdAdvanceEQ applies the EQ predicate on the "third_advance" field.
func ThirdAdvanceEQ(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldThirdAdvance, v))
}

// ThirdAdvanceNEQ applies the NEQ predicate on the "third_advance" field.
func ThirdAdvanceNEQ(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldNEQ(FieldThirdAdvance, v))
}

// ThirdAdvanceIn applies the In predicate on the "third_advance" field.
func ThirdAdvanceIn(vs ...string) predicate.Thesis {
	return predicate.Thesis(sql.FieldIn(FieldThirdAdvance, vs...))
}

// ThirdAdvanceNotIn applies the NotIn predicate on the "third_advance" field.
func ThirdAdvanceNotIn(vs ...string) predicate.Thesis {
	return predicate.Thesis(sql.FieldNotIn(FieldThirdAdvance, vs...))
}

// ThirdAdvanceGT applies the GT predicate on the "third_advance" field.
func ThirdAdvanceGT(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldGT(FieldThirdAdvance, v))
}

// ThirdAdvanceGTE applies the GTE predicate on the "third_advance" field.
func ThirdAdvanceGTE(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldGTE(FieldThirdAdvance, v))
}

// ThirdAdvanceLT applies the LT predicate on the "third_advance" field.
func ThirdAdvanceLT(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldLT(FieldThirdAdvance, v))
}

// ThirdAdvanceLTE applies the LTE predicate on the "third_advance" field.
func ThirdAdvanceLTE(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldLTE(FieldThirdAdvance, v))
}

// ThirdAdvanceContains applies the Contains predicate on the "third_advance" field.
func ThirdAdvanceContains(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldContains(FieldThirdAdvance, v))
}

// ThirdAdvanceHasPrefix applies the HasPrefix predicate on the "third_advance" field.
func ThirdAdvanceHasPrefix(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldHasPrefix(FieldThirdAdvance, v))
}

// ThirdAdvanceHasSuffix applies the HasSuffix predicate on the "third_advance" field.
func ThirdAdvanceHasSuffix(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldHasSuffix(FieldThirdAdvance, v))
}

// ThirdAdvanceEqualFold applies the EqualFold predicate on the "third_advance" field.
func ThirdAdvanceEqualFold(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEqualFold(FieldThirdAdvance, v))
}

// ThirdAdvanceContainsFold applies the ContainsFold predicate on the "third_advance" field.
func ThirdAdvanceContainsFold(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldContainsFold(FieldThirdAdvance, v))
}

// DrawbackEQ applies the EQ predicate on the "drawback" field.
func DrawbackEQ(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldDrawback, v))
}

// DrawbackNEQ applies the NEQ predicate on the "drawback" field.
func DrawbackNEQ(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldNEQ(FieldDrawback, v))
}

// DrawbackIn applies the In predicate on the "drawback" field.
func DrawbackIn(vs ...string) predicate.Thesis {
	return predicate.Thesis(sql.FieldIn(FieldDrawback, vs...))
}

// DrawbackNotIn applies the NotIn predicate on the "drawback" field.
func DrawbackNotIn(vs ...string) predicate.Thesis {
	return predicate.Thesis(sql.FieldNotIn(FieldDrawback, vs...))
}

// DrawbackGT applies the GT predicate on the "drawback" field.
func DrawbackGT(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldGT(FieldDrawback, v))
}

// DrawbackGTE applies the GTE predicate on the "drawback" field.
func DrawbackGTE(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldGTE(FieldDrawback, v))
}

// DrawbackLT applies the LT predicate on the "drawback" field.
func DrawbackLT(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldLT(FieldDrawback, v))
}

// DrawbackLTE applies the LTE predicate on the "drawback" field.
func DrawbackLTE(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldLTE(FieldDrawback, v))
}

// DrawbackContains applies the Contains predicate on the "drawback" field.
func DrawbackContains(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldContains(FieldDrawback, v))
}

// DrawbackHasPrefix applies the HasPrefix predicate on the "drawback" field.
func DrawbackHasPrefix(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldHasPrefix(FieldDrawback, v))
}

// DrawbackHasSuffix applies the HasSuffix predicate on the "drawback" field.
func DrawbackHasSuffix(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldHasSuffix(FieldDrawback, v))
}

// DrawbackEqualFold applies the EqualFold predicate on the "drawback" field.
func DrawbackEqualFold(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldEqualFold(FieldDrawback, v))
}

// DrawbackContainsFold applies the ContainsFold predicate on the "drawback" field.
func DrawbackContainsFold(v string) predicate.Thesis {
	return predicate.Thesis(sql.FieldContainsFold(FieldDrawback, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Thesis {
	return predicate.Thesis(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Thesis {
	return predicate.Thesis(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Thesis {
	return predicate.Thesis(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Thesis {
	return predicate.Thesis(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Thesis {
	return predicate.Thesis(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Thesis {
	return predicate.Thesis(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Thesis {
	return predicate.Thesis(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Thesis {
	return predicate.Thesis(sql.FieldLTE(FieldCreateTime, v))
}

// HasUploaders applies the HasEdge predicate on the "uploaders" edge.
func HasUploaders() predicate.Thesis {
	return predicate.Thesis(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UploadersTable, UploadersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUploadersWith applies the HasEdge predicate on the "uploaders" edge with a given conditions (other predicates).
func HasUploadersWith(preds ...predicate.User) predicate.Thesis {
	return predicate.Thesis(func(s *sql.Selector) {
		step := newUploadersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasExamine applies the HasEdge predicate on the "examine" edge.
func HasExamine() predicate.Thesis {
	return predicate.Thesis(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ExamineTable, ExamineColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasExamineWith applies the HasEdge predicate on the "examine" edge with a given conditions (other predicates).
func HasExamineWith(preds ...predicate.User) predicate.Thesis {
	return predicate.Thesis(func(s *sql.Selector) {
		step := newExamineStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Thesis) predicate.Thesis {
	return predicate.Thesis(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Thesis) predicate.Thesis {
	return predicate.Thesis(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Thesis) predicate.Thesis {
	return predicate.Thesis(sql.NotPredicates(p))
}
