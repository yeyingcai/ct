// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package cttask

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type RetVal int64

const (
	RetVal_SUCCESS RetVal = 1
	RetVal_FAILED  RetVal = 2
)

func (p RetVal) String() string {
	switch p {
	case RetVal_SUCCESS:
		return "SUCCESS"
	case RetVal_FAILED:
		return "FAILED"
	}
	return "<UNSET>"
}

func RetValFromString(s string) (RetVal, error) {
	switch s {
	case "SUCCESS":
		return RetVal_SUCCESS, nil
	case "FAILED":
		return RetVal_FAILED, nil
	}
	return RetVal(0), fmt.Errorf("not a valid RetVal string")
}

func RetValPtr(v RetVal) *RetVal { return &v }

func (p RetVal) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *RetVal) UnmarshalText(text []byte) error {
	q, err := RetValFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

type TaskType int64

const (
	TaskType_TIME  TaskType = 1
	TaskType_EVENT TaskType = 2
)

func (p TaskType) String() string {
	switch p {
	case TaskType_TIME:
		return "TIME"
	case TaskType_EVENT:
		return "EVENT"
	}
	return "<UNSET>"
}

func TaskTypeFromString(s string) (TaskType, error) {
	switch s {
	case "TIME":
		return TaskType_TIME, nil
	case "EVENT":
		return TaskType_EVENT, nil
	}
	return TaskType(0), fmt.Errorf("not a valid TaskType string")
}

func TaskTypePtr(v TaskType) *TaskType { return &v }

func (p TaskType) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *TaskType) UnmarshalText(text []byte) error {
	q, err := TaskTypeFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

type TaskAction int64

const (
	TaskAction_ONLINE  TaskAction = 1
	TaskAction_OFFLINE TaskAction = 2
	TaskAction_PAUSE   TaskAction = 3
	TaskAction_RESUME  TaskAction = 4
	TaskAction_REMOVE  TaskAction = 5
	TaskAction_EXECUTE TaskAction = 6
)

func (p TaskAction) String() string {
	switch p {
	case TaskAction_ONLINE:
		return "ONLINE"
	case TaskAction_OFFLINE:
		return "OFFLINE"
	case TaskAction_PAUSE:
		return "PAUSE"
	case TaskAction_RESUME:
		return "RESUME"
	case TaskAction_REMOVE:
		return "REMOVE"
	case TaskAction_EXECUTE:
		return "EXECUTE"
	}
	return "<UNSET>"
}

func TaskActionFromString(s string) (TaskAction, error) {
	switch s {
	case "ONLINE":
		return TaskAction_ONLINE, nil
	case "OFFLINE":
		return TaskAction_OFFLINE, nil
	case "PAUSE":
		return TaskAction_PAUSE, nil
	case "RESUME":
		return TaskAction_RESUME, nil
	case "REMOVE":
		return TaskAction_REMOVE, nil
	case "EXECUTE":
		return TaskAction_EXECUTE, nil
	}
	return TaskAction(0), fmt.Errorf("not a valid TaskAction string")
}

func TaskActionPtr(v TaskAction) *TaskAction { return &v }

func (p TaskAction) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *TaskAction) UnmarshalText(text []byte) error {
	q, err := TaskActionFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

// Attributes:
//  - TaskId
//  - TaskType
//  - CmdLine
//  - TrigerTime
//  - RetValue
//  - WaitTime
//  - ExecTimeout
//  - RetryCounter
//  - Account
type TaskInfo struct {
	TaskId       string   `thrift:"taskId,1,required" json:"taskId"`
	TaskType     TaskType `thrift:"taskType,2,required" json:"taskType"`
	CmdLine      string   `thrift:"cmdLine,3,required" json:"cmdLine"`
	TrigerTime   *string  `thrift:"trigerTime,4" json:"trigerTime,omitempty"`
	RetValue     *string  `thrift:"retValue,5" json:"retValue,omitempty"`
	WaitTime     *int32   `thrift:"waitTime,6" json:"waitTime,omitempty"`
	ExecTimeout  *int32   `thrift:"execTimeout,7" json:"execTimeout,omitempty"`
	RetryCounter *int32   `thrift:"retryCounter,8" json:"retryCounter,omitempty"`
	Account      *string  `thrift:"account,9" json:"account,omitempty"`
}

func NewTaskInfo() *TaskInfo {
	return &TaskInfo{}
}

func (p *TaskInfo) GetTaskId() string {
	return p.TaskId
}

func (p *TaskInfo) GetTaskType() TaskType {
	return p.TaskType
}

func (p *TaskInfo) GetCmdLine() string {
	return p.CmdLine
}

var TaskInfo_TrigerTime_DEFAULT string

func (p *TaskInfo) GetTrigerTime() string {
	if !p.IsSetTrigerTime() {
		return TaskInfo_TrigerTime_DEFAULT
	}
	return *p.TrigerTime
}

var TaskInfo_RetValue_DEFAULT string

func (p *TaskInfo) GetRetValue() string {
	if !p.IsSetRetValue() {
		return TaskInfo_RetValue_DEFAULT
	}
	return *p.RetValue
}

var TaskInfo_WaitTime_DEFAULT int32

func (p *TaskInfo) GetWaitTime() int32 {
	if !p.IsSetWaitTime() {
		return TaskInfo_WaitTime_DEFAULT
	}
	return *p.WaitTime
}

var TaskInfo_ExecTimeout_DEFAULT int32

func (p *TaskInfo) GetExecTimeout() int32 {
	if !p.IsSetExecTimeout() {
		return TaskInfo_ExecTimeout_DEFAULT
	}
	return *p.ExecTimeout
}

var TaskInfo_RetryCounter_DEFAULT int32

func (p *TaskInfo) GetRetryCounter() int32 {
	if !p.IsSetRetryCounter() {
		return TaskInfo_RetryCounter_DEFAULT
	}
	return *p.RetryCounter
}

var TaskInfo_Account_DEFAULT string

func (p *TaskInfo) GetAccount() string {
	if !p.IsSetAccount() {
		return TaskInfo_Account_DEFAULT
	}
	return *p.Account
}
func (p *TaskInfo) IsSetTrigerTime() bool {
	return p.TrigerTime != nil
}

func (p *TaskInfo) IsSetRetValue() bool {
	return p.RetValue != nil
}

func (p *TaskInfo) IsSetWaitTime() bool {
	return p.WaitTime != nil
}

func (p *TaskInfo) IsSetExecTimeout() bool {
	return p.ExecTimeout != nil
}

func (p *TaskInfo) IsSetRetryCounter() bool {
	return p.RetryCounter != nil
}

func (p *TaskInfo) IsSetAccount() bool {
	return p.Account != nil
}

func (p *TaskInfo) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetTaskId bool = false
	var issetTaskType bool = false
	var issetCmdLine bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetTaskId = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetTaskType = true
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
			issetCmdLine = true
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.readField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.readField7(iprot); err != nil {
				return err
			}
		case 8:
			if err := p.readField8(iprot); err != nil {
				return err
			}
		case 9:
			if err := p.readField9(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetTaskId {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field TaskId is not set"))
	}
	if !issetTaskType {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field TaskType is not set"))
	}
	if !issetCmdLine {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field CmdLine is not set"))
	}
	return nil
}

func (p *TaskInfo) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.TaskId = v
	}
	return nil
}

