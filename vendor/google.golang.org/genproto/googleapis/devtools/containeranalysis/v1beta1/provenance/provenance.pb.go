// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/containeranalysis/v1beta1/provenance/provenance.proto

package provenance

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	source "google.golang.org/genproto/googleapis/devtools/containeranalysis/v1beta1/source"
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

// Specifies the hash algorithm, if any.
type Hash_HashType int32

const (
	// Unknown.
	Hash_HASH_TYPE_UNSPECIFIED Hash_HashType = 0
	// A SHA-256 hash.
	Hash_SHA256 Hash_HashType = 1
)

var Hash_HashType_name = map[int32]string{
	0: "HASH_TYPE_UNSPECIFIED",
	1: "SHA256",
}

var Hash_HashType_value = map[string]int32{
	"HASH_TYPE_UNSPECIFIED": 0,
	"SHA256":                1,
}

func (x Hash_HashType) String() string {
	return proto.EnumName(Hash_HashType_name, int32(x))
}

func (Hash_HashType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4ab6101dd6eab75c, []int{3, 0}
}

// Provenance of a build. Contains all information needed to verify the full
// details about the build from source to completion.
type BuildProvenance struct {
	// Unique identifier of the build.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the project.
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Commands requested by the build.
	Commands []*Command `protobuf:"bytes,3,rep,name=commands,proto3" json:"commands,omitempty"`
	// Output of the build.
	BuiltArtifacts []*Artifact `protobuf:"bytes,4,rep,name=built_artifacts,json=builtArtifacts,proto3" json:"built_artifacts,omitempty"`
	// Time at which the build was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Time at which execution of the build was started.
	StartTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Time at which execution of the build was finished.
	EndTime *timestamp.Timestamp `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// E-mail address of the user who initiated this build. Note that this was the
	// user's e-mail address at the time the build was initiated; this address may
	// not represent the same end-user for all time.
	Creator string `protobuf:"bytes,8,opt,name=creator,proto3" json:"creator,omitempty"`
	// URI where any logs for this provenance were written.
	LogsUri string `protobuf:"bytes,9,opt,name=logs_uri,json=logsUri,proto3" json:"logs_uri,omitempty"`
	// Details of the Source input to the build.
	SourceProvenance *Source `protobuf:"bytes,10,opt,name=source_provenance,json=sourceProvenance,proto3" json:"source_provenance,omitempty"`
	// Trigger identifier if the build was triggered automatically; empty if not.
	TriggerId string `protobuf:"bytes,11,opt,name=trigger_id,json=triggerId,proto3" json:"trigger_id,omitempty"`
	// Special options applied to this build. This is a catch-all field where
	// build providers can enter any desired additional details.
	BuildOptions map[string]string `protobuf:"bytes,12,rep,name=build_options,json=buildOptions,proto3" json:"build_options,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Version string of the builder at the time this build was executed.
	BuilderVersion       string   `protobuf:"bytes,13,opt,name=builder_version,json=builderVersion,proto3" json:"builder_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildProvenance) Reset()         { *m = BuildProvenance{} }
func (m *BuildProvenance) String() string { return proto.CompactTextString(m) }
func (*BuildProvenance) ProtoMessage()    {}
func (*BuildProvenance) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ab6101dd6eab75c, []int{0}
}

func (m *BuildProvenance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildProvenance.Unmarshal(m, b)
}
func (m *BuildProvenance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildProvenance.Marshal(b, m, deterministic)
}
func (m *BuildProvenance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildProvenance.Merge(m, src)
}
func (m *BuildProvenance) XXX_Size() int {
	return xxx_messageInfo_BuildProvenance.Size(m)
}
func (m *BuildProvenance) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildProvenance.DiscardUnknown(m)
}

var xxx_messageInfo_BuildProvenance proto.InternalMessageInfo

func (m *BuildProvenance) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BuildProvenance) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *BuildProvenance) GetCommands() []*Command {
	if m != nil {
		return m.Commands
	}
	return nil
}

func (m *BuildProvenance) GetBuiltArtifacts() []*Artifact {
	if m != nil {
		return m.BuiltArtifacts
	}
	return nil
}

func (m *BuildProvenance) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *BuildProvenance) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *BuildProvenance) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *BuildProvenance) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *BuildProvenance) GetLogsUri() string {
	if m != nil {
		return m.LogsUri
	}
	return ""
}

func (m *BuildProvenance) GetSourceProvenance() *Source {
	if m != nil {
		return m.SourceProvenance
	}
	return nil
}

func (m *BuildProvenance) GetTriggerId() string {
	if m != nil {
		return m.TriggerId
	}
	return ""
}

func (m *BuildProvenance) GetBuildOptions() map[string]string {
	if m != nil {
		return m.BuildOptions
	}
	return nil
}

func (m *BuildProvenance) GetBuilderVersion() string {
	if m != nil {
		return m.BuilderVersion
	}
	return ""
}

// Source describes the location of the source used for the build.
type Source struct {
	// If provided, the input binary artifacts for the build came from this
	// location.
	ArtifactStorageSourceUri string `protobuf:"bytes,1,opt,name=artifact_storage_source_uri,json=artifactStorageSourceUri,proto3" json:"artifact_storage_source_uri,omitempty"`
	// Hash(es) of the build source, which can be used to verify that the original
	// source integrity was maintained in the build.
	//
	// The keys to this map are file paths used as build source and the values
	// contain the hash values for those files.
	//
	// If the build source came in a single package such as a gzipped tarfile
	// (.tar.gz), the FileHash will be for the single path to that file.
	FileHashes map[string]*FileHashes `protobuf:"bytes,2,rep,name=file_hashes,json=fileHashes,proto3" json:"file_hashes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// If provided, the source code used for the build came from this location.
	Context *source.SourceContext `protobuf:"bytes,3,opt,name=context,proto3" json:"context,omitempty"`
	// If provided, some of the source code used for the build may be found in
	// these locations, in the case where the source repository had multiple
	// remotes or submodules. This list will not include the context specified in
	// the context field.
	AdditionalContexts   []*source.SourceContext `protobuf:"bytes,4,rep,name=additional_contexts,json=additionalContexts,proto3" json:"additional_contexts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Source) Reset()         { *m = Source{} }
func (m *Source) String() string { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()    {}
func (*Source) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ab6101dd6eab75c, []int{1}
}

func (m *Source) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Source.Unmarshal(m, b)
}
func (m *Source) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Source.Marshal(b, m, deterministic)
}
func (m *Source) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Source.Merge(m, src)
}
func (m *Source) XXX_Size() int {
	return xxx_messageInfo_Source.Size(m)
}
func (m *Source) XXX_DiscardUnknown() {
	xxx_messageInfo_Source.DiscardUnknown(m)
}

var xxx_messageInfo_Source proto.InternalMessageInfo

func (m *Source) GetArtifactStorageSourceUri() string {
	if m != nil {
		return m.ArtifactStorageSourceUri
	}
	return ""
}

func (m *Source) GetFileHashes() map[string]*FileHashes {
	if m != nil {
		return m.FileHashes
	}
	return nil
}

func (m *Source) GetContext() *source.SourceContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Source) GetAdditionalContexts() []*source.SourceContext {
	if m != nil {
		return m.AdditionalContexts
	}
	return nil
}

// Container message for hashes of byte content of files, used in Source
// messages to verify integrity of source input to the build.
type FileHashes struct {
	// Collection of file hashes.
	FileHash             []*Hash  `protobuf:"bytes,1,rep,name=file_hash,json=fileHash,proto3" json:"file_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileHashes) Reset()         { *m = FileHashes{} }
func (m *FileHashes) String() string { return proto.CompactTextString(m) }
func (*FileHashes) ProtoMessage()    {}
func (*FileHashes) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ab6101dd6eab75c, []int{2}
}

