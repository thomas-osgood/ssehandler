package messages

const ERR_EMPTY_MAP string = "client map cannot be nil"
const ERR_EMPTY_HEADER_ALLOW string = "allow header must be a non-zero length string"
const ERR_EMPTY_HEADER_EXPOSE string = "expose header must be a non-zero length string"
const ERR_EMPTY_METHOD string = "method must be a non-zero length string"
const ERR_EMPTY_ORIGIN string = "origin must be a non-zero length string"
const ERR_GENERATEID_MAXATTEMPTS string = "unable to generate a client id in max number of attempts"
const ERR_MAXAGE_NOTPOS string = "max age must be a positive value"
const ERR_MAXMIN_LEN string = "min and max lengths must be greater than zero"
const ERR_MIN_LEN string = "min length must be less than or equal to max length"
const ERR_RANDSTR_LEN string = "error generating the length: %s"
const ERR_UNSUPPORTED string = "SSE not supported"
