// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/binaryauthorization/v1beta1/resources.proto

package binaryauthorization

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Policy_GlobalPolicyEvaluationMode int32

const (
	// Not specified: DISABLE is assumed.
	Policy_GLOBAL_POLICY_EVALUATION_MODE_UNSPECIFIED Policy_GlobalPolicyEvaluationMode = 0
	// Enables global policy evaluation.
	Policy_ENABLE Policy_GlobalPolicyEvaluationMode = 1
	// Disables global policy evaluation.
	Policy_DISABLE Policy_GlobalPolicyEvaluationMode = 2
)

var Policy_GlobalPolicyEvaluationMode_name = map[int32]string{
	0: "GLOBAL_POLICY_EVALUATION_MODE_UNSPECIFIED",
	1: "ENABLE",
	2: "DISABLE",
}

var Policy_GlobalPolicyEvaluationMode_value = map[string]int32{
	"GLOBAL_POLICY_EVALUATION_MODE_UNSPECIFIED": 0,
	"ENABLE":  1,
	"DISABLE": 2,
}

func (x Policy_GlobalPolicyEvaluationMode) String() string {
	return proto.EnumName(Policy_GlobalPolicyEvaluationMode_name, int32(x))
}

func (Policy_GlobalPolicyEvaluationMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1b83e716c45ff98c, []int{0, 0}
}

type AdmissionRule_EvaluationMode int32

const (
	// Do not use.
	AdmissionRule_EVALUATION_MODE_UNSPECIFIED AdmissionRule_EvaluationMode = 0
	// This rule allows all all pod creations.
	AdmissionRule_ALWAYS_ALLOW AdmissionRule_EvaluationMode = 1
	// This rule allows a pod creation if all the attestors listed in
	// 'require_attestations_by' have valid attestations for all of the
	// images in the pod spec.
	AdmissionRule_REQUIRE_ATTESTATION AdmissionRule_EvaluationMode = 2
	// This rule denies all pod creations.
	AdmissionRule_ALWAYS_DENY AdmissionRule_EvaluationMode = 3
)

var AdmissionRule_EvaluationMode_name = map[int32]string{
	0: "EVALUATION_MODE_UNSPECIFIED",
	1: "ALWAYS_ALLOW",
	2: "REQUIRE_ATTESTATION",
	3: "ALWAYS_DENY",
}

var AdmissionRule_EvaluationMode_value = map[string]int32{
	"EVALUATION_MODE_UNSPECIFIED": 0,
	"ALWAYS_ALLOW":                1,
	"REQUIRE_ATTESTATION":         2,
	"ALWAYS_DENY":                 3,
}

func (x AdmissionRule_EvaluationMode) String() string {
	return proto.EnumName(AdmissionRule_EvaluationMode_name, int32(x))
}

func (AdmissionRule_EvaluationMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1b83e716c45ff98c, []int{2, 0}
}

// Defines the possible actions when a pod creation is denied by an admission
// rule.
type AdmissionRule_EnforcementMode int32

const (
	// Do not use.
	AdmissionRule_ENFORCEMENT_MODE_UNSPECIFIED AdmissionRule_EnforcementMode = 0
	// Enforce the admission rule by blocking the pod creation.
	AdmissionRule_ENFORCED_BLOCK_AND_AUDIT_LOG AdmissionRule_EnforcementMode = 1
	// Dryrun mode: Audit logging only.  This will allow the pod creation as if
	// the admission request had specified break-glass.
	AdmissionRule_DRYRUN_AUDIT_LOG_ONLY AdmissionRule_EnforcementMode = 2
)

var AdmissionRule_EnforcementMode_name = map[int32]string{
	0: "ENFORCEMENT_MODE_UNSPECIFIED",
	1: "ENFORCED_BLOCK_AND_AUDIT_LOG",
	2: "DRYRUN_AUDIT_LOG_ONLY",
}

var AdmissionRule_EnforcementMode_value = map[string]int32{
	"ENFORCEMENT_MODE_UNSPECIFIED": 0,
	"ENFORCED_BLOCK_AND_AUDIT_LOG": 1,
	"DRYRUN_AUDIT_LOG_ONLY":        2,
}

func (x AdmissionRule_EnforcementMode) String() string {
	return proto.EnumName(AdmissionRule_EnforcementMode_name, int32(x))
}

func (AdmissionRule_EnforcementMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1b83e716c45ff98c, []int{2, 1}
}

// Represents a signature algorithm and other information necessary to verify
// signatures with a given public key.
// This is based primarily on the public key types supported by Tink's
// PemKeyType, which is in turn based on KMS's supported signing algorithms.
// See https://cloud.google.com/kms/docs/algorithms. In the future, BinAuthz
// might support additional public key types independently of Tink and/or KMS.
type PkixPublicKey_SignatureAlgorithm int32

const (
	// Not specified.
	PkixPublicKey_SIGNATURE_ALGORITHM_UNSPECIFIED PkixPublicKey_SignatureAlgorithm = 0
	// RSASSA-PSS 2048 bit key with a SHA256 digest.
	PkixPublicKey_RSA_PSS_2048_SHA256 PkixPublicKey_SignatureAlgorithm = 1
	// RSASSA-PSS 3072 bit key with a SHA256 digest.
	PkixPublicKey_RSA_PSS_3072_SHA256 PkixPublicKey_SignatureAlgorithm = 2
	// RSASSA-PSS 4096 bit key with a SHA256 digest.
	PkixPublicKey_RSA_PSS_4096_SHA256 PkixPublicKey_SignatureAlgorithm = 3
	// RSASSA-PSS 4096 bit key with a SHA512 digest.
	PkixPublicKey_RSA_PSS_4096_SHA512 PkixPublicKey_SignatureAlgorithm = 4
	// RSASSA-PKCS1-v1_5 with a 2048 bit key and a SHA256 digest.
	PkixPublicKey_RSA_SIGN_PKCS1_2048_SHA256 PkixPublicKey_SignatureAlgorithm = 5
	// RSASSA-PKCS1-v1_5 with a 3072 bit key and a SHA256 digest.
	PkixPublicKey_RSA_SIGN_PKCS1_3072_SHA256 PkixPublicKey_SignatureAlgorithm = 6
	// RSASSA-PKCS1-v1_5 with a 4096 bit key and a SHA256 digest.
	PkixPublicKey_RSA_SIGN_PKCS1_4096_SHA256 PkixPublicKey_SignatureAlgorithm = 7
	// RSASSA-PKCS1-v1_5 with a 4096 bit key and a SHA512 digest.
	PkixPublicKey_RSA_SIGN_PKCS1_4096_SHA512 PkixPublicKey_SignatureAlgorithm = 8
	// ECDSA on the NIST P-256 curve with a SHA256 digest.
	PkixPublicKey_ECDSA_P256_SHA256 PkixPublicKey_SignatureAlgorithm = 9
	// ECDSA on the NIST P-384 curve with a SHA384 digest.
	PkixPublicKey_ECDSA_P384_SHA384 PkixPublicKey_SignatureAlgorithm = 10
	// ECDSA on the NIST P-521 curve with a SHA512 digest.
	PkixPublicKey_ECDSA_P521_SHA512 PkixPublicKey_SignatureAlgorithm = 11
)