func (m *FileHashes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileHashes.Unmarshal(m, b)
}
func (m *FileHashes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileHashes.Marshal(b, m, deterministic)
}
func (m *FileHashes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileHashes.Merge(m, src)
}
func (m *FileHashes) XXX_Size() int {
	return xxx_messageInfo_FileHashes.Size(m)
}
func (m *FileHashes) XXX_DiscardUnknown() {
	xxx_messageInfo_FileHashes.DiscardUnknown(m)
}

var xxx_messageInfo_FileHashes proto.InternalMessageInfo

func (m *FileHashes) GetFileHash() []*Hash {
	if m != nil {
		return m.FileHash
	}
	return nil
}

// Container message for hash values.
type Hash struct {
	// The type of hash that was performed.
	Type Hash_HashType `protobuf:"varint,1,opt,name=type,proto3,enum=grafeas.v1beta1.provenance.Hash_HashType" json:"type,omitempty"`
	// The hash value.
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hash) Reset()         { *m = Hash{} }
func (m *Hash) String() string { return proto.CompactTextString(m) }
func (*Hash) ProtoMessage()    {}
func (*Hash) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ab6101dd6eab75c, []int{3}
}

func (m *Hash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hash.Unmarshal(m, b)
}
func (m *Hash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hash.Marshal(b, m, deterministic)
}
func (m *Hash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hash.Merge(m, src)
}
func (m *Hash) XXX_Size() int {
	return xxx_messageInfo_Hash.Size(m)
}
func (m *Hash) XXX_DiscardUnknown() {
	xxx_messageInfo_Hash.DiscardUnknown(m)
}

