// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"fmt"
	"time"

	"github.com/go-faster/errors"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

type CreateProjectReq struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

// GetName returns the value of Name.
func (s *CreateProjectReq) GetName() string {
	return s.Name
}

// GetColor returns the value of Color.
func (s *CreateProjectReq) GetColor() string {
	return s.Color
}

// SetName sets the value of Name.
func (s *CreateProjectReq) SetName(val string) {
	s.Name = val
}

// SetColor sets the value of Color.
func (s *CreateProjectReq) SetColor(val string) {
	s.Color = val
}

type CreateTaskReq struct {
	Title    string  `json:"title"`
	Content  string  `json:"content"`
	Priority int     `json:"priority"`
	DueOn    OptDate `json:"dueOn"`
}

// GetTitle returns the value of Title.
func (s *CreateTaskReq) GetTitle() string {
	return s.Title
}

// GetContent returns the value of Content.
func (s *CreateTaskReq) GetContent() string {
	return s.Content
}

// GetPriority returns the value of Priority.
func (s *CreateTaskReq) GetPriority() int {
	return s.Priority
}

// GetDueOn returns the value of DueOn.
func (s *CreateTaskReq) GetDueOn() OptDate {
	return s.DueOn
}

// SetTitle sets the value of Title.
func (s *CreateTaskReq) SetTitle(val string) {
	s.Title = val
}

// SetContent sets the value of Content.
func (s *CreateTaskReq) SetContent(val string) {
	s.Content = val
}

// SetPriority sets the value of Priority.
func (s *CreateTaskReq) SetPriority(val int) {
	s.Priority = val
}

// SetDueOn sets the value of DueOn.
func (s *CreateTaskReq) SetDueOn(val OptDate) {
	s.DueOn = val
}

// DeleteProjectNoContent is response for DeleteProject operation.
type DeleteProjectNoContent struct{}

// DeleteTaskNoContent is response for DeleteTask operation.
type DeleteTaskNoContent struct{}

type Error struct {
	// エラーコード.
	Code int `json:"code"`
	// ユーザ向けの大まかなエラーの説明.
	Message string `json:"message"`
}

// GetCode returns the value of Code.
func (s *Error) GetCode() int {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *Error) SetCode(val int) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

type GetHealthOK struct {
	// Opepe APIのバージョン.
	Version string `json:"version"`
	// Opepe APIのリビジョン.
	Revision string `json:"revision"`
}

// GetVersion returns the value of Version.
func (s *GetHealthOK) GetVersion() string {
	return s.Version
}

// GetRevision returns the value of Revision.
func (s *GetHealthOK) GetRevision() string {
	return s.Revision
}

// SetVersion sets the value of Version.
func (s *GetHealthOK) SetVersion(val string) {
	s.Version = val
}

// SetRevision sets the value of Revision.
func (s *GetHealthOK) SetRevision(val string) {
	s.Revision = val
}

type IsAuthorized struct {
	APIKey string
}

// GetAPIKey returns the value of APIKey.
func (s *IsAuthorized) GetAPIKey() string {
	return s.APIKey
}

// SetAPIKey sets the value of APIKey.
func (s *IsAuthorized) SetAPIKey(val string) {
	s.APIKey = val
}

type ListProjectsSort string

const (
	ListProjectsSortCreatedAt      ListProjectsSort = "createdAt"
	ListProjectsSortMinusCreatedAt ListProjectsSort = "-createdAt"
)

// MarshalText implements encoding.TextMarshaler.
func (s ListProjectsSort) MarshalText() ([]byte, error) {
	switch s {
	case ListProjectsSortCreatedAt:
		return []byte(s), nil
	case ListProjectsSortMinusCreatedAt:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *ListProjectsSort) UnmarshalText(data []byte) error {
	switch ListProjectsSort(data) {
	case ListProjectsSortCreatedAt:
		*s = ListProjectsSortCreatedAt
		return nil
	case ListProjectsSortMinusCreatedAt:
		*s = ListProjectsSortMinusCreatedAt
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type ListTasksSort string

const (
	ListTasksSortCreatedAt      ListTasksSort = "createdAt"
	ListTasksSortMinusCreatedAt ListTasksSort = "-createdAt"
)

// MarshalText implements encoding.TextMarshaler.
func (s ListTasksSort) MarshalText() ([]byte, error) {
	switch s {
	case ListTasksSortCreatedAt:
		return []byte(s), nil
	case ListTasksSortMinusCreatedAt:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *ListTasksSort) UnmarshalText(data []byte) error {
	switch ListTasksSort(data) {
	case ListTasksSortCreatedAt:
		*s = ListTasksSortCreatedAt
		return nil
	case ListTasksSortMinusCreatedAt:
		*s = ListTasksSortMinusCreatedAt
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// NewOptBool returns new OptBool with value set to v.
func NewOptBool(v bool) OptBool {
	return OptBool{
		Value: v,
		Set:   true,
	}
}

// OptBool is optional bool.
type OptBool struct {
	Value bool
	Set   bool
}

// IsSet returns true if OptBool was set.
func (o OptBool) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptBool) Reset() {
	var v bool
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptBool) SetTo(v bool) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptBool) Get() (v bool, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptBool) Or(d bool) bool {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptDate returns new OptDate with value set to v.
func NewOptDate(v time.Time) OptDate {
	return OptDate{
		Value: v,
		Set:   true,
	}
}

// OptDate is optional time.Time.
type OptDate struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDate was set.
func (o OptDate) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDate) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDate) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDate) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDate) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptDateTime returns new OptDateTime with value set to v.
func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}

