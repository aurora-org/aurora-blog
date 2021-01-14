package status

const (
	OK      = 200
	Created = 201
	Updated = 202
	Deleted = 203

	BadRequest = 400
	Forbidden  = 403
	NotFound   = 404

	InternalServerError = 500
	NotImplemented      = 501
	ServiceUnavailable  = 503
)

var CodeMap = map[int]string{
	200: "OK",
	201: "Created",
	202: "Updated",
	203: "Deleted",
	400: "BadRequest",
	403: "Forbidden",
	404: "NotFound",
	500: "InternalServerError",
	501: "NotImplemented",
	503: "ServiceUnavailable",
}