var xxx_messageInfo_Hash proto.InternalMessageInfo

func (m *Hash) GetType() Hash_HashType {
	if m != nil {
		return m.Type
	}
	return Hash_HASH_TYPE_UNSPECIFIED
}

func (m *Hash) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// Command describes a step performed as part of the build pipeline.
type Command struct {
	// Name of the command, as presented on the command line, or if the command is
	// packaged as a Docker container, as presented to `docker pull`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Environment variables set before running this command.
	Env []string `protobuf:"bytes,2,rep,name=env,proto3" json:"env,omitempty"`
	// Command-line arguments used when executing this command.
	Args []string `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	// Working directory (relative to project source root) used when running this
	// command.
	Dir string `protobuf:"bytes,4,opt,name=dir,proto3" json:"dir,omitempty"`
	// Optional unique identifier for this command, used in wait_for to reference
	// this command as a dependency.
	Id string `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	// The ID(s) of the command(s) that this command depends on.
	WaitFor              []string `protobuf:"bytes,6,rep,name=wait_for,json=waitFor,proto3" json:"wait_for,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ab6101dd6eab75c, []int{4}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Command) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *Command) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Command) GetDir() string {
	if m != nil {
		return m.Dir
	}
	return ""
}

func (m *Command) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Command) GetWaitFor() []string {
	if m != nil {
		return m.WaitFor
	}
	return nil
}

// Artifact describes a build product.
type Artifact struct {
	// Hash or checksum value of a binary, or Docker Registry 2.0 digest of a
	// container.
	Checksum string `protobuf:"bytes,1,opt,name=checksum,proto3" json:"checksum,omitempty"`
	// Artifact ID, if any; for container images, this will be a URL by digest
	// like `gcr.io/projectID/imagename@sha256:123456`.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Related artifact names. This may be the path to a binary or jar file, or in
	// the case of a container build, the name used to push the container image to
	// Google Container Registry, as presented to `docker push`. Note that a
	// single Artifact ID can have multiple names, for example if two tags are
	// applied to one image.
	Names                []string `protobuf:"bytes,3,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Artifact) Reset()         { *m = Artifact{} }
func (m *Artifact) String() string { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()    {}
func (*Artifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ab6101dd6eab75c, []int{5}
}

func (m *Artifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifact.Unmarshal(m, b)
}
func (m *Artifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifact.Marshal(b, m, deterministic)
}
func (m *Artifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifact.Merge(m, src)
}
func (m *Artifact) XXX_Size() int {
	return xxx_messageInfo_Artifact.Size(m)
}
func (m *Artifact) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifact.DiscardUnknown(m)
}

var xxx_messageInfo_Artifact proto.InternalMessageInfo

func (m *Artifact) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

func (m *Artifact) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Artifact) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func init() {
	proto.RegisterEnum("grafeas.v1beta1.provenance.Hash_HashType", Hash_HashType_name, Hash_HashType_value)
	proto.RegisterType((*BuildProvenance)(nil), "grafeas.v1beta1.provenance.BuildProvenance")
	proto.RegisterMapType((map[string]string)(nil), "grafeas.v1beta1.provenance.BuildProvenance.BuildOptionsEntry")
	proto.RegisterType((*Source)(nil), "grafeas.v1beta1.provenance.Source")
	proto.RegisterMapType((map[string]*FileHashes)(nil), "grafeas.v1beta1.provenance.Source.FileHashesEntry")
	proto.RegisterType((*FileHashes)(nil), "grafeas.v1beta1.provenance.FileHashes")
	proto.RegisterType((*Hash)(nil), "grafeas.v1beta1.provenance.Hash")
	proto.RegisterType((*Command)(nil), "grafeas.v1beta1.provenance.Command")
	proto.RegisterType((*Artifact)(nil), "grafeas.v1beta1.provenance.Artifact")
}

