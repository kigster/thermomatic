package imei

type fixture struct {
	imei        string
	decoded     uint64
	valid       bool
	decodeError error
}

var validIMEI1 = fixture{
	imei:        "379537021417898",
	valid:       true,
	decodeError: nil,
	decoded:     379537021417898,
}


var validIMEI2 = fixture{
	imei:        "000000000000000",
	valid:       true,
	decodeError: nil,
	decoded:     0,
}

var invalidIMEI1 = fixture{
	imei:        "01234567890XYZA",
	valid:       false,
	decodeError: ErrInvalid,
	decoded:     0,
}

var invalidIMEI2 = fixture{
	imei:        "01234",
	valid:       false,
	decodeError: ErrInvalid,
	decoded:     0,
}

var invalidIMEI3 = fixture{
	imei:        "379537021417897",
	valid:       false,
	decodeError: ErrChecksum,
	decoded:     0,
}


var fixtures = []fixture{
	validIMEI1,
	validIMEI2,
	invalidIMEI1,
	invalidIMEI2,
	invalidIMEI3,
}