// OptDateTime is optional time.Time.
type OptDateTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDateTime was set.
func (o OptDateTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptListProjectsSort returns new OptListProjectsSort with value set to v.
func NewOptListProjectsSort(v ListProjectsSort) OptListProjectsSort {
	return OptListProjectsSort{
		Value: v,
		Set:   true,
	}
}

// OptListProjectsSort is optional ListProjectsSort.
type OptListProjectsSort struct {
	Value ListProjectsSort
	Set   bool
}

// IsSet returns true if OptListProjectsSort was set.
func (o OptListProjectsSort) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptListProjectsSort) Reset() {
	var v ListProjectsSort
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptListProjectsSort) SetTo(v ListProjectsSort) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptListProjectsSort) Get() (v ListProjectsSort, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptListProjectsSort) Or(d ListProjectsSort) ListProjectsSort {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptListTasksSort returns new OptListTasksSort with value set to v.
func NewOptListTasksSort(v ListTasksSort) OptListTasksSort {
	return OptListTasksSort{
		Value: v,
		Set:   true,
	}
}

// OptListTasksSort is optional ListTasksSort.
type OptListTasksSort struct {
	Value ListTasksSort
	Set   bool
}

// IsSet returns true if OptListTasksSort was set.
func (o OptListTasksSort) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptListTasksSort) Reset() {
	var v ListTasksSort
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptListTasksSort) SetTo(v ListTasksSort) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptListTasksSort) Get() (v ListTasksSort, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptListTasksSort) Or(d ListTasksSort) ListTasksSort {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/Project
type Project struct {
	// プロジェクトID.
	ID string `json:"id"`
	// プロジェクト名.
	Name string `json:"name"`
	// プロジェクトカラー.
	Color string `json:"color"`
	// アーカイブ済みか.
	IsArchived bool `json:"isArchived"`
	// 作成日時.
	CreatedAt time.Time `json:"createdAt"`
	// 更新日時.
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetID returns the value of ID.
func (s *Project) GetID() string {
	return s.ID
}

// GetName returns the value of Name.
func (s *Project) GetName() string {
	return s.Name
}

// GetColor returns the value of Color.
func (s *Project) GetColor() string {
	return s.Color
}

// GetIsArchived returns the value of IsArchived.
func (s *Project) GetIsArchived() bool {
	return s.IsArchived
}

// GetCreatedAt returns the value of CreatedAt.
func (s *Project) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *Project) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

// SetID sets the value of ID.
func (s *Project) SetID(val string) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *Project) SetName(val string) {
	s.Name = val
}

// SetColor sets the value of Color.
func (s *Project) SetColor(val string) {
	s.Color = val
}

// SetIsArchived sets the value of IsArchived.
func (s *Project) SetIsArchived(val bool) {
	s.IsArchived = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *Project) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *Project) SetUpdatedAt(val time.Time) {
	s.UpdatedAt = val
}

// Ref: #/components/schemas/Projects
type Projects struct {
	// プロジェクト一覧.
	Projects []Project `json:"projects"`
	// まだ取得可能なプロジェクトが存在するか.
	HasNext bool `json:"hasNext"`
}

// GetProjects returns the value of Projects.
func (s *Projects) GetProjects() []Project {
	return s.Projects
}

// GetHasNext returns the value of HasNext.
func (s *Projects) GetHasNext() bool {
	return s.HasNext
}

// SetProjects sets the value of Projects.
func (s *Projects) SetProjects(val []Project) {
	s.Projects = val
}

// SetHasNext sets the value of HasNext.
func (s *Projects) SetHasNext(val bool) {
	s.HasNext = val
}

// Ref: #/components/schemas/Task
type Task struct {
	// タスクID.
	ID string `json:"id"`
	// 紐づくプロジェクトのID.
	ProjectID string `json:"project_id"`
	// タイトル.
	Title string `json:"title"`
	// 内容.
	Content  string  `json:"content"`
	Priority int     `json:"priority"`
	DueOn    OptDate `json:"dueOn"`
	// 完了日時.
	CompletedAt OptDateTime `json:"completedAt"`
	// 作成日時.
	CreatedAt time.Time `json:"createdAt"`
	// 更新日時.
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetID returns the value of ID.
func (s *Task) GetID() string {
	return s.ID
}

// GetProjectID returns the value of ProjectID.
func (s *Task) GetProjectID() string {
	return s.ProjectID
}

// GetTitle returns the value of Title.
func (s *Task) GetTitle() string {
	return s.Title
}

// GetContent returns the value of Content.
func (s *Task) GetContent() string {
	return s.Content
}

// GetPriority returns the value of Priority.
func (s *Task) GetPriority() int {
	return s.Priority
}

// GetDueOn returns the value of DueOn.
func (s *Task) GetDueOn() OptDate {
	return s.DueOn
}

// GetCompletedAt returns the value of CompletedAt.
func (s *Task) GetCompletedAt() OptDateTime {
	return s.CompletedAt
}

// GetCreatedAt returns the value of CreatedAt.
func (s *Task) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *Task) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