var PkixPublicKey_SignatureAlgorithm_name = map[int32]string{
	0:  "SIGNATURE_ALGORITHM_UNSPECIFIED",
	1:  "RSA_PSS_2048_SHA256",
	2:  "RSA_PSS_3072_SHA256",
	3:  "RSA_PSS_4096_SHA256",
	4:  "RSA_PSS_4096_SHA512",
	5:  "RSA_SIGN_PKCS1_2048_SHA256",
	6:  "RSA_SIGN_PKCS1_3072_SHA256",
	7:  "RSA_SIGN_PKCS1_4096_SHA256",
	8:  "RSA_SIGN_PKCS1_4096_SHA512",
	9:  "ECDSA_P256_SHA256",
	10: "ECDSA_P384_SHA384",
	11: "ECDSA_P521_SHA512",
}

var PkixPublicKey_SignatureAlgorithm_value = map[string]int32{
	"SIGNATURE_ALGORITHM_UNSPECIFIED": 0,
	"RSA_PSS_2048_SHA256":             1,
	"RSA_PSS_3072_SHA256":             2,
	"RSA_PSS_4096_SHA256":             3,
	"RSA_PSS_4096_SHA512":             4,
	"RSA_SIGN_PKCS1_2048_SHA256":      5,
	"RSA_SIGN_PKCS1_3072_SHA256":      6,
	"RSA_SIGN_PKCS1_4096_SHA256":      7,
	"RSA_SIGN_PKCS1_4096_SHA512":      8,
	"ECDSA_P256_SHA256":               9,
	"ECDSA_P384_SHA384":               10,
	"ECDSA_P521_SHA512":               11,
}

func (x PkixPublicKey_SignatureAlgorithm) String() string {
	return proto.EnumName(PkixPublicKey_SignatureAlgorithm_name, int32(x))
}

func (PkixPublicKey_SignatureAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1b83e716c45ff98c, []int{5, 0}
}

// A [policy][google.cloud.binaryauthorization.v1beta1.Policy] for container image binary authorization.
type Policy struct {
	// Output only. The resource name, in the format `projects/*/policy`. There is
	// at most one policy per project.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. A descriptive comment.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// Optional. Controls the evaluation of a Google-maintained global admission
	// policy for common system-level images. Images not covered by the global
	// policy will be subject to the project admission policy. This setting
	// has no effect when specified inside a global admission policy.
	GlobalPolicyEvaluationMode Policy_GlobalPolicyEvaluationMode `protobuf:"varint,7,opt,name=global_policy_evaluation_mode,json=globalPolicyEvaluationMode,proto3,enum=google.cloud.binaryauthorization.v1beta1.Policy_GlobalPolicyEvaluationMode" json:"global_policy_evaluation_mode,omitempty"`
	// Optional. Admission policy whitelisting. A matching admission request will
	// always be permitted. This feature is typically used to exclude Google or
	// third-party infrastructure images from Binary Authorization policies.
	AdmissionWhitelistPatterns []*AdmissionWhitelistPattern `protobuf:"bytes,2,rep,name=admission_whitelist_patterns,json=admissionWhitelistPatterns,proto3" json:"admission_whitelist_patterns,omitempty"`
	// Optional. Per-cluster admission rules. Cluster spec format:
	// `location.clusterId`. There can be at most one admission rule per cluster
	// spec.
	// A `location` is either a compute zone (e.g. us-central1-a) or a region
	// (e.g. us-central1).
	// For `clusterId` syntax restrictions see
	// https://cloud.google.com/container-engine/reference/rest/v1/projects.zones.clusters.
	ClusterAdmissionRules map[string]*AdmissionRule `protobuf:"bytes,3,rep,name=cluster_admission_rules,json=clusterAdmissionRules,proto3" json:"cluster_admission_rules,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Required. Default admission rule for a cluster without a per-cluster, per-
	// kubernetes-service-account, or per-istio-service-identity admission rule.
	DefaultAdmissionRule *AdmissionRule `protobuf:"bytes,4,opt,name=default_admission_rule,json=defaultAdmissionRule,proto3" json:"default_admission_rule,omitempty"`
	// Output only. Time when the policy was last updated.
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b83e716c45ff98c, []int{0}
}

func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Policy) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Policy) GetGlobalPolicyEvaluationMode() Policy_GlobalPolicyEvaluationMode {
	if m != nil {
		return m.GlobalPolicyEvaluationMode
	}
	return Policy_GLOBAL_POLICY_EVALUATION_MODE_UNSPECIFIED
}

func (m *Policy) GetAdmissionWhitelistPatterns() []*AdmissionWhitelistPattern {
	if m != nil {
		return m.AdmissionWhitelistPatterns
	}
	return nil
}

func (m *Policy) GetClusterAdmissionRules() map[string]*AdmissionRule {
	if m != nil {
		return m.ClusterAdmissionRules
	}
	return nil
}

func (m *Policy) GetDefaultAdmissionRule() *AdmissionRule {
	if m != nil {
		return m.DefaultAdmissionRule
	}
	return nil
}

func (m *Policy) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

// An [admission whitelist pattern][google.cloud.binaryauthorization.v1beta1.AdmissionWhitelistPattern] exempts images
// from checks by [admission rules][google.cloud.binaryauthorization.v1beta1.AdmissionRule].
type AdmissionWhitelistPattern struct {
	// An image name pattern to whitelist, in the form `registry/path/to/image`.
	// This supports a trailing `*` as a wildcard, but this is allowed only in
	// text after the `registry/` part.
	NamePattern          string   `protobuf:"bytes,1,opt,name=name_pattern,json=namePattern,proto3" json:"name_pattern,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdmissionWhitelistPattern) Reset()         { *m = AdmissionWhitelistPattern{} }
func (m *AdmissionWhitelistPattern) String() string { return proto.CompactTextString(m) }
func (*AdmissionWhitelistPattern) ProtoMessage()    {}
func (*AdmissionWhitelistPattern) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b83e716c45ff98c, []int{1}
}