func init() {
	proto.RegisterFile("google/devtools/containeranalysis/v1beta1/provenance/provenance.proto", fileDescriptor_4ab6101dd6eab75c)
}

var fileDescriptor_4ab6101dd6eab75c = []byte{
	// 848 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x4b, 0x6f, 0x23, 0x45,
	0x10, 0x66, 0xfc, 0x1c, 0x97, 0xb3, 0x49, 0xb6, 0x59, 0xa4, 0x8e, 0x51, 0x44, 0x64, 0x5e, 0xe1,
	0x32, 0x56, 0x8c, 0x16, 0x01, 0x8b, 0x15, 0x65, 0x83, 0x43, 0x22, 0x1e, 0x1b, 0x8d, 0xb3, 0x2b,
	0xc1, 0x81, 0x51, 0x7b, 0xba, 0x3d, 0x69, 0x76, 0x3c, 0x6d, 0x75, 0xb7, 0x0d, 0xbe, 0x71, 0xe3,
	0xc6, 0x2f, 0xe0, 0xc4, 0x4f, 0xe4, 0x17, 0xa0, 0x7e, 0x8c, 0x6d, 0x76, 0xc1, 0x09, 0x97, 0xa4,
	0xab, 0xe6, 0xfb, 0xbe, 0xae, 0xf9, 0xaa, 0xa6, 0x0c, 0xc3, 0x4c, 0x88, 0x2c, 0x67, 0x3d, 0xca,
	0x16, 0x5a, 0x88, 0x5c, 0xf5, 0x52, 0x51, 0x68, 0xc2, 0x0b, 0x26, 0x49, 0x41, 0xf2, 0xa5, 0xe2,
	0xaa, 0xb7, 0x38, 0x19, 0x33, 0x4d, 0x4e, 0x7a, 0x33, 0x29, 0x16, 0xac, 0x20, 0x45, 0xca, 0x36,
	0x8e, 0xd1, 0x4c, 0x0a, 0x2d, 0x50, 0x27, 0x93, 0x64, 0xc2, 0x88, 0x8a, 0x3c, 0x38, 0x5a, 0x23,
	0x3a, 0x83, 0xfb, 0x5f, 0xa1, 0xc4, 0x5c, 0xa6, 0xcc, 0xff, 0x73, 0xd2, 0x9d, 0x77, 0x3c, 0xdd,
	0x46, 0xe3, 0xf9, 0xa4, 0xa7, 0xf9, 0x94, 0x29, 0x4d, 0xa6, 0x33, 0x07, 0xe8, 0xfe, 0x55, 0x87,
	0xbd, 0xa7, 0x73, 0x9e, 0xd3, 0xeb, 0xd5, 0x9d, 0x68, 0x17, 0x2a, 0x9c, 0xe2, 0xe0, 0x28, 0x38,
	0x6e, 0xc5, 0x15, 0x4e, 0xd1, 0x21, 0xc0, 0x4c, 0x8a, 0x9f, 0x58, 0xaa, 0x13, 0x4e, 0x71, 0xc5,
	0xe6, 0x5b, 0x3e, 0x73, 0x45, 0xd1, 0x29, 0x84, 0xa9, 0x98, 0x4e, 0x49, 0x41, 0x15, 0xae, 0x1e,
	0x55, 0x8f, 0xdb, 0xfd, 0x77, 0xa3, 0xff, 0x7e, 0xa3, 0xe8, 0xdc, 0x61, 0xe3, 0x15, 0x09, 0x7d,
	0x0b, 0x7b, 0xe3, 0x39, 0xcf, 0x75, 0x42, 0xa4, 0xe6, 0x13, 0x92, 0x6a, 0x85, 0x6b, 0x56, 0xe7,
	0xbd, 0x6d, 0x3a, 0x67, 0x1e, 0x1c, 0xef, 0x5a, 0x72, 0x19, 0x2a, 0xf4, 0x04, 0xda, 0xa9, 0x64,
	0x44, 0xb3, 0xc4, 0xbc, 0x2c, 0xae, 0x1f, 0x05, 0xc7, 0xed, 0x7e, 0x27, 0x72, 0x4e, 0x44, 0xa5,
	0x13, 0xd1, 0x4d, 0xe9, 0x44, 0x0c, 0x0e, 0x6e, 0x12, 0xe8, 0x33, 0x00, 0xa5, 0x89, 0xd4, 0x8e,
	0xdb, 0xb8, 0x93, 0xdb, 0xb2, 0x68, 0x4b, 0x7d, 0x0c, 0x21, 0x2b, 0xa8, 0x23, 0x36, 0xef, 0x24,
	0x36, 0x59, 0x41, 0x2d, 0x0d, 0x43, 0xd3, 0xde, 0x2f, 0x24, 0x0e, 0xad, 0xb5, 0x65, 0x88, 0x0e,
	0x20, 0xcc, 0x45, 0xa6, 0x92, 0xb9, 0xe4, 0xb8, 0xe5, 0x1e, 0x99, 0xf8, 0xb9, 0xe4, 0xe8, 0x19,
	0x3c, 0x74, 0x7d, 0x4e, 0xd6, 0x8e, 0x60, 0xb0, 0x97, 0x76, 0xb7, 0x99, 0x36, 0xb2, 0xa4, 0x78,
	0xdf, 0x91, 0x37, 0x7a, 0x7e, 0x08, 0xa0, 0x25, 0xcf, 0x32, 0x26, 0x4d, 0x8f, 0xdb, 0xae, 0xc7,
	0x3e, 0x73, 0x45, 0xd1, 0x18, 0x1e, 0x18, 0x97, 0x69, 0x22, 0x66, 0x9a, 0x8b, 0x42, 0xe1, 0x1d,
	0xdb, 0xa0, 0xc1, 0xb6, 0xbb, 0x5e, 0x19, 0x2b, 0x17, 0x3f, 0x73, 0xfc, 0x61, 0xa1, 0xe5, 0x32,
	0xde, 0x19, 0x6f, 0xa4, 0xd0, 0x87, 0x6e, 0x0c, 0x28, 0x93, 0xc9, 0x82, 0x49, 0xc5, 0x45, 0x81,
	0x1f, 0xd8, 0x3a, 0x76, 0x7d, 0xfa, 0x85, 0xcb, 0x76, 0x4e, 0xe1, 0xe1, 0x6b, 0x5a, 0x68, 0x1f,
	0xaa, 0x2f, 0xd9, 0xd2, 0x4f, 0xad, 0x39, 0xa2, 0x47, 0x50, 0x5f, 0x90, 0x7c, 0xce, 0xfc, 0xc4,
	0xba, 0xe0, 0xf3, 0xca, 0xa7, 0x41, 0xf7, 0x8f, 0x2a, 0x34, 0x9c, 0x13, 0x68, 0x00, 0x6f, 0x97,
	0x53, 0x97, 0x28, 0x2d, 0x24, 0xc9, 0x58, 0xe2, 0x9d, 0x35, 0xb6, 0x3b, 0x39, 0x5c, 0x42, 0x46,
	0x0e, 0xe1, 0xb8, 0xa6, 0x0f, 0x23, 0x68, 0x4f, 0x78, 0xce, 0x92, 0x5b, 0xa2, 0x6e, 0x99, 0xc2,
	0x15, 0xeb, 0x4a, 0xff, 0xee, 0x0e, 0x44, 0x17, 0x3c, 0x67, 0x97, 0x96, 0xe4, 0xac, 0x80, 0xc9,
	0x2a, 0x81, 0x4e, 0xa1, 0x69, 0xbe, 0x72, 0xf6, 0x8b, 0xc6, 0x55, 0xdb, 0xd2, 0xf7, 0x5f, 0x13,
	0xf4, 0x1f, 0xb9, 0x13, 0x3b, 0x77, 0xe0, 0xb8, 0x64, 0xa1, 0x17, 0xf0, 0x26, 0xa1, 0x94, 0x1b,
	0x77, 0x48, 0x9e, 0xf8, 0x6c, 0xf9, 0x51, 0xdd, 0x53, 0x0c, 0xad, 0x15, 0x7c, 0x4a, 0x75, 0x18,
	0xec, 0xbd, 0x52, 0xf7, 0xbf, 0xd8, 0xfe, 0xc5, 0xa6, 0xed, 0xed, 0xfe, 0x07, 0xdb, 0xcc, 0x58,
	0xab, 0x6d, 0xb6, 0xe7, 0x6b, 0x80, 0xf5, 0x03, 0x34, 0x80, 0xd6, 0xca, 0x62, 0x1c, 0xd8, 0x57,
	0x38, 0xda, 0xa6, 0x69, 0x68, 0x71, 0x58, 0xda, 0xd9, 0xfd, 0x3d, 0x80, 0x9a, 0x39, 0xa0, 0x01,
	0xd4, 0xf4, 0x72, 0xc6, 0x6c, 0xa9, 0xbb, 0xfd, 0x8f, 0xee, 0x92, 0xb0, 0x7f, 0x6e, 0x96, 0x33,
	0x16, 0x5b, 0xda, 0x3f, 0xa7, 0x69, 0xc7, 0x97, 0xdb, 0x3d, 0x81, 0xb0, 0xc4, 0xa1, 0x03, 0x78,
	0xeb, 0xf2, 0x6c, 0x74, 0x99, 0xdc, 0x7c, 0x7f, 0x3d, 0x4c, 0x9e, 0x7f, 0x37, 0xba, 0x1e, 0x9e,
	0x5f, 0x5d, 0x5c, 0x0d, 0xbf, 0xdc, 0x7f, 0x03, 0x01, 0x34, 0x46, 0x97, 0x67, 0xfd, 0xc7, 0x9f,
	0xec, 0x07, 0xdd, 0x5f, 0x03, 0x68, 0xfa, 0x1d, 0x88, 0x10, 0xd4, 0x0a, 0x32, 0x65, 0xde, 0x3e,
	0x7b, 0x36, 0x8e, 0xb2, 0x62, 0x61, 0x47, 0xa9, 0x15, 0x9b, 0xa3, 0x41, 0x11, 0x99, 0xb9, 0xe5,
	0xda, 0x8a, 0xed, 0xd9, 0xa0, 0x28, 0x97, 0xb8, 0xe6, 0x7c, 0xa7, 0x5c, 0xfa, 0xad, 0x5d, 0x5f,
	0x6d, 0xed, 0x03, 0x08, 0x7f, 0x26, 0x5c, 0x27, 0x13, 0x21, 0x71, 0xc3, 0x32, 0x9b, 0x26, 0xbe,
	0x10, 0xb2, 0xfb, 0x0d, 0x84, 0xe5, 0xba, 0x44, 0x1d, 0x08, 0xd3, 0x5b, 0x96, 0xbe, 0x54, 0xf3,
	0xa9, 0x2f, 0x63, 0x15, 0x7b, 0xc9, 0xca, 0x4a, 0xf2, 0x11, 0xd4, 0x4d, 0x89, 0x65, 0x25, 0x2e,
	0x78, 0xfa, 0x5b, 0x00, 0x87, 0x5c, 0x6c, 0xf1, 0xf3, 0x3a, 0xf8, 0xe1, 0x47, 0xbf, 0x08, 0x33,
	0x91, 0x93, 0x22, 0x8b, 0x84, 0xcc, 0x7a, 0x19, 0x2b, 0xec, 0x5a, 0xec, 0xb9, 0x47, 0x64, 0xc6,
	0xd5, 0xff, 0xfb, 0x21, 0x7d, 0xb2, 0x3e, 0xfe, 0x59, 0xa9, 0x7e, 0x15, 0x9f, 0x8d, 0x1b, 0x56,
	0xf0, 0xe3, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x79, 0x79, 0x09, 0xaa, 0x98, 0x07, 0x00, 0x00,
}
