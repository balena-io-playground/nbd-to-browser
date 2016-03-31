package nbd

import ()

const (
	NBD_CMD_READ  = 0
	NBD_CMD_WRITE = 1
	NBD_CMD_DISC  = 2
	NBD_CMD_FLUSH = 3
	NBD_CMD_TRIM  = 4
)

const (
	NBD_CMD_FLAG_FUA = 1 << 0
)

const (
	NBD_FLAG_HAS_FLAGS  = 1 << 0
	NBD_FLAG_READ_ONLY  = 1 << 1
	NBD_FLAG_SEND_FLUSH = 1 << 2
	NBD_FLAG_SEND_FUA   = 1 << 3
	NBD_FLAG_ROTATIONAL = 1 << 4
	NBD_FLAG_SEND_TRIM  = 1 << 5
)

const (
	NBD_MAGIC         = 0x4e42444d41474943
	NBD_REQUEST_MAGIC = 0x25609513
	NBD_REPLY_MAGIC   = 0x67446698
	NBD_CLISERV_MAGIC = 0x00420281861253
	NBD_OPTS_MAGIC    = 0x49484156454F5054
	NBD_REP_MAGIC     = 0x3e889045565a9
)

const (
	NBD_DEFAULT_PORT = 10809
)

const (
	NBD_OPT_EXPORT_NAME = 1
	NBD_OPT_ABORT       = 2
	NBD_OPT_LIST        = 3
)

const (
	NBD_REP_ACK    = 1
	NBD_REP_SERVER = 2
)

const (
	NBD_REP_FLAG_ERROR   = 1 << 31
	NBD_REP_ERR_UNSUP    = 1 | NBD_REP_FLAG_ERROR
	NBD_REP_ERR_POLICY   = 2 | NBD_REP_FLAG_ERROR
	NBD_REP_ERR_INVALID  = 3 | NBD_REP_FLAG_ERROR
	NBD_REP_ERR_PLATFORM = 4 | NBD_REP_FLAG_ERROR
)

const (
	NBD_FLAG_FIXED_NEWSTYLE   = 1 << 0
	NBD_FLAG_NO_ZEROES        = 1 << 1
	NBD_FLAG_C_FIXED_NEWSTYLE = NBD_FLAG_FIXED_NEWSTYLE
	NBD_FLAG_C_NO_ZEROES      = NBD_FLAG_NO_ZEROES
)

const (
	NBD_EPERM  = 1
	NBD_EIO    = 5
	NBD_ENOMEM = 12
	NBD_EINVAL = 22
	NBD_ENOSPC = 28
)

type nbdNewStyleHeader struct {
	NbdMagic       uint64
	NbdOptsMagic   uint64
	NbdGlobalFlags uint16
}

type nbdClientFlags struct {
	NbdClientFlags uint32
}

type nbdClientOpt struct {
	NbdOptMagic uint64
	NbdOptId    uint32
	NbdOptLen   uint32
}

type nbdExportDetails struct {
	NbdExportSize  uint64
	NbdExportFlags uint16
}

type nbdOptReply struct {
	NbdOptReplyMagic  uint64
	NbdOptId          uint32
	NbdOptReplyType   uint32
	NbdOptReplyLength uint32
}

type nbdRequest struct {
	NbdRequestMagic uint32
	NbdCommandFlags uint16
	NbdCommandType  uint16
	NbdHandle       uint64
	NbdOffset       uint64
	NbdLength       uint32
}

type nbdReply struct {
	NbdReplyMagic uint32
	NbdError      uint32
	NbdHandle     uint64
}
