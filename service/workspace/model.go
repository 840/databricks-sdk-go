// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

// all definitions in this file are in alphabetical order

type Delete struct {
	// The absolute path of the notebook or directory.
	Path string `json:"path"`
	// The flag that specifies whether to delete the object recursively. It is
	// ``false`` by default. Please note this deleting directory is not atomic.
	// If it fails in the middle, some of objects under this directory may be
	// deleted and cannot be undone.
	Recursive bool `json:"recursive,omitempty"`
}

type ExportRequest struct {
	// Flag to enable direct download. If it is ``true``, the response will be
	// the exported file itself. Otherwise, the response contains content as
	// base64 encoded string.
	DirectDownload bool `json:"-" url:"direct_download,omitempty"`
	// This specifies the format of the exported file. By default, this is
	// ``SOURCE``. However it may be one of: ``SOURCE``, ``HTML``, ``JUPYTER``,
	// ``DBC``.
	//
	// The value is case sensitive.
	Format string `json:"-" url:"format,omitempty"`
	// The absolute path of the notebook or directory. Exporting directory is
	// only support for ``DBC`` format.
	Path string `json:"-" url:"path,omitempty"`
}

type ExportResponse struct {
	// The base64-encoded content. If the limit (10MB) is exceeded, exception
	// with error code **MAX_NOTEBOOK_SIZE_EXCEEDED** will be thrown.
	Content string `json:"content,omitempty"`
}

type GetStatusRequest struct {
	// The absolute path of the notebook or directory.
	Path string `json:"-" url:"path,omitempty"`
}

type Import struct {
	// The base64-encoded content. This has a limit of 10 MB.
	//
	// If the limit (10MB) is exceeded, exception with error code
	// **MAX_NOTEBOOK_SIZE_EXCEEDED** will be thrown. This parameter might be
	// absent, and instead a posted file will be used.
	Content string `json:"content,omitempty"`
	// This specifies the format of the file to be imported. By default, this is
	// ``SOURCE``. However it may be one of: ``SOURCE``, ``HTML``, ``JUPYTER``,
	// ``DBC``. The value is case sensitive.
	Format ImportFormat `json:"format,omitempty"`
	// The language. If format is set to ``SOURCE``, this field is required;
	// otherwise, it will be ignored.
	Language ImportLanguage `json:"language,omitempty"`
	// The flag that specifies whether to overwrite existing object. It is
	// ``false`` by default. For ``DBC`` format, ``overwrite`` is not supported
	// since it may contain a directory.
	Overwrite bool `json:"overwrite,omitempty"`
	// The absolute path of the notebook or directory. Importing directory is
	// only support for ``DBC`` format.
	Path string `json:"path"`
}

// This specifies the format of the file to be imported. By default, this is
// “SOURCE“. However it may be one of: “SOURCE“, “HTML“, “JUPYTER“,
// “DBC“. The value is case sensitive.
type ImportFormat string

const ImportFormatDbc ImportFormat = `DBC`

const ImportFormatHtml ImportFormat = `HTML`

const ImportFormatJupyter ImportFormat = `JUPYTER`

const ImportFormatRMarkdown ImportFormat = `R_MARKDOWN`

const ImportFormatSource ImportFormat = `SOURCE`

// The language. If format is set to “SOURCE“, this field is required;
// otherwise, it will be ignored.
type ImportLanguage string

const ImportLanguagePython ImportLanguage = `PYTHON`

const ImportLanguageR ImportLanguage = `R`

const ImportLanguageScala ImportLanguage = `SCALA`

const ImportLanguageSql ImportLanguage = `SQL`

type ListRequest struct {
	// <content needed>
	NotebooksModifiedAfter int `json:"-" url:"notebooks_modified_after,omitempty"`
	// The absolute path of the notebook or directory.
	Path string `json:"-" url:"path,omitempty"`
}

type ListResponse struct {
	// List of objects.
	Objects []ObjectInfo `json:"objects,omitempty"`
}

type Mkdirs struct {
	// The absolute path of the directory. If the parent directories do not
	// exist, it will also create them. If the directory already exists, this
	// command will do nothing and succeed.
	Path string `json:"path"`
}

type ObjectInfo struct {
	// The location (bucket and prefix) enum value of the content blob. This
	// field is used in conjunction with the blob_path field to determine where
	// the blob is located.
	BlobLocation ObjectInfoBlobLocation `json:"blob_location,omitempty"`
	// ========= File metadata. These values are set only if the object type is
	// ``FILE``. ===========//
	BlobPath string `json:"blob_path,omitempty"`
	// <content needed>
	ContentSha256Hex string `json:"content_sha256_hex,omitempty"`
	// <content needed>
	CreatedAt int64 `json:"created_at,omitempty"`
	// The language of the object. This value is set only if the object type is
	// ``NOTEBOOK``.
	Language ObjectInfoLanguage `json:"language,omitempty"`
	// <content needed>
	MetadataVersion int `json:"metadata_version,omitempty"`
	// <content needed>
	ModifiedAt int64 `json:"modified_at,omitempty"`
	// <content needed>
	ObjectId int64 `json:"object_id,omitempty"`
	// <content needed>
	ObjectType ObjectInfoObjectType `json:"object_type,omitempty"`
	// The absolute path of the object.
	Path string `json:"path,omitempty"`
	// <content needed>
	Size int64 `json:"size,omitempty"`
}

// The location (bucket and prefix) enum value of the content blob. This field
// is used in conjunction with the blob_path field to determine where the blob
// is located.
type ObjectInfoBlobLocation string

const ObjectInfoBlobLocationDbfsRoot ObjectInfoBlobLocation = `DBFS_ROOT`

const ObjectInfoBlobLocationInternalDbfsJobs ObjectInfoBlobLocation = `INTERNAL_DBFS_JOBS`

// The language of the object. This value is set only if the object type is
// “NOTEBOOK“.
type ObjectInfoLanguage string

const ObjectInfoLanguagePython ObjectInfoLanguage = `PYTHON`

const ObjectInfoLanguageR ObjectInfoLanguage = `R`

const ObjectInfoLanguageScala ObjectInfoLanguage = `SCALA`

const ObjectInfoLanguageSql ObjectInfoLanguage = `SQL`

// <content needed>
type ObjectInfoObjectType string

const ObjectInfoObjectTypeDirectory ObjectInfoObjectType = `DIRECTORY`

const ObjectInfoObjectTypeFile ObjectInfoObjectType = `FILE`

const ObjectInfoObjectTypeLibrary ObjectInfoObjectType = `LIBRARY`

const ObjectInfoObjectTypeMlflowExperiment ObjectInfoObjectType = `MLFLOW_EXPERIMENT`

const ObjectInfoObjectTypeNotebook ObjectInfoObjectType = `NOTEBOOK`

const ObjectInfoObjectTypeProject ObjectInfoObjectType = `PROJECT`

const ObjectInfoObjectTypeRepo ObjectInfoObjectType = `REPO`
