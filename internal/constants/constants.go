package constants

const DATA_EVENTID_FORMAT string = "id: %s\n"
const DATA_TRANSMIT_FORMAT string = "data: %s\n\n"
const DATA_TYPE_FORMAT string = "event: %s\n"

// default size of the comms channel that will be
// initialized in EndpointFunc.
const DEFAULT_BUFFER_SIZE int = 256

// default value for the CORS MaxAge option. this value will
// be used to indicate the header be left blank.
const DEFAULT_MAXAGE int = 0

const GENERATE_ATTEMPT_MAX int = 10

const HEADER_ACALLOW_NAM string = "Access-Control-Allow-Origin"
const HEADER_ACALLOW_VAL string = "*"
const HEADER_ACCREDS_NAM string = "Access-Control-Allow-Credentials"
const HEADER_ACEXPOSE_NAM string = "Access-Control-Expose-Headers"
const HEADER_ACEXPOSE_VAL string = "Content-Type"
const HEADER_ACMAXAGE_NAM string = "Access-Control-Max-Age"
const HEADER_ACMETHODS_NAM string = "Access-Control-Allow-Methods"
const HEADER_ACCELBUFFER_NAM string = "X-Accel-Buffering"
const HEADER_ACCELBUFFER_VAL string = "no"
const HEADER_CACHE_NAM string = "Cache-Control"
const HEADER_CACHE_VAL string = "no-cache"
const HEADER_CONTENTTYPE_NAM string = "Content-Type"
const HEADER_CONTENTTYPE_VAL string = "text/event-stream"

const TYPE_MESSAGE string = "message"
