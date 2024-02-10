package types

import (
	"time"
)

type Model struct {
	Model        string     `json:"model"`
	SystemPrompt *string    `json:"systemPrompt,omitempty"`
	Temperature  *float64   `json:"temperature,omitempty"`
	Functions    []Function `json:"functions,omitempty"`
	Provider     string     `json:"provider"`
	Url          *string    `json:"url,omitempty"`
}

type Function struct {
	Name        string      `json:"name"`
	Async       *bool       `json:"async,omitempty"`
	Description *string     `json:"description,omitempty"`
	Parameters  interface{} `json:"parameters,omitempty"` // No direct equivalent for FunctionDefinition or any, using interface{}
}

type PlayHTEmotion string

const (
	FemaleHappy     PlayHTEmotion = "female_happy"
	FemaleSad       PlayHTEmotion = "female_sad"
	FemaleAngry     PlayHTEmotion = "female_angry"
	FemaleFearful   PlayHTEmotion = "female_fearful"
	FemaleDisgust   PlayHTEmotion = "female_disgust"
	FemaleSurprised PlayHTEmotion = "female_surprised"
)

type Voice struct {
	Provider        string         `json:"provider"`
	VoiceId         string         `json:"voiceId"`
	Speed           *float64       `json:"speed,omitempty"`
	Stability       *float64       `json:"stability,omitempty"`
	SimilarityBoost *float64       `json:"similarityBoost,omitempty"`
	Style           *float64       `json:"style,omitempty"`
	UseSpeakerBoost *bool          `json:"useSpeakerBoost,omitempty"`
	Temperature     *float64       `json:"temperature,omitempty"`
	Emotion         *PlayHTEmotion `json:"emotion,omitempty"`
	VoiceGuidance   *float64       `json:"voiceGuidance,omitempty"`
	StyleGuidance   *float64       `json:"styleGuidance,omitempty"`
	TextGuidance    *float64       `json:"textGuidance,omitempty"`
}

type Assistant struct {
	Name                      *string       `json:"name,omitempty"`
	Transcriber               *Transcriber  `json:"transcriber,omitempty"`
	Model                     *Model        `json:"model,omitempty"`
	Voice                     *Voice        `json:"voice,omitempty"`
	Language                  *string       `json:"language,omitempty"`
	ForwardingPhoneNumber     *string       `json:"forwardingPhoneNumber,omitempty"`
	FirstMessage              *string       `json:"firstMessage,omitempty"`
	VoicemailMessage          *string       `json:"voicemailMessage,omitempty"`
	EndCallMessage            *string       `json:"endCallMessage,omitempty"`
	EndCallPhrases            []string      `json:"endCallPhrases,omitempty"`
	InterruptionsEnabled      *bool         `json:"interruptionsEnabled,omitempty"`
	RecordingEnabled          *bool         `json:"recordingEnabled,omitempty"`
	EndCallFunctionEnabled    *bool         `json:"endCallFunctionEnabled,omitempty"`
	DialKeypadFunctionEnabled *bool         `json:"dialKeypadFunctionEnabled,omitempty"`
	FillersEnabled            *bool         `json:"fillersEnabled,omitempty"`
	ClientMessages            []interface{} `json:"clientMessages,omitempty"` // No direct equivalent for any, using interface{}
	ServerMessages            []interface{} `json:"serverMessages,omitempty"` // No direct equivalent for any, using interface{}
	SilenceTimeoutSeconds     *int          `json:"silenceTimeoutSeconds,omitempty"`
	ResponseDelaySeconds      *int          `json:"responseDelaySeconds,omitempty"`
	LiveTranscriptsEnabled    *bool         `json:"liveTranscriptsEnabled,omitempty"`
	Keywords                  []string      `json:"keywords,omitempty"`
	ParentId                  *string       `json:"parentId,omitempty"`
	ServerUrl                 *string       `json:"serverUrl,omitempty"`
	ServerUrlSecret           *string       `json:"serverUrlSecret,omitempty"`
	Id                        *string       `json:"id,omitempty"`
	OrgId                     *string       `json:"orgId,omitempty"`
	CreatedAt                 *time.Time    `json:"createdAt,omitempty"`
	UpdatedAt                 *time.Time    `json:"updatedAt,omitempty"`
}