func (p *TaskInfo) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		temp := TaskType(v)
		p.TaskType = temp
	}
	return nil
}

func (p *TaskInfo) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.CmdLine = v
	}
	return nil
}

func (p *TaskInfo) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.TrigerTime = &v
	}
	return nil
}

func (p *TaskInfo) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.RetValue = &v
	}
	return nil
}

func (p *TaskInfo) readField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.WaitTime = &v
	}
	return nil
}

func (p *TaskInfo) readField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.ExecTimeout = &v
	}
	return nil
}

func (p *TaskInfo) readField8(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 8: ", err)
	} else {
		p.RetryCounter = &v
	}
	return nil
}

func (p *TaskInfo) readField9(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 9: ", err)
	} else {
		p.Account = &v
	}
	return nil
}

func (p *TaskInfo) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TaskInfo"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := p.writeField7(oprot); err != nil {
		return err
	}
	if err := p.writeField8(oprot); err != nil {
		return err
	}
	if err := p.writeField9(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TaskInfo) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("taskId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:taskId: ", p), err)
	}
	if err := oprot.WriteString(string(p.TaskId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.taskId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:taskId: ", p), err)
	}
	return err
}

func (p *TaskInfo) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("taskType", thrift.I32, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:taskType: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.TaskType)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.taskType (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:taskType: ", p), err)
	}
	return err
}