// SetID sets the value of ID.
func (s *Task) SetID(val string) {
	s.ID = val
}

// SetProjectID sets the value of ProjectID.
func (s *Task) SetProjectID(val string) {
	s.ProjectID = val
}

// SetTitle sets the value of Title.
func (s *Task) SetTitle(val string) {
	s.Title = val
}

// SetContent sets the value of Content.
func (s *Task) SetContent(val string) {
	s.Content = val
}

// SetPriority sets the value of Priority.
func (s *Task) SetPriority(val int) {
	s.Priority = val
}

// SetDueOn sets the value of DueOn.
func (s *Task) SetDueOn(val OptDate) {
	s.DueOn = val
}

// SetCompletedAt sets the value of CompletedAt.
func (s *Task) SetCompletedAt(val OptDateTime) {
	s.CompletedAt = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *Task) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *Task) SetUpdatedAt(val time.Time) {
	s.UpdatedAt = val
}

// Ref: #/components/schemas/Tasks
type Tasks struct {
	// タスク一覧.
	Tasks []Task `json:"tasks"`
	// まだ取得可能なタスクが存在するか.
	HasNext bool `json:"hasNext"`
}

// GetTasks returns the value of Tasks.
func (s *Tasks) GetTasks() []Task {
	return s.Tasks
}

// GetHasNext returns the value of HasNext.
func (s *Tasks) GetHasNext() bool {
	return s.HasNext
}

// SetTasks sets the value of Tasks.
func (s *Tasks) SetTasks(val []Task) {
	s.Tasks = val
}

// SetHasNext sets the value of HasNext.
func (s *Tasks) SetHasNext(val bool) {
	s.HasNext = val
}

type UpdateProjectReq struct {
	Name       OptString `json:"name"`
	Color      OptString `json:"color"`
	IsArchived OptBool   `json:"isArchived"`
}

// GetName returns the value of Name.
func (s *UpdateProjectReq) GetName() OptString {
	return s.Name
}

// GetColor returns the value of Color.
func (s *UpdateProjectReq) GetColor() OptString {
	return s.Color
}

// GetIsArchived returns the value of IsArchived.
func (s *UpdateProjectReq) GetIsArchived() OptBool {
	return s.IsArchived
}

// SetName sets the value of Name.
func (s *UpdateProjectReq) SetName(val OptString) {
	s.Name = val
}

// SetColor sets the value of Color.
func (s *UpdateProjectReq) SetColor(val OptString) {
	s.Color = val
}

// SetIsArchived sets the value of IsArchived.
func (s *UpdateProjectReq) SetIsArchived(val OptBool) {
	s.IsArchived = val
}

type UpdateTaskReq struct {
	Title       OptString `json:"title"`
	Content     OptString `json:"content"`
	Priority    OptInt    `json:"priority"`
	DueOn       OptDate   `json:"dueOn"`
	IsCompleted OptBool   `json:"isCompleted"`
}

// GetTitle returns the value of Title.
func (s *UpdateTaskReq) GetTitle() OptString {
	return s.Title
}

// GetContent returns the value of Content.
func (s *UpdateTaskReq) GetContent() OptString {
	return s.Content
}

// GetPriority returns the value of Priority.
func (s *UpdateTaskReq) GetPriority() OptInt {
	return s.Priority
}

// GetDueOn returns the value of DueOn.
func (s *UpdateTaskReq) GetDueOn() OptDate {
	return s.DueOn
}

// GetIsCompleted returns the value of IsCompleted.
func (s *UpdateTaskReq) GetIsCompleted() OptBool {
	return s.IsCompleted
}

// SetTitle sets the value of Title.
func (s *UpdateTaskReq) SetTitle(val OptString) {
	s.Title = val
}

// SetContent sets the value of Content.
func (s *UpdateTaskReq) SetContent(val OptString) {
	s.Content = val
}

// SetPriority sets the value of Priority.
func (s *UpdateTaskReq) SetPriority(val OptInt) {
	s.Priority = val
}

// SetDueOn sets the value of DueOn.
func (s *UpdateTaskReq) SetDueOn(val OptDate) {
	s.DueOn = val
}

// SetIsCompleted sets the value of IsCompleted.
func (s *UpdateTaskReq) SetIsCompleted(val OptBool) {
	s.IsCompleted = val
}