type Transcriber struct {
	Provider string   `json:"provider"`
	Model    *string  `json:"model,omitempty"`
	Keywords []string `json:"keywords,omitempty"`
}

type VapiCallStatus string

const (
	Queued     VapiCallStatus = "queued"
	Ringing    VapiCallStatus = "ringing"
	InProgress VapiCallStatus = "in-progress"
	Forwarding VapiCallStatus = "forwarding"
	Ended      VapiCallStatus = "ended"
)

type VapiWebhookEnum string

const (
	AssistantRequest VapiWebhookEnum = "assistant-request"
	FunctionCall     VapiWebhookEnum = "function-call"
	StatusUpdate     VapiWebhookEnum = "status-update"
	EndOfCallReport  VapiWebhookEnum = "end-of-call-report"
	Hang             VapiWebhookEnum = "hang"
	SpeechUpdate     VapiWebhookEnum = "speech-update"
	Transcript       VapiWebhookEnum = "transcript"
)

type ConversationMessage struct {
	Role             string  `json:"role"`
	Message          *string `json:"message,omitempty"`
	Name             *string `json:"name,omitempty"`
	Args             *string `json:"args,omitempty"`
	Result           *string `json:"result,omitempty"`
	Time             int64   `json:"time"`
	EndTime          *int64  `json:"endTime,omitempty"`
	SecondsFromStart int     `json:"secondsFromStart"`
}
type BaseVapiPayload struct {
	Type VapiWebhookEnum `json:"type"`
	Call VapiCall        `json:"call"`
}

type AssistantRequestPayload struct {
	BaseVapiPayload
}

func (p AssistantRequestPayload) GetCallType() VapiWebhookEnum {
	return p.Type
}

type StatusUpdatePayload struct {
	BaseVapiPayload
	Status   VapiCallStatus        `json:"status"`
	Messages []ConversationMessage `json:"messages,omitempty"`
}

func (p StatusUpdatePayload) GetCallType() VapiWebhookEnum {
	return p.Type
}

type FunctionCallPayload struct {
	BaseVapiPayload
	FunctionCall OpenAIFunctionCall `json:"functionCall"`
}

type OpenAIFunctionCall struct {
	Name       string      `json:"name"`
	Parameters interface{} `json:"parameters"`
}

func (p FunctionCallPayload) GetCallType() VapiWebhookEnum {
	return p.Type
}

type EndOfCallReportPayload struct {
	BaseVapiPayload
	EndedReason  string                `json:"endedReason"`
	Transcript   string                `json:"transcript"`
	Messages     []ConversationMessage `json:"messages"`
	Summary      string                `json:"summary"`
	RecordingUrl *string               `json:"recordingUrl,omitempty"`
}

func (p EndOfCallReportPayload) GetCallType() VapiWebhookEnum {
	return p.Type
}

type HangPayload struct {
	BaseVapiPayload
}

func (p HangPayload) GetCallType() VapiWebhookEnum {
	return p.Type
}

type SpeechUpdatePayload struct {
	BaseVapiPayload
	Status string `json:"status"`
	Role   string `json:"role"`
}

func (p SpeechUpdatePayload) GetCallType() VapiWebhookEnum {
	return p.Type
}

type TranscriptPayload struct {
	BaseVapiPayload
	Role           string `json:"role"`
	TranscriptType string `json:"transcriptType"`
	Transcript     string `json:"transcript"`
}

func (p TranscriptPayload) GetCallType() VapiWebhookEnum {
	return p.Type
}

type VapiCall struct{} // Define the struct fields based on the TypeScript definition

type VapiPayload interface {
	GetCallType() VapiWebhookEnum
}
type FunctionCallMessageResponse struct {
	Result string `json:"result"`
	// Add any other fields that might be part of the union type
}

type AssistantRequestMessageResponse struct {
	Assistant *Assistant `json:"assistant,omitempty"`
	Error     *string    `json:"error,omitempty"`
}

// Define other response structs if needed

type VapiResponse interface{} // Use an empty interface to represent a union type