func (p *TaskInfo) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("cmdLine", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:cmdLine: ", p), err)
	}
	if err := oprot.WriteString(string(p.CmdLine)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.cmdLine (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:cmdLine: ", p), err)
	}
	return err
}

func (p *TaskInfo) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetTrigerTime() {
		if err := oprot.WriteFieldBegin("trigerTime", thrift.STRING, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:trigerTime: ", p), err)
		}
		if err := oprot.WriteString(string(*p.TrigerTime)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.trigerTime (4) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:trigerTime: ", p), err)
		}
	}
	return err
}

func (p *TaskInfo) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetRetValue() {
		if err := oprot.WriteFieldBegin("retValue", thrift.STRING, 5); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:retValue: ", p), err)
		}
		if err := oprot.WriteString(string(*p.RetValue)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.retValue (5) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 5:retValue: ", p), err)
		}
	}
	return err
}

func (p *TaskInfo) writeField6(oprot thrift.TProtocol) (err error) {
	if p.IsSetWaitTime() {
		if err := oprot.WriteFieldBegin("waitTime", thrift.I32, 6); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:waitTime: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.WaitTime)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.waitTime (6) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 6:waitTime: ", p), err)
		}
	}
	return err
}

func (p *TaskInfo) writeField7(oprot thrift.TProtocol) (err error) {
	if p.IsSetExecTimeout() {
		if err := oprot.WriteFieldBegin("execTimeout", thrift.I32, 7); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:execTimeout: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.ExecTimeout)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.execTimeout (7) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 7:execTimeout: ", p), err)
		}
	}
	return err
}

func (p *TaskInfo) writeField8(oprot thrift.TProtocol) (err error) {
	if p.IsSetRetryCounter() {
		if err := oprot.WriteFieldBegin("retryCounter", thrift.I32, 8); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:retryCounter: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.RetryCounter)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.retryCounter (8) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 8:retryCounter: ", p), err)
		}
	}
	return err
}

func (p *TaskInfo) writeField9(oprot thrift.TProtocol) (err error) {
	if p.IsSetAccount() {
		if err := oprot.WriteFieldBegin("account", thrift.STRING, 9); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 9:account: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Account)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.account (9) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 9:account: ", p), err)
		}
	}
	return err
}

func (p *TaskInfo) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TaskInfo(%+v)", *p)
}

// Attributes:
//  - TaskInfo
//  - Action
type TaskEntity struct {
	TaskInfo *TaskInfo   `thrift:"taskInfo,1,required" json:"taskInfo"`
	Action   *TaskAction `thrift:"action,2" json:"action,omitempty"`
}

func NewTaskEntity() *TaskEntity {
	return &TaskEntity{}
}

var TaskEntity_TaskInfo_DEFAULT *TaskInfo

func (p *TaskEntity) GetTaskInfo() *TaskInfo {
	if !p.IsSetTaskInfo() {
		return TaskEntity_TaskInfo_DEFAULT
	}
	return p.TaskInfo
}

var TaskEntity_Action_DEFAULT TaskAction

func (p *TaskEntity) GetAction() TaskAction {
	if !p.IsSetAction() {
		return TaskEntity_Action_DEFAULT
	}
	return *p.Action
}
func (p *TaskEntity) IsSetTaskInfo() bool {
	return p.TaskInfo != nil
}

func (p *TaskEntity) IsSetAction() bool {
	return p.Action != nil
}

func (p *TaskEntity) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetTaskInfo bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetTaskInfo = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetTaskInfo {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field TaskInfo is not set"))
	}
	return nil
}

func (p *TaskEntity) readField1(iprot thrift.TProtocol) error {
	p.TaskInfo = &TaskInfo{}
	if err := p.TaskInfo.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.TaskInfo), err)
	}
	return nil
}

func (p *TaskEntity) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		temp := TaskAction(v)
		p.Action = &temp
	}
	return nil
}

func (p *TaskEntity) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TaskEntity"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TaskEntity) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("taskInfo", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:taskInfo: ", p), err)
	}
	if err := p.TaskInfo.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.TaskInfo), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:taskInfo: ", p), err)
	}
	return err
}

func (p *TaskEntity) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetAction() {
		if err := oprot.WriteFieldBegin("action", thrift.I32, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:action: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.Action)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.action (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:action: ", p), err)
		}
	}
	return err
}

func (p *TaskEntity) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TaskEntity(%+v)", *p)
}