func (m *AdmissionWhitelistPattern) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdmissionWhitelistPattern.Unmarshal(m, b)
}
func (m *AdmissionWhitelistPattern) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdmissionWhitelistPattern.Marshal(b, m, deterministic)
}
func (m *AdmissionWhitelistPattern) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdmissionWhitelistPattern.Merge(m, src)
}
func (m *AdmissionWhitelistPattern) XXX_Size() int {
	return xxx_messageInfo_AdmissionWhitelistPattern.Size(m)
}
func (m *AdmissionWhitelistPattern) XXX_DiscardUnknown() {
	xxx_messageInfo_AdmissionWhitelistPattern.DiscardUnknown(m)
}

var xxx_messageInfo_AdmissionWhitelistPattern proto.InternalMessageInfo

func (m *AdmissionWhitelistPattern) GetNamePattern() string {
	if m != nil {
		return m.NamePattern
	}
	return ""
}

// An [admission rule][google.cloud.binaryauthorization.v1beta1.AdmissionRule] specifies either that all container images
// used in a pod creation request must be attested to by one or more
// [attestors][google.cloud.binaryauthorization.v1beta1.Attestor], that all pod creations will be allowed, or that all
// pod creations will be denied.
//
// Images matching an [admission whitelist pattern][google.cloud.binaryauthorization.v1beta1.AdmissionWhitelistPattern]
// are exempted from admission rules and will never block a pod creation.
type AdmissionRule struct {
	// Required. How this admission rule will be evaluated.
	EvaluationMode AdmissionRule_EvaluationMode `protobuf:"varint,1,opt,name=evaluation_mode,json=evaluationMode,proto3,enum=google.cloud.binaryauthorization.v1beta1.AdmissionRule_EvaluationMode" json:"evaluation_mode,omitempty"`
	// Optional. The resource names of the attestors that must attest to
	// a container image, in the format `projects/*/attestors/*`. Each
	// attestor must exist before a policy can reference it.  To add an attestor
	// to a policy the principal issuing the policy change request must be able
	// to read the attestor resource.
	//
	// Note: this field must be non-empty when the evaluation_mode field specifies
	// REQUIRE_ATTESTATION, otherwise it must be empty.
	RequireAttestationsBy []string `protobuf:"bytes,2,rep,name=require_attestations_by,json=requireAttestationsBy,proto3" json:"require_attestations_by,omitempty"`
	// Required. The action when a pod creation is denied by the admission rule.
	EnforcementMode      AdmissionRule_EnforcementMode `protobuf:"varint,3,opt,name=enforcement_mode,json=enforcementMode,proto3,enum=google.cloud.binaryauthorization.v1beta1.AdmissionRule_EnforcementMode" json:"enforcement_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *AdmissionRule) Reset()         { *m = AdmissionRule{} }
func (m *AdmissionRule) String() string { return proto.CompactTextString(m) }
func (*AdmissionRule) ProtoMessage()    {}
func (*AdmissionRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b83e716c45ff98c, []int{2}
}

func (m *AdmissionRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdmissionRule.Unmarshal(m, b)
}
func (m *AdmissionRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdmissionRule.Marshal(b, m, deterministic)
}
func (m *AdmissionRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdmissionRule.Merge(m, src)
}
func (m *AdmissionRule) XXX_Size() int {
	return xxx_messageInfo_AdmissionRule.Size(m)
}
func (m *AdmissionRule) XXX_DiscardUnknown() {
	xxx_messageInfo_AdmissionRule.DiscardUnknown(m)
}

var xxx_messageInfo_AdmissionRule proto.InternalMessageInfo

func (m *AdmissionRule) GetEvaluationMode() AdmissionRule_EvaluationMode {
	if m != nil {
		return m.EvaluationMode
	}
	return AdmissionRule_EVALUATION_MODE_UNSPECIFIED
}

func (m *AdmissionRule) GetRequireAttestationsBy() []string {
	if m != nil {
		return m.RequireAttestationsBy
	}
	return nil
}

func (m *AdmissionRule) GetEnforcementMode() AdmissionRule_EnforcementMode {
	if m != nil {
		return m.EnforcementMode
	}
	return AdmissionRule_ENFORCEMENT_MODE_UNSPECIFIED
}

// An [attestor][google.cloud.binaryauthorization.v1beta1.Attestor] that attests to container image
// artifacts. An existing attestor cannot be modified except where
// indicated.
type Attestor struct {
	// Required. The resource name, in the format:
	// `projects/*/attestors/*`. This field may not be updated.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. A descriptive comment.  This field may be updated.
	// The field may be displayed in chooser dialogs.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// Required. Identifies an [attestor][google.cloud.binaryauthorization.v1beta1.Attestor] that attests to a
	// container image artifact. This determines how an attestation will
	// be stored, and how it will be used during policy
	// enforcement. Updates may not change the attestor type, but individual
	// attestor fields may be updated
	//
	// Types that are valid to be assigned to AttestorType:
	//	*Attestor_UserOwnedDrydockNote
	AttestorType isAttestor_AttestorType `protobuf_oneof:"attestor_type"`
	// Output only. Time when the attestor was last updated.
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Attestor) Reset()         { *m = Attestor{} }
func (m *Attestor) String() string { return proto.CompactTextString(m) }
func (*Attestor) ProtoMessage()    {}
func (*Attestor) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b83e716c45ff98c, []int{3}
}

func (m *Attestor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attestor.Unmarshal(m, b)
}
func (m *Attestor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attestor.Marshal(b, m, deterministic)
}
func (m *Attestor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attestor.Merge(m, src)
}
func (m *Attestor) XXX_Size() int {
	return xxx_messageInfo_Attestor.Size(m)
}
func (m *Attestor) XXX_DiscardUnknown() {
	xxx_messageInfo_Attestor.DiscardUnknown(m)
}

var xxx_messageInfo_Attestor proto.InternalMessageInfo

func (m *Attestor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Attestor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type isAttestor_AttestorType interface {
	isAttestor_AttestorType()
}

type Attestor_UserOwnedDrydockNote struct {
	UserOwnedDrydockNote *UserOwnedDrydockNote `protobuf:"bytes,3,opt,name=user_owned_drydock_note,json=userOwnedDrydockNote,proto3,oneof"`
}

func (*Attestor_UserOwnedDrydockNote) isAttestor_AttestorType() {}

func (m *Attestor) GetAttestorType() isAttestor_AttestorType {
	if m != nil {
		return m.AttestorType
	}
	return nil
}

func (m *Attestor) GetUserOwnedDrydockNote() *UserOwnedDrydockNote {
	if x, ok := m.GetAttestorType().(*Attestor_UserOwnedDrydockNote); ok {
		return x.UserOwnedDrydockNote
	}
	return nil
}

func (m *Attestor) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Attestor) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Attestor_UserOwnedDrydockNote)(nil),
	}
}

// An [user owned drydock note][google.cloud.binaryauthorization.v1beta1.UserOwnedDrydockNote] references a Drydock
// ATTESTATION_AUTHORITY Note created by the user.
type UserOwnedDrydockNote struct {
	// Required. The Drydock resource name of a ATTESTATION_AUTHORITY Note,
	// created by the user, in the format: `projects/*/notes/*` (or the legacy
	// `providers/*/notes/*`). This field may not be updated.
	//
	// An attestation by this attestor is stored as a Drydock
	// ATTESTATION_AUTHORITY Occurrence that names a container image and that
	// links to this Note. Drydock is an external dependency.
	NoteReference string `protobuf:"bytes,1,opt,name=note_reference,json=noteReference,proto3" json:"note_reference,omitempty"`
	// Optional. Public keys that verify attestations signed by this
	// attestor.  This field may be updated.
	//
	// If this field is non-empty, one of the specified public keys must
	// verify that an attestation was signed by this attestor for the
	// image specified in the admission request.
	//
	// If this field is empty, this attestor always returns that no
	// valid attestations exist.
	PublicKeys []*AttestorPublicKey `protobuf:"bytes,2,rep,name=public_keys,json=publicKeys,proto3" json:"public_keys,omitempty"`
	// Output only. This field will contain the service account email address
	// that this Attestor will use as the principal when querying Container
	// Analysis. Attestor administrators must grant this service account the
	// IAM role needed to read attestations from the [note_reference][Note] in
	// Container Analysis (`containeranalysis.notes.occurrences.viewer`).
	//
	// This email address is fixed for the lifetime of the Attestor, but callers
	// should not make any other assumptions about the service account email;
	// future versions may use an email based on a different naming pattern.
	DelegationServiceAccountEmail string   `protobuf:"bytes,3,opt,name=delegation_service_account_email,json=delegationServiceAccountEmail,proto3" json:"delegation_service_account_email,omitempty"`
	XXX_NoUnkeyedLiteral          struct{} `json:"-"`
	XXX_unrecognized              []byte   `json:"-"`
	XXX_sizecache                 int32    `json:"-"`
}

func (m *UserOwnedDrydockNote) Reset()         { *m = UserOwnedDrydockNote{} }
func (m *UserOwnedDrydockNote) String() string { return proto.CompactTextString(m) }
func (*UserOwnedDrydockNote) ProtoMessage()    {}
func (*UserOwnedDrydockNote) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b83e716c45ff98c, []int{4}
}

func (m *UserOwnedDrydockNote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserOwnedDrydockNote.Unmarshal(m, b)
}
func (m *UserOwnedDrydockNote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserOwnedDrydockNote.Marshal(b, m, deterministic)
}
func (m *UserOwnedDrydockNote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserOwnedDrydockNote.Merge(m, src)
}
func (m *UserOwnedDrydockNote) XXX_Size() int {
	return xxx_messageInfo_UserOwnedDrydockNote.Size(m)
}
func (m *UserOwnedDrydockNote) XXX_DiscardUnknown() {
	xxx_messageInfo_UserOwnedDrydockNote.DiscardUnknown(m)
}

var xxx_messageInfo_UserOwnedDrydockNote proto.InternalMessageInfo

func (m *UserOwnedDrydockNote) GetNoteReference() string {
	if m != nil {
		return m.NoteReference
	}
	return ""
}

func (m *UserOwnedDrydockNote) GetPublicKeys() []*AttestorPublicKey {
	if m != nil {
		return m.PublicKeys
	}
	return nil
}

func (m *UserOwnedDrydockNote) GetDelegationServiceAccountEmail() string {
	if m != nil {
		return m.DelegationServiceAccountEmail
	}
	return ""
}

// A public key in the PkixPublicKey format (see
// https://tools.ietf.org/html/rfc5280#section-4.1.2.7 for details).
// Public keys of this type are typically textually encoded using the PEM
// format.
type PkixPublicKey struct {
	// A PEM-encoded public key, as described in
	// https://tools.ietf.org/html/rfc7468#section-13
	PublicKeyPem string `protobuf:"bytes,1,opt,name=public_key_pem,json=publicKeyPem,proto3" json:"public_key_pem,omitempty"`
	// The signature algorithm used to verify a message against a signature using
	// this key.
	// These signature algorithm must match the structure and any object
	// identifiers encoded in `public_key_pem` (i.e. this algorithm must match
	// that of the public key).
	SignatureAlgorithm   PkixPublicKey_SignatureAlgorithm `protobuf:"varint,2,opt,name=signature_algorithm,json=signatureAlgorithm,proto3,enum=google.cloud.binaryauthorization.v1beta1.PkixPublicKey_SignatureAlgorithm" json:"signature_algorithm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *PkixPublicKey) Reset()         { *m = PkixPublicKey{} }
func (m *PkixPublicKey) String() string { return proto.CompactTextString(m) }
func (*PkixPublicKey) ProtoMessage()    {}
func (*PkixPublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b83e716c45ff98c, []int{5}
}

func (m *PkixPublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PkixPublicKey.Unmarshal(m, b)
}
func (m *PkixPublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PkixPublicKey.Marshal(b, m, deterministic)
}
func (m *PkixPublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PkixPublicKey.Merge(m, src)
}
func (m *PkixPublicKey) XXX_Size() int {
	return xxx_messageInfo_PkixPublicKey.Size(m)
}
func (m *PkixPublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PkixPublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_PkixPublicKey proto.InternalMessageInfo

func (m *PkixPublicKey) GetPublicKeyPem() string {
	if m != nil {
		return m.PublicKeyPem
	}
	return ""
}

func (m *PkixPublicKey) GetSignatureAlgorithm() PkixPublicKey_SignatureAlgorithm {
	if m != nil {
		return m.SignatureAlgorithm
	}
	return PkixPublicKey_SIGNATURE_ALGORITHM_UNSPECIFIED
}

// An [attestor public key][google.cloud.binaryauthorization.v1beta1.AttestorPublicKey] that will be used to verify
// attestations signed by this attestor.
type AttestorPublicKey struct {
	// Optional. A descriptive comment. This field may be updated.
	Comment string `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
	// The ID of this public key.
	// Signatures verified by BinAuthz must include the ID of the public key that
	// can be used to verify them, and that ID must match the contents of this
	// field exactly.
	// Additional restrictions on this field can be imposed based on which public
	// key type is encapsulated. See the documentation on `public_key` cases below
	// for details.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Required. A public key reference or serialized instance. This field may be
	// updated.
	//
	// Types that are valid to be assigned to PublicKey:
	//	*AttestorPublicKey_AsciiArmoredPgpPublicKey
	//	*AttestorPublicKey_PkixPublicKey
	PublicKey            isAttestorPublicKey_PublicKey `protobuf_oneof:"public_key"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *AttestorPublicKey) Reset()         { *m = AttestorPublicKey{} }
func (m *AttestorPublicKey) String() string { return proto.CompactTextString(m) }
func (*AttestorPublicKey) ProtoMessage()    {}
func (*AttestorPublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b83e716c45ff98c, []int{6}
}

func (m *AttestorPublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttestorPublicKey.Unmarshal(m, b)
}
func (m *AttestorPublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttestorPublicKey.Marshal(b, m, deterministic)
}
func (m *AttestorPublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestorPublicKey.Merge(m, src)
}
func (m *AttestorPublicKey) XXX_Size() int {
	return xxx_messageInfo_AttestorPublicKey.Size(m)
}
func (m *AttestorPublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestorPublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_AttestorPublicKey proto.InternalMessageInfo

func (m *AttestorPublicKey) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *AttestorPublicKey) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type isAttestorPublicKey_PublicKey interface {
	isAttestorPublicKey_PublicKey()
}

type AttestorPublicKey_AsciiArmoredPgpPublicKey struct {
	AsciiArmoredPgpPublicKey string `protobuf:"bytes,3,opt,name=ascii_armored_pgp_public_key,json=asciiArmoredPgpPublicKey,proto3,oneof"`
}

type AttestorPublicKey_PkixPublicKey struct {
	PkixPublicKey *PkixPublicKey `protobuf:"bytes,5,opt,name=pkix_public_key,json=pkixPublicKey,proto3,oneof"`
}

func (*AttestorPublicKey_AsciiArmoredPgpPublicKey) isAttestorPublicKey_PublicKey() {}

func (*AttestorPublicKey_PkixPublicKey) isAttestorPublicKey_PublicKey() {}

func (m *AttestorPublicKey) GetPublicKey() isAttestorPublicKey_PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *AttestorPublicKey) GetAsciiArmoredPgpPublicKey() string {
	if x, ok := m.GetPublicKey().(*AttestorPublicKey_AsciiArmoredPgpPublicKey); ok {
		return x.AsciiArmoredPgpPublicKey
	}
	return ""
}

func (m *AttestorPublicKey) GetPkixPublicKey() *PkixPublicKey {
	if x, ok := m.GetPublicKey().(*AttestorPublicKey_PkixPublicKey); ok {
		return x.PkixPublicKey
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AttestorPublicKey) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AttestorPublicKey_AsciiArmoredPgpPublicKey)(nil),
		(*AttestorPublicKey_PkixPublicKey)(nil),
	}
}

func init() {
	proto.RegisterEnum("google.cloud.binaryauthorization.v1beta1.Policy_GlobalPolicyEvaluationMode", Policy_GlobalPolicyEvaluationMode_name, Policy_GlobalPolicyEvaluationMode_value)
	proto.RegisterEnum("google.cloud.binaryauthorization.v1beta1.AdmissionRule_EvaluationMode", AdmissionRule_EvaluationMode_name, AdmissionRule_EvaluationMode_value)
	proto.RegisterEnum("google.cloud.binaryauthorization.v1beta1.AdmissionRule_EnforcementMode", AdmissionRule_EnforcementMode_name, AdmissionRule_EnforcementMode_value)
	proto.RegisterEnum("google.cloud.binaryauthorization.v1beta1.PkixPublicKey_SignatureAlgorithm", PkixPublicKey_SignatureAlgorithm_name, PkixPublicKey_SignatureAlgorithm_value)
	proto.RegisterType((*Policy)(nil), "google.cloud.binaryauthorization.v1beta1.Policy")
	proto.RegisterMapType((map[string]*AdmissionRule)(nil), "google.cloud.binaryauthorization.v1beta1.Policy.ClusterAdmissionRulesEntry")
	proto.RegisterType((*AdmissionWhitelistPattern)(nil), "google.cloud.binaryauthorization.v1beta1.AdmissionWhitelistPattern")
	proto.RegisterType((*AdmissionRule)(nil), "google.cloud.binaryauthorization.v1beta1.AdmissionRule")
	proto.RegisterType((*Attestor)(nil), "google.cloud.binaryauthorization.v1beta1.Attestor")
	proto.RegisterType((*UserOwnedDrydockNote)(nil), "google.cloud.binaryauthorization.v1beta1.UserOwnedDrydockNote")
	proto.RegisterType((*PkixPublicKey)(nil), "google.cloud.binaryauthorization.v1beta1.PkixPublicKey")
	proto.RegisterType((*AttestorPublicKey)(nil), "google.cloud.binaryauthorization.v1beta1.AttestorPublicKey")
}

func init() {
	proto.RegisterFile("google/cloud/binaryauthorization/v1beta1/resources.proto", fileDescriptor_1b83e716c45ff98c)
}

var fileDescriptor_1b83e716c45ff98c = []byte{
	// 1214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x5f, 0x73, 0xdb, 0xc4,
	0x17, 0x8d, 0x9c, 0x7f, 0xcd, 0x75, 0xe2, 0xb8, 0xdb, 0xe6, 0x17, 0xd7, 0xbf, 0x94, 0x1a, 0x03,
	0x33, 0xe1, 0x01, 0xbb, 0x71, 0x92, 0x36, 0xd0, 0x99, 0x0e, 0xb2, 0xad, 0x3a, 0x26, 0x8e, 0x6d,
	0x64, 0x9b, 0x4e, 0x80, 0x99, 0x65, 0x2d, 0x6d, 0x54, 0x4d, 0x24, 0xad, 0x58, 0x49, 0x6d, 0x0d,
	0x4f, 0x0c, 0xc3, 0xf0, 0x04, 0xaf, 0x7c, 0x3f, 0x86, 0x0f, 0xc0, 0x0c, 0x2f, 0x3c, 0x32, 0xab,
	0x3f, 0xb1, 0xe3, 0x3a, 0x25, 0x09, 0x6f, 0xde, 0x7b, 0xee, 0x9e, 0x7b, 0xf6, 0xee, 0xb9, 0x2b,
	0xc3, 0x81, 0xc1, 0x98, 0x61, 0xd1, 0xb2, 0x66, 0xb1, 0x40, 0x2f, 0x0f, 0x4d, 0x87, 0xf0, 0x11,
	0x09, 0xfc, 0x17, 0x8c, 0x9b, 0xdf, 0x11, 0xdf, 0x64, 0x4e, 0xf9, 0xe5, 0xce, 0x90, 0xfa, 0x64,
	0xa7, 0xcc, 0xa9, 0xc7, 0x02, 0xae, 0x51, 0xaf, 0xe4, 0x72, 0xe6, 0x33, 0xb4, 0x1d, 0xed, 0x2c,
	0x85, 0x3b, 0x4b, 0x33, 0x76, 0x96, 0xe2, 0x9d, 0xf9, 0xad, 0xb8, 0x06, 0x71, 0xcd, 0x32, 0x71,
	0x1c, 0xe6, 0x87, 0x70, 0xcc, 0x93, 0x7f, 0x10, 0xa3, 0xe1, 0x6a, 0x18, 0x9c, 0x96, 0x7d, 0xd3,
	0xa6, 0x9e, 0x4f, 0x6c, 0x37, 0x4a, 0x28, 0xfe, 0xb5, 0x04, 0x4b, 0x5d, 0x66, 0x99, 0xda, 0x08,
	0x21, 0x58, 0x70, 0x88, 0x4d, 0x73, 0x52, 0x41, 0xda, 0x5e, 0x51, 0xc3, 0xdf, 0xa8, 0x00, 0x69,
	0x9d, 0x7a, 0x1a, 0x37, 0x5d, 0xc1, 0x9a, 0x5b, 0x0a, 0xa1, 0xc9, 0x10, 0xfa, 0x55, 0x82, 0xfb,
	0x86, 0xc5, 0x86, 0xc4, 0xc2, 0x6e, 0xc8, 0x83, 0xe9, 0x4b, 0x62, 0x05, 0xa1, 0x0a, 0x6c, 0x33,
	0x9d, 0xe6, 0x96, 0x0b, 0xd2, 0x76, 0xa6, 0x72, 0x54, 0xba, 0xea, 0x91, 0x4a, 0x91, 0x9e, 0x52,
	0x23, 0x64, 0x8d, 0x16, 0xca, 0x39, 0xe7, 0x31, 0xd3, 0xa9, 0x9a, 0x37, 0x2e, 0xc5, 0xd0, 0x4f,
	0x12, 0x6c, 0x11, 0xdd, 0x36, 0x3d, 0x4f, 0x28, 0x78, 0xf5, 0xc2, 0xf4, 0xa9, 0x65, 0x7a, 0x3e,
	0x76, 0x89, 0xef, 0x53, 0xee, 0x78, 0xb9, 0x54, 0x61, 0x7e, 0x3b, 0x5d, 0xa9, 0x5d, 0x5d, 0x8f,
	0x9c, 0xb0, 0x3d, 0x4f, 0xc8, 0xba, 0x11, 0x97, 0x9a, 0x27, 0x97, 0x41, 0x1e, 0xfa, 0x51, 0x82,
	0x4d, 0xcd, 0x0a, 0x3c, 0x9f, 0x72, 0x3c, 0xd6, 0xc3, 0x03, 0x8b, 0x7a, 0xb9, 0xf9, 0x50, 0xc2,
	0xf5, 0x5b, 0x52, 0x8b, 0xf8, 0xce, 0x05, 0xa9, 0x82, 0x4d, 0x71, 0x7c, 0x3e, 0x52, 0x37, 0xb4,
	0x59, 0x18, 0xb2, 0xe1, 0x7f, 0x3a, 0x3d, 0x25, 0x81, 0xe5, 0x4f, 0x89, 0xc8, 0x2d, 0x14, 0xa4,
	0xed, 0x74, 0xe5, 0xf1, 0x0d, 0xda, 0x20, 0x98, 0xd5, 0xbb, 0x31, 0xed, 0x85, 0x28, 0x7a, 0x02,
	0xe9, 0xc0, 0xd5, 0x89, 0x4f, 0xb1, 0x30, 0x5a, 0x6e, 0x31, 0xac, 0x91, 0x4f, 0x6a, 0x24, 0x2e,
	0x2c, 0xf5, 0x13, 0x17, 0xaa, 0x10, 0xa5, 0x8b, 0x40, 0xfe, 0x07, 0x09, 0xf2, 0x97, 0x9f, 0x10,
	0x65, 0x61, 0xfe, 0x8c, 0x8e, 0x62, 0x7b, 0x8a, 0x9f, 0xe8, 0x18, 0x16, 0xc5, 0xdd, 0xd3, 0x5c,
	0xea, 0xbf, 0x9d, 0x25, 0x62, 0xf9, 0x24, 0x75, 0x20, 0x15, 0x75, 0xc8, 0x5f, 0xee, 0x3b, 0xf4,
	0x11, 0x7c, 0xd8, 0x68, 0x75, 0xaa, 0x72, 0x0b, 0x77, 0x3b, 0xad, 0x66, 0xed, 0x04, 0x2b, 0x5f,
	0xc8, 0xad, 0x81, 0xdc, 0x6f, 0x76, 0xda, 0xf8, 0xb8, 0x53, 0x57, 0xf0, 0xa0, 0xdd, 0xeb, 0x2a,
	0xb5, 0xe6, 0xb3, 0xa6, 0x52, 0xcf, 0xce, 0x21, 0x80, 0x25, 0xa5, 0x2d, 0x57, 0x5b, 0x4a, 0x56,
	0x42, 0x69, 0x58, 0xae, 0x37, 0x7b, 0xe1, 0x22, 0x55, 0x7c, 0x0a, 0xf7, 0x2e, 0x35, 0x15, 0x7a,
	0x17, 0x56, 0xc5, 0xec, 0x25, 0x86, 0x8d, 0x0f, 0x9c, 0x16, 0xb1, 0x38, 0xa5, 0xf8, 0xcb, 0x02,
	0xac, 0x5d, 0x6c, 0x3c, 0x83, 0xf5, 0xe9, 0xb9, 0x93, 0xc2, 0xb9, 0x7b, 0x76, 0xc3, 0xa6, 0x94,
	0xa6, 0x46, 0x2e, 0x43, 0x2f, 0xb6, 0xe2, 0x11, 0x6c, 0x72, 0xfa, 0x6d, 0x60, 0x72, 0x8a, 0x85,
	0x28, 0x2f, 0x7e, 0x77, 0xf0, 0x70, 0x14, 0x0e, 0xd8, 0x8a, 0xba, 0x11, 0xc3, 0xf2, 0x04, 0x5a,
	0x1d, 0x21, 0x0e, 0x59, 0xea, 0x9c, 0x32, 0xae, 0x51, 0x9b, 0x3a, 0x7e, 0xa4, 0x74, 0x3e, 0x54,
	0xda, 0xb8, 0xb1, 0xd2, 0x31, 0x5f, 0x28, 0x75, 0x9d, 0x5e, 0x0c, 0x14, 0x6d, 0xc8, 0x4c, 0x5d,
	0xe4, 0x03, 0xf8, 0xff, 0xdb, 0xaf, 0x2e, 0x0b, 0xab, 0x72, 0xeb, 0xb9, 0x7c, 0xd2, 0xc3, 0x72,
	0xab, 0xd5, 0x79, 0x9e, 0x95, 0xd0, 0x26, 0xdc, 0x51, 0x95, 0xcf, 0x07, 0x4d, 0x55, 0xc1, 0x72,
	0xbf, 0xaf, 0xf4, 0xfa, 0xe1, 0xde, 0x6c, 0x0a, 0xad, 0x43, 0x3a, 0x4e, 0xad, 0x2b, 0xed, 0x93,
	0xec, 0x7c, 0xd1, 0x85, 0xf5, 0x29, 0x49, 0xa8, 0x00, 0x5b, 0x4a, 0xfb, 0x59, 0x47, 0xad, 0x29,
	0xc7, 0x4a, 0xbb, 0x3f, 0xab, 0xe0, 0x38, 0xa3, 0x8e, 0xab, 0xad, 0x4e, 0xed, 0x08, 0xcb, 0xed,
	0x3a, 0x96, 0x07, 0xf5, 0x66, 0x1f, 0xb7, 0x3a, 0x8d, 0xac, 0x84, 0xee, 0xc1, 0x46, 0x5d, 0x3d,
	0x51, 0x07, 0xed, 0x71, 0x14, 0x77, 0xda, 0xad, 0x93, 0x6c, 0xaa, 0xf8, 0x73, 0x0a, 0x6e, 0x45,
	0x7d, 0x66, 0xfc, 0x86, 0xef, 0xf8, 0x2b, 0xd8, 0x0c, 0x3c, 0xca, 0x31, 0x7b, 0xe5, 0x50, 0x1d,
	0xeb, 0x7c, 0xa4, 0x33, 0xed, 0x0c, 0x3b, 0xcc, 0x8f, 0xae, 0x27, 0x5d, 0x79, 0x7a, 0xf5, 0xeb,
	0x19, 0x78, 0x94, 0x77, 0x04, 0x4f, 0x3d, 0xa2, 0x69, 0x33, 0x9f, 0x1e, 0xce, 0xa9, 0x77, 0x83,
	0x19, 0xf1, 0xe9, 0x27, 0x63, 0xe1, 0x3a, 0x4f, 0x46, 0x75, 0x1d, 0xd6, 0x48, 0x7c, 0x6e, 0xec,
	0x8f, 0x5c, 0x5a, 0xfc, 0x43, 0x82, 0xbb, 0xb3, 0xca, 0xa3, 0x0f, 0x20, 0x23, 0x0e, 0x83, 0x39,
	0x3d, 0xa5, 0x9c, 0x3a, 0x5a, 0xd2, 0x9f, 0x35, 0x11, 0x55, 0x93, 0x20, 0xfa, 0x1a, 0xd2, 0x6e,
	0x30, 0xb4, 0x4c, 0x0d, 0x9f, 0xd1, 0x51, 0xf2, 0xad, 0x78, 0x72, 0x0d, 0x67, 0xc6, 0x6a, 0xba,
	0x21, 0xc9, 0x11, 0x1d, 0xa9, 0xe0, 0x26, 0x3f, 0x3d, 0xd4, 0x80, 0x82, 0x4e, 0x2d, 0x6a, 0x44,
	0x53, 0xea, 0x51, 0xfe, 0xd2, 0xd4, 0x28, 0x26, 0x9a, 0xc6, 0x02, 0xc7, 0xc7, 0xd4, 0x26, 0xa6,
	0x15, 0x76, 0x7b, 0x45, 0xbd, 0x3f, 0xce, 0xeb, 0x45, 0x69, 0x72, 0x94, 0xa5, 0x88, 0xa4, 0xe2,
	0x6f, 0x0b, 0xb0, 0xd6, 0x3d, 0x33, 0x5f, 0x9f, 0x97, 0x41, 0xef, 0x43, 0x66, 0x2c, 0x1c, 0xbb,
	0xd4, 0x8e, 0xcf, 0xb7, 0x7a, 0x5e, 0xbe, 0x4b, 0x6d, 0xf4, 0x3d, 0xdc, 0xf1, 0x4c, 0xc3, 0x21,
	0x7e, 0x20, 0xe6, 0xd6, 0x32, 0x18, 0x37, 0xfd, 0x17, 0x76, 0xf8, 0x7e, 0x66, 0x2a, 0x9f, 0x5d,
	0xe3, 0x7b, 0x34, 0x59, 0xbb, 0xd4, 0x4b, 0x28, 0xe5, 0x84, 0x51, 0x45, 0xde, 0x1b, 0xb1, 0xe2,
	0xef, 0x29, 0x40, 0x6f, 0xa6, 0xa2, 0xf7, 0xe0, 0x41, 0xaf, 0xd9, 0x68, 0xcb, 0xfd, 0x81, 0x18,
	0xad, 0x56, 0xa3, 0xa3, 0x36, 0xfb, 0x87, 0xc7, 0x53, 0xe3, 0x21, 0xa6, 0xaf, 0x27, 0xe3, 0x6e,
	0xaf, 0x87, 0x2b, 0x0f, 0xf7, 0x0e, 0x70, 0xef, 0x50, 0xae, 0xec, 0x3f, 0x8a, 0xc7, 0x32, 0x06,
	0x76, 0x1f, 0x3e, 0xae, 0x24, 0x40, 0x6a, 0x12, 0xd8, 0x7b, 0xf8, 0xf1, 0xa3, 0x04, 0x98, 0x9f,
	0x05, 0xec, 0xef, 0x54, 0xb2, 0x0b, 0xe8, 0x1d, 0xc8, 0x0b, 0x40, 0x88, 0xc1, 0xdd, 0xa3, 0x5a,
	0x6f, 0xe7, 0x42, 0xa9, 0xc5, 0x19, 0xf8, 0x64, 0xc5, 0xa5, 0x19, 0xf8, 0x64, 0xe1, 0xe5, 0xb7,
	0xe0, 0xa2, 0xfe, 0x2d, 0xb4, 0x01, 0xb7, 0x95, 0x5a, 0x5d, 0x48, 0xab, 0xec, 0x9f, 0x6f, 0x5b,
	0x99, 0x08, 0xef, 0x1e, 0xec, 0x89, 0xf0, 0xee, 0xc1, 0x5e, 0x16, 0x26, 0xc2, 0xfb, 0x95, 0x9d,
	0x84, 0x24, 0x5d, 0xfc, 0x53, 0x82, 0xdb, 0x6f, 0x98, 0x10, 0xe5, 0x60, 0x59, 0x63, 0xb6, 0x78,
	0x8e, 0x62, 0x5b, 0x24, 0x4b, 0x94, 0x81, 0x94, 0xa9, 0x87, 0x06, 0x58, 0x51, 0x53, 0xa6, 0x8e,
	0x3e, 0x85, 0x2d, 0xe2, 0x69, 0xa6, 0x89, 0x09, 0xb7, 0x19, 0xa7, 0x3a, 0x76, 0x0d, 0x17, 0x8f,
	0x9d, 0x15, 0xd9, 0xf3, 0x70, 0x4e, 0xcd, 0x85, 0x59, 0x72, 0x94, 0xd4, 0x35, 0xdc, 0x71, 0x2d,
	0x02, 0xeb, 0xee, 0x99, 0xf9, 0x7a, 0x72, 0xd3, 0xe2, 0x75, 0xbf, 0xcf, 0x17, 0xfc, 0x75, 0x38,
	0xa7, 0xae, 0xb9, 0x93, 0x81, 0xea, 0x2a, 0xc0, 0x98, 0xbd, 0xfa, 0xcd, 0x97, 0x5f, 0xc5, 0xc4,
	0x06, 0xb3, 0x88, 0x63, 0x94, 0x18, 0x37, 0xca, 0x06, 0x75, 0xc2, 0xb7, 0xa3, 0x1c, 0x41, 0xc4,
	0x35, 0xbd, 0x7f, 0xff, 0x1f, 0xfe, 0x64, 0x06, 0xf6, 0xb7, 0x24, 0x0d, 0x97, 0x42, 0xaa, 0xdd,
	0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x82, 0x0c, 0x36, 0xd1, 0x0b, 0x00, 0x00,
}
